package validator

import (
	"fmt"
	"strings"

	"github.com/LanfordCai/ava/cashaddr"
)

// Classzz ...
type Classzz struct{}

var _ BitcoinLike = (*Classzz)(nil)
var _ CashAddress = (*Classzz)(nil)

// ValidateAddress returns validate result of classzz address
func (v *Classzz) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := v.CashAddrType(addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns classzz addres