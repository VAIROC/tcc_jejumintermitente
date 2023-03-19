package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeepinValidateAddress(t *testing.T) {
	validator := &Zeepin{}

	var mainnetCases = map[string]*Result{
		"ZQfyK6J8sRtKT5XECyrPDV4KsAXPWfYARd": {Success, true, Normal, ""},
		"ZJSmkuhUL6GUaVwzapCaXyYPT2cBYgMQ4n": {Success, true, Normal, ""},
		"ZD8HgfakgSXjk9XN2c5SS5qbEZp94TEmE6": {Success, true, Normal, ""},
		"ZHeWYS7Kc2Nti6gT8KKec6a4ohGNNcTxVF": {Success, true, Normal, ""},
		"ZakH4H7PG3uXMj1EFFmaz3s1ULYvgTuSRU": {Success, true, Normal, ""},
		"ZEJQUH8MhVRcNY8o4zECS8ZGyqDJYEBTGJ": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCas