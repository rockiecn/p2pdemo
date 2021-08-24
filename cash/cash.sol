// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


contract Cash  {

    event ShowString(string);
    event Showuint(uint256);
    event Showbytes(bytes);
    
    constructor() payable {

    }

    // called by storage
    function apply_cheque(string[] memory stringParams, uint256[] memory intParams, bytes memory bytesParam) external payable returns(bool) {
        
        emit ShowString(stringParams[0]);
        emit Showuint(intParams[0]);
        emit Showbytes(bytesParam);
    
        return true;
    }

}