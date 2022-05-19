package repo

import "testing"

func TestIn(t *testing.T) {
	if err := Init("../data"); err != nil {
		t.Errorf("err: %v", err)
	}
}
