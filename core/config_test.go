package core

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sut := GetOtherReveals("")
		assertInt(t, int64(len(sut)), 0)
	})

	t.Run("non empty", func(t *testing.T) {
		sut := GetOtherReveals("100-200,1200-1400")

		assertInt(t, int64(len(sut)), 2)
		assertInt(t, sut[0][0], 100)
		assertInt(t, sut[0][1], 200)
		assertInt(t, sut[1][0], 1200)
		assertInt(t, sut[1][1], 1400)
	})
}

func assertInt(t *testing.T, got, expected int64) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}
