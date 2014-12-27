package valid

import (
	"errors"
	"net/mail"
	"regexp"
)

// StringValidator interface is implemented by all string validators.
type StringValidator interface {
	// Validate validates the given parameter and returns a validation error, or nil
	// if the input is valid.
	Validate(string) error
}

// A StringFucn takes a value to validate and returns a validation error.
//
// This type implements the StringValidator interface, thus any functions with this
// signature can be casted to StringFunc and used as a StringValidator.
type StringFunc func(val string) error

// Validate function of StringFunc
func (s StringFunc) Validate(val string) error {
	return s(val)
}

// String applies a list of StringValidators to a string value and returns a list of
// aggregated errors.
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

// Regexp creates a regular expression validator, using the pattern given as string.
func Regexp(pattern, message string) StringFunc {
	reg := regexp.MustCompile(pattern)
	return RegexpCompiled(reg, message)
}

// RegexpCompiled creates a regular expression validator, using the given already-compiled *regexp.Regexp.
func RegexpCompiled(reg *regexp.Regexp, message string) StringFunc {
	f := func(val string) error {
		if reg.MatchString(val) {
			return nil
		}
		return errors.New(message)
	}
	return StringFunc(f)
}

// Len creates a length validator, that checks if the length of given string is in the closed interval [min, max].
// It includes min and max: all val that satisfy max >= len(val) >= min are considered valid.
func Len(min, max int, message string) StringFunc {
	f := func(val string) error {
		if max < len(val) || len(val) < min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// MinLen creates a minimum length string validator that consideres all strings val valid if they
// satisfy len(val) >= min.
func MinLen(min int, message string) StringValidator {
	f := func(val string) error {
		if len(val) < min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// MaxLen creates a maximum length string validator that consideres all strings val valid if they
// satisfy len(val) <= max.
func MaxLen(max int, message string) StringValidator {
	f := func(val string) error {
		if len(val) > max {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// LenStrict creates a length validator, that checks if the length of given string is in the open interval (min, max).
// It does not include min and max: all val that satisfy max > len(val) > min are considered valid.
func LenStrict(min, max int, message string) StringFunc {
	f := func(val string) error {
		if max <= len(val) || len(val) <= min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// MinLenStrict creates a strict minimum length string validator that consideres all strings val valid if they
// satisfy len(val) > min.
func MinLenStrict(min int, message string) StringValidator {
	f := func(val string) error {
		if len(val) <= min {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// MaxLenStrict creates a maximum length string validator that consideres all strings val valid if they
// satisfy len(val) < max.
func MaxLenStrict(max int, message string) StringValidator {
	f := func(val string) error {
		if len(val) >= max {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Nonempty creates a validator that checks whether the given string is not empty.
func Nonempty(message string) StringValidator {
	f := func(val string) error {
		if val == "" {
			return errors.New(message)
		}
		return nil
	}
	return StringFunc(f)
}

// Common regular expressions used for validation.
var (
	// RegAlphanumeric matches alphanumeric characters
	RegAlphanumeric = regexp.MustCompile("^[a-zA-Z0-9]*$")

	// RegAlphanumericPermissive matches numbers, letters,  -, _, and .
	RegAlphanumericPermissive = regexp.MustCompile("^[a-zA-Z0-9-_.]*$")

	// RegEmail is the HTML5 E-mail address regular expression, according to W3C http://www.w3.org/TR/html-markup/input.email.html#input.email.attrs.value.single
	RegEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
)

// Alphanumeric creates a validator that consideres all alphanumeric inputs (including the empty string) valid.
func Alphanumeric(message string) StringValidator {
	return RegexpCompiled(RegAlphanumeric, message)
}

// AlphanumericPermissive creates a validator that consideres all inputs containing only letters (a-z, A-Z), numbers (0-9),
// underscore ("_"), minus sign ("-") and period (".") as valid. (including the empty string) valid.
func AlphanumericPermissive(message string) StringValidator {
	return RegexpCompiled(RegAlphanumericPermissive, message)
}

// EmailRFC creates a validator for e-mail address according to the net/mail package. Non-empty Address.Name addresses are
// considered invalid (e.g. "John <john@example.org>" is invalid).
//
// For validation of uniqueness, consider normalising the addresses. (e.g. "john+tag@example.com" is the same as
// "john@example.com".
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

// Email creates a validator that uses the HTML5 e-mail field regexp as defined by W3C to
// validate e-mail addresses.
//
// The regular expression is defined by W3C here:
// http://www.w3.org/TR/html-markup/input.email.html#input.email.attrs.value.single
func Email(message string) StringValidator {
	return RegexpCompiled(RegEmail, message)
}
