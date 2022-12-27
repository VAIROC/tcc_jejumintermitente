package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcacoinValidateAddress(t *testing.T) {
	validator := &Ucacoin{}

	var mainnetCases = map[string]*Result{
		"UaFYCsCHXtbW4EWmckYXP8BmGgezJrRdem": {Success, true, P2PKH, ""},
		"Uak5JfpmLyjJkGJP37eBWHVPQmYo9wLX7w": {Success, true, P2PKH, ""},
		"UbehDN7bgkJ8VgCsM96r2W2UpRoiHS