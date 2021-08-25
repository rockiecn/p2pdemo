// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


struct Cheque {
    address opAddr;
	address fromAddr;
	address toAddr;
	address tokenAddr;

	uint256 value ;
	uint256 nodeNonce;
}

struct PayCheque {
	Cheque cheque;
	bytes chequeSig;

	address cashAddr;
	address fromAddr;
	address toAddr;

	uint256 payValue;
}


contract Cash  {


    event ShowFrom(address);
    event ShowNonce(uint256);
    event ShowHash(bytes32);
    event ShowSigner(address);
    event ShowSig(bytes);
    event ShowPack(bytes);
    
    constructor() payable {

    }

    // called by storage
    function apply_cheque(PayCheque memory paycheque) public payable returns(bool) {
        
      
    emit ShowFrom(paycheque.cheque.fromAddr);
    emit ShowNonce(paycheque.cheque.nodeNonce);
    emit ShowSig(paycheque.chequeSig);
    
    bytes memory pack = abi.encodePacked(paycheque.cheque.fromAddr, paycheque.cheque.nodeNonce,"",uint256(0));
    emit ShowPack(pack);
    
    // hash =  cheque.from + cheque.nonce 
    bytes32 hash = keccak256(abi.encodePacked(paycheque.cheque.fromAddr, paycheque.cheque.nodeNonce,"",uint256(0)));
    emit ShowHash(hash);
    
    address signer = Recover.recover(hash,paycheque.chequeSig);
    emit ShowSigner(signer);
    
    // address tt = signer;
    // tt;
    
    if(paycheque.cheque.opAddr == signer) {
        uint256 weiPay;
        weiPay = paycheque.payValue * 1000000000000000000; // eth to wei
        payable(paycheque.toAddr).transfer(weiPay); //pay value to storage
        return true;
    }

    return false;

    
    // require(send == channelSender, "illegal sig");

    }


}