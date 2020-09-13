package main

//go가 찾을 수 있게 package를 명시, 소문자로만 사용한다.

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sungjun1313/goScrapper/accounts"
	"github.com/sungjun1313/goScrapper/mydict"
	"github.com/sungjun1313/goScrapper/scrapper"
	"github.com/sungjun1313/goScrapper/utils"
)

var errRequestFailed = errors.New("Request failed")

const fileName string = "jobs.csv"

type urlResultStruct struct {
	url    string
	status string
}

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

	//goRoutine
	//기본적으로 코드는 한 코드가 실행되서 끝난 뒤에 다음 코드가 실행되는데
	//goRoutine은 코드를 동시에 실행할 수 있다.
	//main method가 실행되는 동안만 동시에 실행 가능하다.
	go utils.SayCount("Kim")
	utils.SayCount("Sung")

	//Channel은 goRoutine과 메인함수 사이에 정보를 전달할 수 있게 해준다.
	//위 설명에서 goRoutine은 main메서드가 실행된 상태에서만 작동하는데
	//Channel을 사용하면 Channel로부터 값을 받을 때까지 먼저 종료하지 않는다.
	c := make(chan string)
	people := [2]string{"Gang", "John"}
	for _, person := range people {
		go utils.SayPerson(person, c)
	}
	for i := 0; i < len(people); i++ {
		//<-c는 blockinig operation 이다.
		fmt.Println(<-c)
	}

	//여기까지 기초개념

	//bank 프로젝트 시작
	account := accounts.NewAccount("nicolas")
	account.Deposit(10)
	err := account.Withdraw(5)
	if err != nil {
		//Fatalln은 에러를 출력하고 더 이상 코드를 실행하지 않는다.
		log.Fatalln(err)
	}

	//그냥 account를 호출하면 python에서 __str__ 을 호출하듯이 자동으로 String method를 호출한다.
	//account의 String method가 있으면 그것을 호출하고 없으면 자동으로 Go에서 정의된 String method가 호출된다.
	fmt.Println(account)
	//bank 프로젝트 끝

	//dictionary 프로젝트 시작
	dictionary := mydict.Dictionary{"first": "First word", "del": "Soon delete"}
	definition, err2 := dictionary.Search("second")
	if err2 != nil {
		//에러를 출력하고 계속 코드를 실행한다.
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}
	err3 := dictionary.Add("hello", "Greeting")
	if err3 != nil {
		fmt.Println(err3)
	}
	err4 := dictionary.Update("first", "First word22")
	if err4 != nil {
		fmt.Println(err4)
	}
	err5 := dictionary.Delete("del")
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(dictionary)
	//dictionary 프로젝트 끝

	//URLChecker 프로젝트 시작
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	urlResults := make(map[string]string) //map을 빈값으로 만들어준다
	ch := make(chan urlResultStruct)

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		result := <-ch
		urlResults[result.url] = result.status
	}

	for url, status := range urlResults {
		fmt.Println(url, status)
	}

	//URLChecker 프로젝트 끝

	//Scrapper 프로젝트 시작
	//go get github.com/PuerkitoBio/goquery 실행해서 라이브러리를 설치한다.
	//go get github.com/labstack/echo 실행해서 라이브러리를 설치한다,
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//go에는 buffalo라는 웹프레임워크가 있어서 참고하면 좋다.

	//Scrapper 프로젝트 끝
}

func hitURL(url string, ch chan<- urlResultStruct) {
	//chan<- 는 보내기만 가능하다는 뜻이다.

	//go lang std library 검색하면 많은 라이브러리를 이용할 수 있다.
	fmt.Println("Checking: ", url)
	//http 라이브러리를 import하면 go build runtime/cgo: copying ..... permission denied 에러 발생
	//sudo chmod -R 777 /usr/local/go 를 실행해서 권한을 부여하면 에러가 사라진다.
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- urlResultStruct{url: url, status: status}
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}
