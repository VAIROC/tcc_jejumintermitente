package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolkadotValidateAddress(t *testing.T) {
	validator := &Polkadot{}

	var validCases = map[string]*Result{
		