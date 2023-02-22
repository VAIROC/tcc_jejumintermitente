
package validator

// Zcash ...
type Zcash struct{}

var _ ZcashLike = (*Zcash)(nil)

// ValidateAddress returns validate result of zcash address
func (v *Zcash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := ZcashlikeNormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns zcash address version according to the address type and
// network type
func (v *Zcash) AddressVersion(addrType AddressType, network NetworkType) []byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return []byte{28, 184}
		}
		return []byte{28, 189}
	}

	if addrType == P2PKH {
		return []byte{29, 37}
	}
	return []byte{28, 186}
}