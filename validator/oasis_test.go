package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisIsAddressFormatValid(t *testing.T) {
	validator := &Oasis{}

	var validCases = []string{
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2",
		"oasis1qrad7s7nqm4gvyzr8yt2rdk0ref489rn3vn400d6",
		"oasis1qp29h8ykmxet46eqzw0wennrmm