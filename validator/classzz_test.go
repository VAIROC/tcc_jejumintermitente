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
		"cp2xxd7q5wzjxkjmvy6jshh2rz3eq5sjey7ja09mwx":         {Success, true, CashAddrP2PKH, ""},
		"classzz:cp2sn5ws3u5epfmvuc4tu5szf32gj78dnqekvpk4mp": {Success, true, CashAddrP2PKH, ""},
		"classzz:cp2ts8jxwwadkc3dgqp08uk29tcmqvsjzsh7jevgfm": {Success, true, CashAddrP2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"CP2XXD7Q5WZJXKJMVY6JSHH2RZ3EQ5SJEY7JA09MWX":             {Success, false, Unknown, ""},
		"qzy9dedpwm53cqgcr2e3z9qy788cca6y0ce8mfkezh":             {Success, false, Unknown, ""},
		"bitcoincash:qzy9dedpwm53cqgcr2e3z9qy788cca6y0ce8mfkezh": {Success, false, Unknown, ""},
		"clbsszz:cp2ts8jxwwadkc3dgqp08uk29tcmqvsjzsh7jevgfm":     {Success, false, Unknown, ""},
		"cp2xxd3q5wzjxkjmvy6jshh2rz3eq5sjey7ja09mwx":             {Success, false, Unknown, ""},
		"cp283r0gjwd7lre962r32prfyhqz0l0l6shwmzn7st ":            {Success, false, Unknown, ""},
	