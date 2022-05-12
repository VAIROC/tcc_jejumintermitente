package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKusamaValidateAddress(t *testing.T) {
	validator := &Kusama{}

	var validCases = map[string]*Result{
		"GZTZiL1F1TRYMcT