package controllers

import (
	"github.com/petetheman79/idnumber/app/routes"
	"github.com/revel/revel"	
)

type Manual struct {
	App
}

func (c *Manual) Entry() revel.Result {
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


func (c *Manual) Capture(idnumber string) revel.Result {
		
	// Make sure the ID number contains more than one character
	c.Validation.MinSize(idnumber, 1).Message("The ID number must have one or more characters")	

	// Handle errors.	
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Manual.Entry())
	}
	
	id := GetID(idnumber)

	WriteIdToFile(id)
	
	err := c.Txn.Insert(&id)
	if err != nil {
		panic(err)
	}
	
	return c.RenderJson(map[string]interface{}{
		"ID": id,
	})
}

