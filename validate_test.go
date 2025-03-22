package validator_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/adityarizkyramadhan/validator-go"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	form := url.Values{}
	form.Add("email", "test@example.com")
	form.Add("username", "user")
	form.Add("password", "123456")

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.PostForm = form
	v := validator.New(r)

	v.Required("email").Email("email")
	v.Required("username").MinLength("username", 3).MaxLength("username", 10)
	v.Required("password").MinLength("password", 6)

	assert.True(t, v.IsValid())
}

func TestValidatorFailures(t *testing.T) {
	form := url.Values{}
	form.Add("email", "invalid-email")
	form.Add("username", "us")
	form.Add("password", "123")

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.PostForm = form
	v := validator.New(r)

	v.Required("email").Email("email")
	v.Required("username").MinLength("username", 3).MaxLength("username", 10)
	v.Required("password").MinLength("password", 6)

	assert.False(t, v.IsValid())
	errors := v.Errors()
	assert.Equal(t, "Invalid email format", errors["email"])
	assert.Equal(t, "Minimum length is 3", errors["username"])
	assert.Equal(t, "Minimum length is 6", errors["password"])
}

func TestMatch(t *testing.T) {
	form := url.Values{}
	form.Add("zipcode", "12345")

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.PostForm = form
	v := validator.New(r)

	v.Match("zipcode", `^\d{5}$`)

	assert.True(t, v.IsValid())
}
