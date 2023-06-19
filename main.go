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
	fmt.Println("\n closure version 2")
	for i2 < 5{
		fmt.Println(P2())
		i2++
	}

	f := fib()
	i3 := 0
	fmt.Println("\n closure version 3")
	for i3 < 8{
		fmt.Println(f())
		i3++
	}

	sliceAdder := BufferThis("start")
	fmt.Println(sliceAdder("Add This"))
	fmt.Println(sliceAdder("Also add this"))
	
	fmt.Println("\n--------Decorator Function--------\n")
	dec := Decorater(AddTheseTwo)
	fmt.Println(dec(2))

	fmt.Println("\n--------Decorator Function 2--------\n")
	dec2 := Decorater2(StringBuilder)
	fmt.Println(dec2("start"))

	fmt.Println("\n--------Decorator Function 2--------\n")
	dec3 := Decorater3(SliceAdder)
	fmt.Println(dec3([]int{1,2,3,4,5}))
}
// ---------------------------------------------------
func BufferThis(start string) func(string)[]string{
	sliceWords := []string{start}
	return func(word string)[]string{
		sliceWords = append(sliceWords, word)
		return sliceWords
	}
}

func SliceAdder(slice []int)[]int{
	slice = append(slice, 999999)
	return slice
}
// ---------------------------------------------------

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

//--------------------- Function and decorator---------------------
func Decorater(fn func(int) int) func(int) int{
	return func(param int) int{
		fmt.Println("We now have other functionality with this")
		return fn(param)
	}
}

func AddTheseTwo(num int) int{
return num + num
}

func Decorater2(fn func(string) string) func(string)string{
	return func(param string)string{
		fmt.Println("string decorator")
		return fn(param)
	}
}

func StringBuilder(str string) string{
	return str + " added this string"
}


func Decorater3(fn func([]int)[]int)func([]int)[]int{
	return func(slice []int)[]int{
		fmt.Println("Slice handler decorater")
		return fn(slice)
	}
}



