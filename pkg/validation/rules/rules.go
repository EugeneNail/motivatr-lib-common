package rules

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

const (
	Alpha    = `^[a-zA-Z]+$`       // alphabet symbols only
	AlphaNum = `^[a-zA-Z0-9]+$`    // alphabet symbols and numbers
	San      = `^[a-zA-Z0-9\s]+$`  // spaces, alphabet symbols and numbers
	Sand     = `^[a-zA-Z0-9\s-]+$` // spaces, alphabet symbols, numbers and dashes
	Email    = `^[\w.+-]+@[\w.+-]+\.[a-zA-Z]{1,10}$`
)

type RuleFunc func(data map[string]any, field string) (message string, err error)

func Required() RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		message := fmt.Sprintf("The %s field is required", field)

		value, exists := data[field]
		if !exists {
			return message, nil
		}

		switch typedValue := value.(type) {
		case string:
			if len(typedValue) == 0 {
				return message, nil
			}
		case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			if typedValue == 0 {
				return message, nil
			}
		case float32:
			delta := 1e-6
			if math.Abs(float64(typedValue)) < delta {
				return message, nil
			}

		case float64:
			delta := 1e-6
			if math.Abs(typedValue) < delta {
				return message, nil
			}
		}

		reflected := reflect.ValueOf(value)
		kind := reflected.Kind()
		if (kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array) && reflected.Len() == 0 {
			return message, nil
		}

		return "", nil
	}
}

func Max(limit int) RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		if _, exists := data[field]; !exists {
			return "", nil
		}

		switch value := data[field].(type) {
		case string:
			if utf8.RuneCountInString(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d characters", field, limit), nil
			}
		case int:
			if value > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case uint:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case int8:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case uint8:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case int16:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case uint16:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case int32:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case uint32:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case int64:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case uint64:
			if int(value) > limit {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case float32:
			if value > float32(limit) {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		case float64:
			if value > float64(limit) {
				return fmt.Sprintf("The %s field must not be greater than %d", field, limit), nil
			}
		}

		reflected := reflect.ValueOf(data[field])
		kind := reflected.Kind()
		isIterable := kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array

		if isIterable && limit < 0 {
			return "", fmt.Errorf("the max rule limit must not be negative while validating arrays")
		}

		if isIterable && reflected.Len() > limit {
			return fmt.Sprintf("The %s field must not have more than %d items", field, limit), nil
		}

		return "", nil
	}
}

func Min(limit int) RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		if _, exists := data[field]; !exists {
			return "", nil
		}

		switch value := data[field].(type) {
		case string:
			if utf8.RuneCountInString(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d characters", field, limit), nil
			}
		case int:
			if value < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case uint:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case int8:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case uint8:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case int16:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case uint16:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case int32:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case uint32:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case int64:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case uint64:
			if int(value) < limit {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case float32:
			if value < float32(limit) {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		case float64:
			if value < float64(limit) {
				return fmt.Sprintf("The %s field must not be less than %d", field, limit), nil
			}
		}

		reflected := reflect.ValueOf(data[field])
		kind := reflected.Kind()
		isIterable := kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array

		if isIterable && limit < 0 {
			return "", fmt.Errorf("the min rule limit must not be negative while validating arrays")
		}

		if isIterable && reflected.Len() < limit {
			return fmt.Sprintf("The %s field must not have less than %d items", field, limit), nil
		}

		return "", nil
	}
}

func Date() RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		value, ok := data[field].(string)
		if !ok {
			return "", fmt.Errorf("the value is not string")
		}

		date, err := time.Parse("2006-01-02", value)
		previousCentury := time.Date(2000, 01, 01, 0, 0, 0, 0, time.UTC)
		nextCentury := time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC)

		if err != nil || date.After(nextCentury) || date.Before(previousCentury) {
			return fmt.Sprintf("The %s field format is invalid", field), nil
		}

		return "", nil
	}
}

func Regex(pattern string) RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		value, ok := data[field].(string)
		if !ok {
			return "", fmt.Errorf("the value is not string")
		}

		if len(value) == 0 {
			return "", nil
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return "", fmt.Errorf("cannot compile pattern %q: %w", pattern, err)
		}

		if !regex.MatchString(value) {
			return fmt.Sprintf("The %s field format is invalid", field), nil
		}

		return "", nil
	}
}

func Password() RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		message, err := Min(8)(data, field)
		if err != nil {
			return "", fmt.Errorf("cannot validate min length: %w", err)
		}

		if len(message) != 0 {
			return message, nil
		}

		message, err = Regex("[a-z]+")(data, field)
		if err != nil {
			return "", fmt.Errorf("cannot validate lower case symbols: %w", err)
		}

		if len(message) != 0 {
			return fmt.Sprintf("The %s field must contain at least one lower case letter", field), nil
		}

		message, err = Regex("[A-Z]+")(data, field)
		if err != nil {
			return "", fmt.Errorf("cannot validate upper case symbols: %w", err)
		}

		if len(message) != 0 {
			return fmt.Sprintf("The %s field must contain at least one upper case letter", field), nil
		}

		message, err = Regex("[0-9]+")(data, field)
		if err != nil {
			return "", fmt.Errorf("cannot validate numbers: %w", err)
		}

		if len(message) != 0 {
			return fmt.Sprintf("The %s field must contain at least one number", field), nil
		}

		return "", nil
	}
}

func Same(fieldToMatch string) RuleFunc {
	return func(data map[string]any, field string) (string, error) {
		message := fmt.Sprintf("The %s field must match %s", field, fieldToMatch)

		valueToMatch, exists := data[fieldToMatch]
		if !exists {
			return message, nil
		}

		if valueToMatch != data[field] {
			return message, nil
		}

		return "", nil
	}
}
