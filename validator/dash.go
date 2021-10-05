package validator

// Dash ...
type Dash struct{}

var _ BitcoinLike = (*Dash)(nil)

// ValidateAddress returns validate result of dash address
fun