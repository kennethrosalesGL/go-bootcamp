package main

import "testing"

func TestPi(t *testing.T) {
    total := Pi()
    if total != 3.141592653589793238 {
       t.Errorf("PI was incorrect, got: %b, want: %d.", total, 10)
    }
}