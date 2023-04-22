package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	r := Roll("1d20")
	if r > 20 {
		t.Errorf("roll failed got %v, expected it to be less then %v", r, 20)
	}
}

func TestRollHelper(t *testing.T) {
	r := rollHelper("20")
	assert.Equal(t, r, 20, "The numbers should be the same")
}
