//go가 찾을 수 있게 package를 명시
package main

import (
	"fmt"

	"github.com/sungjun1313/goScrapper/accounts"
	"github.com/sungjun1313/goScrapper/utils"
)

//entry point
func main() {
	//작은 따옴표를 쓰면 에러 발생
	fmt.Println("Hello World!")

	//utils 패키지로 정의된 파일들을 찾아서 필요한 함수 등을 바로 사용 가능
	utils.SayHello()

	//const 새로운 값으로 변경 불가, var 새로운 값으로 변경 가능
	const firstName string = "sung"
	var lastName string = "jun"
	lastName = "junhee"
	utils.SayAllArgs(firstName, lastName)

	//var name string = "sungjun13" 와 같음. 타입은 자동으로 잡아준다. 타입은 변경 불가
	//하지만 이 표현방식은 함수 안 변수에서만 작동한다.
	name := "sungjun"
	utils.SaySomething(name)

	utils.Multiply(2, 2)
	totalLength, upperName := utils.LenAndUpper("nice")
	_, upperName2 := utils.LenAndUpper("bad") // _는 무시할 수 있다.
	_, upperName3 := utils.LenAndUpperUpdate("good")
	utils.SaySomething(upperName)
	utils.SaySomething(upperName2)
	utils.SaySomething(upperName3)
	utils.Sum(totalLength, 3)
	utils.SuperSum(4, 5, 6, 7, 8)

	age := 17
	if utils.CanIDrink(age) {
		utils.SaySomething("I can drink")
	} else {
		utils.SaySomething("I can not drink")
	}

	if utils.CanIDrinkKorea(age) {
		utils.SaySomething("I can drink Korean version")
	} else {
		utils.SaySomething("I can not drink Korean version")
	}

	if utils.CanIDrinkSwitch(age) {
		utils.SaySomething("I can drink Switch version")
	} else {
		utils.SaySomething("I can not drink Switch version")
	}

	utils.ShowAddress()
	utils.ShowArray()
	utils.ShowSlices()
	utils.ShowMap()
	utils.ShowStructs()

	//여기까지 기초개념

	//bank 프로젝트 시작
	account := accounts.NewAccount("nicolas")
	fmt.Println(account)
}
