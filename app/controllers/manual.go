package controllers

import (
	"github.com/petetheman79/idnumber/app/routes"

	"github.com/revel/revel"
	
	"github.com/petetheman79/idnumber/app/idnumberutil"
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

	return c.RenderJson(map[string]interface{}{
		"ID": id,
	})
}

