package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQtumValidateAddress(t *testing.T) {
	validator := &Qtum{}

	var mainnetCases = map[string]*Result{
		"QhzgLgC9DY