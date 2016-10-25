package base

type Department struct {
	Base
	Name    string  //团队名称
	Leader  *User   `orm:"rel(fk)"`       //团队领导者
	Members []*User `orm:"reverse(many)"` //成员
}
