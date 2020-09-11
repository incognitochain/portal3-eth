pragma solidity ^0.6.6;

contract Portal {

    event DepositETH(string custodianIncAddress, uint64 amount);
    event DepositTokens(string custodianIncAddress, uint64 amount, string tokenID);


    function DepositETH(string memory custodianIncAddress, uint64 amount) payable public {
        require(msg.value >= amount);
        emit DepositETH(custodianIncAddress, amount);
    }
}