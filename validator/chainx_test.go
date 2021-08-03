package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainXValidateAddress(t *testing.T) {
	validator := &ChainX{}

	var validCases = map[string]*Result{
		"5QxFTz6oeBbgPtGMg531zUZ4JJz2RkbiCjkUu9z1Uf7Wi92d": {Success, true, Normal, ""},
		"5SdW3TbRpnsbvTDa9QDXftM8mt