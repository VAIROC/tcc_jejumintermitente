package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinSVValidateAddress(t *testing.T) {
	validator := &BitcoinSV{}

	var mainnetC