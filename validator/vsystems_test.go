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
		"ARBQTCYws5FZAVtA1ZFLsGhBtPymr4Hp5CX": {Success, true, No