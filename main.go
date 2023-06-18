package main

import "github.com/prashantacharya/js-like-functions/arrays"

func main() {
	numArray := arrays.Array[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stringArray := arrays.Array[string]{"Prashant", "World"}

	double := numArray.Map(func(v int) int {
		return v * 2
	})

	greeting := stringArray.Map(func(v string) string {
		return "Hello " + v
	})

	even := numArray.Filter(func(v int) bool {
		return v%2 == 0
	})

	odd := numArray.Filter(func(v int) bool {
		return v%2 != 0
	})

	println("Double: ")
	for _, val := range double {
		println(val)
	}

	println("Greeting: ")
	for _, val := range greeting {
		println(val)
	}

	println("Even: ")
	for _, val := range even {
		println(val)
	}

	println("Odd: ")
	for _, val := range odd {
		println(val)
	}

	println("Includes: ")
	println(numArray.Includes(5))
	println(numArray.Includes(10))
}
