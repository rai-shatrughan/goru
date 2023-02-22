package base

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func Hello() {
	fmt.Println("Hello")
}

func Add(a, b int64) {
	fmt.Println("Sum = ", a+b)
}

func Modulo(a, b int64) {
	fmt.Println("Mod = ", a%b)
}

func StrJoin(str1, str2 string) {
	fmt.Println("Joined String - ", str1+" "+str2)
}

func MultiLineString() {
	str := `Hi
	this 
	is 
	multi 
	line String`
	fmt.Println("Multi Line String - ", str)
}

type Person struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func JsonFromStruct() {
	p := &Person{"sr", 29}
	b, _ := json.Marshal(p)
	fmt.Println("struct print - ", string(b))
}

func JsonToStruct() {
	per := []byte(`{"name":"rr","age":22}`)
	var p2 Person
	json.Unmarshal(per, &p2)
	fmt.Println("struct from byte - ", p2)
}

func RandomInt() {
	show("random no. between 0 - 100 :", rand.Intn(100))
}

func RandomString() {
	rand.Seed(time.Now().UnixNano())
	show("random char:", string(rand.Intn(175-41)+41))
}

func show(name string, v1 any) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "%s\t%v\n", name, v1)
}

func RangeOverString(str string) {
	for _, c := range str {
		fmt.Println("char/rune", string(c))
	}
}

func SumOfSeries(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("Sum - ", total)
}
