package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Chris")
    want := "Hello, Chris"

    if got != want {
		// `f`は、プレースホルダー値％qに値が挿入された文字列を作成できる形式
        t.Errorf("got %q want %q", got, want)
    }
}