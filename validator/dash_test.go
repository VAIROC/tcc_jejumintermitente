package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashValidateAddress(t *testing.T) {
	validator := &Dash{}

	var mainnetCases = map[string]*Result{
		"XpMfHazu1y8bgnP8csW73LWULGKgs4XH6U": {Success, true, P2PKH, ""},
		"Xdk9TLNrELrM4jpu6m15zZvurG66sZKi6C": {Success, true, P2PKH, ""},
		"XripctRDfASrofg4BKco41Zqvm1KaCdQpk": {Success,