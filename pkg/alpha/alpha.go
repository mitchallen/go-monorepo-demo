// Author: Mitch Allen
// File: alpha.go

package alpha

import (
	"fmt"

	"github.com/mitchallen/coin"
)

func Hello() {
	fmt.Println("[alpha]: Hello!")
}

func CoinCount(limit int) map[bool]int {

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < limit; i++ {
		m[coin.Flip()]++
	}

	return m
}
