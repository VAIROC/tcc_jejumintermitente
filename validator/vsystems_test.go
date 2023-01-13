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
		"AR95ZdHV3Wx759r5io7gmb2GUz9vRtFF1F8": {Success, true, Normal, ""},
		"AREExiJHmLb15ePMTyajnt4wb2bD4BENsM4": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, r