package helpers

import (
	"fmt"
	"testing"
)

func TestShortenURL(t *testing.T) {
	got, _ := ShortenURL("test", 6)
	want := "xxxx"

	if got == want {

	}
	got, _ = ShortenURL("https://detik.com", 6)
	fmt.Println("got1 ", got)
	want = "correct"

	got, _ = ShortenURL("https://detik.com", 0)
	fmt.Println("got2 ", got)
	want = "correct"

	got, _ = ShortenURL("detik.com", 6)
	fmt.Println("got3 ", got)
	want = "correct"

}
