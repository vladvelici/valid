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

