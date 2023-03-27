package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZilliqaValidateAddress(t *testing.T) {
	validator := &Zilliqa{}

	var main