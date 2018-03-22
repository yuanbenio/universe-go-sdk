package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"net/http"
	"universe-go-sdk/app"
	kts "universe-go-sdk/types"
	uts "universe-go-sdk/utils"
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
	if conBs, err := ioutil.ReadFile("test2"); err != nil {
		panic(err)
	} else {
		content = string(conBs)
	}

	priBs, _ := hex.DecodeString(priKey)

	md := &kts.Metadata{
		BlockHash: "afsf",
		Content:   content,
		Title:     "dsfdf",
		Type:      "article",
		PubKey:    uts.GetPubKeyFromPri(priBs),
		License: struct {
			Type       string            `json:"type,omitempty" binding:"required"`
			Parameters map[string]string `json:"parameters,omitempty" binding:"required"`
		}{Type: "cc", Parameters: map[string]string{
			"y": "4",
			"b": "2",
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

	//fmt.Println(md.DumpsRmSignSort())
	fmt.Println(string(md.Dumps()))

	//metadata post test
	//res := app.Metadata_post(url,"",md)
	//if r,err := json.Marshal(res); err != nil {
	//	fmt.Println(err.Error())
	//}else {
	//	fmt.Println(string(r))
	//}

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
	if resp, err := http.Get(fmt.Sprintf("%s/ping", url)); err != nil {
		fmt.Println("error", err.Error())
	} else {
		d, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(d))
	}
}

func test4() {
	dna := "2ed3fde80b5868a00eb31dd2bb0debcafeef9f6f"

	res := app.Metadata_get(url, "", dna)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
}

func test5() {
	uts.GenPrivKeySecp256k1()
}

func test6() {
	privKey := "af7a3e2004b35f73e9bada891563d6522847b98e0edf7c12af6f40bf859da24c"
	pubKey := "0264d6d0b926a6382eb227f7b233bdaa98c4b51998322831cf197f3ebe760e096b"

	content := "有关所有现有选项的参考，其描述和默认值，可以参考默认配置ethereumj.conf（可以在库jar或源树中找到它ethereum-core/src/test/resources）。要覆盖需要的选项，可以使用以下方法之一："
	contentHash := "58449825e30c3c9865321d21912818ea25e879dfedefa0461518eb0c82bb6059"
	sign_msg := "010147bc3a3b896aeb9be06456e1a1f6dac582f242b97318ff12df607d8b2ce97f60728a18a6be7e3bd6f5faf53724aa9511fa1f7bfa6b6544dd489f2c7b764700"

	pi1, _ := hex.DecodeString(privKey)
	pp1 := uts.GetPubKeyFromPri(pi1)
	if pp1 == pubKey {
		fmt.Println("公钥推导成功")
	} else {
		fmt.Println("公钥推导失败")
	}

	h1 := []byte(content)
	h2 := uts.Hasher(h1)
	if h2 == contentHash {
		fmt.Println("keccak256计算成功")
	} else {
		fmt.Println("keccak256计算失败")
	}

	h3, _ := hex.DecodeString(h2)
	sBs, _ := uts.Sign(h3, pi1)
	if hex.EncodeToString(sBs) == sign_msg {
		fmt.Println("sign success")
	} else {
		fmt.Println("sign fail")
	}

	h4, _ := hex.DecodeString(contentHash)
	s1, _ := hex.DecodeString(sign_msg)

	pp2, _ := hex.DecodeString(pubKey)

	if crypto.VerifySignature(pp2, h4, s1[:len(s1)-1]) {
		fmt.Println("verify success")
	} else {
		fmt.Println("verify fail")
	}

}

func test7() {
	md := &kts.Metadata{
		Abstract: "sagdsfg",
		License: struct {
			Type       string            `json:"type,omitempty" binding:"required"`
			Parameters map[string]string `json:"parameters,omitempty" binding:"required"`
		}{Type: "cc", Parameters: map[string]string{"data": "afgad", "cata": "afgad", "adta": "afgad", "bcdta": "afgad"}},
		DNA: "afsdfdfa",
	}
	bs, _ := json.Marshal(md)
	fmt.Println(string(bs))

}

func test8() {

}

func main() {

	test7()
}
