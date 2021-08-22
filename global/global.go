package global

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
)

const DEBUG bool = true

// package level variable
var Ctx context.Context
var Peerid peer.ID = ""
var Index []string
var RemoteExist bool = false
var UserSK = []byte("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
var OperatorSK = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")

//var strOperatorSk string = string("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
var Increase int64 = 10