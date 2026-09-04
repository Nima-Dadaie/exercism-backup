package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
    if len(a) != len(b) {
        return 0, errors.New("a and b must be of equal length")
    } else if a == b {
        return count, nil
    } else {
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
               count += 1 
            }
        }
         return count, nil
    }
}
