package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQtumValidateAddress(t *testing.T) {
	validator := &Qtum{}

	var mainnetCases = map[string]*Result{
		"QhzgLgC9DYxn9sTXWfcsFdKbFbbx2DKDL4": {Success, true, P2PKH, ""},
		"QfFE2wjWdqPDX4TQESSdEyb5y9DqL5wYHA": {Success, true, P2PKH, ""},
		"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6QY": {Success, true, P2PKH, ""},
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg": {Success, true, P2SH, ""},
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT": {Success, true, P2SH, ""},
		"