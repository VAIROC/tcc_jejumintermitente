package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinValidateAddress(t *testing.T) {
	validator := &Bitcoin{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF":                             {Success, true, P2PKH, ""},
		"1