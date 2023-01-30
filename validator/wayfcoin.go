
package validator

import "github.com/btcsuite/btcutil/base58"

// Wayfcoin ...
type Wayfcoin struct{}

var _ BitcoinLike = (*Wayfcoin)(nil)

// ValidateAddress returns validate result of qtum address
// NOTICE: only support mainnet P2PKH now