package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pinyougou/models"
)

type SellerController struct {
	beego.Controller
}

func (c *SellerController) ShowSeller() {


	c.TplName = "manager/seller.html"
}

func (c *SellerController) ShowSeller1() {


	c.TplName = "manager/seller_1.html"
}

func (c *SellerController)FindPage()  {
	rows, _ := c.GetInt("rows")
	page, _ := c.GetInt("page")

	defer c.ServeJSON()

	var seller  []models.TbSeller

	o := orm.NewOrm()

	o.QueryTable("tb_seller").Limit(rows,(page-1)*rows).All(&seller)
	count, _ := o.QueryTable("tb_seller").Count()

	resp := make(map[string]interface{})

	resp["total"]=count
	resp["rows"]=seller
	c.Data["json"]=resp

}

func (c *SellerController)FindOne()  {
	id := c.GetString("id")

	defer c.ServeJSON()

	var seller models.TbSeller

	o := orm.NewOrm()

	o.QueryTable("tb_seller").Filter("SellerId",id).All(&seller)

	c.Data["json"]=seller

}

func (c *SellerController)UpdateStatus()  {
	sellerId := c.GetString("sellerId")
	status := c.GetString("status")

	defer c.ServeJSON()

	var seller models.TbSeller

	o := orm.NewOrm()

	o.QueryTable("tb_seller").Filter("SellerId",sellerId).All(&seller)



	seller.Status=status

	o.Update(&seller)

	c.Data["json"]=JsonResponse(true,"")

}
