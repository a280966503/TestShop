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

type Specification struct {
	beego.Controller
}

func (c *Specification)ShowSpecification()  {


	c.TplName = "manager/specification.html"
}

func (c *Specification)FindPage()  {

	page, pageErr := c.GetInt("page")
	rows, rowsErr := c.GetInt("rows")

	requestBody := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&requestBody)



	resp := make(map[string]interface{})

	defer c.ServeJSON()

	if pageErr != nil || rowsErr != nil {

	}

	if page < 1 {
		page=1
	}
	if rows<10 {
		rows=10
	}

	var specifications []models.TbSpecification
	o := orm.NewOrm()

	var total int64
	if requestBody["name"]!= nil {

		count, _ := o.QueryTable("tb_specification").Filter("SpecName__icontains",requestBody["name"]).Count()
		if int64(page) >= count/int64(rows) {
			page=1
		}
		o.QueryTable("tb_specification").Filter("SpecName__icontains",requestBody["name"]).Limit(rows, (page-1)*rows).All(&specifications)

		total=count

	}else {
		o.QueryTable("tb_specification").Limit(rows, (page-1)*rows).All(&specifications)
		count, _ := o.QueryTable("tb_specification").Count()
		total=count
	}


	resp["rows"]=specifications
	resp["total"]=total
	c.Data["json"] = resp


	return

}



//{"specificationOptionList":[{"optionName":"16G","orders":"1"},{"optionName":"8G","orders":"2"}],"specification":{"specName":"内存"}}
func (c *Specification) Add()  {
	//接收网页端传过来的json数据,并转换model
	var requestBody models.Spec
	json.Unmarshal(c.Ctx.Input.RequestBody,&requestBody)
	//response格式定义
	resp := make(map[string]interface{})

	defer c.ServeJSON()

	specificationOptionList := requestBody.SpecificationOptionList

	//这里开始报错
	//把model插入TbSpecification表
	o := orm.NewOrm()
	o.Begin()//标识事务的开始

	//对specification赋值
	var specification models.TbSpecification
	specification.SpecName = requestBody.Specification.SpecName

	id, specificationErr := o.Insert(&specification)
	if specificationErr!=nil {
		resp["success"]=false
		resp["message"]="插入specification表失败"
		c.Data["json"]=resp
		//事务回滚
		o.Rollback()
		return
	}

	//把model插入TbSpecificationOption表
	for _,value := range specificationOptionList{
		var specificationOption models.TbSpecificationOption

		specificationOption.SpecId=int(id)
		specificationOption.OptionName=value.OptionName
		specificationOption.Orders=value.Orders

		_, specificationOptionErr := o.Insert(&specificationOption)

		if specificationOptionErr != nil {

			resp["success"]=false
			resp["message"]="插入specificationOption表失败"
			c.Data["json"]=resp
			//事务回滚
			o.Rollback()
			return
		}
	}

	//插入成功
	resp["success"]=true
	resp["message"]="添加成功"

	c.Data["json"] =resp
	o.Commit()  //提交事务

	return

}

//查找
func (c *Specification) FindOne()  {
	id, _ := c.GetInt("id")

	defer c.ServeJSON()

	resp := make(map[string]interface{})
	var specification models.TbSpecification
	specification.Id=id
	o := orm.NewOrm()
	o.Read(&specification)

	var specificationOptionList []models.TbSpecificationOption
	_, specificationOptionListErr := o.QueryTable("tb_specification_option").Filter("SpecId", id).All(&specificationOptionList)

	if specificationOptionListErr != nil {

	}

	resp["specification"]=specification
	resp["specificationOptionList"]=specificationOptionList

	c.Data["json"]=resp

}

//更新
func (c *Specification) Update()  {

	var requestBody models.Spec
	json.Unmarshal(c.Ctx.Input.RequestBody,&requestBody)

	resp := make(map[string]interface{})

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.Begin()
	_, specificationErr := o.Update(requestBody.Specification)

	if specificationErr != nil {
		resp["success"]=false
		resp["message"]="Specification更新失败"
		o.Rollback()
		return
	}


	for _,value := range requestBody.SpecificationOptionList {
		_, specificationOptionErr := o.Update(value)
		fmt.Println(value.Id)
		if specificationOptionErr != nil {
			resp["success"]=false
			resp["message"]="SpecificationOptionList更新失败"
			o.Rollback()
			fmt.Println("SpecificationOptionList更新失败")
			return
		}
	}
	fmt.Println("update2")
	o.Commit()
	resp["success"]=true

	c.Data["json"]=resp
	fmt.Println("-----------------------")
	fmt.Println(resp)

	return
}

//删除
func (c *Specification)Delete()  {
	ids := strings.Split(c.GetStrings("ids")[0],",")
	fmt.Println(ids[0])

	resp := make(map[string]interface{})

	defer c.ServeJSON()
	o := orm.NewOrm()
	o.Begin()//添加事物


	for _,value := range ids{
		var specification models.TbSpecification

		id, idErr := strconv.Atoi(value)
		if idErr != nil {

		}
		specification.Id = id
		var specificationOption models.TbSpecificationOption
		specificationOption.SpecId=id

		_, specificationErr := o.Delete(&specification)
		fmt.Println(specification.Id)
		if specificationErr != nil {
			o.Rollback()
			resp["flag"]=false
			fmt.Println(specificationErr)
			return
		}

		_, specificationOptionErr := o.Delete(&specificationOption)
		fmt.Println(specificationOptionErr)
		if specificationOptionErr!=nil {
			o.Rollback()
			resp["flag"]=false
			fmt.Println(specificationOptionErr)
			return
		}

	}

	o.Commit()
	resp["flag"]=true
	c.Data["json"]=resp
	return

}

func (c *Specification)SelectOptionList()  {
	defer c.ServeJSON()

	var specifications []models.TbSpecification

	o := orm.NewOrm()

	o.QueryTable("tb_specification").All(&specifications)

	resp := make([]map[string]interface{},len(specifications))

	for i,specification := range specifications{
		value := make(map[string]interface{})
		value["id"]=specification.Id
		value["text"]=specification.SpecName
		resp[i]=value
	}

	c.Data["json"]=resp
}

