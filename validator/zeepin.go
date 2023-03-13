package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Zeepin ...
type Zeepin struct{}

var _ Validator = (*Zeepin)(nil)

const zeepinAddrVersion = 80

// ValidateAddress returns validate result of zeepin address
// mainnet address and testnet address are the same.
func (v *Zeepin) ValidateAddress(addr string, network NetworkTy