package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRippleValidateAddress(t *testing.T) {
	validator := &Ri