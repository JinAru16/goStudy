package main

func main() {
	msg := "Hello Google" // 변경전의 메세지
	say(&msg)             // &를 붙이게 되면 Hello Google이라는 글자를 복사하는게 아니라 이 문자열을 갖는 메모리 영역의 주소를 갖는거임.
	println(msg)
}
func say(msg *string) {
	println(*msg)
	*msg = "changed" // 변경후의 메세지
}
