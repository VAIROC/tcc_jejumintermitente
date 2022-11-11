package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTezosValidateAddress(t *testing.T) {
	validator := &Tezos{}

	var validCases = map[string]*Result{
		"tz1