package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainXValidateAddress(t *testing.T) {
	validator := &ChainX{}

	var validCases = map[string]*Result{
		"5QxFTz6oeBbgPtGMg531zUZ4JJz2RkbiCjkUu9z1Uf7Wi92d": {Success, true, Normal, ""},
		"5SdW3TbRpnsbvTDa9QDXftM8mtdw86ou8Z1Lyd3D4XgjSEii": {Success, true, Normal, ""},
		"5QbeyxMUMhEeukBaxAqHip9E1u93kMqxf2dTBPDrPeY52B9o": {Success, true, Normal, ""},
		"5R7sxzEYGGSGSCQd3DQLyYwkA29jYGYaecHUHCsQ8s1Vxeo7": {Success, true, Normal, ""},
		"5VQLcbTNn6UPYqLK87kPBHQvhvYL24z9bCa9XBtkkWekGrwj": {Success, true, Normal, ""},
		"5UBLKqQ7efPHxwVQCT82Vf86PxsDRCRv8nVZZQ9YpUrkg25M": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"15SdW3TbRpnsbvTDa9QDXftM8mtdw86ou8Z1Ly3D4XgjSEii":  {Success, 