package valid

import (
	"testing"
)

type stringTest struct {
	input      string
	validator  StringValidator
	shouldPass bool
	message    string
}

var strTests = []stringTest{
	// Uncompiled regexp
	{"nonmatch003", Regexp("^[a-z]+$", "a"), false, "Non-match regexp."},
	{"match", Regexp("^[a-z]+$", "a"), true, "Match regexp."},

	// MinLen
	{"abcd", MinLen(2, "a"), true, "abcd has length greater than 2."},
	{"abcd", MinLen(5, "a"), false, "abcd has length less han 5."},
	{"abcd", MinLen(4, "a"), true, "abcd has length 4, shuold be OK with MinLen(4, ...)"},

	// MinLenStrict
	{"abcd", MinLenStrict(2, "a"), true, "abcd has length strictly greater than 2."},
	{"abcd", MinLenStrict(5, "a"), false, "abcd has length strictly less than 5."},
	{"abcd", MinLenStrict(4, "a"), false, "abcd has length 4, shuold fail with MinLenStrict(4, ...)"},

	// MaxLen
	{"abcd", MaxLen(8, "a"), true, "(MaxLen) abcd has length less than 8."},
	{"abcd", MaxLen(3, "a"), false, "(MaxLen) abcd has length greater han 3."},
	{"abcd", MaxLen(4, "a"), true, "(MaxLen) abcd has length 4, shuold be OK with MaxLen(4, ...)"},

	// MaxLenStrict
	{"abcd", MaxLenStrict(8, "a"), true, "(MaxLenStrict) abcd has length strictly greater than 2."},
	{"abcd", MaxLenStrict(3, "a"), false, "(MaxLenStrict) abcd has length strictly less than 5."},
	{"abcd", MaxLenStrict(4, "a"), false, "(MaxLenStrict) abcd has length 4, shuold fail with MaxLenStrict(4, ...)"},

	// Len
	{"abcd", Len(2, 9, "a"), true, "Len: len(\"abcd\")=4 is in range [2,9]."},
	{"abcd", Len(2, 4, "a"), true, "Len: len(\"abcd\")=4 is in range [2,4]."},
	{"abcd", Len(4, 4, "a"), true, "Len: len(\"abcd\")=4 is in range [4,4]."},
	{"abcd", Len(4, 9, "a"), true, "Len: len(\"abcd\")=4 is in range [4,9]."},
	{"abcd", Len(5, 9, "a"), false, "Len: len(\"abcd\")=4 is not in range [5,9]."},
	{"abcd", Len(5, 3, "a"), false, "Len: len(\"abcd\")=4 is not in range [5,3]."},
	{"abcd", Len(2, 3, "a"), false, "Len: len(\"abcd\")=4 is not in range [2,3]."},

	// LenStrict
	{"abcd", LenStrict(2, 9, "a"), true, "Len: len(\"abcd\")=4 is in range (2,9)."},
	{"abcd", LenStrict(2, 4, "a"), false, "Len: len(\"abcd\")=4 is in range (2,4)."},
	{"abcd", LenStrict(4, 4, "a"), false, "Len: len(\"abcd\")=4 is in range (4,4)."},
	{"abcd", LenStrict(4, 9, "a"), false, "Len: len(\"abcd\")=4 is in range (4,9)."},
	{"abcd", LenStrict(5, 9, "a"), false, "Len: len(\"abcd\")=4 is not in range (5,9)."},
	{"abcd", LenStrict(5, 3, "a"), false, "Len: len(\"abcd\")=4 is not in range (5,3)."},
	{"abcd", LenStrict(2, 3, "a"), false, "Len: len(\"abcd\")=4 is not in range (2,3)."},

	// Alphanumeric
	{"", Alphanumeric("a"), true, "Empty string is alphanumeric."},
	{"afdSADSgdfgfds", Alphanumeric("a"), true, "Letters are alphanumeric."},
	{"054360243", Alphanumeric("a"), true, "Numbers are alphanumeric."},
	{"fdsa0543hgfd602HGF43YTR", Alphanumeric("a"), true, "Letters and numbers are alphanumeric."},
	{"__%$abc012def", Alphanumeric("a"), false, "Alphanumeric prefixed with nonaphpanumeric string is not alphanumeric."},
	{"abc012def_$#@", Alphanumeric("a"), false, "Alphanumeric suffixed with nonaphpanumeric string is not alphanumeric."},
	{"_$#@", Alphanumeric("a"), false, "Completely nonalphanumeric string."},

	// AlphanumericPermissive
	{"", AlphanumericPermissive("a"), true, "Empty string is permissive alphanumeric."},
	{"afdSADSgdfgfds", AlphanumericPermissive("a"), true, "Letters are permissive alphanumeric."},
	{"054360243", AlphanumericPermissive("a"), true, "Numbers are permissive alphanumeric."},
	{"_.-._", AlphanumericPermissive("a"), true, ". - and _ are permissive alphanumeric."},
	{"fds_a054.3hgf--d60__2HGF43YTR", AlphanumericPermissive("a"), true, "Letters, numbers and allowed symbols are permissive alphanumeric."},
	{"$$#@__abc012def", AlphanumericPermissive("a"), false, "Permissive alphanumeric prefixed with nonaphpanumeric string is not alphanumeric."},
	{"abc-01_2d.ef_$$#@", AlphanumericPermissive("a"), false, "Permissive alphanumeric suffixed with nonaphpanumeric string is not alphanumeric."},
	{"$#@", AlphanumericPermissive("a"), false, "Completely nonalphanumeric (permissive) string."},

	// Nonempty
	{"", Nonempty("a"), false, "Empty string shuold false."},
	{"hello", Nonempty("a"), true, "Non-empty string should pass."},

	// Email RFC
	{"jo.hn+tag@subdomain.example.com", EmailRFC("a"), true, "Valid e-mail addresss."},
	{"%$)@jo.hn+tag@subdomain.example.com", EmailRFC("a"), false, "Invalid e-mail address."},
	{"John <jo.hn+tag@subdomain.example.com>", EmailRFC("a"), false, "Valid with-name email address should not pass."},

	// Email regexp
	{"jo.hn+tag@subdomain.example.com", Email("a"), true, "Valid e-mail addresss (regexp)."},
	{"%$)@jo.hn+tag@subdomain.example.com", Email("a"), false, "Invalid e-mail address (regexp)."},
	{"John <jo.hn+tag@subdomain.example.com>", Email("a"), false, "Valid with-name email address should not pass (regexp)."},
}

func TestStringValidators(t *testing.T) {
	var err error
	for i, test := range strTests {
		err = test.validator.Validate(test.input)
		pass := err == nil
		if pass != test.shouldPass {
			t.Errorf("String #%d failed [%v/%v]: %s", i, pass, test.shouldPass, test.message)
		}
	}
}

func TestString(t *testing.T) {
	errs := String("john", Email("is not email"), MaxLen(10, "this should pass"), MinLen(5, "this should fail"))
	if len(errs) != 2 {
		t.Errorf("String function returns the wrong number of errors: %d instead of 2.", len(errs))
		t.FailNow()
	}
	messages := []string{"is not email", "this should fail"}
	for i, msg := range messages {
		if errs[i].Error() != msg {
			t.Errorf("String() returns the wrong message: errs[%d] = %#v instead of %#v", i, errs[i].Error(), msg)
		}
	}
}
