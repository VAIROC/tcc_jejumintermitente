package validator

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
)

// ZcashLike ...
type ZcashLike interface {
	Validator
	AddressVersion(addrType AddressType, network NetworkType) []byte
}

// ZcashlikeNormalAddrType .