package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWayfcoinValidateAddress(t *testing.T) {
	validator := &Wayfcoin{}

	var mainnetCases = map[string]*Result{
		"WaaKv3fUm2s3H4qwpDTVdHNwjhxPoLSQFH": {Success, true, P2PKH, ""},
		"WcPLF1mvbjNvYaS12kjaK9Gz5Gg9ZHyMGu": {Success, true, P2PKH, ""},
		"WfdWw5VoKTgxAYYZdsMHvaJWYCy2n9gido": {Success, true, P2PKH, ""},
		"WgxcsZERSMwTT37dUMgngt7FoyUsJuh7sL": {Success, true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {Success, false, Unknown, ""},
		"WbxcsZERSMwTT37dUMgngt7FoyUsJuh7sL":         {Success, false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {Success, false, Unknown, ""},
		"UcehDN7bgkJ8VgCsM96r2W2UpRoiHSQizY ":        {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":         {Success, false, Unknown, ""},
		"WgxcsZErSMwTT37dUMgngt7Fo