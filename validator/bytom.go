package validator

// Bytom ...
type Bytom struct{}

var _ Validator = (*Bytom)(nil)
var _ SegwitAddress = (*Bytom)(nil)

// ValidateAddress returns validate result o