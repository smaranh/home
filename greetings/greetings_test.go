package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello with name and checks the valid return
func TestHelloName(t *testing.T) {
	name := "Sam"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Sam") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//TestHelloEmpty calls greetings.Hello with an empty string and checks if error is thrown
func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
