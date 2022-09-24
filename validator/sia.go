package validator

import (
	"encoding/hex"

	"github.com/LanfordCai/ava/crypto"
)

// Sia ...
type Sia struct{}

var _ Validator = (*Sia)(nil)

// ValidateAddress - Check a Sia addres