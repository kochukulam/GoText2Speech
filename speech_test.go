package main

import (
  "testing"
)

func TestSpeech(t *testing.T) {
  want := "Success"
  ifgot := Speech(); got != want {
    t.Errorf("Speech() = %q, want %q",got,want)
  }
}