package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/stacktic/dropbox"
)

var Clientid string
var Clientsecret string
var TokenId string

func ListDirectories(path string) []string {
	dropBoxObject := dropbox.NewDropbox()
	fmt.Println("In DropBox:")
	dropBoxObject.SetAccessToken(TokenId)
	dropBoxMetaData, _ := dropBoxObject.Metadata(path, true, true, "", "", 1000)
	directories := []string{}
	for index, _ := range dropBoxMetaData.Contents {
		if dropBoxMetaData.Contents[index].IsDir == true {
			directories = append(directories, dropBoxMetaData.Contents[index].Path)
		}

	}
	directories = append(directories, path)
	return directories
}
func HashAndWrite(path string, hashValueMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	dropBoxObject := dropbox.NewDropbox()
	dropBoxObject.SetAccessToken(TokenId)
	dropBoxMetaData, _ := dropBoxObject.Metadata(path, true, true, "", "", 1000)
	for index, _ := range dropBoxMetaData.Contents {
		if dropBoxMetaData.Contents[index].IsDir == false {
			absolutePath := dropBoxMetaData.Contents[index].Path
			downloadedFile, size, _ := dropBoxObject.Download(absolutePath, "", 0)
			if size > 0 {
				hashValue := md5.New()
				if _, err := io.Copy(hashValue, downloadedFile); err != nil {

					log.Fatal("exiting in copy", err)
				}

				stringValueOfHash := hex.EncodeToString(hashValue.Sum(nil))
				if value, ok := hashValueMap.Load(stringValueOfHash); !ok {
					hashValueMap.Store(stringValueOfHash, []string{absolutePath})
				} else {
					fileArray, ok := value.([]string)
					if ok {
						fileArray = append(fileArray, absolutePath)
						hashValueMap.Store(stringValueOfHash, fileArray)
					}
				}

			}
		}

	}
}

/*func HashAndWrite(path string, hashValueMap *sync.Map, wg *sync.WaitGroup) {
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
}*/
func main() {
	hashValueMap := new(sync.Map)
	var wg sync.WaitGroup

	/*listDirs := ListDirectories("/Users/akshaydeo/Downloads")*/
	Clientid = "31rmr26bffk3ij8"
	Clientsecret = "n0rlqt27iuf7scp"
	TokenId = "KeymFkX_8yAAAAAAAAACQ9aPx3fvPufbDi6MIvsYheIQtmidTq9MkJKYTXfKpeIv"
	listDirs := ListDirectories("/")
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
