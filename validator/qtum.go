package validator

// Qtum ...
type Qtum struct{}

var _ BitcoinLike = (*Qtum)(nil)

// ValidateAddress returns validate result of qtum address
func (v *Qtum) Validat