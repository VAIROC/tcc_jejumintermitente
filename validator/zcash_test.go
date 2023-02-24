package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZcashValidateAddress(t *testing.T) {
	validator := &Zcash{}

	var mainnetCases = map[string]*Result{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6zbiG": {Success, true, P2PKH, ""},
		"t1J1YojHoLb5L9pvRi1sCSNgyRe