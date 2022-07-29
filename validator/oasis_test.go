package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisIsAddressFormatValid(t *testing.T) {
	validator := &Oasis{}

	var validCases = [