package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: no testnet cases
func TestVsystemsValidateAddress(t *testing.T) {
	validator := &Vsyste