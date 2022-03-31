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
func (e *IOST) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := e.IsAddressFormatValid(addr, network); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := e.Client.GetAccount(addr)
	if err != nil {
		return &Result{Failure, false, Unknown, err.Error()}
	}

	if addrType == Unknown {
		