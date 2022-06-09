package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLitecoinValidateAddress(t *testing.T) {
	validator := &Litecoin{}

	var mainnetCases = map[string]*Result{
		"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ": {Success, true, P2PKH, ""},
		"LfzMXaeTd1ZvThWdTHvAiucgLBy461DdCf": {Success, true, P2PKH, ""},
		"LY1cKc26dhtGg6EJutjSH1QXbXPbRNdVZ8": {Success, true, P2PKH, ""},
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg": {Success, true, P2SH, ""},
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT": {Success, true, P2SH, ""},
		"MTf4tP1TCNBn8dNkyxeBVoPrFCcVzxJvvh": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"mhKYHWXLmQb93RuzyVmTK2KmM8KJPFrzdb": {Success, true, P2PKH, ""},
		"muAGGFMnSCes3hyixTDw1abTSrQDqpbVsP": {Success, true, P2PKH, ""},
		"mvvuJvego4AeD3gqk665JVF2Y