package validator

import "github.com/btcsuite/btcutil/base58"

// Tron - Tron address validator
type Tron struct{}

var _ Validator = (*Tron)(nil)

// ValidateAddress - Check a Ripple address is valid or not
func (r *Tron) ValidateAddress(addr string, network NetworkType) *Res