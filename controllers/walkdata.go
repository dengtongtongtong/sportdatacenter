package controllers

import (
	"encoding/json"
	"sportdatacenter/models"
	"strconv"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
)

// WalkdataController operations for Walkdata
type WalkdataController struct {
	beego.Controller
}

// URLMapping ...
func (c *WalkdataController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Walkdata
// @Param	body		body 	models.Walkdata	true		"body for Walkdata content"
// @Success 201 {object} models.Walkdata
// @Failure 403 body is empty
// @router / [post]
func (c *WalkdataController) Post() {
	var walkdata models.Walkdata
	json.Unmarshal(c.Ctx.Input.RequestBody, &walkdata)
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Walkdata by id
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Walkdata
// @Failure 403 :uid is empty
// @router /:uid/:date [get]
func (c *WalkdataController) GetOne() {
	uid := c.GetString(":uid")
	date, _ := strconv.ParseInt(c.GetString(":date"), 10, 64)
	c.Data["json"] = map[string]string{"name": "dengtongtong"}
	walkdata, _ := models.GetWalkdataByUId(uid, date)
	c.Data["json"], _ = json.Marshal(walkdata)
	c.ServeJSON()
}

// GetBatch ...
// @Title GetBatchData
// @Description get Walkdata by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Walkdata
// @Failure 403 :uid is empty
// @router /:uid/:startdate/:enddate [get]
func (c *WalkdataController) GetBatchData() {
	// uid := c.GetString(":uid")
}

// GetAll ...
// @Title GetAll
// @Description get Walkdata
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Walkdata
// @Failure 403
// @router / [get]
func (c *WalkdataController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Walkdata
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Walkdata	true		"body for Walkdata content"
// @Success 200 {object} models.Walkdata
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WalkdataController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Walkdata
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WalkdataController) Delete() {

}
