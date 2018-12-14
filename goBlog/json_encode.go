/*
	根据map序列化json
*/
package lightGoJson

import (
	"encoding/json"
	"fmt"
	"log"
)

type LightJsonMap map[string]interface{}

func LightEncode(elememt interface{}) string {
	//拼接的结果字符串
	var s string
	//若为对象，则拼接字符串
	if LightJson, err := elememt.(LightJsonMap); !err {
		s = string("{")
		for key, val := range LightJson {
			s += fmt.Sprintf("\"%s\":\"%s\",", key, LightEncode(val))
		}
		s += string("}")
	} else {
		jsonStr, err := json.Marshal(elememt)
		if err != nil {
			log.Fatal("Can't transform jsonString,Because ", err)
		}
		s = string(jsonStr)
	}
	return s
}
