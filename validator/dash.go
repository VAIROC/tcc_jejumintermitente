package validator

// Dash ...
type Dash struct{}

var _ BitcoinLike = (*Dash)(nil)

// ValidateAddress returns validate result of dash address
func (v *Dash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns dash address version according to the address type and
// networ