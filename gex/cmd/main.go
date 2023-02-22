package main

import (
	"gex/base"
)

func main() {
	// baseExp()

}

func baseExp() {
	base.Hello()
	base.Add(10, 15)
	base.Modulo(15, 2)
	base.StrJoin("Hi", "there!")
	base.MultiLineString()
	base.JsonFromStruct()
	base.JsonToStruct()
	base.RandomInt()
	base.RandomString()
	base.RangeOverString("hello")
	a := []int{1, 2, 3, 4, 5}
	base.SumOfSeries(a...)
	base.ReadFileAtOnce()
	base.ListDirs()
	base.ListFilesRecursive()
	base.BufferedReader()
	base.BufferedWriter()
}
