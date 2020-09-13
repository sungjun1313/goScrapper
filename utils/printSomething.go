package utils

import (
	"fmt"
	"time"
)

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

//SayCount prints count
func SayCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, " is ", i)
	}
}

//SayPerson prints person
func SayPerson(person string, c chan string) {
	//c는 channel 타입이며  channel을 통해 main method에 bool타입을 전달한다.
	time.Sleep(time.Second * 5)
	c <- person + " is here"
}
