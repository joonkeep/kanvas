import { promises as fs } from 'fs'

import '@nomiclabs/hardhat-ethers'
import { providers, Wallet } from 'ethers'
import { task, types } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  MessageStatus,
  ContractsLike,
} from '../src'

task('finalize-withdrawal', 'Finalize a withdrawal')
  .addParam(
    'transactionHash',
    'L2 Transaction hash to finalize',
    '',
    types.string
  )
  .addParam('l2Url', 'L2 HTTP URL', 'http://localhost:9545', types.string)
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre: HardhatRuntimeEnvironment) => {
    const txHash = args.transactionHash
    if (txHash === '') {
      console.log('No tx hash')
    }

    const signers = await hre.ethers.getSigners()
    if (signers.length === 0) {
      throw new Error('No configured signers')
    }
    const signer = signers[0]
    const address = await signer.getAddress()
    console.log(`Using signer: ${address}`)

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2Url)
    const l2Signer = new Wallet(hre.network.config.accounts[0], l2Provider)

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]
    if (args.l1ContractsJsonPath) {
      const data = await fs.readFile(args.l1ContractsJsonPath)
      contractAddrs = {
        l1: JSON.parse(data.toString()),
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as ContractsLike
    } else if (!contractAddrs) {
      const Deployment__L1CrossDomainMessenger = await hre.deployments.get(
        'L1CrossDomainMessengerProxy'
      )

      const Deployment__L1StandardBridge = await hre.deployments.get(
        'L1StandardBridgeProxy'
      )

      const Deployment__KanvasPortal = await hre.deployments.get(
        'KanvasPortalProxy'
      )

      const Deployment__L2OutputOracle = await hre.deployments.get(
        'L2OutputOracleProxy'
      )

      contractAddrs = {
        l1: {
          L1CrossDomainMessenger: Deployment__L1CrossDomainMessenger,
          L1StandardBridge: Deployment__L1StandardBridge,
          KanvasPortal: Deployment__KanvasPortal.address,
          L2OutputOracle: Deployment__L2OutputOracle.address,
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      }
    }

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await signer.getChainId(),
      l2ChainId,
      contracts: contractAddrs,
    })

    console.log('Waiting to be able to prove withdrawal')

    let receipt = await l2Provider.getTransactionReceipt(txHash)

    const proveInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(receipt)
      console.log(
        `Message status: ${
          MessageStatus[currentStatus]
        } (${await messenger.getLatestBlockNumber()} / ${receipt.blockNumber})`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        receipt,
        MessageStatus.READY_TO_PROVE
      )
    } finally {
      clearInterval(proveInterval)
    }

    const proveTx = await messenger.proveMessage(txHash)
    const proveReceipt = await proveTx.wait()
    console.log('Prove receipt', proveReceipt)

    const proveBlock = await hre.ethers.provider.getBlock(
      proveReceipt.blockHash
    )
    const finalizationPeriodSeconds =
      await messenger.getChallengePeriodSeconds()

    console.log(`Withdrawal proven at ${proveBlock.timestamp}`)
    console.log(`Finalization period is ${finalizationPeriodSeconds} seconds`)

    const proveBlockFinalizedSeconds =
      proveBlock.timestamp + finalizationPeriodSeconds
    console.log(
      `Waiting to be able to finalize withdrawal until ${proveBlockFinalizedSeconds}`
    )

    const finalizeInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(txHash)
      const lastBlock = await hre.ethers.provider.getBlock(-1)
      console.log(
        `Message status: ${MessageStatus[currentStatus]} (${lastBlock.number}:${lastBlock.timestamp})`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        txHash,
        MessageStatus.READY_FOR_RELAY
      )
    } finally {
      clearInterval(finalizeInterval)
    }

    const tx = await messenger.finalizeMessage(txHash)
    receipt = await tx.wait()
    console.log(receipt)
    console.log('Finalized withdrawal')
  })
