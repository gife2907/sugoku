package sugoutils_test

import (
	"sugoutils"
	"testing"
)

func TestNextLocation01(t *testing.T) {
	var el []sugoutils.Location
	el = sugoutils.EmptyLocations(sugoutils.Grid99())
	if len(el) != 17 || el[0].ValCount != 6 {
		t.Errorf("EmptyLocations[0].ValCount = %d, len(el) = %d; wanted 6, 17", el[0].ValCount, len(el))
	}
}
