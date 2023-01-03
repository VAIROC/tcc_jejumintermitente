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

// ZcashlikeNormalAddrType ...
func ZcashlikeNormalAddrType(v ZcashLike, addr string, network NetworkType) AddressType {
	decoded, version, err := base58.CheckDecode(addr)
	if 