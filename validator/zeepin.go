package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Zeepin ...
type Zeepin struct{}

var _ Validator = (*Zeepin)(nil)

const zeepinAddrVersion = 80

// ValidateAddres