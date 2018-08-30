/*
Package hamming calculates the Hamming difference between two DNA strands.
*/
package hamming

import "github.com/pkg/errors"

// Distance calculates the Hamming difference between two DNA strands. It returns an error if the two strands do not
// have the same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("two dna strands must have the same length")
	}

	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
