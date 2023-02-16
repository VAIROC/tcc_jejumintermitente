package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: testnet cases
func TestYcashValidateAddress(t *testing.T) {
	validator := &Ycash{}

	var mainnetCases = map[string]*Result{
		"s1jF4pKgr8izDWNLZ9D1AtJfWFhgCY1Qexd": {Success, true, P2PKH, ""},
		"s1MLh26NK6cMAPQqxBbuJcKVgo1C1tBkT3j": {Success, true, P2PKH, ""},
		"s1gsQmQeieHAMh8caW9f5jDTVqZUmTsQBsX": {Success, true, P2PKH, ""},
		"s35qM2AnDARUhfFnWYirrcXbeTd4f7SE7T5": {Success, true, P2SH, ""},
		"s3KJ6JzL72o2rLXdwdpW9HjqW94NHCssbq8": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"s1MBh26NK6cMAPQqxBbuJcKVgo1C1tBkT3j":                            {Success, false, Unknown, ""},
		"s1jF4pKgr8izDWNLZ9D1AtJfWFMgCY1Qexd":                            {Success, false, Unknown, ""},
		"t1gsQmQeieHAMh8caW9f5jDTVqZUmTsQBsX":                            {Success, false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxii