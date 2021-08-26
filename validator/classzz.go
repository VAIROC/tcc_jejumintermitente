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

// AddressVersion returns classzz address version according to the address type and
// network type
func (v *Classzz) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Testnet || addrType == P2SH {
		panic(ErrUnsupported.Error())
	}

	return 192
}

// CashAddrType ...
// NOTE: only support CashAddrP2PKH
func (v *Classzz) CashAddrType(addr string, network NetworkType) AddressType {
	if network == Testnet {
		panic(ErrUnsupported.Error())
	}

	if !strings.HasPrefix(addr, "classzz:") {
		addr = fmt