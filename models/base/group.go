package base

import "github.com/astaxie/beego/orm"

type Group struct {
	Base
	Name    string  `orm:"unique" xml:"name"` //组名称
	Members []*User `orm:"reverse(many)"`     //组员
}

func AddGroup(obj Group, user User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	group := new(Group)
	group.Name = obj.Name
	group.CreateUser = &user
	group.UpdateUser = &user
	id, err := o.Insert(group)
	return id, err
}
