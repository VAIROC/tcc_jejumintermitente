package validator

import "regexp"

// Bitshares address validator
type Bitshares struct {
	Client *BitsharesClient
}

var _ OnchainValidator = (*Bitshares)(nil)

// ValidateAddress ...
func (e *Bitshares)