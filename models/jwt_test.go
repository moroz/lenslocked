package models_test

import (
	"testing"

	"github.com/moroz/lenslocked/models"
)

func TestIssueTokenForUser(t *testing.T) {
	user := &models.User{
		ID:           4,
		Email:        "user@example.com",
		PasswordHash: "test",
	}
	token, err := models.IssueTokenForUser(user)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	if token == "" {
		t.Errorf("Expected non-empty string")
	}

	id := models.DecodeAccessTokenClaims(token)
	if id != user.ID {
		t.Errorf("Expected JWT %s to decode to original user ID, got: %q", token, id)
	}
}
