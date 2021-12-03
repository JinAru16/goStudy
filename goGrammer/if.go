package main

var a int = 3

const b int = 5

const c int = 4

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	println(sum)

}
