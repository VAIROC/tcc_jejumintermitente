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
		"12xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW": {Success, true, Normal,