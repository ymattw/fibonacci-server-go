// Provides a fibonacci sequence (starts from zero)

package fibonacci

import (
	"math/big"
)

var results = []*big.Int{big.NewInt(0), big.NewInt(1)}

func Sequence(n int) []*big.Int {
	length := len(results)

	if n > 2 {
		for i := length; i < n; i++ {
			r := big.NewInt(0)
			r.Add(results[i-1], results[i-2])
			results = append(results, r)
		}
	}

	return results[0:n]
}
