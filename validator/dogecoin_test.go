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
		"DCZHAVah2dCyD1vmjn