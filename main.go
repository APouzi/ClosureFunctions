package main

import "fmt"

func main() {
	P := Adder(0)
	i := 0
	for i < 5{
		fmt.Println(P())
		i++
	}

	P2 := AdderV2()
	i2 := 0
	fmt.Println("closure version 2")
	for i2 < 5{
		fmt.Println(P2())
		i2++
	}

	f := fib()
	i3 := 0
	fmt.Println("closure version 3")
	for i3 < 8{
		fmt.Println(f())
		i3++
	} 
	
}

func Adder(num int) func() int {
	return func() int {
		num += 1
		return num
	}
}

func AdderV2() func() int{
	init := 0
	return func() int{
		init += 1
		return init
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b 
	}
}