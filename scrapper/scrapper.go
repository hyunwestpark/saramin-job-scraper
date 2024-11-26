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
	id 			string
	company		string
	title 		string
	location 	string
}

// Scrape Saramin with term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	var jobs []extractedJob
	ch := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, ch)
	}
	for i := 0; i < totalPages; i++ {
		job := <- ch
		jobs = append(jobs, job...)
	}
	writeJobs(jobs)
	fmt.Println("Done extracting ", len(jobs), " jobs")
}



func getPage(page int, url string, ch chan<- []extractedJob) {
	c := make(chan extractedJob)
	var jobs []extractedJob
	pageURL := url + "&recruitPage" + strconv.Itoa(page)
	fmt.Println("Requesting: ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	recruitItem := doc.Find(".item_recruit")
	// s는 각 아이템임
	recruitItem.Each(func(i int, card *goquery.Selection){
		// job 추출하는 goroutine 생성
		go extractJob(card, c)
	})
	// 채널에 보내진 값 받기
	for i := 0; i < recruitItem.Length(); i++ {
		job := <- c
		jobs = append(jobs, job)
	}
	ch <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	company := CleanString(card.Find(".area_corp>strong>a").Text())
	title := CleanString(card.Find(".area_job>h2").Text())
	location := CleanString(card.Find(".area_job>.job_condition>span>a").Text())
	// 추출한 job 채널로 보내기
	c <- extractedJob{id:id, company: company, title:title, location:location}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	// prevent memory leaks
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})
	fmt.Println(doc)
	return pages
}

func writeJobs(jobs []extractedJob) {
	// jobs.csv 파일 생성
	file, err := os.Create("jobs.csv")
	checkErr(err)

	// jobs.csv에 작성한다고 명시
	w := csv.NewWriter(file)
	// 모든 작업 종료 후 저장
	defer w.Flush()

	headers := []string{"ID", "Company", "Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.company, job.title, job.location}
		wErr := w.Write(jobSlice)
		checkErr(wErr)
	}
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