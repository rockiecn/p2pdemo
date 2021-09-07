package global

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
)

const DEBUG bool = true

// package level variable
var Ctx context.Context
var Peerid peer.ID = ""
var RemoteExist bool = false
