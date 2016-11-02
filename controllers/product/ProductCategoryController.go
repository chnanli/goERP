package product

import (
	"pms/controllers/base"
	"pms/models/product"
	"pms/utils"
	"strconv"
)

const (
	productCategoryListCellLength = 3
)

type ProductCategoryController struct {
	base.BaseController
}

func (this *ProductCategoryController) Get() {
	action := this.GetString(":action")
	viewType := this.Input().Get("view_type")
	switch action {
	case "list":
		switch viewType {
		case "list":
			this.List()

		default:
			this.List()
		}
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	default:
		this.List()
	}
	this.Data["searchKeyWords"] = "产品类别"
	this.Data["Action"] = "create"
	this.Data["formName"] = "产品类别"
	this.Data["productRootActive"] = "active"
	this.Data["productCategoryActive"] = "active"
	this.URL = "/product/category"
	this.Data["URL"] = this.URL
	this.Layout = "base/base.html"
}
func (this *ProductCategoryController) Edit() {

}
func (this *ProductCategoryController) Create() {
	this.Layout = "base/base.html"
	this.TplName = "product/product_category_form.html"

}
func (this *ProductCategoryController) List() {
	this.Data["listName"] = "产品类别"

	this.TplName = "product/product_category_list.html"

	condArr := make(map[string]interface{})
	page := this.Input().Get("page")
	offset := this.Input().Get("offset")
	var (
		err         error
		pageInt64   int64
		offsetInt64 int64
	)
	if pageInt, ok := strconv.Atoi(page); ok == nil {
		pageInt64 = int64(pageInt)
	}
	if offsetInt, ok := strconv.Atoi(offset); ok == nil {
		offsetInt64 = int64(offsetInt)
	}
	var productcategories []product.ProductCategory
	paginator, err, productcategories := product.ListProductCategory(condArr, pageInt64, offsetInt64)

	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [productCategoryListCellLength]string{"类别", "上级类别", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, productcategory := range productcategories {
			oneLine := make([]interface{}, productCategoryListCellLength, productCategoryListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			edit := make(map[string]string)
			detail := make(map[string]string)
			id := int(productcategory.Id)

			lineInfo["id"] = id
			oneLine[0] = productcategory.Name

			edit["name"] = "编辑"
			edit["url"] = this.URL + "/edit/" + strconv.Itoa(id)
			detail["name"] = "详情"
			detail["url"] = this.URL + "/detail/" + strconv.Itoa(id)
			action["edit"] = edit
			action["detail"] = detail
			if productcategory.Parent != nil {
				oneLine[1] = productcategory.Parent.Name
			} else {
				oneLine[1] = "-"
			}

			oneLine[2] = action

			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = productCategoryListCellLength
		tableInfo.TitleIndexLen = productCategoryListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
