package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSesison(t *testing.T) {
	ts, gitlab := Stub("stubs/session/index.json")
	session, err := gitlab.GetSession("john@example.com", "samplepassword")

	assert.Equal(t, err, nil)
	assert.Equal(t, session.UserName, "john_smith")
	assert.Equal(t, session.Email, "john@example.com")
	defer ts.Close()
}
