package gojson

import (
	"encoding/json"
)

type myJson struct {
	Data interface{}
}

func new_myJson(data interface{}) *myJson {
	return &myJson{
		Data: data,
	}
}

func (m myJson) Get(keys ...string) *myJson {
	// 多级map索引
	vv, ok := m.Data.(map[string]interface{})
	if ok {
		var result interface{}
		for _, key := range keys {
			result = vv[key]
			vv, _ = result.(map[string]interface{})
		}
		// if result == nil {
		// 	return new_myJson(nil)
		// }
		return new_myJson(result)
	}
	panic("Data is Not Like map .")
}

func (m myJson) Default(data interface{}) interface{} {
	// 默认值取值
	if m.Data == nil {
		m.Data = data
	}
	return m.Data
}

func (m myJson) Index(index int) *myJson {
	vv, ok := m.Data.([]interface{})
	if ok {
		result := vv[index]
		return new_myJson(result)
	}
	panic("Data is Not Like slice.")
}

func Loads(data *string) *myJson {
	/*
		_, ok := v.(int)
		fmt.Println(ok)

		switch v.(type) {
		case string:
		case interface{}:
			fmt.Println(v)
		default:
		}
	*/
	var my_json_data interface{}
	my_json := new_myJson(my_json_data)
	err := json.Unmarshal([]byte(*data), &my_json.Data)
	if err != nil {
		panic(err)
	}

	return my_json
}

func Dumps(data interface{}) string {
	k, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(k)
}

/*
func main() {
	data := `{"code":200,"data":{"ext":{"utm_medium":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase"},"size":1,"items":[{"ext":null,"resourceId":"","mediaAssetInfo":null,"productId":"golang make","reportData":{"eventClick":true,"data":{"mod":"popu_895","extra":"{\"utm_medium\":\"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase\",\"dist_request_id\":\"1646296542475_57712\",\"hotword\":\"golang make\"}","dist_request_id":"1646296542475_57712","ab_strategy":"default","index":"1","strategy":"alirecmd"},"urlParams":{"utm_medium":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase","depth_1-utm_source":"distribute.pc_search_hot_word.none-task-hot_word-alirecmd-1.nonecase"},"eventView":true},"recommendType":"ali","index":1,"style":"word_1","strategyId":"alirecmd","productType":"hot_word"}]},"message":"success"}`
	d := Loads(&data)
	m := d.Get("data", "items").Index(0).Get("recommendType")
	fmt.Println(m.Data)
	fmt.Println(m.Default(1))

	m_data := map[int]int{1: 1, 2: 211}
	data = Dumps(m_data)
	fmt.Println(data)
}
*/
