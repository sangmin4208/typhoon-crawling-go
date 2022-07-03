package write

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sangmin4208/typhoon-crawling-go/typhoon"
)

func TyphoonList(path string, title string, typhoons []typhoon.Typhoon) {
	filename := fmt.Sprintf("%v/%v.txt", path, filePrefix())
	data := title + "\n" + toStringTyphoons(typhoons)
	err := ioutil.WriteFile(filename, []byte(data), os.FileMode(0666))
	if err != nil {
		log.Fatal(err)
	}
}
func toStringTyphoons(typhoons []typhoon.Typhoon) string {
	s := make([]string, 0, len(typhoons))
	for _, typhoon := range typhoons {
		s = append(s, toStringTyphoon(typhoon))
	}
	return strings.Join(s, "")
}

func toStringTyphoon(typhoon typhoon.Typhoon) string {
	date := formatDate(typhoon.Date)
	latitude := int(typhoon.Latitude * 10)
	longitude := int(typhoon.Longitude * 10)
	atm := typhoon.Atm
	velocity := typhoon.Velocity
	formatted := fmt.Sprintf("%v 005 5 %v %v  %v  %v\n", date, latitude, longitude, atm, velocity)
	return formatted
}

func formatDate(d string) string {
	now := time.Now()
	year := now.Local().Year()
	month := int(now.Local().Month())
	temp := strings.Split(d, "일")
	date := temp[0]
	hour := strings.TrimSpace(strings.Split(temp[1], "시")[0])
	if month < 10 {
		return fmt.Sprintf("%v0%v%v%v", year, month, date, hour)
	}
	return fmt.Sprintf("%v%v%v%v", year, month, date, hour)
}

func filePrefix() int64 {
	return time.Now().UnixMilli()
}
