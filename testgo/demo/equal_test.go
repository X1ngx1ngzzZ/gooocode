package demo

import (
	"testing"
  )
func TestEqual(t *testing.T) {
  a := 1
  b := 1
  shouldBe := true
  real := equal(a, b) 
  if  real != shouldBe {
    t.Errorf( "should be %v, but is:%v", a, b)
  }
}	