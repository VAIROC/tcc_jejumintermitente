package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeraIsAddressFormatValid(t *testing.T) {
	validator := &Tera{}

	var validCas