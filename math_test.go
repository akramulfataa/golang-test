package math_test

import (
	"testing"
)

func AddTest(t *testing.T) {
	tambah := Add(1, 2)
	if result != 3 {
		log.Logf("Got= %d but expect %d", tambah, 3)
		t.Fail()
	}
}
