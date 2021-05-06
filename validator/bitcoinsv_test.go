package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinSVValidateAddress(t *testing.T) {
	validator := &BitcoinSV{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF": {Success, true, P2PKH, ""},
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw": {Success, true, P2PKH, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854": {Success, true, P2PKH, ""},
		"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9": {Success, true, P2SH, ""},
		"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga": {Success, true, P2SH, ""},
		"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	