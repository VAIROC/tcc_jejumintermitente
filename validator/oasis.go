package validator

import (
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

// Oasis ...
type Oasis struct {
	Client *RosettaClient
}

var _ RosettaValidator = (*Oasis)(nil)
var _ Bech32Address = (*Oasis)(nil)

// ValidateAddress returns validate result of stellar address
func (v *Oasis)