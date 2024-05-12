package validator

type errorMap map[string]ValidatorError

func (e errorMap) Error() map[string]string {
	var errors = make(map[string]string)
	for key, value := range e {
		errors[key] = value.Error()
	}
	return errors
}
