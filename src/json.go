package main

import "encoding/json"
import "fmt"

//import "os"

type SReq struct {
	id   int
	reqs []string
}

type SReqex struct {
	id   int      `json: "id"` //注意不是单引号
	reqs []string `json:"reqs"`
}

func main() {

	bolB, _ := json.Marshal(true)

	fmt.Println(string(bolB)) // 转化为字符串打印，上面的Marshal将true转化字符串"true"

	intB, _ := json.Marshal(30)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(43.6)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("hello")
	fmt.Println(string(strB))

	slcD := []string{"docker", "cdn", "vedio"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"docker": 10, "cdn": 8, "video": 9}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	reqD := &SReq{
		id:   100,
		reqs: []string{"cdn", "docker", "video"}}

	reqB, _ := json.Marshal(reqD) // 为空的原因
	fmt.Println(string(reqB))
	fmt.Println(reqD)

	reqexD := &SReqex{
		id:   100,
		reqs: []string{"cdn", "docker", "video"}}

	reqexB, _ := json.Marshal(reqexD) // 为空的原因
	fmt.Println(string(reqexB))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//   map[string]interface{} 将保存一个 string 为键，值为任意值的map
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	str2 := strs[1].(string)
	fmt.Println(str2)

	strjson := `{"id": 95, "reqs": ["live","cdn"]}`
	reqex := &SReqex{}
	json.Unmarshal([]byte(strjson), &reqex)
	fmt.Println(reqex)
	fmt.Println(reqex.reqs[0])

}
