package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	nums := []int{1, 2, 3, 4, 10, 20, 30, 50, 80, 99, 100, 110, 220, 330, 400}
	for _, v := range nums {
		eFizz, eBuzz, eFizzBuzz := expected(v)
		gFizz, gBuzz, gFizzBuzz := FizzBuzz(v)

		assert.Equal(t, eFizz, gFizz)
		assert.Equal(t, eBuzz, gBuzz)
		assert.Equal(t, eFizzBuzz, gFizzBuzz)
	}

}

func expected(num int) (fizz int, buzz int, fizzBuzz int) {
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			fizz++
		}
		if i%5 == 0 {
			buzz++
		}
		if i%15 == 0 {
			fizzBuzz++
		}
	}
	return
}
