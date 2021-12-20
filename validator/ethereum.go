package validator

import (
	"bytes"
	"encoding/hex"
	"strings"
	"unicode"

	"github.com/LanfordCai/ava/crypto"
)

// Ethereum ..
type Ethereum struct{}

var _ Validator = (*Ethereum)(nil)
var _ Prefixer = (*Ethereum)(nil)

// ValidateAddress - Check an Ethereum address is valid or not
func (e *Ethereum) ValidateAddress(addr string, network NetworkType) *Result {
	if !e.isValidUncheckedAddress(addr) {
		return &Result{Success, false, Unknown, ""}
	}

	if !e.withChecksum(addr) {
		return &Result{Success, true, HexWithoutChecksum, ""}
	}

	checksumedAddr := e.toChecksumedAddress(addr)
	if checksumedAddr != addr {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, HexWithChecksum, ""}
}

func (e *Ethereum) withChecksum(addr string) bool {
	noPrefixAddr := AddressWithoutPrefix(e, addr, Mainnet)

	return (strings.ToUpper(noPrefixAddr) != noPrefixAddr) &&
		(strings.ToLower(noPrefixAddr) != noPrefixAddr)
}

func (e *Ethereum) isValidUnche