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
		"WcPL