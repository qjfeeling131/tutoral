package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "pari"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("pari")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("pari") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
