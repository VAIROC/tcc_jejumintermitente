package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIrisValidateAddress(t *testing.T) {
	validator := &Iris{}

	var mainnetCases = map[string]*Result{
		"iaa1qjqvwsmqhm6qjgs0qsyhqc3z5f5slctp6efhhw": {Success, true, Normal, ""},
		"iaa1rpdhedqrmwtjk4lluqj6m2pvhm544h2vfu2mq8": {Success, true, Normal, ""},
		"iaa1lz9nwk6y5eeafr5c54k5dw8rwje2sj2e5pcqd6": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"iaa1lz9nwk6y3eeafr5c54k5dw8rwje2sj2e5pcqd6":                      {Success, false, Unknown, ""},
		"baa1lz9nwk6y5eeafr5c54k5dw8rwje2s