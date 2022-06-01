package validator

// Litecoin ...
type Litecoin struct{}

var _ BitcoinLike = (*Litecoin)(nil)

// ValidateAddress returns validate result of litecoin address
func (v *Litecoin) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns litecoin address version according to the address type and
// network type
func (v *Litecoi