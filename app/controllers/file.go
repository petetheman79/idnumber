package controllers

import (
	"strings"
	"fmt"
	//"io/ioutil"	

	"github.com/petetheman79/idnumber/app/routes"

	"github.com/revel/revel"
	
	"github.com/petetheman79/idnumber/app/util/idnumberutil"
	"github.com/petetheman79/idnumber/app/util/fileutil"
	"github.com/petetheman79/idnumber/app/util/dbutil"
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
	return c.Render()
}

func (c File) HandleUpload(idnumberlist []byte) revel.Result {
	// Validation rules.
	c.Validation.Required(idnumberlist)
	//c.Validation.MinSize(idnumberlist, 2*KB).
	//	Message("Minimum a file size of 2KB expected")
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
	var results []idnumberutil.ID
	
	for i := 0; i < len(listOfId); i++ {
		idnumber := listOfId[i]
		id := idnumberutil.GetID(idnumber)
		results = append(results, id);
		fmt.Println("-----");
		fmt.Println(results)
	}
	
	fileutil.WriteIdListToFile(results)
	dbutil.InsertIDList(results)
		
	result := c.RenderJson(map[string]interface{}{		
		"Status":      "Successfully uploaded",
		"List":		   results,
	});
	
	fmt.Println("Results:");
	fmt.Println(result)
	
	return result
}
