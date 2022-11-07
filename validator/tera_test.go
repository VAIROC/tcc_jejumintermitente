package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeraIsAddressFormatValid(t *testing.T) {
	validator := &Tera{}

	var validCases = []string{
		"189862",
		"224577",
		"189007",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddressFormatValid(addr, Mainnet), addr)
	}
