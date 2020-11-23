
package base58check

import (
	"bytes"
	"math/big"
	"strings"

	"github.com/LanfordCai/ava/crypto"
)

// BitcoinEncoder ...
var BitcoinEncoder = &Encoder{
	enc:          Encoding{alphabet: "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"},
	ChecksumLen:  4,
	ChecksumType: ChecksumDoubleSha256,
}

// RippleEncoder ...
var RippleEncoder = &Encoder{
	enc:          Encoding{alphabet: "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"},
	ChecksumLen:  4,
	ChecksumType: ChecksumDoubleSha256,
}
