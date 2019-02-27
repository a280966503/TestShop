package routers

import (
	"pinyougou/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//brand
	beego.Router("/manager/brand.html", &controllers.BrandController{},"get:ShowBrand")
	beego.Router("/brand/search.do", &controllers.BrandController{},"post:FindPage")
	beego.Router("/brand/findById.do", &controllers.BrandController{},"get:FindById")
	beego.Router("/brand/update.do", &controllers.BrandController{},"post:Update")
	beego.Router("/brand/save.do", &controllers.BrandController{},"post:Save")
	beego.Router("/brand/delete.do", &controllers.BrandController{},"get:Delete")
	beego.Router("/brand/selectOptionList.do", &controllers.BrandController{},"get:SelectOptionList")

	//content
	beego.Router("/manager/content.html", &controllers.ContentController{},"get:ShowContent")
	beego.Router("/content/search.do", &controllers.ContentController{},"post:FindPage")
	beego.Router("/content/findOne.do", &controllers.ContentController{},"get:FindOne")
	beego.Router("/content/update.do", &controllers.ContentController{},"post:Update")
	beego.Router("/content/add.do", &controllers.ContentController{},"post:Add")
	beego.Router("/content/delete.do", &controllers.ContentController{},"get:Delete")

	//contentCategory

	beego.Router("/manager/contentCategory.html", &controllers.ContentCategoryController{},"get:ShowContentCategory")
	beego.Router("/contentCategory/findAll.do", &controllers.ContentCategoryController{},"get:FindAll")
	beego.Router("/contentCategory/search.do", &controllers.ContentCategoryController{},"post:FindPage")
	beego.Router("/contentCategory/add.do", &controllers.ContentCategoryController{},"post:Add")
	beego.Router("/contentCategory/findOne.do", &controllers.ContentCategoryController{},"get:FindOne")
	beego.Router("/contentCategory/update.do", &controllers.ContentCategoryController{},"post:Update")
	beego.Router("/contentCategory/delete.do", &controllers.ContentCategoryController{},"get:Delete")

	//goods
	beego.Router("/manager/goods.html", &controllers.GoodsController{},"get:ShowGoods")
	beego.Router("/goods/search.do", &controllers.GoodsController{},"post:FindPage")
	beego.Router("/goods/updateStatus.do", &controllers.GoodsController{},"get:UpdateStatus")
	beego.Router("/goods/delete.do", &controllers.GoodsController{},"get:Delete")

	//home
	beego.Router("/manager/home.html", &controllers.HomeController{},"get:ShowHome")


	//index
	beego.Router("/manager/index.html", &controllers.MainController{},)

	//itemCat
	beego.Router("/manager/item_cat.html", &controllers.ItemCatController{},"get:ShowItemCat")
	beego.Router("/itemCat/findAll.do", &controllers.ItemCatController{},"get:FindAll")
	beego.Router("/itemCat/findByParentId.do", &controllers.ItemCatController{},"get:FindByParentId")


	//seller
    beego.Router("/manager/seller_1.html", &controllers.SellerController{},"get:ShowSeller1")
    beego.Router("/manager/seller.html", &controllers.SellerController{},"get:ShowSeller")
    beego.Router("/seller/search.do", &controllers.SellerController{},"post:FindPage")
    beego.Router("/seller/findOne.do", &controllers.SellerController{},"GET:FindOne")
    beego.Router("/seller/updateStatus.do", &controllers.SellerController{},"GET:UpdateStatus")





	//specification
    beego.Router("/manager/specification.html", &controllers.Specification{},"get:ShowSpecification")
    beego.Router("/specification/search.do", &controllers.Specification{},"post:FindPage")
    beego.Router("/specification/add.do", &controllers.Specification{},"post:Add")
    beego.Router("/specification/findOne.do", &controllers.Specification{},"get:FindOne")
    beego.Router("specification/update.do", &controllers.Specification{},"post:Update")
    beego.Router("specification/delete.do", &controllers.Specification{},"get:Delete")
    beego.Router("specification/selectOptionList.do", &controllers.Specification{},"get:SelectOptionList")

	//type_template
	beego.Router("/manager/type_template.html", &controllers.TypeTemplateController{},"get:ShowTypeTemplate")
	beego.Router("/typeTemplate/search.do", &controllers.TypeTemplateController{},"post:FindPage")
	beego.Router("/typeTemplate/findOne.do", &controllers.TypeTemplateController{},"get:FindOne")
	beego.Router("/typeTemplate/add.do", &controllers.TypeTemplateController{},"post:Add")
	beego.Router("/typeTemplate/update.do", &controllers.TypeTemplateController{},"post:Update")
	beego.Router("/typeTemplate/delete.do", &controllers.TypeTemplateController{},"get:Delete")


	//upload
	beego.Router("upload/uploadFile.do", &controllers.UploadController{},"post:UploadFile")



}
