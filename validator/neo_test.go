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
		"ASax7usrW1qwVWpW3eG14mxvP7uGQUL1eM": {S