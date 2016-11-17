package base

import (
	"pms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Base
	Name       string      `orm:"size(20)" xml:"name" form:"name"`                   //用户名
	NameZh     string      `orm:"size(20)"  form:"namezh"`                           //中文用户名
	Department *Department `orm:"rel(fk);null;"`                                     //部门
	Email      string      `orm:"size(20)" xml:"email" form:"email"`                 //邮箱
	Mobile     string      `orm:"size(20);default(\"\")" xml:"mobile" form:"mobile"` //手机号码
	Tel        string      `orm:"size(20);default(\"\")" form:"tel"`                 //固定号码
	Password   string      `xml:"password"`                                          //密码
	Group      []*Group    `orm:"rel(m2m);rel_table(user_groups)"`                   //权限组
	IsAdmin    bool        `orm:"default(false)" xml:"isAdmin"`                      //是否为超级用户
	Active     bool        `orm:"default(true)" xml:"active"`                        //有效
	Qq         string      `orm:"default(\"\")" xml:"qq" form:"qq"`                  //QQ
	WeChat     string      `orm:"default(\"\")" xml:"wechat" form:"wechat"`          //微信
	Position   *Position   `orm:"rel(fk);null;"`                                     //职位

}

//多字段唯一
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Email", "Mobile"},
	}
}

// 多字段索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Email", "Mobile"},
	}
}
func (u *User) TableName() string {
	return "auth_user"
}
func ListUser(condArr map[string]interface{}, user User, page, offset int64) (utils.Paginator, error, []User) {

	if page < 1 {
		page = 1
	}

	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(User))
	// qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if active, ok := condArr["active"]; ok {
		cond = cond.And("active", active)

	} else {
		cond = cond.And("active", true)
	}
	if departmentId, ok := condArr["departmentId"]; ok {
		cond = cond.And("department__id", departmentId)
	}
	var (
		users []User
		num   int64
		err   error
	)
	var paginator utils.Paginator

	//后面再考虑查看权限的问题
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(page, offset, cnt)
	}
	start := (page - 1) * offset
	if num, err = qs.Limit(offset, start).All(&users); err == nil {
		paginator.CurrentPageSize = num
	}

	return paginator, err, users
}

//添加用户
func AddUser(obj User, cUser User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Name = obj.Name
	user.Email = obj.Email
	user.Mobile = obj.Mobile
	user.Tel = obj.Tel
	user.IsAdmin = obj.IsAdmin
	user.Active = obj.Active
	user.CreateUser = &cUser
	user.UpdateUser = &cUser
	user.Password = utils.PasswordMD5(obj.Password, obj.Mobile)

	id, err := o.Insert(user)
	return id, err
}

//获得某一个用户信息
func GetUserById(id int64) (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Base: Base{Id: id}}
	err := o.Read(&user)
	return user, err
}
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	err := qs.One(&user)
	return user, err
}
func CheckUserByName(name, password string) (User, error, bool) {
	o := orm.NewOrm()
	var (
		user User
		err  error
		ok   bool
	)
	ok = false
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {

		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			ok = true
		}
	}
	return user, err, ok

}
func UserNameExsit(name string) bool {
	var (
		user User
	)
	o := orm.NewOrm()
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("name", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err := qs.One(&user); err == nil {
		return true
	}
	return false
}
func UpdateUser(user User) (int64, error) {
	o := orm.NewOrm()
	return o.Update(&user)
}
