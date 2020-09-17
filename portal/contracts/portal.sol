pragma solidity 0.6.6;
pragma experimental ABIEncoderV2;

import "./IERC20.sol";
import "./pause.sol";

/**
 * Math operations with safety checks
 */
library SafeMath {
    function safeMul(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a * b;
        require(a == 0 || c / a == b);
        return c;
    }

    function safeDiv(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0);
        uint256 c = a / b;
        require(a == b * c + a % b);
        return c;
    }

    function safeSub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        return a - b;
    }

    function safeAdd(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c>=a && c>=b);
        return c;
    }
}


/**
 * @dev Interface of the contract capable of checking if an instruction is
 * confirmed over at Incognito Chain
 */
interface Incognito {
    function instructionApproved(
        bool,
        bytes32,
        uint,
        bytes32[] calldata,
        bool[] calldata,
        bytes32,
        bytes32,
        uint[] calldata,
        uint8[] calldata,
        bytes32[] calldata,
        bytes32[] calldata
    ) external view returns (bool);
}

contract PortalV3 is AdminPausable {
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    address public delegator;
    Incognito public incognito;
    bool notEntered = true;

    using SafeMath for uint;
    mapping(bytes32 => bool) public withdrawed;
    struct BurnInstData {
        uint8 meta; // type of the instruction
        uint8 shard; // ID of the Incognito shard containing the instruction, must be 1
        address token; // ETH address of the token contract (0x0 for ETH)
        address payable to; // ETH address of the receiver of the token
        uint amount; // burned amount (on Incognito)
        bytes32 itx; // Incognito's burning tx
    }

    event Deposit(address tokenID, string custodianIncAddress, uint amount);
    event Withdraw(address token, address to, uint amount);
    event Delegator(address);
    event IncognitoProxy(address);

    function deposit(string calldata custodianIncAddress) isNotPaused  payable external {
        require(address(this).balance <= 10 ** 27, "max value reached");

        emit Deposit(ETH_TOKEN, custodianIncAddress, msg.value);
    }

    function depositERC20(address token, uint amount, string calldata custodianIncAddress) isNotPaused external {
        IERC20 erc20Interface = IERC20(token);
        uint8 decimals = getDecimals(address(token));
        uint tokenBalance = erc20Interface.balanceOf(address(this));
        uint beforeTransfer = tokenBalance;
        uint emitAmount = amount;
        if (decimals > 9) {
            emitAmount = emitAmount / (10 ** (uint(decimals) - 9));
            tokenBalance = tokenBalance / (10 ** (uint(decimals) - 9));
        }
        require(emitAmount <= 10 ** 18 && tokenBalance <= 10 ** 18 && emitAmount.safeAdd(tokenBalance) <= 10 ** 18, "max value reached");
        erc20Interface.transferFrom(msg.sender, address(this), amount);
        require(checkSuccess(), "transfer from got error");
        require(balanceOf(token).safeSub(beforeTransfer) == amount, "the input amount not equal to amount received");

        emit Deposit(token, custodianIncAddress, emitAmount);
    }

    /**
     * @dev Verifies that a burn instruction is valid
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @notice All params are the same as in `withdraw`
     */
    function verifyInst(
        bytes memory inst,
        uint heights,
        bytes32[] memory instPaths,
        bool[] memory instPathIsLefts,
        bytes32 instRoots,
        bytes32 blkData,
        uint[] memory sigIdxs,
        uint8[] memory sigVs,
        bytes32[] memory sigRs,
        bytes32[] memory sigSs
    ) view internal {
        // Each instruction can only by redeemed once
        bytes32 beaconInstHash = keccak256(abi.encodePacked(inst, heights));

        // Verify instruction on beacon
        require(incognito.instructionApproved(
                true, // Only check instruction on beacon
                beaconInstHash,
                heights,
                instPaths,
                instPathIsLefts,
                instRoots,
                blkData,
                sigIdxs,
                sigVs,
                sigRs,
                sigSs
            ), "invalid instruction data");
    }

    function withdrawLockedTokens(
        bytes memory inst,
        uint heights,
        bytes32[] memory instPaths,
        bool[] memory instPathIsLefts,
        bytes32 instRoots,
        bytes32 blkData,
        uint[] memory sigIdxs,
        uint8[] memory sigVs,
        bytes32[] memory sigRs,
        bytes32[] memory sigSs
    ) isNotPaused public {
        BurnInstData memory data = parseBurnInst(inst);
        require((data.meta == 119 || data.meta == 200) && data.shard == 1); // Check instruction type
        // Not withdrawed
        require(!withdrawed[data.itx], "withdraw transaction already used");
        withdrawed[data.itx] = true;

        // Update decimal if not ETH coin
        if (data.token != ETH_TOKEN) {
            uint8 decimals = getDecimals(data.token);
            if (decimals > 9) {
                data.amount = data.amount * (10 ** (uint(decimals) - 9));
            }
        }

        verifyInst(
            inst,
            heights,
            instPaths,
            instPathIsLefts,
            instRoots,
            blkData,
            sigIdxs,
            sigVs,
            sigRs,
            sigSs
        );

        // Send and notify
        if (data.token == ETH_TOKEN) {
            (bool success, ) =  data.to.call{value: data.amount}("");
            require(success, "internal transaction error");
        } else {
            IERC20(data.token).transfer(data.to, data.amount);
            require(checkSuccess(), "internal transaction error");
        }

        emit Withdraw(data.token, data.to, data.amount);
    }

    /**
     * @dev Parses a burn instruction and returns the components
     * @param inst: the full instruction, containing both metadata and body
     */
    function parseBurnInst(bytes memory inst) public pure returns (BurnInstData memory) {
        BurnInstData memory data;
        data.meta = uint8(inst[0]);
        data.shard = uint8(inst[1]);
        address token;
        address payable to;
        uint amount;
        bytes32 itx;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x22)) // [3:34]
            to := mload(add(inst, 0x42)) // [34:66]
            amount := mload(add(inst, 0x62)) // [66:98]
            itx := mload(add(inst, 0x82)) // [98:130]
        }
        data.token = token;
        data.to = to;
        data.amount = amount;
        data.itx = itx;
        return data;
    }

    /**
     * @dev Update delegator address
     * @param _delegator: delegator address
     */
    function updateDelegatorAddress(address _delegator) external onlyAdmin isPaused {
        delegator = _delegator;

        Delegator(delegator);
    }

    /**
     * @dev Update incognito proxy address
     * @param _incognitoProxy: incognito proxy address
     */
    function updateIncognitoAddress(address _incognitoProxy) external onlyAdmin isPaused {
        incognito = Incognito(_incognitoProxy);

        IncognitoProxy(delegator);
    }

    /**
    * @dev Check if transfer() and transferFrom() of ERC20 succeeded or not
    * This check is needed to fix https://github.com/ethereum/solidity/issues/4116
    * This function is copied from https://github.com/AdExNetwork/adex-protocol-eth/blob/master/contracts/libs/SafeERC20.sol
    */
    function checkSuccess() private pure returns (bool) {
        uint256 returnValue = 0;
        assembly {
        // check number of bytes returned from last function call
            switch returndatasize()

            // no bytes returned: assume success
            case 0x0 {
                returnValue := 1
            }

            // 32 bytes returned: check if non-zero
            case 0x20 {
            // copy 32 bytes into scratch space
                returndatacopy(0x0, 0x0, 0x20)

            // load those bytes into returnValue
                returnValue := mload(0x0)
            }

            // not sure what was returned: don't mark as success
            default { }
        }
        return returnValue != 0;
    }

    /**
     * @dev Get the decimals of an ERC20 token, return 0 if it isn't defined
     * We check the returndatasize to covert both cases that the token has
     * and doesn't have the function decimals()
     */
    function getDecimals(address token) public view returns (uint8) {
        IERC20 erc20 = IERC20(token);
        return uint8(erc20.decimals());
    }

    /**
     * @dev Get the amount of coin deposited to this smartcontract
     */
    function balanceOf(address token) public view returns (uint) {
        if (token == ETH_TOKEN) {
            return address(this).balance;
        }
        return IERC20(token).balanceOf(address(this));
    }
}