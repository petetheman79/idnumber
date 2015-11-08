package fileutil


import (	
		"fmt"
		"os"
		"syscall"
		"github.com/petetheman79/idnumber/app/util/idnumberutil"
)


func WriteIdListToFile(idList []idnumberutil.ID) {
	for i := 0; i < len(idList); i++ {
		WriteIdToFile(idList[i])
	}
}

func WriteIdToFile(id idnumberutil.ID) {
	if id.Vadility == "Valid" {
		writeToFile("valid.txt", id.IDNumber)
	} else if id.Vadility == "Invalid" {
		writeToFile("invalid.txt", id.IDNumber)
	} 	
}

func writeToFile(fileName string, data string) {
	fo, err := os.OpenFile(fileName, syscall.O_APPEND, 0)
	
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