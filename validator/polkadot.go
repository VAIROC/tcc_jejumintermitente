package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
	"github.com/LanfordCai/ava/crypto"
)

// Polkadot ...
type Polkadot struct{}

var _ SS58 = (*Polkadot)(n