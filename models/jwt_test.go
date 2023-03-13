package models_test

import (
	"testing"

	"github.com/moroz/lenslocked/models"
)

func TestIssueTokenForUser(t *testing.T) {
	id := 2137
	token, err := models.IssueTokenForUserID(id)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	if token == "" {
		t.Errorf("Expected non-empty string")
	}

	actual := models.DecodeSubjectFromAccessToken(token)
	if id != actual {
		t.Errorf("Expected JWT %s to decode to original user ID, got: %q", token, id)
	}
}
