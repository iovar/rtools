package tools

import "testing"

func TestNewUuid(t *testing.T) {
	wantLength := 36

	uuid := NewUuid()

	if len(uuid) != wantLength {
		t.Errorf("Want length of %d, got length of %d", wantLength, len(uuid))
	}
}
