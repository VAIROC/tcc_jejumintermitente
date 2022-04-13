package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTIsAddressFormatValid(t *testing.T) {
	validator := &IOST{}

	var validCases = []string{
		"kainiost",
		"ukay66",
		"huobi_iost",
		"13177913280",
		"_iost",
		"y2a_yfec2st",
	}

	for _, addr := range validCases {
		assert.True(t, valida