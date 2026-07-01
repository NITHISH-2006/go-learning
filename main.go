package main

import ("fmt")
func main(){
	fmt.Println("Go prep")

	// var x int = 5 // explicit
	y := 10 // inferred — "short declaration"
	// var a float64 = 2.3
	// s := "hello world"
	// u := true
	z := 23
	n := 10

	fmt.Println(sumToN(n))
	fmt.Println(add(y,z))
	fmt.Println(max(y,z))
	fmt.Println(isEven(y))

	for i:= 1 ; i<= 20; i++{
		if i%3 == 0 && i%5 == 0{
			fmt.Println("FizzBuzz")
		}else if i%3 == 0{
			fmt.Println("Fizz")
		}else if i%5 == 0{
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}

	// fmt.Println(x+y)
	// fmt.Println(s)
	// fmt.Println(u)
}

func sumToN(n int) int{
	sum := 0
	for i := 0 ; i <= n ; i++{
		sum += i
	}
	return sum
}

func isEven(a int) bool{
	if a % 2 == 0{
		return true
	}
	return false
}

func max(a int, b int) int{
	if a > b {
		return a
	}
		return b	
}






func add(a int, b int) int{
	return a + b
}


