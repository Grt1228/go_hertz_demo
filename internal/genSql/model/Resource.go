package model

type Resource struct {
	BaseModel
	Id       int
	Code     string
	Name     string
	Type     string
	Path     string
	Resource string
	Button   string
	Icon     string
	Visible  int8
	ParentId int
	System   int
	OrderNo  int
}

func (Resource) TableName() string {
	return "common_resource"
}
