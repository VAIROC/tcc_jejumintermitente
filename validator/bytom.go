package validator

// Bytom ...
type Bytom struct{}

var _ Validator = (*Bytom)(nil)
var _ SegwitAddress = (*Bytom)(nil)

// ValidateAddress returns validate result of bytom address
func (v *Bytom) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := SegwitAddrType(v, addr, network); addrType != Unknown {