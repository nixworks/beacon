package beacon_test

import (
	beacon "."
	"testing"
)

func TestFilterEmpty(t *testing.T) {
	f := beacon.NewFilter(nil)
	containers := []*beacon.Container{
		{Labels: map[string]string{}},
		{Labels: map[string]string{"a": "aye"}},
		{Labels: map[string]string{"a": "aye", "b": "bee"}},
	}

	for n, container := range containers {
		if !f.MatchContainer(container) {
			t.Errorf("container %d did not match", n)
		}
	}
}
