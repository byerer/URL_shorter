package shorter

import (
	"fmt"
	"testing"
)

func TestShorten(t *testing.T) {
	url := "https://www.google.com"
	shortURL := Shorten(url)
	if shortURL == "" {
		t.Errorf("Shorten() = %v; want non-empty string", shortURL)
	}
	fmt.Println(shortURL)
}
