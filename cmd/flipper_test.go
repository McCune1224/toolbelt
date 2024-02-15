package main

import "testing"

func TestFlipName(t *testing.T) {
	type test struct {
		name string
		want string
	}

	tests := []test{
		{name: "First, Last", want: "Last, First"},
		{name: "First Middle, Last", want: "Last, First Middle"},
		{name: "First", want: ""},
	}

	for _, tc := range tests {
		got, err := FlipName(tc.name)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}

}

func TestEmailToName(t *testing.T) {
	type test struct {
		email string
		want  string
	}

	tests := []test{
		{email: "Alex_McCune@URMC.Rochester.edu", want: "Alex McCune"},
		{email: "First_D'angelo@URMC.Rochester.edu", want: "First D'angelo"},
		{email: "Foo@gmail.com", want: ""},
	}

	for _, tc := range tests {
		got, err := EmailToName(tc.email)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}
}
