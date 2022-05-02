package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
	"github.com/LanfordCai/ava/crypto"
)

// Kusama ...
type Kusama struct{}

var _ SS58 = (*Kusama)(nil)

// ValidateAddress returns validate result of aeternity address
// see: https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)
func (v *Kusama) ValidateAddress(addr stri