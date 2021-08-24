// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


struct Cheque {
    string opAddr;
	string fromAddr;
	string toAddr;
	string tokenAddr;

	uint256 value ;
	uint256 nodeNonce;
}

struct PayCheque {
	Cheque cheque;
	bytes chequeSig;

	string cashAddr;
	string fromAddr;
	string toAddr;

	uint256 payValue;
}


contract Cash  {

    event ShowString(string);
    event Showuint(uint256);
    event Showbytes(bytes);
    
    constructor() payable {

    }

    // called by storage
    //function apply_cheque(string[] memory stringParams, uint256[] memory intParams, bytes memory bytesParam) external payable returns(bool) {
    function apply_cheque(PayCheque memory paycheque) public payable returns(bool) {
        
        // emit ShowString(stringParams[0]);
        // emit Showuint(intParams[0]);
        // emit Showbytes(bytesParam);
        
        emit ShowString(paycheque.cashAddr);
        emit Showuint(paycheque.cheque.nodeNonce);

    
        return true;
    }

}