package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClasszzValidateAddress(t *testing.T) {
	validator := &Classzz{}

	var mainnetCases = map[string]*Result{
		"cp283r0gjwd7lre962r32prfyhqz0l0l6shwmzn7st":         {Success, true, CashAddrP2PKH, ""},
		"cp2xxd7q5wzjxkjmvy6jshh2rz3eq5sjey7ja09mwx":         {Suc