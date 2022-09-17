package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRippleValidateAddress(t *testing.T) {
	validator := &Ripple{}

	var validCases = map[string]*Result{
		"r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV": {Success, true, Normal, ""},
		"rKLpjpCoXgLQQYQyj13zgay73rsgmzNH13": {Succe