package main

import (
	"fmt"
	"os"

	"github.com/LanfordCai/ava/validator"
)

func validateBitcoinAddr() {
	v := &validator.Bitcoin{}
	addr := "19JeUHUvw23fwKeK1zZD4moKyxj1xn4Kxi"
	result := v.ValidateAddress(addr, validator.Mainnet)

	fmt.Printf("Address is valid?: %t\n", result.IsValid)
	fmt.Printf("Address type: %s\n", result.Type)
}

func validateEOSAddr(