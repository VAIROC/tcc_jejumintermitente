package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Vsystems ...
type Vsystems struct{}

var _ Validator = (*Vsystems)(nil)

const vsystemsAddrVersion = 5

// ValidateAddress returns validate result of vsystems address
func (v *Vsystems) ValidateAddress(addr strin