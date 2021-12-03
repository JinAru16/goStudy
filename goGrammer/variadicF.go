package main

func main() {
	count, totoal := sum(1, 3, 5, 7, 9)
	println(count, totoal)
}

func sum(nums ...int) (int, int) { // (int, int)라고 써주는 건 리턴값이 여러개 일때 처음 리턴값은 int, 두번째 리턴값도 int라는 말.
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}
