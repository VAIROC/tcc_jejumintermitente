package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArweaveValidateAddress(t *testing.T) {
	validator := &Arweave{}

	var validCases = map[string]*Result{
		"kEGgWKgTtojj-TGdbIUPPHmntIWQEs1IPegkCH3SyZ8": {Success, true, Normal, ""},
		"181f-LjI1ooRZhj3wLK2bCjOpgUmfOo_BplG4mCGuxU": {Success, true, Normal, ""},
		"Tk1NuG7Jxr9Ecgva5tWOJya2QGDOoS6hMZP0paB129c": {Success, true, Normal, ""},
		"XepqBRwnq8SNfple3WHh0VWo06P8KQ0hVlNOjOgrJ5w": {Success, true, Normal, ""},
		"tO710xPXwCGwTPGKtOkEq3PMbWqXs9jOGiL8TCpDuw0": {Success, true, Normal, "