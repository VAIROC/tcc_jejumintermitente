package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAeternityValidateAddress(t *testing.T) {
	validator := &Aeternity{}

	var validCases = map[string]*Result{
		"ak_2QkttUgEyPixKzqXkJ4LX7ugbRjwCDWPBT4p4M2r8brjxUxUYd": {Success, true, Normal, ""},
		"ak_KHfXhF2J6VBt3sUgFygdbpEkWi6AKBkr9jNKUCHbpwwagzHUs":  {Success, true, Normal, ""},
		"ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR":  {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ak_4h3c6RH52R": {Success, false, Unknown, ""},
		"akzvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR":    {Success, false, Unknown, ""},
		"ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJV":    {Success, false, Unknown, ""},
		"aK_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR":   {Success, false, Unknown, ""},
		"