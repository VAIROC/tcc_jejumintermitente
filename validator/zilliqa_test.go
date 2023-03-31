package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZilliqaValidateAddress(t *testing.T) {
	validator := &Zilliqa{}

	var mainnetCases = map[string]*Result{
		"zil1pyvfpng2rv2p47ha67yh0edd7eyxwyv0w8hzt6": {Success, true, Normal, ""},
		"zil144fujqgsuvv5zyr2v23m7ndyt3artkh78kyd92": {Success, true, Normal, ""},
		"zil13gpgtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"zib13gpgtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa":                      {Success, false, Unknown, ""},
		"zil13gegtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa":                      {S