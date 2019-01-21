package vsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitState(t *testing.T) {
	m := New(DefaultTransitions)
	assert.Equal(t, m.state, StateReady)
}
