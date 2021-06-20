package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitsharesIsAddressFormatValid(t *testing.T) {
	validator := &Bitshares{}

	var validCases = []string{
		"zbbts001",
		"beos.gateway",
		