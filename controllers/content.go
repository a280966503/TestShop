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

type ContentController struct {
	beego.Controller
}

func (c *ContentController)ShowContent()  {
	c.TplName="manager/content.html"
}

func (c *ContentController)FindPage()  {

	rows, _ := c.GetInt("rows")
	page, _ := c.GetInt("page")

	if page < 1 {
		page=1
	}
	if rows<10 {
		rows=10
	}
	defer c.ServeJSON()

	var content []models.TbContent

	o := orm.NewOrm()

	o.QueryTable("tb_content").Limit(rows,(page-1)*rows).All(&content)
	total, _ := o.QueryTable("tb_content").Count()

	resp := make(map[string]interface{})

	resp["rows"]=content
	resp["total"]=total
	c.Data["json"]=resp
	return

}

func (c *ContentController)FindOne()  {
	id, _ := c.GetInt("id")

	defer c.ServeJSON()

	var content models.TbContent


	o := orm.NewOrm()

	o.QueryTable("tb_content").Filter("Id",id).All(&content)

	c.Data["json"]=content
	return

}

//{"id":9,"categoryId":1,"title":"1元秒月饼","url":"http://www.163.com","pic":"http://192.168.25.133/group1/M00/00/00/wKgZhVnJ1pyAFbWqAAFyVK2N7Ig973.jpg","status":"1","sortOrder":"3"}
func (c *ContentController)Update()  {

	var content models.TbContent
	json.Unmarshal(c.Ctx.Input.RequestBody,&content)

	value := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&value)
	fmt.Println(value)
	fmt.Println(content)

	defer c.ServeJSON()

	o := orm.NewOrm()

	_, err := o.Update(&content)

	resp := make(map[string]interface{})

	if err != nil {
		resp["flag"]=false
		resp["message"]="修改失败"
		return
	}

	resp["flag"]=true
	c.Data["json"]=resp
}

func (c *ContentController)Add()  {
	var content models.TbContent
	json.Unmarshal(c.Ctx.Input.RequestBody,&content)

	value := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,value)
	fmt.Println(value)

	defer c.ServeJSON()

	resp := make(map[string]interface{})

	o := orm.NewOrm()

	_, err := o.Insert(&content)
	if err != nil {

	}

	resp["flag"]=true
	c.Data["json"]=resp


}

func (c *ContentController)Delete()  {
	ids := strings.Split(c.GetString("ids"),",")

	defer c.ServeJSON()

	o := orm.NewOrm()
	for _,id := range ids {
		var content models.TbContent
		value, _ :=strconv.Atoi(id)
		content.Id=value
		o.Delete(&content)
	}

	c.Data["json"]=	JsonResponse(true,"删除成功")

	return
}