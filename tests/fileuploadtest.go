package tests

import (
	"net/url"
	"path"

	//"github.com/petetheman79/idnumber/app/routes"
	//"github.com/petetheman79/idnumber/app/routes"

	"github.com/revel/revel"
	"github.com/revel/revel/testing"
)

// FileUpload is a test suite of File controller.
type FileUploadTest struct {
	testing.TestSuite
}


func (t *FileUploadTest) TestThatNoFilesUploadFails() {
	// Make sure it is not allowed to submit no file.
	t.PostFile("/File/HandleUpload", url.Values{}, url.Values{
		"idnumberlist": {
			//path.Join(revel.BasePath, "public/img/favicon.png"),
		},
	})
	t.AssertOk()
	t.AssertContains("Required")

}

func (t *FileUploadTest) TestThatSampleFileUploadWorks() {
	// Make sure it is not allowed to submit no file.
	t.PostFile("/File/HandleUpload", url.Values{}, url.Values{
		"idnumberlist": {
			path.Join(revel.BasePath, "public/sample/idnumberlist.txt"),
		},
	})
	t.AssertOk()	
	t.AssertOk()
	t.AssertContains("Successfully uploaded")	
}
