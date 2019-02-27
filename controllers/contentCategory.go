package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pinyougou/models"
	"strconv"
	"strings"
)

type ContentCategoryController struct {
	beego.Controller
}

func (c *ContentCategoryController)ShowContentCategory()  {
	c.TplName="manager/contentCategory.html"
}

func (c *ContentCategoryController)FindAll()  {

	defer c.ServeJSON()


	var contentCategorys []models.TbContentCategory

	o := orm.NewOrm()
	o.QueryTable("tb_content_category").All(&contentCategorys)


	c.Data["json"]=contentCategorys
	return
}

func (c *ContentCategoryController)FindPage()  {
	page, _ := c.GetInt("page")
	rows, _ := c.GetInt("rows")

	defer c.ServeJSON()

	var contentCategorys []models.TbContentCategory
	o := orm.NewOrm()

	o.QueryTable("tb_content_category").Limit(rows,(page-1)*rows).All(&contentCategorys)
	total, _ := o.QueryTable("tb_content_category").Count()

	resp := make(map[string]interface{})

	resp["total"]=total
	resp["rows"]=contentCategorys

	c.Data["json"]=resp
}

//{name: "dyy"}
func (c *ContentCategoryController)Add()  {

	var contentCategory models.TbContentCategory

	json.Unmarshal(c.Ctx.Input.RequestBody,&contentCategory)

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.Insert(&contentCategory)

	c.Data["json"]=JsonResponse(true,"添加成功")
}

func (c *ContentCategoryController)FindOne()  {
	id, _ := c.GetInt("id")

	defer c.ServeJSON()

	var contentCategory models.TbContentCategory

	o := orm.NewOrm()

	o.QueryTable("tb_content_category").Filter("Id",id).All(&contentCategory)

	c.Data["json"]=contentCategory
}

func (c *ContentCategoryController)Update()  {

	var contentCategory models.TbContentCategory
	json.Unmarshal(c.Ctx.Input.RequestBody,&contentCategory)

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.Update(&contentCategory)

	c.Data["json"]=JsonResponse(true,"修改成功")
}

func (c *ContentCategoryController)Delete()  {
	ids := strings.Split(c.GetStrings("ids")[0],",")

	defer c.ServeJSON()

	o := orm.NewOrm()

	for _,id := range ids{
		var contentCategory models.TbContentCategory
		value, _ := strconv.Atoi(id)
		contentCategory.Id=value
		fmt.Println(value)
		fmt.Println(id)

		o.Delete(&contentCategory)

	}

	c.Data["json"]=JsonResponse(true,"删除成功")
}