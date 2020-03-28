package main

import "testing"

const expected_greeting_string = "Hello World";

func TestExpectedGreeting(t *testing.T) {
    test_greeting := GetGreetingString();
    if test_greeting != expected_greeting_string {
       t.Errorf("Greeting string is incorrect, got: %s, want: %s.", test_greeting, expected_greeting_string)
    }
}
//const not_expected_greeting_string = "Hello World!";
//func TestUnexpectedGreeting(t *testing.T) {
//    test_greeting := GetGreetingString();
//    if test_greeting != not_expected_greeting_string {
//       t.Errorf("Greeting string is incorrect, got: %s, want: %s.", test_greeting, not_expected_greeting_string)
//    }
//}
