# services
SDK服务

## 项目初始化

1. 请安装task项目管理工具

```
go get -u -v github.com/go-task/task/cmd/task
```

2. 初始化项目

```
task init
```

3. 安装依赖

```
task deps
```
其他依赖请把依赖添加到`scripts/deps.sh中`

4. Goland作为开发工具配置
请把当前目录`pwd`配置为`Project GOPATH`



## API说明文档
# 原本链SDK介绍-golang版本
> 这个版本的SDK用来给go语言开发者提供便捷生成metadata的服务。

**NOTE** 原本链中所有字节数组都以16进制的字符串存储，公钥为压缩格式。

## git路径

```
https://git.dev.yuanben.org/scm/unv/universe-go-sdk.git
```

### 服务方法分布
 > go-sdk提供三个文件用来生成metadata相关参数：app/headler.key.go、app/headler.dtcp.go、app/headler.node.go


```
1.  headler.key.go
    这个文件主要提供对公私钥对生成、签名方法，开发者通过这个文件中的方法可以便捷的获取原本链所应用格式的公私钥。更多哈希函数的使用请参考util/crypto.go。
2.  headler.dtcp.go
    这个文件主要提供metadata各项参数的生成方法，开发者可以通过这个文件的方法来对metadata中对各项参数进行补全，生成可被原本链node节点接受的metadata。
3.  headler.node.go
    这个文件主要提供对node节点进行请求对方法，开发者可以通过这个文件中的方法来访问node接口。
```


### metadata介绍

| name           | type    | comment                                  |source|
| -------------- | ------- | ---------------------------------------- |------|
| type           | string  | 类型, image(图片类型),article(文档类型) , others(其他类型) |用户传入|
| language       | string  | 语言 'zh-CN',                              |默认zh-cn,可用户传入|
| title          | string  | 内容标题                                       |用户传入|
| signature      | string  | 内容签名, 算法(secp256k1)                      |系统生成|
| abstract       | string  | 描述,内容摘要                                       |用户传入，为空时，系统自动取内容的前200个字符|
| category       | string  | 分类集, 以逗号分隔 "新闻, 商业"                      |用户传入，如果有传入content，则系统会追加五个|
| dna            | string  | metadata dna                             |系统生成|
| parent_dna            | string  | 该metadata修改前的dna                             |用户传入，如果时修改前一个metadata的数据，则需要传入前一个metadata的dna|
| block_hash            | string  | 区块链上的一个block_hash值                             |用户传入，会到链上做校验|
| block_height            | long  | block_hash对应的区块的height值                             |用户传入，会到链上做校验|
| created        | integer | 创建的时间,时间戳,10位长度, 1506302092               |系统生成|
| content_hash   | string  | 内容哈希,hash算法(keccak256)                   |可用户传入，如果没有，系统根据content生成|
| extra          | object  | 扩展内容,自定义内容                               |用户传入|
| license        | object  | 许可证                                      |用户传入|
| license.type   | string  | 许可证类型                                    |用户传入|
| license.parameters | object  | 许可证参数对象,自定义内容                            |用户传入|
| source         | string  | 原内容的链接, article,image的官网或者内容的连接                |用户传入|
| data           | object  | 根据type存放相关的article，image的下面定义的内容           |用户传入|


### GenPrivKeySecp256k1
> 生成公私密钥对

```golang
// GenPrivKeySecp256k1 生成一对公私密钥
// return:16进制编码对私钥，公钥
func GenPrivKeySecp256k1() (string, string) {
	return uts.GenPrivKeySecp256k1()
}

// 该方法位于headler.key.go,使用secp256k1曲线，生成16进制对公私钥，其中公钥为压缩形式，公钥字符串长度为66，私钥的字符串长度为64。
```

### Sign
> 签名函数

```golang
// Sign  签名函数
// return 签字字节数组
func Sign(hash, prv []byte) (sig []byte, err error) {
	return uts.Sign(hash, prv)
}
// 该方法位于headler.key.go,需要传入需要签名的内容的hash(keccak256)，prv为私钥字节数组。hex.DecodeString(prv_str)可以将私钥字符串转为字节数组。
```


### GenContentHash
> 生成content_hash

```golang
func GenContentHash(content string) string {
	return uts.Hasher([]byte(content))
}

// 该方法位于headler.dtcp.go,需要传入content，用于对content进行keccak256哈希运算。
```

### GenMetadataSignature
> 生成metadata的签名。该签名是metadata中的Signature

```golang
func GenMetadataSignature(private_key string, md *kts.Metadata) (string, error) {
	if private_key == "" || md == nil {
		return "", errors.New("there must be a private key and license")
	}
	prvBs, _ := hex.DecodeString(private_key)

	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := uts.Sign(h, prvBs); err != nil {
		return "", err
	} else {
		return hex.EncodeToString(signBs), nil
	}
}

// 该方法位于headler.dtcp.go,需要传入16进制的私钥和metadata，输出metadata的签名。
// md.DumpsRmSignSort()返回去除metadata中的Content、Signature和DNA，然后对字段对名称进行hash升序排列，再转成json，方法源码见：types/metadata.go

```

### VerifySignature
> 验证metadata的签名。该签名是metadata中的Signature

```golang
func VerifySignature(md *kts.Metadata) (bool, error) {
	if md == nil || md.PubKey == "" {
		return false, errors.New("public key is empty or metadata is nil")
	}
	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := hex.DecodeString(md.Signature); err != nil {
		return false, err
	} else {
		d1, _ := hex.DecodeString(md.PubKey)
		return crypto.VerifySignature(d1, h, signBs[:len(signBs)-1]), nil 
	}
}

// 该方法位于headler.dtcp.go,需要传入metadata，输出metadata的签名验证结果。
```


### GenerateDNA
> 验证metadata的DNA

```golang
//GenerateDNA
func GenerateDNA(md_sign string) string {
	return uts.GenerateDNA(md_sign)
}

// 该方法位于headler.dtcp.go,需要传入metadata的签名，输出metadata的DNA。
```

### FullMetadata
> 对metadata进行信息补全。

**NOTE** metadata中BlockHash、Type、Title、License为必填字段，如果contentHash为空，则content的值不能为空；如果类型不是article，则需要传入contentHash的值，该方法会对这些属性做非空判断。

```golang
//FullMetadata
func FullMetadata(private_key string, md *kts.Metadata) (err error) {
	......
}

// 该方法位于headler.dtcp.go,需要传入16进制编码对私钥和metadata。
// Abstract：如果为空，则截取content的前200个字符做摘要
// Category：如果类型是文章，系统会在用户写入的基础上，在添加5个类别，类别使用jieba分词抽取

```

### QueryMetadata
> 通过DNA向node节点查询metadata

```golang
func QueryMetadata(url string, version string, dna string) (res *kts.MetadataQueryResp) {
	......
}

// 该方法位于headler.node.go,需要传入node节点的url、node节点的版本（不传默认为v1)以及metadata的dna。该方法会返回查询到的metadata，code为error表示错误。

```

### SaveMetadata
> 向node节点提交metadata，请求注册

```golang
func SaveMetadata(url string, version string, async bool, md *kts.Metadata) (res *kts.MetadataSaveResp){
	......
}

// 该方法位于headler.node.go,需要传入node节点的url、node节点的版本（不传默认为v1)、async以及metadata。该方法会返回在原本链上注册的metadata的DNA，code为error表示错误。
// async=true为异步发送,如果async=false为同步发送，选择异步发送node节点会将请求存储队列。
```

### QueryLicense
> 根据license的type向node节点查询license

```golang
func QueryLicense(url string, version string, license_type string) (res *kts.LicenseQueryResp) {
	......
}
// 该方法位于headler.node.go,需要传入node节点的url、node节点的版本（不传默认为v1)以及license_type。该方法会返回对应的license信息，code为error表示错误。
```

### QueryLatestBlockHash
> 向node节点查询最新的blockHash

```golang
func QueryLatestBlockHash(url string,version string) (res *kts.BlockHashQueryResp) {
	......
}
// 该方法位于headler.node.go,需要传入node节点的url、node节点的版本（不传默认为v1)。该方法会返回原本链最新的区块hash，code为error表示错误。
```

### CheckBlockHash
> 向node节点查询blockHash是否在链上，并且高度正确

```golang
func CheckBlockHash(url string,version string,req *kts.BlockHashCheckReq) (res *kts.BlockHashCheckResp) {
	......
}
// 该方法位于headler.node.go,需要传入node节点的url、node节点的版本（不传默认为v1)以及要检查的blockHash和blockHeight。该方法会返回bool类型的查询结果(BlockHashCheckResp.Data)，code为error表示错误。
```

**NOTE** 方法的使用请查看test包下的[样例](https://git.dev.yuanben.org/projects/UNV/repos/universe-go-sdk/browse/test)
