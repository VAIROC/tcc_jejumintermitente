package validator

import "strconv"

// Tera ...
type Tera struct {
	Client *TeraClient
}

var _ OnchainValidator = (*Tera)(nil)

// ValidateAddress returns validate result of stellar address
func (v *Tera) ValidateAddress(addr string, network NetworkType) *Result {
	if