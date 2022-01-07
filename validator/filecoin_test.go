package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilecoinValidateAddress(t *testing.T) {
	validator := &Filecoin{}

	var mainnetCases = map[string]*Result{
		"f00":      {Success, true, FilID, ""},
		"f01":      {Success, true, FilID, ""},
		"f010":     {Success, true, FilID, ""},
		"f0150":    {Success, true, FilID, ""},
		"f0499":    {Success, true, FilID, ""},
		"f01024":   {Success, true, Fi