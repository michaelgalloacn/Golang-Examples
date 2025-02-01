package exported_test

import (
	"testing"

	"github.com/michaelgalloacn/Golang-Examples/exported"
	"github.com/stretchr/testify/assert"
)

func TestPublicFunction(t *testing.T) {
	assert.Equal(t, "public", exported.PublicFunction())

	// Uncommenting this will cause the test to not compile
	// assert.Equal(t, "private", exported.privateFunction())
}
