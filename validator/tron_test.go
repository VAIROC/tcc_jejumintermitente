package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTronValidateAddress(t *testing.T) {
	validator := &Tron{}

	var validCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY": {Success, true, Normal, ""},
		"TW73UWp1a5s8Zs36D6LkQuBp27dz2VRg1a": {Success, true, Normal, ""},
		"TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH": {Success, true, Normal, ""},
		"TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var