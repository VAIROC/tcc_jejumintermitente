package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLitecoinValidateAddress(t *testing.T) {
	validator := &Litecoin{}

	var mainnetCases = map[string]*Result{
		"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ": {Success, true, 