package app_test

import (
	"encoding/json"
	"github.com/yuanbenio/universe-go-sdk/app"
	kts "github.com/yuanbenio/universe-go-sdk/types"
	"testing"
)

func getProcessor() *app.NodeProcessor {
	return app.InitNodeProcessor("https://testnet.yuanbenlian.com", app.DefaultChainVersion)
}

func TestNodeProcessor_QueryMetadata(t *testing.T) {
	res, err := getProcessor().QueryMetadata("5GIKRZ2E4Q2M9U2ZVEL3N6QK2IZKDHLAG6VAUA13ICM58TSPI3")
	if err != nil {
		t.Fatal(err)
	} else {
		if !res.Success() {
			t.Fatal(res.Msg)
		} else {
			t.Log(string(res.Data.Dumps()))
		}
	}
}

func TestNodeProcessor_SaveMetadata(t *testing.T) {
	md := &kts.Metadata{
		Content: "原本链是一个分布式的底层数据网络；" +
			"原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；" +
			"原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；" +
			"原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；" +
			"原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别",
		BlockHash:   "4D36473D2FF1FE0772A6C0C55D7911295D8E1E27",
		BlockHeight: "12345",
		Type:        "custom",
		Title:       "原本链测试",
		Category:    "测试,custom",
		License:     kts.NoneLicense,
	}
	priKey := "50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade"
	err := app.FullMetadata(priKey, md)
	if err != nil {
		t.Fatalf("full metadata fail:%s", err.Error())
	} else {
		res, err := getProcessor().SaveMetadata(md)
		if err != nil {
			t.Fatal(err)
		} else {
			if !res.Success() {
				t.Fatal(res.Msg)
			} else {
				t.Log(res.Data.Dna)
			}
		}
	}

}

func TestNodeProcessor_QueryLicense(t *testing.T) {
	res, err := getProcessor().QueryLicense("cc", "4.0")
	if err != nil {
		t.Fatal(err)
	} else {
		if !res.Success() {
			t.Fatal(res.Msg)
		} else {
			_d, err := json.Marshal(res.Data)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(_d))
		}
	}
}

func TestNodeProcessor_QueryLatestBlockHash(t *testing.T) {
	res, err := getProcessor().QueryLatestBlockHash()
	if err != nil {
		t.Fatal(err)
	} else {
		if !res.Success() {
			t.Fatal(res.Msg)
		} else {
			_d, err := json.Marshal(res.Data)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(_d))
		}
	}
}

func TestNodeProcessor_CheckBlockHash(t *testing.T) {
	req := &kts.BlockHashCheckReq{
		Hash:   "325D652724600F3748391C54AC9A1A9D1D5C671D",
		Height: 33333,
	}

	res, err := getProcessor().CheckBlockHash(req)
	if err != nil {
		t.Fatal(err)
	} else {
		if !res.Success() {
			t.Fatal(res.Msg)
		} else {
			_d, err := json.Marshal(res.Data)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(_d))
		}
	}
}

func TestNodeProcessor_RegisterAccount(t *testing.T) {
	subKeys := make([]string, 0)

	for i := 0; i < 5; i++ {
		_, pubKey := app.GenPrivKeySecp256k1()
		subKeys = append(subKeys, pubKey)
	}

	req, err := app.GenRegisterAccountReq(pri_key, subKeys)
	if err != nil {
		t.Fatal(err)
	}

	res, err := getProcessor().RegisterAccount(req)
	if err != nil {
		t.Fatal(err)
	} else {
		if !res.Success() {
			t.Fatal(res.Msg)
		}
		t.Log("register success~")
	}
}
