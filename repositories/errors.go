package repositories

// Describes generic error returned by the data layer
// Implements the "error" interface
type InternalError struct {
	message string
}

func (e *InternalError) Error() string {
	if e.message == "" {
		return "internal server error"
	}
	return e.message
}

func NewInternalError(message string) error {
	return &InternalError{
		message: message,
	}
}

// Error returned when a data entry is not found by the data layer
// Implements the "error" interface
type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	if e.message == "" {
		return "not found"
	}
	return e.message
}

func NewNotFoundError(message string) error {
	return &NotFoundError{
		message: message,
	}
}

// Error returned when a user supplied ID is correctly formatted
// Implements the "error" interface
type InvalidIdError struct {
	message string
}

func (e *InvalidIdError) Error() string {
	if e.message == "" {
		return "invalid supplied ID"
	}
	return e.message
}

func NewInvalidIdError(message string) error {
	return &InvalidIdError{
		message: message,
	}
}