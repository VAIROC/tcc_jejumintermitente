
package validator

// Ycash ...
type Ycash struct{}

var _ ZcashLike = (*Ycash)(nil)

// ValidateAddress returns validate result of ycash address
func (v *Ycash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := ZcashlikeNormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}