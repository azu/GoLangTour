/**
 * Created with IntelliJ IDEA.
 * User: azu
 * Date: 2013/04/14
 * Time: 12:26
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func add(x , y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
