package validator

// Cosmos ...
type Cosmos struct{}

var _ Validator = (*Cosmos)(nil)
var _ Bech32Address = (*Cosmos)(nil)

// ValidateAddress returns validate result of cosmos address
func (v *Cosmos) ValidateAd