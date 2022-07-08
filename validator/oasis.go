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
func (v *Oasis) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := v.IsAddressFormatValid(addr, network); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := v.Client.GetAccount(v.NetworkIdentifier(network), addr)
	if err != nil {
		return &Result{Failure, false, Unk