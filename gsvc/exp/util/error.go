package util

import "fmt"

func HandleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}
