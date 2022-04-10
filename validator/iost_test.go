package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTIsAddressFormatValid(t *testing.T) {
	validator := &IOST{}

	var validCas