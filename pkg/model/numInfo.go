package model

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

//Redis类似序列化操作
func (info NumInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(info)
}

func (info NumInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &info)
}
