package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTronValidateAddress(t *testing.T) {
	validator := &Tron{}

	var validCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY": {Success, t