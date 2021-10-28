
package validator

// Dogecoin ...
type Dogecoin struct{}

var _ BitcoinLike = (*Dogecoin)(nil)

// ValidateAddress returns validate result of dogecoin address
func (v *Dogecoin) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns dogecoin address version according to the address type and
// network type
func (v *Dogecoin) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 30
		}
		return 22
	}

	if addrType == P2PKH {
		return 113
	}
	return 196
}