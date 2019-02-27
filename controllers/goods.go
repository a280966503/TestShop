package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pinyougou/models"
	"strconv"
	"strings"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController)ShowGoods()  {
	c.TplName="manager/goods.html"
}

func (c *GoodsController)FindPage()  {

	page, _ := c.GetInt("page")
	rows, _ := c.GetInt("rows")


	var goods []models.TbGoods

	defer c.ServeJSON()

	o := orm.NewOrm()

	fmt.Println("========================================")
	_, err := o.QueryTable("tb_goods").Limit(rows, (page-1)*rows).All(&goods)

	fmt.Println("-----------------------")
	fmt.Println(err)

	count, _ := o.QueryTable("tb_goods").Count()



	resp := make(map[string]interface{})

	resp["total"]=count
	resp["rows"]=goods
	c.Data["json"]=resp
	return
}

//ids=149187842867912&status=2
func (c *GoodsController)UpdateStatus()  {
	ids := strings.Split(c.GetString("ids"), ",")
	status:= c.GetString("status")

	defer c.ServeJSON()
	o := orm.NewOrm()




	for _,id := range ids {
		var goods models.TbGoods
		o.QueryTable("tb_goods").Filter("Id",id).All(&goods)

		if goods.AuditStatus=="0" {
			goods.AuditStatus=status
			o.Update(&goods)
		}

	}

	c.Data["json"]=JsonResponse(true,"操作成功")

}

func (c *GoodsController)Delete()  {
	ids := strings.Split(c.GetString("ids"),",")

	defer c.ServeJSON()

	o := orm.NewOrm()
	for _,id := range ids {
		var goods models.TbGoods
		value, _ :=strconv.Atoi(id)
		goods.Id=uint64(value)
		o.Delete(&goods)
	}

	c.Data["json"]=	JsonResponse(true,"删除成功")

	return
}