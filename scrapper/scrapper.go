package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

//Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	scrappCh := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, scrappCh)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-scrappCh
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Links", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.location,
			job.salary,
			job.summary,
		}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, url string, scrappCh chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err2 := goquery.NewDocumentFromReader(res.Body)
	checkErr(err2)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	scrappCh <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	//각각 element에서 data-jk 속성이 있는 지 확인
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() //함수가 끝나면 바디를 닫아준다,

	//res.Body로부터 도큐먼트를 읽어온다.
	doc, err2 := goquery.NewDocumentFromReader(res.Body)
	checkErr(err2)

	//클래스가 pagination인 element들을 찾아서 각각 함수를 실행시켜준다,
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		//각 element 자식의 a 태그가 몇 개인지 파악한다.
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

//CleanString cleans a string
func CleanString(str string) string {
	//str의 좌우공백을 제거한 뒤에 단어 단위로 분리하여 배열로 만든다.
	//이후 다시 하나의 문장으로 합친다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
