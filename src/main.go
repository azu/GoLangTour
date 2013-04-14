/**
 * Created with IntelliJ IDEA.
 * User: azu
 * Date: 2013/04/14
 * Time: 12:26
 */
package main

import (
	"fmt"
)

func add(x , y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = (sum*4)/9
	y = sum - x
	return
}

func varFunc() {
	var x, y, z int = 1, 2, 3
	var c, python, java = true, false, "no!!"
	fmt.Println(x, y, z, c, python, java)
}

func shortVarFunc() {
	// in func
	x, y, z := 1, 2, 3
	c, python, java := true, false, "no!!"
	fmt.Println(x, y, z, c, python, java)
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// split
	doubleDigit, singleDigit := split(17)
	fmt.Println(doubleDigit, singleDigit)

	// varFunc
	varFunc()
	// alt shortVarFunc
	shortVarFunc()
}
