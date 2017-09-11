package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type result map[string][]string

func fileHashToHexString(file *os.File) string {
	p1 := md5.New()
	if _, err := io.Copy(p1, file); err != nil {
		log.Fatal(err)
	}

	q := hex.EncodeToString(p1.Sum(nil))
	return q
}
func compareAndReturn(input []string) result {
	solution := make(map[string][]string)
	for i := 0; i < len(input); i++ {
		f1, _ := os.Open(input[i])
		r := fileHashToHexString(f1)
		if _, ok := solution[r]; !ok {
			solution[r] = append(solution[r], input[i])
		} else {

			solution[r] = append(solution[r], input[i])
		}

	}
	return solution
}
func f1(path string) []result {
	path = strings.Replace(path, "\n", "", -1)
	if path == "" {
		path = "./"
	}

	return listFiles(path)
}

func listFiles(path string) []result {
	fileNames, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	files := []string{}
	dirs := []string{}

	for _, f := range fileNames {
		if !f.IsDir() {
			files = append(files, filepath.Join(path, f.Name()))
		} else {
			dirs = append(dirs, filepath.Join(path, f.Name()))
		}
	}

	results := []result{}

	r := compareAndReturn(files)
	results = append(results, r)

	for _, d := range dirs {
		rd := listFiles(d)
		results = append(results, rd...)
	}

	return results
}
func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter Directory Path or press enter for default directory ./:")
	text := "/Users/akshaydeo/Downloads/" //, _ := reader.ReadString('\n')
	solution := f1(text)
	fmt.Println("The Duplicate Files are:\n")
	for _, r := range solution {
		for _, v := range r {
			if len(v) > 1 {
				fmt.Println("Files:\n", v)
			}
		}
	}
}
