// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";


contract Cash  {

    event show(bytes);
    
    constructor() payable {

    }

    // called by storage
    function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes memory sign) external payable returns(bool) {

        // bytes32 proof = keccak256(abi.encodePacked(address(this), value, nonce, msg.sender));
        // require(proof == hash, "illegal hash");

        bytes32 hash = keccak256(abi.encodePacked(userAddr, nonce, stAddr, payAmount));

        address signer = Recover.recover(hash,sign);
        
        address tt = signer;
        tt;
        
        if(userAddr == signer) {
            uint256 weiPay;
            weiPay = payAmount * 1000000000000000000; // eth to wei
            payable(stAddr).transfer(weiPay); //pay value to storage
            return true;
        }
        else {
            return false;
        }
        
        // require(send == channelSender, "illegal sig");

    }

}
