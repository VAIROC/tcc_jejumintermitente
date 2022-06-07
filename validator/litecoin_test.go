package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLitecoinValidateAddress(t *testing.T) {
	validator := &Litecoin{}

	var mainnetCases = map[string]*Result{
		"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ": {Success, true, P2PKH, ""},
		"LfzMXaeTd1ZvThWdTHvAiucgLBy461DdCf": {Success, true, P2PKH, ""},
		"LY1cKc26dhtGg6EJutjSH1QXbXPbRNdVZ8": {Success, true, P2PKH, ""},
	