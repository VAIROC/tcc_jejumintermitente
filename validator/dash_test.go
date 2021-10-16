package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashValidateAddress(t *testing.T) {
	validator := &Dash{}

	var mainnetCases = map[string]*Result{
		"XpMfHazu1y8bgnP8csW73LWULGKgs4XH6U": {