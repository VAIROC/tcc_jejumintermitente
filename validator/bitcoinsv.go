package validator

// BitcoinSV ...
type BitcoinSV struct{}

var _ BitcoinLike = (*BitcoinSV)(nil)

// ValidateAddress returns validate result of bitcoin address
func (v *BitcoinSV) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != 