// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { stdError } from "forge-std/Test.sol";

import { Types } from "../libraries/Types.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { Proxy } from "../universal/Proxy.sol";
import { L2OutputOracle_Initializer, NextImpl } from "./CommonTest.t.sol";

contract L2OutputOracleTest is L2OutputOracle_Initializer {
    bytes32 submittedOutput1 = keccak256(abi.encode(1));

    function test_constructor_succeeds() external {
        assertEq(oracle.VALIDATOR(), asserter);
        assertEq(oracle.CHALLENGER(), challenger);
        assertEq(oracle.SUBMISSION_INTERVAL(), submissionInterval);
        assertEq(oracle.latestBlockNumber(), startingBlockNumber);
        assertEq(oracle.startingBlockNumber(), startingBlockNumber);
        assertEq(oracle.startingTimestamp(), startingTimestamp);
    }

    function test_constructor_badTimestamp_reverts() external {
        vm.expectRevert("L2OutputOracle: starting L2 timestamp must be less than current time");

        new L2OutputOracle({
            _submissionInterval: submissionInterval,
            _l2BlockTime: l2BlockTime,
            _startingBlockNumber: startingBlockNumber,
            _startingTimestamp: block.timestamp + 1,
            _validator: asserter,
            _challenger: challenger,
            _finalizationPeriodSeconds: 7 days
        });
    }

    function test_constructor_l2BlockTimeZero_reverts() external {
        vm.expectRevert("L2OutputOracle: L2 block time must be greater than 0");
        new L2OutputOracle({
            _submissionInterval: submissionInterval,
            _l2BlockTime: 0,
            _startingBlockNumber: startingBlockNumber,
            _startingTimestamp: block.timestamp,
            _validator: asserter,
            _challenger: challenger,
            _finalizationPeriodSeconds: 7 days
        });
    }

    function testFuzz_constructor_submissionIntervalLteL2BlockTime_reverts(
        uint256 _submissionInterval,
        uint256 _l2BlockTime
    ) external {
        // Bound the _l2blockTime to be in the range of [1, type(uint256).max]
        _l2BlockTime = bound(_l2BlockTime, 1, type(uint256).max);
        // Roll the block number to _l2blockTime (the starting L2 timestamp must be less than or equal to the current time)
        vm.roll(_l2BlockTime);
        // Bound _submissionInterval to be less than or equal to _l2BlockTime
        _submissionInterval = bound(_submissionInterval, 0, _l2BlockTime);

        vm.expectRevert("L2OutputOracle: submission interval must be greater than L2 block time");
        new L2OutputOracle({
            _submissionInterval: _submissionInterval,
            _l2BlockTime: _l2BlockTime,
            _startingBlockNumber: startingBlockNumber,
            _startingTimestamp: block.timestamp,
            _validator: asserter,
            _challenger: challenger,
            _finalizationPeriodSeconds: 7 days
        });
    }

    /****************
     * Getter Tests *
     ****************/

    // Test: latestBlockNumber() should return the correct value
    function test_latestBlockNumber_succeeds() external {
        uint256 submittedNumber = oracle.nextBlockNumber();

        // Roll to after the block number we'll submit
        warpToSubmitTime(submittedNumber);
        vm.prank(asserter);
        oracle.submitL2Output(submittedOutput1, submittedNumber, 0, 0);
        assertEq(oracle.latestBlockNumber(), submittedNumber);
    }

    // Test: getL2Output() should return the correct value
    function test_getL2Output_succeeds() external {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        oracle.submitL2Output(submittedOutput1, nextBlockNumber, 0, 0);

        Types.CheckpointOutput memory output = oracle.getL2Output(nextOutputIndex);
        assertEq(output.outputRoot, submittedOutput1);
        assertEq(output.timestamp, block.timestamp);

        // The block number is larger than the latest submitted output:
        vm.expectRevert(stdError.indexOOBError);
        oracle.getL2Output(nextOutputIndex + 1);
    }

    // Test: getL2OutputIndexAfter() returns correct value when input is exact block
    function test_getL2OutputIndexAfter_sameBlock_succeeds() external {
        bytes32 output1 = keccak256(abi.encode(1));
        uint256 nextBlockNumber1 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber1);
        vm.prank(asserter);
        oracle.submitL2Output(output1, nextBlockNumber1, 0, 0);

        // Querying with exact same block as submitted returns the output.
        uint256 index1 = oracle.getL2OutputIndexAfter(nextBlockNumber1);
        assertEq(index1, 0);
    }

    // Test: getL2OutputIndexAfter() returns correct value when input is previous block
    function test_getL2OutputIndexAfter_previousBlock_succeeds() external {
        bytes32 output1 = keccak256(abi.encode(1));
        uint256 nextBlockNumber1 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber1);
        vm.prank(asserter);
        oracle.submitL2Output(output1, nextBlockNumber1, 0, 0);

        // Querying with previous block returns the output too.
        uint256 index1 = oracle.getL2OutputIndexAfter(nextBlockNumber1 - 1);
        assertEq(index1, 0);
    }

    // Test: getL2OutputIndexAfter() returns correct value during binary search
    function test_getL2OutputIndexAfter_multipleOutputsExist_succeeds() external {
        bytes32 output1 = keccak256(abi.encode(1));
        uint256 nextBlockNumber1 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber1);
        vm.prank(asserter);
        oracle.submitL2Output(output1, nextBlockNumber1, 0, 0);

        bytes32 output2 = keccak256(abi.encode(2));
        uint256 nextBlockNumber2 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber2);
        vm.prank(asserter);
        oracle.submitL2Output(output2, nextBlockNumber2, 0, 0);

        bytes32 output3 = keccak256(abi.encode(3));
        uint256 nextBlockNumber3 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber3);
        vm.prank(asserter);
        oracle.submitL2Output(output3, nextBlockNumber3, 0, 0);

        bytes32 output4 = keccak256(abi.encode(4));
        uint256 nextBlockNumber4 = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber4);
        vm.prank(asserter);
        oracle.submitL2Output(output4, nextBlockNumber4, 0, 0);

        // Querying with a block number between the first and second output
        uint256 index1 = oracle.getL2OutputIndexAfter(nextBlockNumber1 + 1);
        assertEq(index1, 1);

        // Querying with a block number between the second and third output
        uint256 index2 = oracle.getL2OutputIndexAfter(nextBlockNumber2 + 1);
        assertEq(index2, 2);

        // Querying with a block number between the third and fourth output
        uint256 index3 = oracle.getL2OutputIndexAfter(nextBlockNumber3 + 1);
        assertEq(index3, 3);
    }

    // Test: getL2OutputIndexAfter() reverts when no output exists yet
    function test_getL2OutputIndexAfter_noOutputsExis_reverts() external {
        vm.expectRevert("L2OutputOracle: cannot get output as no outputs have been submitted yet");
        oracle.getL2OutputIndexAfter(0);
    }

    // Test: nextBlockNumber() should return the correct value
    function test_nextBlockNumber_succeeds() external {
        assertEq(
            oracle.nextBlockNumber(),
            // The return value should match this arithmetic
            oracle.latestBlockNumber() + oracle.SUBMISSION_INTERVAL()
        );
    }

    function test_computeL2Timestamp_succeeds() external {
        // reverts if timestamp is too low
        vm.expectRevert(stdError.arithmeticError);
        oracle.computeL2Timestamp(startingBlockNumber - 1);

        // returns the correct value...
        // ... for the very first block
        assertEq(oracle.computeL2Timestamp(startingBlockNumber), startingTimestamp);

        // ... for the first block after the starting block
        assertEq(
            oracle.computeL2Timestamp(startingBlockNumber + 1),
            startingTimestamp + l2BlockTime
        );

        // ... for some other block number
        assertEq(
            oracle.computeL2Timestamp(startingBlockNumber + 96024),
            startingTimestamp + l2BlockTime * 96024
        );
    }

    /*****************************
     * Submit Tests - Happy Path *
     *****************************/

    // Test: submitL2Output succeeds when given valid input, and no block hash and number are
    // specified.
    function test_submitL2Output_submitAnotherOutput_succeeds() public {
        bytes32 submittedOutput2 = keccak256(abi.encode());
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        warpToSubmitTime(nextBlockNumber);
        uint256 submittedNumber = oracle.latestBlockNumber();

        // Ensure the submissionInterval is enforced
        assertEq(nextBlockNumber, submittedNumber + submissionInterval);

        vm.roll(nextBlockNumber + 1);

        vm.expectEmit(true, true, true, true);
        emit OutputSubmitted(submittedOutput2, nextOutputIndex, nextBlockNumber, block.timestamp);

        vm.prank(asserter);
        oracle.submitL2Output(submittedOutput2, nextBlockNumber, 0, 0);
    }

    // Test: submitL2Output succeeds when given valid input, and when a block hash and number are
    // specified for reorg protection.
    function test_submitWithBlockhashAndHeight_succeeds() external {
        // Get the number and hash of a previous block in the chain
        uint256 prevL1BlockNumber = block.number - 1;
        bytes32 prevL1BlockHash = blockhash(prevL1BlockNumber);

        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        oracle.submitL2Output(nonZeroHash, nextBlockNumber, prevL1BlockHash, prevL1BlockNumber);
    }

    /***************************
     * Submit Tests - Sad Path *
     ***************************/

    // Test: submitL2Output fails if called by a party that is not the validator.
    function test_submitL2Output_notValidator_reverts() external {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);

        vm.prank(address(128));
        vm.expectRevert("L2OutputOracle: only the validator address can submit new outputs");
        oracle.submitL2Output(nonZeroHash, nextBlockNumber, 0, 0);
    }

    // Test: submitL2Output fails given a zero blockhash.
    function test_submitL2Output_emptyOutput_reverts() external {
        bytes32 outputToSubmit = bytes32(0);
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        vm.expectRevert("L2OutputOracle: L2 checkpoint output cannot be the zero hash");
        oracle.submitL2Output(outputToSubmit, nextBlockNumber, 0, 0);
    }

    // Test: submitL2Output fails if the block number doesn't match the next expected number.
    function test_submitL2Output_unexpectedBlockNumber_reverts() external {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        vm.expectRevert("L2OutputOracle: block number must be equal to next expected block number");
        oracle.submitL2Output(nonZeroHash, nextBlockNumber - 1, 0, 0);
    }

    // Test: submitL2Output fails if it would have a timestamp in the future.
    function test_submitL2Output_futureTimetamp_reverts() external {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        uint256 nextTimestamp = oracle.computeL2Timestamp(nextBlockNumber);
        vm.warp(nextTimestamp);
        vm.prank(asserter);
        vm.expectRevert("L2OutputOracle: cannot submit L2 output in the future");
        oracle.submitL2Output(nonZeroHash, nextBlockNumber, 0, 0);
    }

    // Test: submitL2Output fails if a non-existent L1 block hash and number are provided for reorg
    // protection.
    function test_submitL2Output_wrongFork_reverts() external {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        vm.expectRevert(
            "L2OutputOracle: block hash does not match the hash at the expected height"
        );
        oracle.submitL2Output(
            nonZeroHash,
            nextBlockNumber,
            bytes32(uint256(0x01)),
            block.number - 1
        );
    }

    // Test: submitL2Output fails when given valid input, but the block hash and number do not
    // match.
    function test_submitL2Output_unmatchedBlockhash_reverts() external {
        // Move ahead to block 100 so that we can reference historical blocks
        vm.roll(100);

        // Get the number and hash of a previous block in the chain
        uint256 l1BlockNumber = block.number - 1;
        bytes32 l1BlockHash = blockhash(l1BlockNumber);

        uint256 nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);

        // This will fail when foundry no longer returns zerod block hashes
        vm.expectRevert(
            "L2OutputOracle: block hash does not match the hash at the expected height"
        );
        oracle.submitL2Output(nonZeroHash, nextBlockNumber, l1BlockHash, l1BlockNumber - 1);
    }

    /*****************************
     * Delete Tests - Happy Path *
     *****************************/

    function test_deleteOutputs_singleOutput_succeeds() external {
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();

        uint256 latestBlockNumber = oracle.latestBlockNumber();
        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.CheckpointOutput memory newLatestOutput = oracle.getL2Output(latestOutputIndex - 1);

        vm.prank(challenger);
        vm.expectEmit(true, true, false, false);
        emit OutputsDeleted(latestOutputIndex + 1, latestOutputIndex);
        oracle.deleteL2Outputs(latestOutputIndex);

        // validate latestBlockNumber has been reduced
        uint256 latestBlockNumberAfter = oracle.latestBlockNumber();
        uint256 latestOutputIndexAfter = oracle.latestOutputIndex();
        assertEq(latestBlockNumber - submissionInterval, latestBlockNumberAfter);

        // validate that the new latest output is as expected.
        Types.CheckpointOutput memory output = oracle.getL2Output(latestOutputIndexAfter);
        assertEq(newLatestOutput.outputRoot, output.outputRoot);
        assertEq(newLatestOutput.timestamp, output.timestamp);
    }

    function test_deleteOutputs_multipleOutputs_succeeds() external {
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();

        uint256 latestBlockNumber = oracle.latestBlockNumber();
        uint256 latestOutputIndex = oracle.latestOutputIndex();
        Types.CheckpointOutput memory newLatestOutput = oracle.getL2Output(latestOutputIndex - 3);

        vm.prank(challenger);
        vm.expectEmit(true, true, false, false);
        emit OutputsDeleted(latestOutputIndex + 1, latestOutputIndex - 2);
        oracle.deleteL2Outputs(latestOutputIndex - 2);

        // validate latestBlockNumber has been reduced
        uint256 latestBlockNumberAfter = oracle.latestBlockNumber();
        uint256 latestOutputIndexAfter = oracle.latestOutputIndex();
        assertEq(latestBlockNumber - submissionInterval * 3, latestBlockNumberAfter);

        // validate that the new latest output is as expected.
        Types.CheckpointOutput memory output = oracle.getL2Output(latestOutputIndexAfter);
        assertEq(newLatestOutput.outputRoot, output.outputRoot);
        assertEq(newLatestOutput.timestamp, output.timestamp);
    }

    /***************************
     * Delete Tests - Sad Path *
     ***************************/

    function test_deleteL2Outputs_ifNotChallenger_reverts() external {
        uint256 latestBlockNumber = oracle.latestBlockNumber();

        vm.expectRevert("L2OutputOracle: only the challenger address can delete outputs");
        oracle.deleteL2Outputs(latestBlockNumber);
    }

    function test_deleteL2Outputs_nonExistent_reverts() external {
        test_submitL2Output_submitAnotherOutput_succeeds();

        uint256 latestBlockNumber = oracle.latestBlockNumber();

        vm.prank(challenger);
        vm.expectRevert("L2OutputOracle: cannot delete outputs after the latest output index");
        oracle.deleteL2Outputs(latestBlockNumber + 1);
    }

    function test_deleteL2Outputs_afterLatest_reverts() external {
        // Start by submitting three outputs
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();
        test_submitL2Output_submitAnotherOutput_succeeds();

        // Delete the latest two outputs
        uint256 latestOutputIndex = oracle.latestOutputIndex();
        vm.prank(challenger);
        oracle.deleteL2Outputs(latestOutputIndex - 2);

        // Now try to delete the same output again
        vm.prank(challenger);
        vm.expectRevert("L2OutputOracle: cannot delete outputs after the latest output index");
        oracle.deleteL2Outputs(latestOutputIndex - 2);
    }

    function test_deleteL2Outputs_finalized_reverts() external {
        test_submitL2Output_submitAnotherOutput_succeeds();

        // Warp past the finalization period + 1 second
        vm.warp(block.timestamp + oracle.FINALIZATION_PERIOD_SECONDS() + 1);

        uint256 latestOutputIndex = oracle.latestOutputIndex();

        // Try to delete a finalized output
        vm.prank(challenger);
        vm.expectRevert("L2OutputOracle: cannot delete outputs that have already been finalized");
        oracle.deleteL2Outputs(latestOutputIndex);
    }
}

contract L2OutputOracleUpgradeable_Test is L2OutputOracle_Initializer {
    Proxy internal proxy;

    function setUp() public override {
        super.setUp();
        proxy = Proxy(payable(address(oracle)));
    }

    function test_initValuesOnProxy_succeeds() external {
        assertEq(submissionInterval, oracleImpl.SUBMISSION_INTERVAL());
        assertEq(l2BlockTime, oracleImpl.L2_BLOCK_TIME());
        assertEq(startingBlockNumber, oracleImpl.startingBlockNumber());
        assertEq(startingTimestamp, oracleImpl.startingTimestamp());

        assertEq(asserter, oracleImpl.VALIDATOR());
        assertEq(challenger, oracleImpl.CHALLENGER());
    }

    function test_initializeProxy_alreadyInitialized_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        L2OutputOracle(payable(proxy)).initialize(startingBlockNumber, startingTimestamp);
    }

    function test_initializeImpl_alreadyInitialized_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        L2OutputOracle(oracleImpl).initialize(startingBlockNumber, startingTimestamp);
    }

    function test_upgrading_succeeds() external {
        // Check an unused slot before upgrading.
        bytes32 slot21Before = vm.load(address(oracle), bytes32(uint256(21)));
        assertEq(bytes32(0), slot21Before);

        NextImpl nextImpl = new NextImpl();
        vm.startPrank(multisig);
        proxy.upgradeToAndCall(
            address(nextImpl),
            abi.encodeWithSelector(NextImpl.initialize.selector)
        );
        assertEq(proxy.implementation(), address(nextImpl));

        // Verify that the NextImpl contract initialized its values according as expected
        bytes32 slot21After = vm.load(address(oracle), bytes32(uint256(21)));
        bytes32 slot21Expected = NextImpl(address(oracle)).slot21Init();
        assertEq(slot21Expected, slot21After);
    }
}
