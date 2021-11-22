package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEOSIsAddressFormatValid(t *testing.T) {
	validator := &EOS{}

	var validCases = []string{
		"eosnationftw",
		"atticlabeosb",
		"helloeoscnbp",
		"eosasia11111",
		"zbeosbp11111",
		"eoshuobipool",
		"big.one",
		"okcap