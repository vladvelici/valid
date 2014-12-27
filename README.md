valid
=====

Simple input validation for Go. This package does not deal with any form of input processing.

This package is all about checking if values satisfy different validation rules. It is outside the goals of this
package to deal with web form input processing. There are already plenty of librbaries doing that.

Usage
=====


    // A variable to validate.
    var input string

    // ... input processing (e.g. from web forms) ...

    // Validation
    errs := valid.String(input, Nonempty("This field is required."), MaxLen(20, "Input is too long."))

    // errs is an []error containing errors with the relevant error messages.


Current validators
==================

### Strings

- **Len** - Length in closed interval `[min, max]`.
- **LenStrict** - Length in open interval `(min, max)`.
- **MinLen** - Length at least `min`.
- **MinLenStrict** - Length at least `min+1`.
- **MaxLen** - Length at most `max`.
- **MaxLenStrict** - Length at most `max-1`.
- **Nonempty** - string must not be empty.
- **EmailRFC** - e-mail address validation using the `net/mail` package.
- **Email** - e-mail address validation using the HTML5 Email field regexp.
- **Regex** - Validate using regular expression, given as string.
- **RegexCompiled** - Validate using regular expression, given as `*regexp.Regexp`.

Create custom validators
========================

Validators are different for each type, but they all have similar interfaces.

### Strings

The interface `StringValidator` has one method `Validate(string) error`. Any struct that
implements this interface can be used as a string validator.

A convenience type is also defined: `type StringFunc func(val string) error`. This type implements the 
`StringValidator` interface, so any function with the above signature can be used as a validator.

### Example

You can see code for a custom validator here https://gist.github.com/vladvelici/00679f8dff9e205cc157.
