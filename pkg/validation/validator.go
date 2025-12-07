package validation

import (
	"fmt"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation/rules"
)

type Validator struct {
	data   map[string]any
	rules  map[string][]rules.RuleFunc
	errors map[string]string
}

func NewValidator(data map[string]any, rules map[string][]rules.RuleFunc) *Validator {
	return &Validator{
		data:   data,
		rules:  rules,
		errors: make(map[string]string),
	}
}

func (validator *Validator) Errors() map[string]string {
	return validator.errors
}

func (validator *Validator) Failed() bool {
	return len(validator.errors) > 0
}

func (validator *Validator) Validate() error {
	for field, ruleFuncs := range validator.rules {
	ruleLoop:
		for _, ruleFunc := range ruleFuncs {
			message, err := ruleFunc(validator.data, field)
			if err != nil {
				return fmt.Errorf("cannot validate the %s field: %w", field, err)
			}

			if len(message) > 0 {
				validator.AddError(field, message)
				break ruleLoop
			}
		}
	}

	return nil
}

func (validator *Validator) AddError(field string, message string) {
	validator.errors[field] = message
}
