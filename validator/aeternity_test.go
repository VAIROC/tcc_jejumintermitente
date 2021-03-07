package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAeternityValidateAddress(t *testing.T) {
	validator := &Aeternity{}

	var 