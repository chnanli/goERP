package product

import (
	"encoding/json"
	"pms/controllers/base"
	mp "pms/models/product"
	"strconv"
	"strings"
)

type ProductAttributeController struct {
	base.BaseController
}

func (ctl *ProductAttributeController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductAttributeController) Get() {
	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "detail":
		ctl.Detail()
	default:
		ctl.GetList()

	}
	ctl.URL = "/product/attribute"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuProductAttributeActive"] = "active"
}
func (ctl *ProductAttributeController) Edit() {

}
func (ctl *ProductAttributeController) Create() {
	method := strings.ToUpper(ctl.Ctx.Request.Method)
	if method == "GET" {
		ctl.Data["Action"] = "create"
		ctl.Data["Readonly"] = false
		ctl.Data["listName"] = "创建属性"
		ctl.TplName = "product/product_attribute_form.html"

	}
}
func (ctl *ProductAttributeController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductAttributeController) PostCreate() {

}
func (ctl *ProductAttributeController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	if _, err := mp.GetProductAttributeByName(name); err != nil {
		result["valid"] = true
	} else {
		result["valid"] = false
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductAttributeController) productAttributeList(start, length int64, condArr map[string]interface{}) (map[string]interface{}, error) {

	var arrs []mp.ProductAttribute
	paginator, arrs, err := mp.ListProductAttribute(condArr, start, length)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["code"] = line.Code
			oneLine["sequence"] = line.Sequence
			mapValues := make(map[int64]string)
			oneLine["Id"] = line.Id
			values := line.ValueIds
			for _, line := range values {
				mapValues[line.Id] = line.Name
			}
			//测试代码
			mapValues[12] = "11231232"
			mapValues[2] = "21231230"
			mapValues[3] = "121321"
			mapValues[4] = "20123"
			mapValues[52] = "12"
			mapValues[72] = "20"
			mapValues[21] = "12"
			mapValues[37] = "20"
			oneLine["values"] = mapValues
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
		if jsonResult, er := json.Marshal(&paginator); er == nil {
			result["paginator"] = string(jsonResult)
			result["total"] = paginator.TotalCount
		}
	}
	return result, err
}
func (ctl *ProductAttributeController) PostList() {
	condArr := make(map[string]interface{})
	start := ctl.Input().Get("offset")
	length := ctl.Input().Get("limit")
	var (
		startInt64  int64
		lengthInt64 int64
	)
	if startInt, ok := strconv.Atoi(start); ok == nil {
		startInt64 = int64(startInt)
	}
	if lengthInt, ok := strconv.Atoi(length); ok == nil {
		lengthInt64 = int64(lengthInt)
	}
	if result, err := ctl.productAttributeList(startInt64, lengthInt64, condArr); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductAttributeController) GetList() {
	ctl.Data["tableId"] = "table-product-attribute"
	ctl.TplName = "base/table_base.html"
}
