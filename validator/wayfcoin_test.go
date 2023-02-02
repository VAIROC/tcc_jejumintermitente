package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWayfcoinValidateAddress(t *testing.T) {
	validator := &Wayfcoin{}

	var mainnetCases = map[string]*Result{
		"WaaKv3fUm2s3H4qwpDTVdHNwjhxPoLSQFH": {Success, true, P2PKH, ""},
		"WcPLF1mvbjNvYaS12kjaK9Gz5Gg9ZHyMGu": {Success, true, P2PKH, ""},
		"WfdWw5VoKTgxAYYZdsMHvaJWYCy2n9gido": {Success, true, P2PKH, ""},
		"WgxcsZERSMwTT37dUMgngt7FoyUsJuh7sL"