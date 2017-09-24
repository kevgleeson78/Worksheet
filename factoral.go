package main

import (
	"fmt"
	"math/big"
)

func factoral(n uint64) (r *big.Int) {

	one, bn := big.NewInt(1), new(big.Int).SetUint64(n)

	r = big.NewInt(1)
	if bn.Cmp(one) <= 0 {
		return
	}
	for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
		r.Mul(r, i)
	}
	return
}

/*func add(number *big.Int) *big.Int {

	ten := big.NewInt(10)
	div := number.Div(number, ten)
	mod := number.Mod(number, ten)

	return number.Add(mod, div)

}
*/
func main() {

	fmt.Println(factoral(20))

}
