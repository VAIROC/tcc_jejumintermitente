package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
	"github.com/LanfordCai/ava/crypto"
)

// ChainX ...
type ChainX struct{}

var _ SS58 = (*ChainX)(nil)

// ValidateAddress returns validate result of aeternity address
// see: https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)
func (v *ChainX) ValidateAddress(addr string, network NetworkType) *Result {
	decoded := base58check.BitcoinEncoder.Decode(addr)
	dataLen := len(decoded)
	// 1 byte type + 3 bytes address format
	if dataLen < 4 {
		return &Result{Success, false, Unknown, ""}
	}
	ss58AddrType := decoded[0]
	if ss58AddrType != v.Addr