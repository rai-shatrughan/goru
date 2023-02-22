package base

import (
	"fmt"
	"time"
)

func GetCurrentTime() {
	t := time.Now()
	fmt.Println("time is - ", t)
}

func GetCurrentTimeUTC() {
	t := time.Now().UTC()
	fmt.Println("time is - ", t)
}

func GetCurrentTimeISO() {
	t := time.Now().Format(time.RFC3339)
	fmt.Println("time is - ", t)
}

func GetCurrentTimeMDSP() {
	t := time.Now().Format("2006-01-02T15:04:05Z")
	fmt.Println("time is - ", t)
}

func AddTime() {
	t := time.Now()
	d := time.Hour * 2
	fmt.Println("time after 2 hrs - ", t.Add(d))
}

func SubTime() {
	t := time.Now()
	d := time.Hour * 2
	fmt.Println("time before 2 hrs - ", t.Add(-d))
}

func ParseTimeFromString(tstr string) {
	t, err := time.Parse("2006-01-02T15:04:05Z", tstr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time - ", t.Format(time.RFC3339))
}
