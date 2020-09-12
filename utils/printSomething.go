package utils

import "fmt"

//func에서 소문자로 시작하면 private 대문자로 시작하면 public이다
//만약 import가 안 되면  go install 을 실행시켜준다.

func sayBye() {
	fmt.Println("Bye")
}

//SayHello prints Hello
func SayHello() {
	fmt.Println("Hello")
}

//SaySomething prints arg p
func SaySomething(p string) {
	fmt.Println(p)
}

//SayAllArgs prints words which have one more args of string
func SayAllArgs(words ...string) {
	//string 타입의 인자들을 모두 받는다.
	fmt.Println(words)
}
