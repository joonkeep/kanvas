import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const L1StandardBridgeProxy = await getContractFromArtifact(
    hre,
    'L1StandardBridgeProxy'
  )

  await deploy(hre, 'KanvasMintableERC20Factory', {
    args: [L1StandardBridgeProxy.address],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        L1StandardBridgeProxy.address
      )
    },
  })
}

deployFn.tags = ['KanvasMintableERC20Factory', 'setup']

export default deployFn
