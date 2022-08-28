package validator

// Qtum ...
type Qtum struct{}

var _ BitcoinLike = (*Qtum)(nil)

// ValidateAddress returns validate result of qtum address
func (v *Qtum) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}