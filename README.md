# services
Yuanben Chain SDK for Gopher

![Yuanben chain](https://github.com/yuanbenio/universe-go-sdk/blob/master/img/yuanbenlian.png)

## Data flow diagram
![Data flow](https://github.com/yuanbenio/universe-go-sdk/blob/master/img/data-flow.png)

## return code

|code|describe|
|---|---|
|ok|Success|
|3001|Invalid parameter|
|3002|Empty parameter|
|3003|Record does not exist|
|3004|Record already exists|
|3005|Permission denied: please register the public key first|
|3006|Permission denied: please contract Yuanben chain support|
|3007|Incorrect data|
|3009|Data storage fail|
|3010|Data not on YuanBen chain |
|3011|block information is empty|
|3012|Query fail|
|3020|Signature verification fail|
|3023|Parameter verification fail|
|3021|Invalid public key|
|3022|License's parameters are empty|
|4001|Error connecting to redis server|
|4002|Error connecting to the first-level node|
|4003|Broadcast transaction fail|
|4004|ABCI query fail|
|4005|Redis handling error|
|5000|Unknown error|


## API Document

> TestNet address: https://testnet.yuanbenlian.com

### Service Introduction

The Golang-SDK provides three processors: app/key.go, app/dtcp.go and app/node.go.

```text
1. app/key.go
    This is a base service, it supports: generating key pair, calculating signatures, verifying signatures, etc
2. app/dtcp.go
    Process metadata
3. app/node.go
    To access the YuanBen Chain Node, supports: query and saving metadata, query licenses, query latest BlockHash、registering sub accounts
```

### Metadata Introduction
| name           | type    |must| comment              |source|
| -------------- | ------- | ----|---------------------|------|
| type           | string  | Y |eg:image,article,audio,video,custom,private |user-defined|
| language       | string  | Y |'zh-CN',                                    |user-defined,default:zh-CN|
| title          | string  | N |title                                       |user-defined|
| signature      | string  | Y |sign by secp256k1                           |generate by system|
| abstract       | string  | N |Content summary                             |default:content[:200],user-defined|
| category       | string  | N |eg:"news"                                   |user-defined，if there is content, the system will add five more|
| dna            | string  | Y |metadata dna                                |generate by system|
| parent_dna     | string  | N |link an other metadata|user-defined|
| block_hash     | string  | Y |block_hash on YuanBen chain                 |user-defined|
| block_height   | string  | Y |block_hash corresponding block_height       |user-defined|
| created        | integer | Y |timestamp, eg:1506302092                    |generate by system|
| content_hash   | string  | Y |Keccak256(content)                          |user-defined. default:Keccak256(content)|
| extra          | TreeMap<String, Object>  | N | user-defined content       |user-defined|
| license        | Metadata.License  | Y |                                 |user-defined|
| license.type   | string  | Y |the type of license                         |user-defined|
| license.parameters | TreeMap<String, Object>  | N | the parameters of license   |user-defined|
| source         | string  | N |source link.                                |user-defined|
| data           | TreeMap<String, Object>  | N |extension data of the type |user-defined|
| id         | string  | N |business id    |user-defined|
| pubkey         | string  | N |public key     |generate by private key|

### API Interface

> app/key.go

#### GeneratorSecp256k1Key
###### Response
|attr| type | info|
|---|---|---|
|privateKey | string| hex encode
|publicKey | string | hex encode

---
#### Sign
###### Request
|attr| type | info|
|---|---|---|
|hash |[]byte | 256 bits
|prv | []byte | private key by hex decode


###### Response
|attr| type | info|
|---|---|---|
|signature | []byte| hex bytes


> app/dtcp.go

#### GenContentHash
###### Request
|attr| type | info|
|---|---|---|
|content | string| source content


###### Response
|attr| type | info|
|---|---|---|
|contentHash | string |  keccak256(content)


---
#### GenMetadataSignature
###### Request
|attr| type | info|
|---|---|---|
|privateKey | string | hex string|
|md | Metadata | metadata |


###### Response
|attr| type | info|
|---|---|---|
|signature | string | matadata's signature |


###### Errors
|attr | info|
|---|---|
|ErrPrivateKeyNil| private key is nil|
|ErrMDNil | metadata is nil|
|others | invalid private key, sign failure|

---
#### VerifySignature
###### Request
|attr| type | info|
|---|---|---|
|md | Metadata | metadata includes all attrs|


###### Response
|attr| type | info|
|---|---|---|
|result | bool| varify result |


###### Errors
|attr | info|
|---|---|
|ErrPublicKeyNil| public key is nil|
|ErrMDNil | metadata is nil|
|others | invalid signature, invalid public key|

---
#### GenerateDNA
###### Request
|attr| type | info|
|---|---|---|
| mdSign | string | metadata's signature|

###### Response
|attr| type | info|
|---|---|---|
|dna | string | matadata's dna|

---
#### FullMetadata
###### Request
|attr| type | info|
|---|---|---|
|privateKey | string | hex string|
|md | Metadata | metadata includes all attrs|


| name           | type    |must| comment              |source|
| -------------- | ------- | ----|---------------------|------|
| type           | string  | Y |eg:image,article,audio,video,custom,private |user-defined|
| title          | string  | Y |content title       |user-defined,private can be empty|
| category       | string  | N |eg:"news,article"    |user-defined,private can be empty|
| block_hash     | string  | Y |block_hash on YuanBen chain                 |user-defined|
| block_height   | string  | Y |block_hash corresponding block_height       |user-defined|
| content   | string  | N |content                        |user-defined,content and content_hash can't be empty at same time|
| content_hash   | string  | N |Keccak256(content)                          |user-defined. default:Keccak256(content),content and content_hash can't be empty at same time|
| data           | TreeMap<String, Object>  | Y |extension data of the type |user-defined,private\custom\article can be nil|
| license        | Metadata.License  | Y |                                 |user-defined|
| license.type   | string  | Y |the type of license                         |user-defined|
| license.parameters | TreeMap<String, Object>  | Y | the parameters of license   |user-defined,none can be nil|
| created        | integer | N |timestamp, eg:1506302092                    |generate by system|
| language       | string  | N |'zh-CN',          |user-defined,default:zh-CN|
| parent_dna     | string  | N |link an other metadata|user-defined|
| abstract       | string  | N |Content summary                             |default:content[:200],user-defined|
| source         | string  | N |source link.                                |user-defined|
| id         | string  | N |business id    |user-defined|
| pubkey         | string  | N |public key                               |generate by private key|
| extra          | TreeMap<String, Object>  | N | more information by user defined  |user-defined|
| signature      | string  | N |sign by secp256k1     |generate by system|
| dna            | string  | N |metadata dna                                |generate by system|


###### Response
|attr| type | info|
|---|---|---|
|err | error | error message|


###### Errors
|attr | info|
|---|---|
|ErrPrivateKeyNil | private key is nil|
|ErrMDNil | metadata is nil|
|ErrBlockInfoNil | block information is nil|
|ErrLicenseNil | invalid licensce|
|ErrContentNil | content and contentHash can't be nil at same time|
|ErrTypeNil | type is nil|
|ErrTypeInvalid | invalid type|
|ErrExtendDataNil | extends data can't be nil for spcical data|
|ErrTitleNil | title is nil. if type is private, it allow|
|ErrCategoryNil | category is nil. if type is private, it allow|
|others | invalid private key, sign failure|

---
#### GenRegisterAccountReq
###### Request
|attr| type | info|
|---|---|---|
|privateKey | string | hex string|
|subKeys | []string | sub accounts. public key == account|


###### Response
|attr| type | info|
|---|---|---|
|req | RegisterAccountReq | rqeuest parameters|
|err | error | error message|


###### Errors
|attr | info|
|---|---|
|ErrPrivateKeyNil| private key is nil|
|ErrSubKeysNil | subKeys is nil|
|others | sign failure, invalid private key|


> app/node.go

#### InitNodeProcessor
###### Request

|attr| type | info|
|---|---|---|
|url | string | Yuanben Chain node URL|
|chainVersion | string | API version ,default: v1|


###### Response
|attr| type | info|
|---|---|---|
|processor | NodeProcessor |node processor instance|

---
#### QueryMetadata
> query the metadata by dna from Yuanben Chain node

###### Request
|attr| type | info|
|---|---|---|
|dna | string | metadata's dna


###### Response
|attr| type | info|
|---|---|---|
|res | MetadataQueryResp | response data, include error log|
|err | error | http failure message|

###### Errors
|attr | info|
|---|---|
|ErrDNAInvalid| DNA is nil or invalid|

---
#### SaveMetadata
> save metadata to Yuanben Chain node

###### Request
|attr| type | info|
|---|---|---|
|md | Metadata | metadata

| name           | type    |must| comment              |
| -------------- | ------- | ----|---------------------|
| content_hash   | string  | Y |content                        |
| created        | integer | Y |timestamp, eg:1506302092       |
| license        | Metadata.License  | Y | license information|
| license.type   | string  | Y |the type of license     |
| license.parameters | TreeMap<String, Object>  | Y | the parameters of license |
| type           | string  | Y |eg:image,article,audio,video,custom,private |
| block_hash     | string  | Y |block_hash on YuanBen chain        |
| block_height   | string  | Y |block_hash corresponding block_height|
| pubkey         | string  | Y |public key   |
| signature      | string  | Y |sign by secp256k1  |
| language       | string  | Y |'zh-CN'|
| dna            | string  | N |metadata dna       |
| title          | string  | N |content title       |
| category       | string  | N |eg:"news,article"    |
| data           | TreeMap<String, Object>  | N |extension data of the type |
| parent_dna     | string  | N |link an other metadata|
| abstract       | string  | N |Content summary  |
| source         | string  | N |source link.    |
| id         | string  | N |business id    |
| extra          | TreeMap<String, Object>  | N | more information by user defined  |


###### Response
|attr| type | info|
|---|---|---|
|res | MetadataSaveResp | response data, include DNA or error log|
|err | error | http failure message|

###### Errors
|attr | info|
|---|---|
|ErrMDNil| metadata is nil|
|ErrSignNil| signature is nil|
|ErrLicenseNil| license is nil|
|others | invalid metadata|

---
#### QueryLicense
> query license to Yuanben Chain node

###### Request
|attr| type | info|
|---|---|---|
|licenseType | string | 
|licenseVersion | string | 


###### Response
|attr| type | info|
|---|---|---|
|res | LicenseQueryResp | response data, include error log|
|err | error | http failure message|


###### Errors
|attr | info|
|---|---|
|ErrLicenseNil| licenseType is empty|

---
#### QueryLatestBlockHash
> query the lasted block information from the Yuaben Chain node

###### Response
|attr| type | info|
|---|---|---|
|res | LicenseQueryResp | response data, include error log|
|err | error | http failure message|

---
#### CheckBlockHash
> check whether block is on the Yuanben Chain

###### Request
|attr| type | info|
|---|---|---|
|hash | string | block hash| 
|height | int64 | block height|


###### Response
|attr| type | info|
|---|---|---|
|res | BlockHashCheckResp | response data, include error log|
|err | error | http failure message|


###### Errors
|attr | info|
|---|---|
|ErrParameters| block hash or block height is empty|


---
#### RegisterAccount
> register sub accounts to special YuanBen chain node<br>
> For nodes that have authentication enabled, if you need to store data, you must registered

###### Request
|attr| type | info|
|---|---|---|
|req | RegisterAccountReq | see dtcp.go ---> GenRegisterAccountReq| 


###### Response
|attr| type | info|
|---|---|---|
|res | RegisterAccountResp | response data, include error log|
|err | error | http failure message|


###### Errors
|attr | info|
|---|---|
|ErrParameters| subkeys is nil|

