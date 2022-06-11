package entity

type NumInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Num  int64  `json:"num"`
}

func (stu NumInfo) TableName() string {
	return "num_info"
}
