package validator

// Bytom ...
type Bytom struct{}

var _ Validator = (*Bytom)(nil)
var _ SegwitAddress = (*Bytom)(nil)

// ValidateAddress returns validate result of bytom address
func (v *Bytom) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := SegwitAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}
	return &Result{Success, false, Unknown, ""}
}

// AddressHrp returns hrps of bytom according to the network
func (v *Bytom) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "bm"
	}
	return "tm"
}

// SegwitProgramLength returns segwit program length of bytom
func (v *Bytom) SegwitProgramLength(addrType AddressType) int {
	if addrType == P2WPKH {
		return 20
	} else if addrType == P2WS