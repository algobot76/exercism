/*
Package twofer implements a simple function that shows who you share with.

It has the format "One for X, one for me."
*/
package twofer

import "fmt"

// ShareWith returns a string that shows who you share with
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
