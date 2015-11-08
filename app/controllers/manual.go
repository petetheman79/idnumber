package controllers

import (
	"github.com/petetheman79/idnumber/app/routes"

	"github.com/revel/revel"
	
	"github.com/petetheman79/idnumber/app/util/idnumberutil"
	"github.com/petetheman79/idnumber/app/util/fileutil"
	"github.com/petetheman79/idnumber/app/util/dbutil"	
)

type Manual struct {
	App
}

func (c *Manual) Entry() revel.Result {
	return c.Render()
}


func (c *Manual) Capture(idnumber string) revel.Result {
		
	// Make sure the ID number contains more than one character
	c.Validation.MinSize(idnumber, 1).Message("The ID number must have one or more characters")	

	// Handle errors.	
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Manual.Entry())
	}
	
	id := idnumberutil.GetID(idnumber)

	fileutil.WriteIdToFile(id)
	dbutil.InsertID(id)
	
	return c.RenderJson(map[string]interface{}{
		"ID": id,
	})
}

