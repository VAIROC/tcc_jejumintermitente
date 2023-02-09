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
		"s1MLh26NK6cMAPQqxBbuJcKVgo1C1tBkT3j": {Success, true