package mac_test

import (
	"testing"

	mac "github.com/go-random/mac"
	"github.com/stretchr/testify/assert"
)

func TestMAC(t *testing.T) {
	// Setup
	randomizer := mac.NewRandomizer()
	randomMAC := randomizer.MAC()

	// Assert
	assert.Equal(t, 17, len(randomMAC))
}

func TestLocalMAC(t *testing.T) {
	// Setup
	randomizer := mac.NewRandomizer()
	randomMAC := randomizer.LocalMAC()

	// Assert
	assert.Equal(t, 17, len(randomMAC))
}

func TestGlobalMAC(t *testing.T) {
	// Setup
	randomizer := mac.NewRandomizer()
	randomMAC := randomizer.GlobalMAC()

	// Assert
	assert.Equal(t, 17, len(randomMAC))
}
