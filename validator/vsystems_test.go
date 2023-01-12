package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: no testnet cases
func TestVsystemsValidateAddress(t *testing.T) {
	validator := &Vsystems{}

	var mainnetCases = map[string]*Result{
		"ARF12jvtjz9caUFmiwBeRe1SPRGQhUWKrtd": {Success, true, Normal, ""},
		"ARBQTCYws5FZAVtA1ZFLsGhBtPymr4Hp5CX": {Success, true, Normal, ""},
		"AR45wyKHZnmt7ujqJRT7b4hSk9wX1bjwDkz": {Success, true, Normal, ""},
		"ARCVpcq2i6rQ7kkzNeJ1jsMec6TLmC7RNHn": {Success, true, Normal, ""},
	