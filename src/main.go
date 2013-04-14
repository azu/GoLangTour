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
func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
