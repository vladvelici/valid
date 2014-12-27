package valid

import (
	"errors"
	"net/mail"
	"regexp"
)

type StringValidator interface {
	Validate(string) error
}

// StringValidator takes a value to validate and returns a validation error.
type StringFunc func(val string) error

// Validate function of StringFunc
func (s StringFunc) Validate(val string) error {
	return s(val)
}

// Validate a string value against the given validators.
func String(val string, v ...StringValidator) []error {
	errors := make([]error, 0)
	for _, validator := range v {
		err := validator.Validate(val)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// Validate against a regexp given as string
func Regexp(pattern, message string) StringFunc {
	reg := regexp.MustCompile(pattern)
	return RegexpCompiled(reg, message)
}

// Validate against a regexp given as *regexp.Regexp
func RegexpCompiled(reg *regexp.Regexp, message string) StringFunc {
	f := func(val string) error {
		if reg.MatchString(val) {
			return nil
		}
		return errors.New(message)
	}
	return StringFunc(f)
}

// Length in interval [min, max]. Includes min and max, max >= len(val) >= min.
func Len(min, max int, message string) StringFunc {
	f := func(val string) error {
		if max < len(val) || len(val) < min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Minimum length of string. (len(val) >= min)
func MinLen(min int, message string) StringValidator {
	f := func(val string) error {
		if len(val) < min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Maximum length of string. (len(val) <= max)
func MaxLen(max int, message string) StringValidator {
	f := func(val string) error {
		if len(val) > max {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Length in interval (min, max). Does not include min and max, max > len(val) > min.
func LenStrict(min, max int, message string) StringFunc {
	f := func(val string) error {
		if max <= len(val) || len(val) <= min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Minimum length of string. (len(val) > min)
func MinLenStrict(min int, message string) StringValidator {
	f := func(val string) error {
		if len(val) <= min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Maximum length of string. (len(val) < max)
func MaxLenStrict(max int, message string) StringValidator {
	f := func(val string) error {
		if len(val) >= max {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// String must be different then "".
func Nonempty(message string) StringValidator {
	f := func(val string) error {
		if val == "" {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Common regular expressions.
var (
	// Alphanumeric characters
	RegAlphanumeric = regexp.MustCompile("^[a-zA-Z0-9]*$")

	// Alphanumeric characters and -, _ and .
	RegAlphanumericPermissive = regexp.MustCompile("^[a-zA-Z0-9-_.]*$")

	// E-mail addresses, according to W3C http://www.w3.org/TR/html-markup/input.email.html#input.email.attrs.value.single
	RegEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
)

// Aplhanumeric validator.
func Alphanumeric(message string) StringValidator {
	return RegexpCompiled(RegAlphanumeric, message)
}

// Alphanumeric permissive validator.
func AlphanumericPermissive(message string) StringValidator {
	return RegexpCompiled(RegAlphanumericPermissive, message)
}

// Validate e-mail address according to net/mail. Address.Name must be empty.
// For validation of unique e-mail addresses, consider normalising the addresses.
func EmailRFC(message string) StringValidator {
	f := func(val string) error {
		addr, err := mail.ParseAddress(val)
		if err != nil || addr.Name != "" {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Validate e-mail according to HTML5 e-mail field regex.
// The regexp is defined by W3C here:
// http://www.w3.org/TR/html-markup/input.email.html#input.email.attrs.value.single
func Email(message string) StringValidator {
	return RegexpCompiled(RegEmail, message)
}
