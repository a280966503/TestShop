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

type BrandController struct {
	beego.Controller
}

func (c *BrandController)ShowBrand()  {

	c.TplName="manager/brand.html"
}

//{firstChar: "a", name: "a"}
func (c *BrandController)FindPage()  {

	page, _ := c.GetInt("page")
	rows, _ := c.GetInt("rows")

	requestBody := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&requestBody)

	if page < 1 {
		page=1
	}
	if rows<10 {
		rows=10
	}
	defer c.ServeJSON()

	var brand []models.TbBrand
	resp := make(map[string]interface{})
	var total int64

	o := orm.NewOrm()

	if requestBody["firstChar"]!=nil && requestBody["name"]!= nil {

		_, brandErr := o.QueryTable("tb_brand").Filter("Name__icontains", requestBody["name"]).Filter("FirstChar__icontains", requestBody["firstChar"]).All(&brand)
		count, _ := o.QueryTable("tb_brand").Filter("Name__icontains", requestBody["name"]).Filter("FirstChar__icontains", requestBody["firstChar"]).Count()


		total=count
		if brandErr!=nil {
			resp["total"]=count
			resp["row"]=""
			return
		}
	}else if requestBody["firstChar"]!=nil {
		_, brandErr := o.QueryTable("tb_brand").Filter("FirstChar__icontains", requestBody["firstChar"]).All(&brand)
		count, _ := o.QueryTable("tb_brand").Filter("FirstChar__icontains", requestBody["firstChar"]).Count()


		total=count
		if brandErr!=nil {
			resp["total"]=count
			resp["row"]=""
			return
		}
	}else if requestBody["name"]!=nil {
		_, brandErr := o.QueryTable("tb_brand").Filter("Name__icontains", requestBody["name"]).All(&brand)
		count, _ := o.QueryTable("tb_brand").Filter("Name__icontains", requestBody["name"]).Count()


		total=count
		if brandErr!=nil {
			resp["total"]=count
			resp["row"]=""
			c.Data["json"]=resp
			return
		}
	}  else {
		_, brandErr := o.QueryTable("tb_brand").Limit(rows, (page-1)*rows).All(&brand)
		count, _ := o.QueryTable("tb_brand").Count()
		total=count
		if brandErr!=nil {
			resp["total"]=0
			resp["row"]=""
			c.Data["json"]=resp
			return
		}
	}

	resp["total"]=total
	resp["rows"]=brand
	c.Data["json"]=resp


	return
}

//{id:xx,name:yy,firstChar:zz}
func (c *BrandController) FindById()  {
	id, _ := c.GetInt("id")

	defer c.ServeJSON()

	var brand models.TbBrand


	o := orm.NewOrm()

	o.QueryTable("tb_brand").Filter("Id", id).All(&brand)



	c.Data["json"]=brand

	return

}

func (c *BrandController)Update()  {

	defer c.ServeJSON()
	var brand models.TbBrand
	json.Unmarshal(c.Ctx.Input.RequestBody,&brand)

	resp := make(map[string]interface{})

	o := orm.NewOrm()

	_, err := o.Update(&brand)
	if err != nil {
		resp["flag"]=false
		resp["message"]="修改失败"
		c.Data["json"]=resp
		return
	}

	//{flag:true,message:xxx}
	resp["flag"]=true
	resp["message"]="修改成功"
	c.Data["json"]=resp

}

func (c *BrandController)Save()  {

	var brand models.TbBrand

	defer c.ServeJSON()

	json.Unmarshal(c.Ctx.Input.RequestBody,&brand)

	resp := make(map[string]interface{})

	o := orm.NewOrm()

	_, err := o.Insert(&brand)

	if err != nil {
		//{flag:true,message:xxx}
		fmt.Println(err)
		resp["flag"]=false
		resp["message"]="新建失败"
		c.Data["json"]=resp
		return
	}
	fmt.Println(brand)

	resp["flag"]=true
	resp["message"]="新建成功"
	c.Data["json"]=resp
	return

}

func (c *BrandController)Delete()  {
	ids := strings.Split(c.GetStrings("ids")[0],",")

	defer c.ServeJSON()
	resp := make(map[string]interface{})

	o := orm.NewOrm()
	for _,id := range ids{
		var brand models.TbBrand
		value, _ := strconv.Atoi(id)
		brand.Id=value
		o.Delete(&brand)
	}

	resp["flag"]=true
	resp["message"]="删除成功"
	c.Data["json"]=resp
	return

}

func (c *BrandController)SelectOptionList()  {

	var brands []models.TbBrand

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.QueryTable("tb_brand").All(&brands)

	resp := make([]map[string]interface{},len(brands))


	for i,brand := range brands{
		value := make(map[string]interface{})
		value["id"]=brand.Id
		value["text"]=brand.Name

		resp[i]=value

	}

	c.Data["json"]=resp
}
