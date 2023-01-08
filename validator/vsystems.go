package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Vsystems ...
type Vsystems struct{}

var _ Validator = (*Vsystems)(nil)

const vsystemsAddrVersion = 5

// ValidateAddress returns validate result of vsystems address
func (v *Vsystems) ValidateAddress(addr string, network NetworkType) *Result {
	encoder := base58check.BitcoinEncoder
	encoder.ChecksumType = base58check.ChecksumBlake2bKeccak256

	decoded, err := e