package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTezosValidateAddress(t *testing.T) {
	validator := &Tezos{}

	var validCases = map[string]*Result{
		"tz1aQWeRw2MJpvvurJb78nwSpK8nvNg6EoCc": {Success, true, Normal, ""},
		"tz1hkzS6pnfnHv9KzX1nbtqXVqUkzcem8FJs": {Success, true, Normal, ""},
		"tz1RuFpeQEp8yC4B7yX4amBiLvpbcW7HHRrS": {Success, true, Normal, ""},
		"tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r": {Success, true, Normal, ""},
		"tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWha