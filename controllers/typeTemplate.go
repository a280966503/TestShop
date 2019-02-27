package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pinyougou/models"
	"strconv"
	"strings"
)

type TypeTemplateController struct {
	beego.Controller
}



func (c *TypeTemplateController)ShowTypeTemplate()  {
	c.TplName="manager/type_template.html"
}

func (c *TypeTemplateController)FindPage()  {
	page,_ := c.GetInt("page")
	rows,_ := c.GetInt("rows")

	defer c.ServeJSON()

	var typeTemplates []models.TbTypeTemplate

	o := orm.NewOrm()

	o.QueryTable("tb_type_template").Limit(rows,(page -1)*rows).All(&typeTemplates)
	total, _ := o.QueryTable("tb_type_template").Count()

	resp := make(map[string]interface{})

	resp["total"]=total
	resp["rows"]=typeTemplates

	c.Data["json"]=resp
}

func (c *TypeTemplateController)FindOne()  {
	id, _ := c.GetInt("id")

	defer c.ServeJSON()

	var typeTemplate models.TbTypeTemplate

	o := orm.NewOrm()

	o.QueryTable("tb_type_template").Filter("Id",id).All(&typeTemplate)

	c.Data["json"]=typeTemplate

}

//add struct
type Add struct {
	ID 		int64 	`json:"id"`
	CustomAttributeItems []struct {
		Text string `json:"text"`
	} `json:"customAttributeItems"`
	Name     string `json:"name"`
	BrandIds []struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"brandIds"`
	SpecIds []struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"specIds"`
}

//{"customAttributeItems":[{"text":"内存"}],"name":"电脑","brandIds":[{"id":1,"text":"联想"},{"id":3,"text":"三星"},{"id":4,"text":"小米"},{"id":2,"text":"华为"},{"id":9,"text":"苹果"},{"id":14,"text":"海尔"},{"id":7,"text":"中兴"}],"specIds":[{"id":26,"text":"尺码"}]}
func (c *TypeTemplateController)Add()  {

	var typeTemplate models.TbTypeTemplate

	var add Add

	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody,&add)

	brandIds, _ := json.Marshal(add.BrandIds)
	customAttributeItems,_ := json.Marshal(add.CustomAttributeItems)
	specIds, _ := json.Marshal(add.SpecIds)

	typeTemplate.Name=add.Name
	typeTemplate.BrandIds=string(brandIds)
	typeTemplate.CustomAttributeItems=string(customAttributeItems)
	typeTemplate.SpecIds=string(specIds)

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.Insert(&typeTemplate)

	c.Data["json"]=JsonResponse(true,"添加成功")

}

func (c *TypeTemplateController)Update()  {
	var typeTemplate models.TbTypeTemplate

	var add Add

	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody,&add)

	brandIds, _ := json.Marshal(add.BrandIds)
	customAttributeItems,_ := json.Marshal(add.CustomAttributeItems)
	specIds, _ := json.Marshal(add.SpecIds)

	typeTemplate.Id=add.ID
	typeTemplate.Name=add.Name
	typeTemplate.BrandIds=string(brandIds)
	typeTemplate.CustomAttributeItems=string(customAttributeItems)
	typeTemplate.SpecIds=string(specIds)

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.Update(&typeTemplate)

	c.Data["json"]=JsonResponse(true,"更新成功")
}

func (c *TypeTemplateController)Delete()  {
	ids := strings.Split(c.GetString("ids"), ",")

	defer c.ServeJSON()



	o := orm.NewOrm()
	for _,id := range ids{
		var typeTemplate models.TbTypeTemplate
		value, _ := strconv.Atoi(id)
		typeTemplate.Id=int64(value)
		o.Delete(&typeTemplate)
	}

	c.Data["json"]=JsonResponse(true,"删除成功")
}