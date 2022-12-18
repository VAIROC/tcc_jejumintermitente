package validator

import "github.com/btcsuite/btcutil/base58"

// Ucacoin ...
type Ucacoin struct{}

var _ BitcoinLike = (*Ucacoin)(nil)

// ValidateAddress returns validate result of qtum address
// NOTICE: only support mainnet P2PKH now
func (v *Ucacoin) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, version, err := base58.CheckD