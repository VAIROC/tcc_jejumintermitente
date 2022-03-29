package validator

import (
	"regexp"
)

// IOST address validator
type IOST struct {
	Client *IOSTClient
}

var _ OnchainValidator = (*IOST)(nil)

// ValidateAddress ...
func (e *IOST) ValidateAddress(addr 