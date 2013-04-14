/**
 * Created with IntelliJ IDEA.
 * User: azu
 * Date: 2013/04/14
 * Time: 12:26
 */
package main

import (
	"fmt"
	"math"
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

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x*0.1
}

func constFunc() {
	const (
		String  = "Const が使えるのは Char,String,Boolean,Numeric"
		Pi      = 3.14
		Boolean = true
		Big     = 1<<100
		Small   = Big>>99
	)

	fmt.Println(Pi, String , Boolean)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func forFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum +=i
	}
	fmt.Println(sum)

	// while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := 1.0
	return z - (z*z - x)/(2*x)
}

func pow(x float64, n float64, limit float64) float64 {
	// let ?
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}

	// ~>
	// doesn't access ``v`` value
	return limit
}

type Vertex struct {
	X  int
	Y  int
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

	// Const
	constFunc()

	// for
	forFunc()

	// sqrt
	fmt.Println("sqrt: ", sqrt(2), sqrt(-4))

	// pow
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Sqrt
	fmt.Println(Sqrt(10))
	// Struct
	vertex := Vertex{1, 2}
	vertex.X = 4
	fmt.Println(vertex.X)
}
