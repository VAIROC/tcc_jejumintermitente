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
	if !e.isValidUn