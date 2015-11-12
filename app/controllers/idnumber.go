package controllers


import (		
		//"strings"
		"time"
		"strconv"
		"fmt"
)

type ID struct {
	Id int
	IDNumber string
	Count int
	DateOfBirth string
	Gender string
	Citizenship string
	Vadility string
}


//Get the DateOfBirth
//DateTime Reference Mon Jan 2 15:04:05 MST 2006
func dateOfBirth(idnumber string) string {
	if len(idnumber) < 6 {
		return "Cannot compute Date of Birth"
	}
	
	datePortion := idnumber[0:6]
	t, err := time.Parse("060102", datePortion)
	
	if err != nil {
		return err.Error()
	}
	
	return t.Format("2 Jan 2006")		
}

//Get Gender
func gender(idnumber string) string {
	if len(idnumber) < 10 {
		return "Cannot compute Gender"
	}
	genderPortion := idnumber[6:10]
		
	var genderValue int
	var gender string
	
	genderValue, err := strconv.Atoi(genderPortion)
	
	if err != nil {
		return err.Error()
	}
	
	if genderValue >= 5000 {
		gender = "Male"
	} else if genderValue < 5000 {
		gender = "Female"
	} else {
		gender = "Invalid"
	}
	
	return gender
}

//Get Citerzenship
func citerzenship(idnumber string) string {
	if len(idnumber) < 11 {
		return "Cannot compute Citizenship"
	}

	countryPortion := idnumber[10:11]
	
	countryVal, err := strconv.Atoi(countryPortion)
		
	if err != nil {
		return err.Error()
	}
	
	var citizenship string
	
	if countryVal == 0 {
		citizenship = "South African"
	} else {
		citizenship = "Other"
	}
	
	return citizenship
}

//Get Vadility
func valid(idnumber string) string {
	//Reference: http://geekswithblogs.net/willemf/archive/2005/10/30/58561.aspx
	
	if(len(idnumber) != 13) {
		return "Invalid"
	}
	
	inputNumber := idnumber[:len(idnumber)-1] //remove the last digit off the list
	checkBit := idnumber[len(idnumber)-1:] //the last digit is the check bit
	check, err := strconv.Atoi(checkBit)
	
	if err != nil {
		return err.Error()
	}
	
	//Add all digits in odd positions
	oddNumbers := 0
	for i := 0; i < 6; i++ {
		number, err := strconv.Atoi(inputNumber[i*2:(i*2)+1])		
		
		if err != nil {			
			return err.Error()
		}		
		
		oddNumbers += number
	}

	evenNumbers := ""
	for i := 0; i < 6; i++ {
		number := inputNumber[i*2+1]
		evenNumbers += string(number)
	}
	evenNumberResult, err := strconv.Atoi(evenNumbers)
	
	if err != nil {
		return err.Error()
	}
	
	evenNumberResult *= 2
	evenNumberString := strconv.Itoa(evenNumberResult)
		
	sumEvenNumbers := 0
	for i := 0; i < len(evenNumberString); i++ {
		fmt.Println(evenNumberString[i:i+1])
		number, err := strconv.Atoi(evenNumberString[i:i+1])
		
		if err != nil {
			return err.Error()
		}
		
		sumEvenNumbers += number
	}
	fmt.Println("Results");
	
	answer := sumEvenNumbers + oddNumbers
	fmt.Println(answer);
		
	secondDigit := strconv.Itoa(answer)[1:2]
	
	secondDigitVal, err := strconv.Atoi(secondDigit)
	if err != nil {
		return err.Error()
	}
		
	finalAnswer := 10 - secondDigitVal
	fmt.Println(finalAnswer);
	
	validator := 0
	if finalAnswer >= 10 {
		validatorStr := strconv.Itoa(finalAnswer)
		validatorInt, err := strconv.Atoi(validatorStr[1:2])
		if err != nil {
			return err.Error()
		}
		validator = validatorInt
	} else {
		validator = finalAnswer
	}
	fmt.Print("Validator: ");
	fmt.Println(validator);
	if check == validator {
		return "Valid"
	} else {
		return "Invalid"
	}	
}

func GetID(idnumber string) ID {
	id := ID {
		0,
		idnumber,
		len(idnumber),
		dateOfBirth(idnumber),
		gender(idnumber),
		citerzenship(idnumber),
		valid(idnumber),		
	}
	return id
}