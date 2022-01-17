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
		"f01024":   {Success, true, FilID, ""},
		"f01729":   {Success, true, FilID, ""},
		"f0999999": {Success, true, FilID, ""},
		"f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq":                                              {Success, true, FilSecp256k1, ""},
		"f12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a":                                              {Success, tr