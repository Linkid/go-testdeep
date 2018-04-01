package testdeep

import (
	"testing"
)

// Edge cases not tested elsewhere...

func TestTestDeepBase(t *testing.T) {
	td := TestDeepBase{}

	td.setLocation(200)
	if td.location.File != "???" && td.location.Line != 0 {
		t.Errorf("Location found! => %s", td.location)
	}
}

func TestTdSetResult(t *testing.T) {
	if tdSetResultKind(199).String() != "?" {
		t.Errorf("tdSetResultKind stringification failed => %s",
			tdSetResultKind(199))
	}
}
