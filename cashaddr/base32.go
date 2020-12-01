package cashaddr

type base32Encoder struct {
	alphabet string
}

func newBase32Encoder(alphabet string) *base32Encoder {
	return &base32Encoder{alphabet: alphabet}
}

func (e *base32Encoder) numToCode() map[byte]rune {
	result := make(map[byte]rune)
	for i, r := range e.alphabet {
		result[byte(i)] = r
	}
	return result
}

func (e *base32Encoder) codeToNum() map[rune]byte {
	result := make(map[rune]byte)
	for i, r := range e.alphabet {
		result[r] = byte(i)
	}
	return result
}

func (e *base32Encoder) encode(