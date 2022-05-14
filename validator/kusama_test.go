package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKusamaValidateAddress(t *testing.T) {
	validator := &Kusama{}

	var validCases = map[string]*Result{
		"GZTZiL1F1TRYMcT7mmkUiujR6R2oVcw7bsWpWawwQJuRzdD": {Success, true, Normal, ""},
		"HEVQFCv5HWcwZhzJEX93gQfXikJQ9VgfHoQ9RvXk3KYHN29": {Success, true, Normal, ""},
		"Ek5FeXtZWxvcS5snsFMdwHZdRnaYvNYRJdBTZMiSur5bde6": {Success, true, Normal, ""},
		"DYMQA7uvnxayYxcSQywPN7jufwZovBHL48hp747me94aZdG": {Success, true, Normal, ""},
		"Ho6PiyJXo5paaZJADX3eCQ