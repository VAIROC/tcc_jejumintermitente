package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogecoinValidateAddress(t *testing.T) {
	validator := &Dogecoin{}

	var mainnetCases = map[string]*Result{
		"D8osUgxgU4J2yEhCzN5xQQ42tuuG7qBdri": {Success, true, P2PKH, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnev": {Success, true, P2PKH, ""},
		"DCZHAVah2dCyD1vmjnUm6t7Aao4YwtYWkM": {Success, true, P2PKH, ""},
		"A6T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"nYF8DkhE3orcq7CypP1oEQTWWcG1yYkqec":  {Success, true, P2PKH, ""},
		"nXiLZ1g1xrcLwYq1s9wVCdiewLeY7Bw1X3":  {Success, true, P2PKH, ""},
		"nj4xWxERcYpXQLZf61R5xzMDfTjv72ajJh":  {Success, true, P2PKH, ""},
		"2NDutUPDhfMYGUFPhHRy4ZRraHVrEK7odr6": {Success, true, P2SH, ""},
		"2MwU3UtzeHG59xV7J7VSoM42U7qwCJ1GSJz": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False