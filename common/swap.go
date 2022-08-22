package common

import (
	"encoding/json"
)

//通过json tag 进行结构体赋值
func SwapTo(request, category interface{}) (err error)  {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, category)
	return
}