package main

import (
	"fmt"
	"strings"
)

// import "./util"

func mul(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lam(fn func() string) func() {
	return func() {
		fmt.Println("inside")
	}
}

// func lam() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func compute(fn func())

func lenAndUpper2(name string) (length int, uppercase string) {
	// return len(name), strings.ToUpper(name)
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
} // naked return

func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("done") //after function finished
	// return len(name), strings.ToUpper(name)
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
} // naked return

func superAdd(numbers ...int) int {
	// fmt.Println(numbers)
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(i, numbers[i])
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

func sw(age int) bool {
	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	// return false
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// switch {
	// case age < 18:
	// 	return false
	// case age >= 18:
	// 	return true
	// }
	// return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"burger"}
	// nic := person{"nico", 22, favFood}
	nic := person{name: "nico", age: 22, favFood: favFood} // struct,,,
	fmt.Println(nic.name)

	nico := map[string]string{"name": "nico", "age": "12"} // map
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}

	names := [5]string{"a", "b", "c"} // array
	names[3] = "d"
	fmt.Println(names)

	aa := []string{"1", "2"} // slice
	aa = append(aa, "3")
	fmt.Println(aa, len(aa))

	a := 2
	// b := a
	b := &a // 메모리 주소
	fmt.Println(a, *b)
	a = 10
	*b = 11 // 해당 메모리 주소의 데이터
	fmt.Println(a, *b)
	fmt.Println(&a, &b)
	// fmt.Println(b)

	fmt.Println(canIDrink(16))
	fmt.Println(sw(16))

	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)
	// fmt.Println("Hello World, go!")
	// util.SayHi()

	// const name string = ""

	// var name string = "yoo"
	// name := "yoo"
	// name = 1
	// fmt.Println(name)
	fmt.Println(mul(2, 2))
	totalLength, upperName := lenAndUpper("yoo")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpper2("yoo2")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpper3("yoo3")
	fmt.Println(totalLength, upperName)

	repeatMe("a", "b")
	// fmt.Println(repeatMe)
	lam(func() string { return "outside" })()
}
