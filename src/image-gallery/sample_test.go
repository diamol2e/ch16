package main

import "testing"

func Test(t *testing.T){

    got := 1
    want := 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}