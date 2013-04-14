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
	"time"
	"runtime"
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
	X  float64
	Y  float64
}

func do_vertex() {
	// Sqrt
	fmt.Println(Sqrt(10))
	// Struct
	vertex := Vertex{1, 2}
	vertex.X = 4
	fmt.Println("vertex.X", vertex.X)
	// pointer
	vertex_p := &vertex
	vertex_p.X = 1e9
	fmt.Println("vertex.X", vertex.X)

	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex
		r = Vertex{X: 1}  // Y:0 is implicit
		s = Vertex{}      // X:0 and Y:0
	)
	fmt.Println(p, q, r, s)

	// new
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

func slice() {
	p := []int{
		1, 2, 3, 5, 7, 9,
	}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])

	fmt.Println("p[4:] ==", p[0:len(p) - 1])

	// nil slice
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil")
	}

	for i, v := range p {
		fmt.Println("2**%d", i, v)
	}
}

func mapFunc() {
	m := make(map[string]Vertex)
	m["key"] = Vertex{
		40, 60,
	}
	fmt.Println(m["key"])
}

func mapCheckState() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func adder() func (int)int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closureFunc() {
	hypot := func(x , y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	// adder
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func switchFunc() {
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s ", os)
	}
}
func timeFunc() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good AfterMoon")
	default:
		fmt.Println("Good evening")
	}
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-1)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func abserFunc() {
	var abser Abser
	f := MyFloat(-math.Sqrt2)

	v := Vertex{3, 4}
	abser = f
	abser = &v
	abser = f
	fmt.Println(abser.Abs())
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

	// struct
	do_vertex()

	// slice
	slice()
	// map
	mapFunc()
	mapCheckState()
	// closure
	closureFunc()
	// switch
	switchFunc()
	timeFunc()
	// interface
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v_interface := &Vertex{3, 4}
	v_interface.Scale(5)
	fmt.Println(v_interface , v_interface.Abs())

	abserFunc()
}

