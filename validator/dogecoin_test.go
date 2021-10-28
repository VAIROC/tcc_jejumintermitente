package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogecoinValidateAddress(t *testing.T) {
	validator := &Dogecoin{}

	var mainnetCases = map[string]*Result{
		"D8