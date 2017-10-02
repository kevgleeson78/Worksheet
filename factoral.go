//Problem Sheet 4
//Factoral
//Author: Kevin Gleeson
//URL: https://github.com/kevgleeson78/Worksheet
//Source:
//https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
//Date 30/09/17

package main

// import fmt and math/big
import (
	"fmt"
	"math/big"
)

// factoral function using BigInts
func factoral(n uint64) (r *big.Int) {
	//declare two variables bn from function input
	one, bn := big.NewInt(1), new(big.Int).SetUint64(n)
	//declare r for multiplying in loop
	//and returning to second function
	r = big.NewInt(1)
	//for loop different to standard
	//i:=1 ; i <= bn; i++;
	//when is  == bn the loop will stop
	for i := big.NewInt(1); i.Cmp(bn) <= 0; i.Add(i, one) {
		//multiply r * i store in r for next itteration
		//until condition is met
		r.Mul(r, i)
	}
	//return final result
	return
}

//function for adding the result of the factoral function
func add(number *big.Int) *big.Int {
	//three variables declared for function
	ten := big.NewInt(10)
	sum := big.NewInt(0)
	mod := big.NewInt(0)
	//loop for adding the digits of teh number
	for ten.Cmp(number) < 0 {
		//get the remainder of the number
		sum.Add(sum, mod.Mod(number, ten))
		number.Div(number, ten)
	}
	//add and retun the result of the loop
	sum.Add(sum, number)
	return sum
}
func main() {
	fmt.Println(add(factoral(10)))

}
