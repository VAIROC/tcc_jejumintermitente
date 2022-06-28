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
		"ATzoCmmsjqPHCDPmY7mxEhSTGzpSK