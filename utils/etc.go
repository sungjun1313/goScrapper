package utils

import (
	"fmt"
	"strings"
)

//LenAndUpper returns length of name and upper string of name
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//LenAndUpperUpdate returns length of name and upper string of name
func LenAndUpperUpdate(name string) (length int, uppercase string) {
	//length와 uppercase를 어떤 타입으로 리턴할 것인지 미리 정해놓았기 때문에
	//값을 대입해주고 return만 하면 된다.

	//defer은 함수가 모두 실행되서 끝난 후에 작동한다.
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

//CanIDrink returns bool
func CanIDrink(age int) bool {
	if age < 18 {
		return false
	}

	return true
}

//CanIDrinkKorea returns bool
func CanIDrinkKorea(age int) bool {
	//if문 안에서 koreanAge 변수를 만들고 조건에 만족하는 지 체크한다.
	//이 변수는 if문 안에서만 사용가능하다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

//CanIDrinkSwitch returns bool
func CanIDrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

//CanIDrinkOtherSwitch returns bool
func CanIDrinkOtherSwitch(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

//CanIDrinkKoreanSwitch returns bool
func CanIDrinkKoreanSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

//ShowAddress show memory address
func ShowAddress() {
	//변수 앞에 &를 붙이면 메모리 주소를 참조할 수 있다.
	a := 3
	b := &a
	fmt.Println(b)
	//메모리 주소를 담은 변수 앞에 *를 붙이면 그 주소의 해당값을 볼 수 있다,
	fmt.Println(*b)

	*b = 20
	fmt.Println(a)
}

//ShowArray shows array
func ShowArray() {
	//[]안에 숫자를 넣으면 배열의 개수를 제한할 수 있으며, 숫자를 안 넣으면 제한이 없다.
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala"
	names[4] = "dlsc"
	fmt.Println(names)
}

//ShowSlices shows slices
func ShowSlices() {
	names := []string{"csis", "fiw"}
	//append는 값을 추가한 새 배열을 리턴하므로 다시 잡아주어야한다.
	names = append(names, "vsiw")
	fmt.Println(names)
}

//ShowMap shows map
func ShowMap() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

//ShowStructs show structs
func ShowStructs() {
	favFood := []string{"kimch", "ramen"}
	nico := person{name: "nico", age: 20, favFood: favFood}
	fmt.Println(nico.name, nico.age)
}
