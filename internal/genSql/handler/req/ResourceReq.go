package req

type InsertReq struct {
	Type     string `json:"type" vd:"in($,'J','N','C','M','DDDD')"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId" vd:"$>1000"`
	Path     string `json:"path"`
	Resource string `json:"resource"`
	Button   string `json:"button"`
	Visible  int8   `json:"visible"`
}

type person struct {
	Age  int    `path:"age" json:"age"`    // 从路径中获取参数
	Name string `query:"name" json:"name"` // 从query中获取参数
	City string `json:"city"`              // 从body中获取参数
}
