package base

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func ReadFileAtOnce() {
	f, err := os.OpenFile("res/test.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error reading file - %s", err)
	}
	defer f.Close()

	fc, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("error reading file - %s", err)
	}
	fmt.Println("--- file ---")
	fmt.Println(string(fc))
	fmt.Println("--- file end ---")
}

func ListDirs() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("error reading dirs - %s", err)
	}
	fmt.Println(" - dir - ")
	for _, dir := range dirs {
		fmt.Printf("name : %s - type : %s \n", dir.Name(), dir.Type())
	}
	fmt.Println(" - dir end - ")
}

func ListFilesRecursive() {
	var files []string
	wdf := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	}
	err := filepath.WalkDir("./", wdf)
	if err != nil {
		fmt.Printf("error walking dirs - %s", err)
	}
	fmt.Println(files)

}

func BufferedReader() {
	f, err := os.OpenFile("res/test.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println("text - ", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func BufferedWriter() {
	f, err := os.Create("res/testWrite.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.WriteString("model,year,maker\n")
	writer.WriteString("innova,2021,toyota\n")
	writer.WriteString("alto,2011,maruti\n")
	writer.Flush()
}
