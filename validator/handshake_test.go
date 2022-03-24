package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: testnet and P2WSH cases
func TestHandshakeValidateAddress(t *testing.T) {
	validator := &Handshake{}

	