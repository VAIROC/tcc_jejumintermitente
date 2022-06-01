package validator

// Litecoin ...
type Litecoin struct{}

var _ BitcoinLike = (*Litecoin)(nil)

// ValidateAddress returns validate result of litecoin address
func (v *Litecoin) ValidateAddress(addr string, network NetworkType) *Result