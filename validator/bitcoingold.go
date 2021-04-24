package validator

// BitcoinGold ...
type BitcoinGold struct{}

var _ BitcoinLike = (*BitcoinGold)(nil)

// ValidateAddress returns validate result of bitcoingold address
func (v *BitcoinGold) ValidateAddress(addr string, 