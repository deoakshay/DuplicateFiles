package main

import (
	"fmt"

	"github.com/stacktic/dropbox"
)

func main() {
	var err error
	var db *dropbox.Dropbox
	//var l *dropbox.List

	var clientid, clientsecret string
	var token string

	clientid = "31rmr26bffk3ij8"
	clientsecret = "n0rlqt27iuf7scp"
	//token = "KeymFkX_8yAAAAAAAAAB4admc5jZ6CrcY7-AMOZaZpk"

	// 1. Create a new dropbox object.
	db = dropbox.NewDropbox()

	// 2. Provide your clientid and clientsecret (see prerequisite).
	db.SetAppInfo(clientid, clientsecret)

	// 3. Provide the user token.
	if err = db.Auth(); err != nil {
		fmt.Println(err)
		return
	}
	// You can now retrieve the token if you want.
	token = db.AccessToken()
	db.SetAccessToken(token)

	// 4. Send your commands.
	// In this example, you will create a new folder named "demo".

	/*	folder := "demo"
		if _, err = db.CreateFolder(folder); err != nil {
			fmt.Printf("Error creating folder %s: %s\n", folder, err)
		} else {
			fmt.Printf("Folder %s successfully created\n", folder)
		}*/
	p, _ := db.GetAccountInfo()
	fmt.Println(p.DisplayName)
	fmt.Println("Root dir is,", db.RootDirectory)
	new1, _ := db.Metadata("/", true, true, "", "", 1000)
	dirs := []string{}
	for n, _ := range new1.Contents {
		//fmt.Println(new1.Contents[n].Path)
		dirs = append(dirs, new1.Contents[n].Path)

	}
	fmt.Println(dirs)
	//fmt.Println(new1.Path, new1.Root, new1.ParentSharedFolderID, new1.IsDir)
	//	fmt.Printf("%T", new1)

	/*down, siz, _ := db.Download("hello/", "", 0)
	//op := make([]byte, siz)
	//down.Read(op)
	//op := NewReader(down)
	//body, err := ioutil.ReadAll(op)
	p1 := md5.New()
	if _, err := io.Copy(p1, down); err != nil {

		log.Fatal("exiting in copy", err)
	}

	q := hex.EncodeToString(p1.Sum(nil))
	fmt.Println(q)
	fmt.Println(siz)
	fmt.Println(new1.Contents)*/
}
