package validator

import (
	"github.com/btcsuite/btcutil/base58"
)

// Aeternity ...
type Aeternity struct{}

var _ Validator = (*Aeternity)(nil)
var _ Prefixer = (*Aeternity)(nil)

// ValidateAddress returns validate result of aeternity address
func (v *Aeternity) ValidateAddress(addr string, network NetworkType) *Result {
	if !IsPrefixValid(v, addr, network) {
		return &Result{Success, false, Unknown, ""}
	}

	encodedPubkey := AddressWithoutPrefix(v, addr, network)
	decoded, _, err := base58.CheckDecode(encodedPubkey)
	if err != nil || len(decoded) != 31 {
		return &Result{Success, false, Unknown, ""}
	}
	return &Result{Success, true, Normal, ""}
}

// GetP