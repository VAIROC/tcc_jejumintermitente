package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolkadotValidateAddress(t *testing.T) {
	validator := &Polkadot{}

	var validCases = map[string]*Result{
		"15j4dg5GzsL1bw2U2AWgeyAk6QTxq43V7ZPbXdAmbVLjvDCK": {Success, true, Normal, ""},
		"12xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW": {Success, true, Normal, ""},
		"14Ns6kKbCoka3MS4Hn6b7oRw9fFejG8RH5rq5j63cWUfpPDJ": {Success, true, Normal, ""},
		"16DGiP6jDwAfkAeqGfkUCtheKgUzTy7UeaiFFBAv8BwX3RhN": {Success, true, Normal, ""},
		"1vTfju3zruADh7sbBznxWCpircNp9ErzJaPQZKyrUknApRu":  {Success, true, Normal, ""},
		"1tZzPmcq8Auisttygmg9g6tPMtrh9i3b22D3tKXvde7ibRB":  {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"22xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW":     {Success, false, Unknown, ""},
		"16DGiP6