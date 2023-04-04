package dog

import (
	"testing"
)

func TestWoof(t *testing.T) {
	res := Woof()

	if res != "woof" {
		t.Error("Incorrect woof!")
	}
}
