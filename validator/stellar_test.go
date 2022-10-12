package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStellarIsAddressFormatValid(t *testing.T) {
	validator := &Stellar{}

	var validCases = []string{
		"GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER",
		"GBLKRATZODTSJNU7XTB5HY5VAAN63CPRT77UYZT2VLCNXE7F3YHSW22M",
		"GDDYS6MS7EBKGC4YNSWZQZGCMQ5LKG