package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArweaveValidateAddress(t *testing.T) {
	validator := &Arweave{}

	var validCases = map[string]*Result{
		"kEGgWKgTtojj-TGdbIUPPHmntIWQEs1IPegkCH3SyZ8": {Success, 