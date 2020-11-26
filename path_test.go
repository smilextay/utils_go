package utils

import (
	"testing"
)

func TestParent(t *testing.T) {
	dir := `/logs/redis/`
	target := `/logs`
	dir = Parent(dir)
	t.Log(dir)
	if dir != target {
		t.Fail()
	}
}
