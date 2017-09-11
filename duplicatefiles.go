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
	"sync"
)

func HashAndWrite(path string, hashValueMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	fileNames, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("exiting in readdir", err)
	}
	for _, f := range fileNames {
		if !f.IsDir() {

			absolutePath := filepath.Join(path, f.Name())
			file, err := os.Open(absolutePath)
			if err != nil {
				log.Fatal("exiting in opening", err)
			}
			p1 := md5.New()

			if _, err := io.Copy(p1, file); err != nil {

				log.Fatal("exiting in copy", err)
			}

			q := hex.EncodeToString(p1.Sum(nil))
			if val, ok := hashValueMap.Load(q); !ok {
				hashValueMap.Store(q, []string{absolutePath})
			} else {
				ip, ok := val.([]string)
				if ok {
					ip = append(ip, absolutePath)
					hashValueMap.Store(q, ip)
				}
			}
		}
	}

}
func ListDirectories(path string) []string {
	fileNames, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	dirs := []string{}
	dirs = append(dirs, path)
	for _, f := range fileNames {

		if !f.IsDir() {
			continue
		} else {
			dirs = append(dirs, filepath.Join(path, f.Name()))
		}
	}
	return dirs
}
func main() {
	hashValueMap := new(sync.Map)
	var wg sync.WaitGroup
	listDirs := ListDirectories("/Users/akshaydeo/Downloads")
	for dir := range listDirs {
		wg.Add(1)
		go HashAndWrite(listDirs[dir], hashValueMap, &wg)

	}
	wg.Wait()
	s := make(map[interface{}]interface{})
	fmt.Println("The Duplicate Files are:\n")
	hashValueMap.Range(func(k, v interface{}) bool {
		s[k] = v
		if len(v.([]string)) > 1 {
			fmt.Println("Files:\n", v.([]string))
		}

		return true
	})
}
