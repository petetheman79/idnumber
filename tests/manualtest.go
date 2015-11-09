package tests

import (
	"net/url"

	//"github.com/revel/revel"
	"github.com/revel/revel/testing"
)

// SingleTest is a test suite of Manual controller.
type ManualTest struct {
	testing.TestSuite
}

// Test that the controller can process a single ID number entered manually - Valid
func (t *ManualTest) TestThatManualEntryWorks() {	
	t.PostForm("/Manual/Capture", url.Values{
		"idnumber": { "7808215185082" },		
	})
	t.AssertOk()
	t.AssertContains("Valid")
	t.AssertContains("Male")
	t.AssertContains("South African")
	t.AssertContains("21 Aug 1978")		
}

// Test that the controller can process a single ID number entered manually - Invalid
func (t *ManualTest) TestThatInvalidManualEntryWorks() {
	t.PostForm("/Manual/Capture", url.Values{
		"idnumber": { "7808215185083" },		
	})
	t.AssertOk()
	t.AssertContains("Invalid")
	t.AssertContains("Male")
	t.AssertContains("South African")
	t.AssertContains("21 Aug 1978")	
}

// Test that the controller can process a single ID number entered manually - Empty Data
func (t *ManualTest) TestThatEmptyManualEntryWorks() {
	t.PostForm("/Manual/Capture", url.Values{
		"idnumber": { "" },		
	})
	t.AssertOk()
	t.AssertContains("The ID number must have one or more characters")	
}

// Test that the controller can process a single ID number entered manually - Cannot Compute Data
func (t *ManualTest) TestThatCannotComputeManualEntryWorks() {
	t.PostForm("/Manual/Capture", url.Values{
		"idnumber": { "jkhds" },		
	})
	t.AssertOk()
	t.AssertContains("Cannot compute")	
}


