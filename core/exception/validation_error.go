package exception

type ErrValidation struct {
	Validations map[string]any `json:"validations,omitempty" `
	Message     string         `json:"message,omitempty"`
}

func NewValidationError(errors map[string]any, msg string) *ErrValidation {
	return &ErrValidation{
		Validations: errors,
		Message:     msg,
	}
}

func (ev *ErrValidation) Error() string {
	return "validation error"
}
