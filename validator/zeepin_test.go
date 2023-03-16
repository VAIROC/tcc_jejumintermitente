package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeepinValidateAddress(t *testing.T) {
	validator := &Zeepin{}

	var mainnetCases = map[string]*Result{
		"ZQfyK6J8sRtKT5XE