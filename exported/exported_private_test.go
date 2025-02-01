package exported

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicAndPrivateFunction(t *testing.T) {
	assert.Equal(t, "public", PublicFunction())
	assert.Equal(t, "private", privateFunction())
}
