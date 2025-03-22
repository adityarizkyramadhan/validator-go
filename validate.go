package validator

import (
	"fmt"
	"net/http"
	"regexp"
)

type Validator struct {
	Request *http.Request
	errors  map[string]string
}

func New(r *http.Request) *Validator {
	return &Validator{
		Request: r,
		errors:  make(map[string]string),
	}
}

func (v *Validator) Required(field string) *Validator {
	value := v.Request.FormValue(field)
	if value == "" {
		v.errors[field] = "Field is required"
	}
	return v
}

func (v *Validator) MinLength(field string, length int) *Validator {
	value := v.Request.FormValue(field)
	if len(value) < length {
		v.errors[field] = "Minimum length is " + fmt.Sprint(length)
	}
	return v
}

func (v *Validator) MaxLength(field string, length int) *Validator {
	value := v.Request.FormValue(field)
	if len(value) > length {
		v.errors[field] = "Maximum length is " + fmt.Sprint(length)
	}
	return v
}

func (v *Validator) Email(field string) *Validator {
	value := v.Request.FormValue(field)
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(emailRegex, value); !match {
		v.errors[field] = "Invalid email format"
	}
	return v
}

func (v *Validator) Match(field string, pattern string) *Validator {
	value := v.Request.FormValue(field)
	if match, _ := regexp.MatchString(pattern, value); !match {
		v.errors[field] = "Invalid format"
	}
	return v
}

func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

func (v *Validator) Errors() map[string]string {
	return v.errors
}
