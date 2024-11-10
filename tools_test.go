package toolkit

import "testing"

func TestRandomString(t *testing.T) {
	testTools := Tools{}

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string")
		return
	}
}
