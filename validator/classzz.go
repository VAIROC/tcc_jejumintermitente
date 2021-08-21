package validator

import (
	"fmt"
	"strings"

	"github.com/LanfordCai/ava/cashaddr"
)

// Classzz ...
type Classzz struct{}

var _ BitcoinLike = (*Classzz)(nil)
var _ CashAddress = (*Classzz)(nil)

// ValidateAddress returns validate result of