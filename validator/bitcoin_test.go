package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinValidateAddress(t *testing.T) {
	validator := &Bitcoin