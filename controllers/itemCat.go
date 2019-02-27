package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pinyougou/models"
)

type ItemCatController struct {
	beego.Controller
}

func (c *ItemCatController)ShowItemCat()  {
	c.TplName="manager/item_cat.html"
}

func (c *ItemCatController)FindAll()  {
	var itemCat []models.TbItemCat

	defer c.ServeJSON()

	o := orm.NewOrm()

	o.QueryTable("tb_item_cat").All(&itemCat)


	c.Data["json"]=itemCat
}

func (c *ItemCatController)FindByParentId()  {

	defer c.ServeJSON()
	parentId, _ := c.GetInt("parentId")

	var itemCats []models.TbItemCat
	defer c.ServeJSON()

	o := orm.NewOrm()

	o.QueryTable("tb_item_cat").Filter("ParentId",parentId).All(&itemCats)

	c.Data["json"]=itemCats



}