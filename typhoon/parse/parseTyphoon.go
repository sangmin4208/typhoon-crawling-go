package parse

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sangmin4208/typhoon-crawling-go/typhoon"
)

const tableSelector = "body > div.container > section > div > div.cont-wrap.cmp-typ-report > div:nth-child(2) > div.typhoon-report > div > div > div.over-scroll > table > tbody > tr"
const titleSelector = "div.typhoon-cont > div.title"

func ParseTyphoonInfo(body io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	data := doc.Find(titleSelector).Text()
	name := getTyphoonName(data)  // AERE
	num := getTyphoonNumber(data) // 04
	date := getReportDate(data)   // 202207031600

	return fmt.Sprintf("%v %v %v", name, num, date)
}
func ParseTyphoonTable(body io.Reader) []typhoon.Typhoon {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	typhoonList := make([]typhoon.Typhoon, 0)

	doc.Find(tableSelector).Each(func(i int, s *goquery.Selection) {
		typhoonList = append(typhoonList, parseRow(s))
	})
	return typhoonList
}

func parseRow(s *goquery.Selection) typhoon.Typhoon {
	date := s.Find("tr>td:nth-child(1)").Text()
	latitude, _ := strconv.ParseFloat(s.Find("tr>td:nth-child(2)").Text(), 64)
	longitude, _ := strconv.ParseFloat(s.Find("tr>td:nth-child(3)").Text(), 64)
	atm, _ := strconv.ParseInt(s.Find("tr>td:nth-child(4)").Text(), 10, 64)
	velocity, _ := strconv.ParseInt(s.Find("tr>td:nth-child(5)").Text(), 10, 64)

	return typhoon.Typhoon{
		Date:      date,
		Latitude:  latitude,
		Longitude: longitude,
		Atm:       int(atm),
		Velocity:  int(velocity),
	}
}

func getTyphoonName(data string) string {
	fmt.Println(data)
	sp := strings.IndexRune(data, '(')
	ep := strings.IndexRune(data, ')')
	return strings.TrimSpace(data[sp+1 : ep])
}
func getTyphoonNumber(data string) string {
	sp := strings.IndexRune(data, '제')
	ep := strings.IndexRune(data, '호')
	ns := strings.TrimSpace(data[sp+3 : ep])
	num, err := strconv.ParseInt(ns, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if num < 10 {
		ns = "0" + ns
	}
	return ns
}
func getReportDate(data string) string {
	sp := strings.IndexRune(data, '|')
	ep := strings.IndexRune(data, '분')
	temp := strings.TrimSpace(data[sp+1 : ep+3])

	yearIdx := strings.IndexRune(temp, '년')
	year := strings.TrimSpace(temp[yearIdx-4 : yearIdx])
	monthIdx := strings.IndexRune(temp, '월')
	month := strings.TrimSpace(temp[monthIdx-2 : monthIdx])
	dateIdx := strings.IndexRune(temp, '일')
	date := strings.TrimSpace(temp[dateIdx-2 : dateIdx])
	hourIdx := strings.IndexRune(temp, '시')
	hour := strings.TrimSpace(temp[hourIdx-2 : hourIdx])
	minuteIdx := strings.IndexRune(temp, '분')
	minute := strings.TrimSpace(temp[minuteIdx-2 : minuteIdx])
	return fmt.Sprintf("%v%v%v%v%v", year, month, date, hour, minute)
}
