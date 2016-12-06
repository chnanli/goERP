package product

import (
	"fmt"
	"pms/models/base"
	"pms/utils"

	"github.com/astaxie/beego/orm"
)

type ProductAttribute struct {
	base.Base
	Name           string                   `orm:"unique"`        //产品属性名称
	Code           string                   `orm:"default(\"\")"` //产品属性编码
	Sequence       int32                    //序列
	ValueIds       []*ProductAttributeValue `orm:"reverse(many)"` //属性值
	AttributeLines []*ProductAttributeLine  `orm:"reverse(many)"`
}

//列出记录
func ListProductAttribute(condArr map[string]interface{}, start, length int64) (utils.Paginator, []ProductAttribute, error) {

	o := orm.NewOrm()

	o.Using("default")
	qs := o.QueryTable(new(ProductAttribute))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()

	var (
		productAttributes []ProductAttribute
		num               int64
		err               error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(start, length, cnt)
	}

	if num, err = qs.OrderBy("-id").Limit(length, start).All(&productAttributes); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, productAttributes, err
}

//添加属性
func AddProductAttribute(obj ProductAttribute, user base.User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	productAttribute := new(ProductAttribute)
	productAttribute.Name = obj.Name
	productAttribute.CreateUser = &user
	productAttribute.UpdateUser = &user
	productAttribute.Code = obj.Code
	productAttribute.Sequence = obj.Sequence
	id, err := o.Insert(productAttribute)
	return id, err
}

//获得某一个属性信息
func GetProductAttributeByID(id int64) (ProductAttribute, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		productAttribute ProductAttribute
		err              error
	)
	cond := orm.NewCondition()
	cond = cond.And("id", id)
	qs := o.QueryTable(new(ProductAttribute))
	qs = qs.RelatedSel()
	err = qs.One(&productAttribute)
	return productAttribute, err
}
func GetProductAttributeByName(name string) (ProductAttribute, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		obj ProductAttribute
		err error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(ProductAttribute))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&obj)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return obj, err
}
