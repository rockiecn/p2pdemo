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

var StrTokenAddr = "b213d01542d129806d664248a380db8b12059061"    // token address
var StrFromAddr = "1ab6a9f2b90004c1269563b5da391250ede3c114"     // user
var StrToAddr = "b213d01542d129806d664248a380db8b12059061"       // storage
var StrOperatorAddr = "9e0153496067c20943724b79515472195a7aedaa" // operator

var StrUserSK = "b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f"
var StrOperatorSK = "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
var StrStorageSK = "aa03c94703e40a3f9e694a002dcb250182970917a7cd2346f2dfd337ada154f5"

var Increase int64 = 1

var ContractAddress string
