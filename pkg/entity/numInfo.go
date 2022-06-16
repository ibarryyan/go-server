package entity

import "encoding/json"

type NumInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	InfoKey string `json:"info_key"`
	InfoNum int64  `json:"info_num"`
}

func (stu NumInfo) TableName() string {
	return "num_info"
}

func (info NumInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":       info.Id,
		"name":     info.Name,
		"info_key": info.InfoKey,
		"info_num": info.InfoNum,
	})
}
