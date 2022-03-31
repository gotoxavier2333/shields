package internal

import "testing"

var exp = "Test shields.io"

func TestPrintString(t *testing.T) {
	res := printString()
	if exp != res {
		t.Fatalf("exp: %s not matched res: %s", exp, res)
	}
	t.Logf("exp: %s, res: %s", exp, res)
}
