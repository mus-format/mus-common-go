package muscom

// Validator performs data validation.
type Validator[T any] interface {
	Validate(t T) error
}

// ValidatorFn is a function implementation of the Validator interface.
type ValidatorFn[T any] func(t T) (err error)

func (fn ValidatorFn[T]) Validate(t T) (err error) {
	return fn(t)
}
