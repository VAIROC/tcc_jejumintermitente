package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
	"github.com/LanfordCai/ava/crypto"
)

// Kusama ...
type Kusama struct{}

var _ SS58 = (*Kusama)(nil)

// ValidateAddress returns validate result of