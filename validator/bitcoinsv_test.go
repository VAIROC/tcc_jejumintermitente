package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinSVValidateAddress(t *testing.T) {
	validator := &BitcoinSV{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF": {Success, true, P2PKH, ""},
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw": {Success, tru