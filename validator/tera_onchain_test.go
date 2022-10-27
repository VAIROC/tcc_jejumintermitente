// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeraValidateAddress(t *testing.T) {
	client := TeraClient{Endpoint: os.Getenv("AVA_TERA_ENDPOINT")}
	validator := &Tera{Client: &client}

	var validCases = map[string]*Result{
		"189862": {Success, true, Normal, ""},
		"189007": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Resu