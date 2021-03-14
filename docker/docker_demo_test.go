package docker

import "testing"

func TestCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skiping")
	}
}