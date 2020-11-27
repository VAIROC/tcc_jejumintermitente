package cashaddr

type base32Encoder struct {
	alphabet string
}

func newBase32Encoder(alphabet string) *base32Encoder {
	return &base32Encoder{alphabet: alphabet}
}

func (e *base32Encoder) numToCode() 