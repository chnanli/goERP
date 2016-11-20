package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Company struct {
	Base
	Name       string        `orm:"unique"`        //公司名称
	Children   []*Company    `orm:"reverse(many)"` //子公司
	Parent     *Company      `orm:"rel(fk);null"`  //上级公司
	Department []*Department `orm:"reverse(many)"` //部门
	Country    *Country      `orm:"rel(fk);null"`  //国家
	Province   *Province     `orm:"rel(fk);null"`  //身份
	City       *City         `orm:"rel(fk);null"`  //城市
	District   *District     `orm:"rel(fk);null"`  //区县
	Street     string        `orm:"default(\"\")"` //街道
}

//添加公司
func AddCompany(obj Company, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	company := new(Company)
	company.Name = obj.Name
	company.CreateUser = &user
	company.UpdateUser = &user
	company.Province = obj.Province
	id, err := o.Insert(company)
	return id, err
}

//获得某一个公司信息
func GetCompanyByID(id int64) (Company, error) {
	o := orm.NewOrm()
	o.Using("default")
	company := Company{Base: Base{Id: id}}
	err := o.Read(&company)
	return company, err
}

//根据名称查询公司
func GetCompanyByName(name string) (Company, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		company Company
		err     error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(Company))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&company)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return company, err
}

//列出记录
func ListCompany(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []Company, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Company))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
	}
	var (
		companys []Company
		num      int64
		err      error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&companys); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, companys, err
}
