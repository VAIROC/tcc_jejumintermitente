package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcacoinValidateAddress(t *testing.T) {
	validator := &Ucacoin{}

	var mainnetCases = map[string]*Result{
		"UaFYCsCHXtbW4EWmckYXP8BmGgezJrRdem": {Success, true, P2PKH, ""},
		"Uak5JfpmLyjJkGJP37eBWHVPQmYo9wLX7w": {Success, true, P2PKH, ""},
		"UbehDN7bgkJ8VgCsM96r2W2UpRoiHSQizY": {Success, true, P2PKH, ""},
		"UcAkUrswRtP6vwZQf2F26k68dyHqWG7Sw5": {Success, true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"UcDkUrswRtP6vwZQf2F26k68dyHqWG7Sw5":         {Success, false, Unknown, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":         {Success, false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {Success, false, Unknown, ""},
		"UcehDN7bgkJ8