package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"project/app"
	kts "project/types"
	uts "project/utils"
	"encoding/json"
	"net/http"
)

var url = "http://119.23.22.129:9000"

func getContent() (content string) {
	if conBs, err := ioutil.ReadFile("test1"); err != nil {
		panic(err)
	} else {
		content = string(conBs)
	}
	return
}
func test1() {
	priKey := "b133fb0fa361a292d37df2f5ac13ea64ba734a6c6319f03ded565bff0dd2c6c3"

	var content string
	if conBs, err := ioutil.ReadFile("test1"); err != nil {
		panic(err)
	} else {
		content = string(conBs)
	}

	priBs, _ := hex.DecodeString(priKey)

	md := &kts.Metadata{
		Content: content,
		Type:    "article",
		PubKey:  uts.GetPubKeyFromPri(priBs),
		License: struct {
			Type   string `json:"type,omitempty" binding:"required"`
			Params map[string]string `json:"params,omitempty" binding:"required"`
		}{Type: "cc", Params: map[string]string{
			"y":"4",
			"b":"2",
		}},
	}

	if err := app.GenerateMetadataFromContent(priKey, md); err != nil {
		fmt.Println(err.Error())
	}

	if isPass, err := app.VerifySignature(md); err != nil {
		fmt.Println("验证失败", err.Error())
	} else {
		fmt.Println("验证结果", isPass)
	}


	//metadata post test
	res := app.Metadata_post(url,"",md)
	if r,err := json.Marshal(res); err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(string(r))
	}



	//jBs,err := json.Marshal(md)
	//if err != nil {
	//	fmt.Println(fmt.Sprint("结构体转json错误:",err.Error()))
	//}
	//
	//fmt.Println(string(jBs))
}
func test2() {
	/*sign_str := "22798cd599538604acf4940d32dce402f344ced39bfc6a322238d7093ac42a1621567631ea4debdb85c68e69d3c29dbda4d482dbe4381b91699132afd6ceff6600"
	priKey := "b133fb0fa361a292d37df2f5ac13ea64ba734a6c6319f03ded565bff0dd2c6c3"

	fmt.Println(uts.VerifySignature())*/

	fmt.Println(fmt.Sprintf("%s/%s/metadata", url, "v1"))
}
func test3() {
	if resp,err := http.Get(fmt.Sprintf("%s/ping",url)); err != nil {
		fmt.Println("error",err.Error())
	}else {
		d, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(d))
	}
}

func test4 () {
	dna := "2ed3fde80b5868a00eb31dd2bb0debcafeef9f6f"

	res := app.Metadata_get(url,"",dna)
	j,_ := json.Marshal(res)
	fmt.Println(string(j))
}

func test5 (){
	uts.GenPrivKeySecp256k1()
}

func main() {

	test5()
}

