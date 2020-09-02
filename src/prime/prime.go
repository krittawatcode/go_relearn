package prime

import (
	"fmt"
	"math"
)

var notPrimes [1000000]bool

func init() {
	fmt.Println("start init")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))

	for i := 2; i < sqrtLen; i++ {
		if notPrimes[i] {
			continue
		}
		notPrimes[i] = false
		for j := i * 2; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("finish init")
}

func IsPrime(x int) bool {
	return !notPrimes[x]
}
