package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZilliqaValidateAddress(t *testing.T) {
	validator := &Zilliqa{}

	var mainnetCases = map[string]*Result{
		"zil1pyvfpng2rv2p47ha67yh0edd7eyxwyv0w8hzt6": {Success, true, Normal, ""},
		"zil144fujqgsuvv5zyr2v23m7ndyt3artkh78kyd9