package validations

import (
	"time"

	"github.com/bearbin/go-age"
	"github.com/go-ozzo/ozzo-validation"
)

func RequiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}

func AgeRange(min int, max int) validation.RuleFunc {
	return func(value interface{}) error {
		date, _ := value.(time.Time)
		a := age.Age(date)
		return validation.Validate(a, validation.Min(min), validation.Max(max))
	}
}
