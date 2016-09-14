package function

import "testing"

func Test_Division(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("Fail to pass")
	} else {
		t.Log("Pass")
	}
}

func Test_Division_Zero(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("Should not go here.")
	} else {
		t.Log("Check zero correctlly.")
	}
}