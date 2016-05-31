// Provides a fibonacci sequence (starts from zero)

package fibonacci

var results = []int{0, 1}

func Sequence(n int) []int {
	var length = len(results)

	if n > 2 {
		for i := length; i < n; i++ {
			results = append(results, results[i-1]+results[i-2])
		}
	}

	return results[0:n]
}
