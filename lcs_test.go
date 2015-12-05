package main

import "testing"

type tests struct {
  s1 string // input 1
  s2 string // input 2
  s3 string // result LCS
}

func TestLCS(t *testing.T) {
  lcsTests := []tests {
      {"ABCD", "BCF", "BC"},
      {"", "", ""},
    }

    for _, ex := range lcsTests {
      actual := LCS(ex.s1, ex.s2)
      if actual != ex.s3 {
        t.Errorf("LCS(%s, %s): expected %s, actual %s", ex.s1, ex.s2, ex.s3, actual)
      }
    }
}
