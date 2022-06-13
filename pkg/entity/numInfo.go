package entity

type NumInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	InfoKey string `json:"info_key"`
	InfoNum int64  `json:"info_num"`
}

func (stu NumInfo) TableName() string {
	return "num_info"
}
