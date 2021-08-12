package sigapi

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/p2pdemo/utils"
)

// Sign msg with privateKey
func Sign(hash []byte, skByte []byte) (sigRet []byte, err error) {

	// byte to string, then string to ecdsa
	privateKeyECDSA, err := crypto.HexToECDSA(utils.Byte2Str(skByte))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// compute digest
	//digest := crypto.Keccak256Hash(msg)

	// sign to bytes
	sigByte, err := crypto.Sign(hash, privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//fmt.Println("len sigByte:", len(sigByte))
	//fmt.Println("len skByte:", len(skByte))

	return sigByte, nil

}

// Verify signature
func Verify(hash []byte, sigByte []byte, fromAddress common.Address) (ok bool, err error) {

	// compute digest
	//digest := crypto.Keccak256Hash(msg)

	// signature to public key
	//pubKeyECDSA, err := crypto.SigToPub(digest.Bytes(), sigByte)
	pubKeyECDSA, err := crypto.SigToPub(hash, sigByte)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}
	//fmt.Println("pubKeyECDSA", pubKeyECDSA)

	// pub key to address
	recoveredAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	fmt.Println("fromAddress", fromAddress)
	fmt.Println("recoveredAddr", recoveredAddr)
	matches := (fromAddress == recoveredAddr)

	//fmt.Println("matches:", matches)

	return matches, nil

}
