package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const validPass = "testPass"
const invalidPass = "a"

func TestGenActivationToken(t *testing.T) {
	u := &Account{}
	assert.Nil(t, u.GenActivationToken())
	assert.NotEmpty(t, u.ActivationToken, "token should not be empty")
}

func TestPasswordHash(t *testing.T) {
	u := &Account{}
	assert.Nil(t, u.SetPassword(validPass))
	assert.Nil(t, u.ComparePassword(validPass), "password should match")
	assert.NotNil(t, u.ComparePassword(invalidPass), "password should not match")
}

func TestInvalidPassword(t *testing.T) {
	u := &Account{}
	assert.EqualError(t, u.SetPassword(invalidPass), ErrInvalidPassword.Error())
}

func TestAccountActive(t *testing.T) {
	nu, _ := NewAccountFor(0, "test@test.com")
	var users = []struct {
		in     Account
		active bool
	}{
		{Account{Disabled: true}, false},
		{Account{Disabled: false}, true},
		{Account{ActivationToken: "", Disabled: false}, true},
		{Account{ActivationToken: "abc", Disabled: false}, false},
		{Account{ActivationToken: "abc", Disabled: true}, false},
		{*nu, false},
	}

	for _, u := range users {
		assert.Equal(t, u.active, u.in.IsActive())

		// Set to active
		u.in.SetActive()
		assert.Equal(t, true, u.in.IsActive())
	}

}
