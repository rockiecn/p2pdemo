// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


struct Cheque {
    uint256 value ;         // value of the cheque, payvalue shoud not exceed value
    address tokenAddr;      // token address, point out which token to pay
    uint256 nonce;          // nonce of the cheque, cheque's nonce should not smaller than it.
	address fromAddr;       // buyer of this cheque, should be cheque's signer
	address toAddr;         // receiver of cheque's money, point out who to pay
	address opAddr;         // operator of this cheuqe, shuould be contract's owner
	address contractAddr;   // contract address, should be this contract
}

struct PayCheque {
	Cheque cheque;
	bytes chequeSig;        // signer of this signature should be fromAddr.

	uint256 payValue;       // money to pay, should not exceed value.
}


contract Cash  {

    event Received(address, uint256);
    event Paid(address, uint256);
    
    address owner;
    mapping(address => uint256) public nodeNonce;
    

    // constructor
    constructor() payable {
        owner = msg.sender;
    }
    
    // receiver
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    // called by storage
    function apply_cheque(PayCheque memory paycheque, bytes memory paychequeSig) public payable returns(bool) {
      
        
        require(paycheque.cheque.nonce >= nodeNonce[paycheque.cheque.toAddr], "cheque.nonce too old");
        require(paycheque.payValue <= paycheque.cheque.value, "payvalue should not exceed value of cheque.");
        require(paycheque.cheque.contractAddr == address(this), "contract address error");
        require(paycheque.cheque.toAddr == msg.sender, "caller shuould be cheque.toAddr");
        require(paycheque.cheque.opAddr == owner, "operator should be owner of this contract");
        
        
        // used for calc hash
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
            chequePack,
            paycheque.payValue
        );
        

        bytes32 chequeHash = keccak256(chequePack);
        bytes32 paychequeHash = keccak256(paychequePack);
        // get signer from signature
        address chequeSigner = Recover.recover(chequeHash,paycheque.chequeSig);
        address paychequeSigner = Recover.recover(paychequeHash,paychequeSig);
        
        require(paycheque.cheque.opAddr == chequeSigner, "illegal cheque sig");
        require(paycheque.cheque.fromAddr == paychequeSigner, "illegal paycheque sig");
        
        // pay
        uint256 weiPay;
        weiPay = paycheque.payValue * 1000000000000000000; // eth to wei
        payable(paycheque.cheque.toAddr).transfer(weiPay); //pay value to storage
        emit Paid(paycheque.cheque.toAddr, weiPay);
        
        
        // update nonce after paid
        nodeNonce[paycheque.cheque.toAddr] = paycheque.cheque.nonce+1;
        
        
        return true;
    }

    // get nonce of a specified node
    function get_node_nonce(address node) public view returns(uint256) {
        return nodeNonce[node];
    }
    
    // get owner of the contract
    function get_owner() public view returns(address) {
        return owner;
    }

}