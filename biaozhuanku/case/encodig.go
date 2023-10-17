package _case

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// 编码
func EncodingCase() {
	type user struct {
		ID   int64
		Name string
		Age  uint8
	}
	u := user{ID: 1, Name: "nick", Age: 18}
	//序列化

	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)
	fmt.Println(string(bytes))
	u1 := user{}
	//反序列化
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)

	//base64
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1)

	//16进制编码
	str1 := hex.EncodeToString(bytes1)
	fmt.Println(str1)
	bytes2, _ := hex.DecodeString(str1)
	fmt.Println(bytes2)
}
