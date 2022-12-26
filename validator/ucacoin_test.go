package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcacoinValidateAddress(t *testing.T) {
	validator := &Ucacoin{}

	var mainnetCases = map[string]*Result{
		"UaFYCsCHXtbW4EW