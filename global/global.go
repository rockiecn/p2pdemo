package global

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
)

const DEBUG bool = true

// package level variable
var Ctx context.Context
var Peerid peer.ID = ""
var UserIndex []string
var StorageIndex []string
var RemoteExist bool = false

var StrTokenAddr = "b213d01542d129806d664248a380db8b12059061" // token address

// var StrOperatorAddr = "9e0153496067c20943724b79515472195a7aedaa" // operator
// var StrFromAddr = "1ab6a9f2b90004c1269563b5da391250ede3c114"     // user
// var StrToAddr = "b213d01542d129806d664248a380db8b12059061"       // storage

// var StrOperatorSK = "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
// var StrUserSK = "b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f"
// var StrStorageSK = "aa03c94703e40a3f9e694a002dcb250182970917a7cd2346f2dfd337ada154f5"

var StrOperatorAddr = "5B38Da6a701c568545dCfcB03FcB875f56beddC4" // operator
var StrFromAddr = "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"     // user
var StrToAddr = "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"       // storage

var StrOperatorSK = "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
var StrUserSK = "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"
var StrStorageSK = "cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4"

var Increase string = "1000000000000000000"

var ContractAddress string
