package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNEOValidateAddress(t *testing.T) {
	validator := &NEO{}

	var mainnetCases = map[string]*Result{
		"AVKeLwu1uRtM7Lf7ckkqrbtkvss7jkN38N": {Success, true, P2PKH, ""},
		"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo": {Success, true, P2PKH, ""},
		"AetXb6U1FMcA3zNak8FuPmY33ovi4xj4wg": {Success, true, P2PKH, ""},
		"ASax7usrW1qwVWpW3eG14mxvP7uGQUL1eM": {Success, true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ASax7usrW1qwVWpW3eG15mxvP7uGQUL1eM":         {Success, false, Unknown, ""},
		"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo ":        {Success, false, Unknown, ""},
		"AEzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo":         {Success, false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":        {Success, false, Unknown, ""},
		"