package validator

import "regexp"

// Bitshares address validator
type Bitshares struct {
	Client *BitsharesClient
}

var _ OnchainValidator = (*Bitshares)(nil)

// ValidateAddress ...
func (e *Bitshares) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := e.IsAddressFormatValid(addr, network); !isValid {
		return &Result{Success, fal