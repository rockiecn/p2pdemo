// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


struct Cheque {
    uint256 value ;
    address tokenAddr;
    uint256 nonce;
	address fromAddr;
	address toAddr;
	address opAddr;
	address contractAddr;
}

struct PayCheque {
	Cheque cheque;
	bytes chequeSig;

	uint256 payValue;
}


contract Cash  {

    event ShowFrom(address);
    event ShowNonce(uint256);
    event ShowChequeHash(bytes32);
    event ShowPayChequeHash(bytes32);
    event ShowChequeSigner(address);
    event ShowChequeSig(bytes);
    event ShowPayChequeSigner(address);
    event ShowPayCheckPack(bytes);
    
    mapping(address => uint256) public nodeNonce;
    

    constructor() payable {

    }
    
    event Received(address, uint);
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    // called by storage
    function apply_cheque(PayCheque memory paycheque, bytes memory paychequeSig) public payable returns(bool) {
      
        // emit ShowFrom(paycheque.cheque.fromAddr);
        // emit ShowNonce(paycheque.cheque.nonce);
        // emit ShowChequeSig(paycheque.chequeSig);
        
        require(paycheque.cheque.nonce >= nodeNonce[paycheque.cheque.toAddr], "cheque.nonce too old");
        require(paycheque.cheque.toAddr == msg.sender, "sender shuould be cheque.toAddr");
        
        //bytes memory pack = abi.encodePacked(paycheque.cheque.fromAddr, paycheque.cheque.nodeNonce,"",uint256(0));
        bytes memory chequePack = 
        abi.encodePacked(
            paycheque.cheque.value,
            paycheque.cheque.tokenAddr,
            paycheque.cheque.nonce,
            paycheque.cheque.fromAddr,
            paycheque.cheque.toAddr,
            paycheque.cheque.opAddr,
            paycheque.cheque.contractAddr
    		);
    	
        bytes memory paychequePack = 
        abi.encodePacked(
            paycheque.cheque.value,
            paycheque.cheque.tokenAddr,
            paycheque.cheque.nonce,
            paycheque.cheque.fromAddr,
            paycheque.cheque.toAddr,
            paycheque.cheque.opAddr,
            paycheque.cheque.contractAddr,
            paycheque.payValue
        );
    		
        // emit ShowPayCheckPack(paychequePack);
        
        // hash =  cheque.from + cheque.nonce 
        bytes32 chequeHash = keccak256(chequePack);
        bytes32 paychequeHash = keccak256(paychequePack);
        // emit ShowChequeHash(chequeHash);
        // emit ShowPayChequeHash(paychequeHash);
        
        address chequeSigner = Recover.recover(chequeHash,paycheque.chequeSig);
        emit ShowChequeSigner(chequeSigner);
        
        address paychequeSigner = Recover.recover(paychequeHash,paychequeSig);
        emit ShowPayChequeSigner(paychequeSigner);
        
        require(paycheque.cheque.opAddr == chequeSigner, "illegal cheque sig");
        require(paycheque.cheque.fromAddr == paychequeSigner, "illegal paycheque sig");
        
        // transfer money to storage node
        uint256 weiPay;
        weiPay = paycheque.payValue * 1000000000000000000; // eth to wei
        payable(paycheque.cheque.toAddr).transfer(weiPay); //pay value to storage
        
        nodeNonce[paycheque.cheque.toAddr] = paycheque.cheque.nonce+1;
        
        emit ShowNonce(nodeNonce[paycheque.cheque.toAddr]);
        
        return true;
    }

    // get nonce of a specified node
    function get_node_nonce(address node) public view returns(uint256) {
        return nodeNonce[node];
    }

}