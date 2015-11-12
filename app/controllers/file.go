package controllers

import (
	"strings"
	"fmt"
	"path"
	"os"
	"syscall"
	"github.com/petetheman79/idnumber/app/routes"
	"github.com/revel/revel"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type File struct {
	App
}

func (c File) Upload() revel.Result {
	results, err := c.Txn.Select(ID{},
		`select * from ID`)
		
	if err != nil {
		panic(err)
	}

	var idlist []*ID
	for _, r := range results {
		id := r.(*ID)
		idlist = append(idlist, id)
	}

	return c.Render(idlist)
}

func (c File) HandleUpload(idnumberlist []byte) revel.Result {
	// Validation rules.
	c.Validation.Required(idnumberlist)
		
	c.Validation.MaxSize(idnumberlist, 2*MB).
		Message("File cannot be larger than 2MB")
		
	// Handle errors.
	if c.Validation.HasErrors() {
		fmt.Println("Has errors");
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.File.Upload())
	}
	
	listOfId := strings.Split(string(idnumberlist), "\r\n")
	
	fmt.Println(listOfId)
	var results []ID
	
	for i := 0; i < len(listOfId); i++ {
		idnumber := listOfId[i]
		id := GetID(idnumber)
		
		err := c.Txn.Insert(&id)
		if err != nil {
			panic(err)
		}
		results = append(results, id);
		fmt.Println("-----");
		fmt.Println(results)
	}
	
	WriteIdListToFile(results)
			
	result := c.RenderJson(map[string]interface{}{		
		"Status":      "Successfully uploaded",
		"List":		   results,
	});
	
	fmt.Println("Results:");
	fmt.Println(result)
	
	return result
}

func WriteIdListToFile(idList []ID) {
	for i := 0; i < len(idList); i++ {
		WriteIdToFile(idList[i])
	}
}

func WriteIdToFile(id ID) {
	toWrite := id.IDNumber + "\r\n"

	if id.Vadility == "Valid" {
		writeToFile(path.Join(revel.BasePath, "valid.txt"), toWrite)
	} else if id.Vadility == "Invalid" {
		writeToFile(path.Join(revel.BasePath,  "invalid.txt"), toWrite)
	} 	
}

func writeToFile(fileName string, data string) {
	fo, err := os.OpenFile(fileName, syscall.O_APPEND|syscall.O_WRONLY|os.O_CREATE, 600)
	
	if err != nil {
		panic(err)
	}
	
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)	
		}
	}()
	
	no, err := fo.WriteString(data)
		
	if err != nil {		
		panic(err)
		fmt.Println(no)
	}	
}
