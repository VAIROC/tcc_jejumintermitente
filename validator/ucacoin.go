package validator

import "github.com/btcsuite/btcutil/base58"

// Ucacoin ...
type Ucacoin struct{}

var _ BitcoinLike = (*Ucacoin)(nil)

// ValidateAddress returns validate result of q