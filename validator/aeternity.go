package validator

import (
	"github.com/btcsuite/btcutil/base58"
)

// Aeternity ...
type Aeternity struct{}

var _ Validator = (*Aeternity)(nil)
var _ Prefixer = (*Aeternity)(nil)

// Validat