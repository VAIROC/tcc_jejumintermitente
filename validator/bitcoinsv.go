package validator

// BitcoinSV ...
type BitcoinSV struct{}

var _ BitcoinLike = (*BitcoinSV)(nil)

// ValidateAddress returns validate result of bitcoin address
func 