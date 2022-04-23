package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIrisValidateAddress(t *testing.T) {
	validator := &Iris{}

	var mainnetCases = map[string]*Result{
		"iaa1qjqvwsmqhm6qjgs0qsyhqc3z5f5slctp6efhhw": {Success, true, Norm