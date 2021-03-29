package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinValidateAddress(t *testing.T) {
	validator := &Bitcoin{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF":                             {Success, true, P2PKH, ""},
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {Success, true, P2PKH, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {Success, true, P2PKH, ""},
		"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9":                             {Success, true, P2SH, ""},
		"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga":                             {Success, true, P2SH, ""},
		"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX":                             {Success, true, P2SH, ""},
		"bc1q2l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, true, P2WPKH, ""},
		"bc1q86ml6tnunc2cs30centm2dnqqammrzqhkflc98":                     {Success, true, P2WPKH, ""},
		"bc1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, true, P2WPKH, ""},
		"bc1qxcjkl0gyffz2tz935cepgetruee7n3kcva80a0xd9wgcyz93r2pqkgk