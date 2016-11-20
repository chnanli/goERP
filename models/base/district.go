package base

import (
	"fmt"
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type District struct {
	Base
	Name string //区县名称
	City *City  `orm:"rel(fk)"` //城市
}

//添加区县
func AddDistrict(obj District, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := new(District)
	district.Name = obj.Name
	district.City = obj.City
	district.CreateUser = &user
	district.UpdateUser = &user
	id, err := o.Insert(district)
	return id, err
}

//获得某一个区县信息
func GetDistrictByID(id int64) (District, error) {
	o := orm.NewOrm()
	o.Using("default")
	district := District{Base: Base{Id: id}}
	err := o.Read(&district)
	return district, err
}
func GetDistrictByName(name string) (District, error) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		district District
		err      error
	)
	cond := orm.NewCondition()
	qs := o.QueryTable(new(District))

	if name != "" {
		cond = cond.And("name", name)
		qs = qs.SetCond(cond)
		qs = qs.RelatedSel()
		err = qs.One(&district)
	} else {
		err = fmt.Errorf("%s", "查询条件不成立")
	}

	return district, err
}

//列出记录
func ListDistrict(condArr map[string]interface{}, page, offset int64) (utils.Paginator, []District, error) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(District))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if name, ok := condArr["name"]; ok {
		cond = cond.And("name_icontains", name)
	}
	var (
		districts []District
		num       int64
		err       error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}

	start := (page - 1) * offset
	if num, err = qs.OrderBy("-id").Limit(offset, start).All(&districts); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, districts, err
}
