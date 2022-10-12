package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStellarIsAddressFormatValid(t *testing.T) {
	validator := &Stellar{}

	var validCases = []