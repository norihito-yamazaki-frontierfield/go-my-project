package main

import "testing"

func TestHello(t *testing.T) {
	// local function
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// このメソッドがヘルパーであることをテストスイートに伝えるために必要です。こうすることで、テストが失敗したときに報告される行番号は、テストヘルパーの中ではなく 呼び出された関数 の中を示します。
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })
    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
	t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
}