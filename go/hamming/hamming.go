package hamming

import "errors"

func Distance(a, b string) (int, error) {
	n := len(a)
	distance := 0
	if n != len(b) {
		return 0, errors.New("Length of strings do not match!")
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
