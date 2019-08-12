package gomonorepo

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("user1", "bob", 50)

	if u.Name != "bob" {
		t.Errorf("expected Bob, got %s", u.Name)
	}
}
