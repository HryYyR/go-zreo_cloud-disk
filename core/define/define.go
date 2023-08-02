package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var JwtKey = "Hyyyh"

var EmailPassword = "ztgxgymchofgdjjj"

// 云存储对象
var OssaccessKeyID = "LTAI5tFVa3i1VKq5F9ZW5iDN"
var OssaccessSecret = "ixqzkmCNkp6pzigEil7CQ0BaGxsiDC"
var OssEndPoint = "oss-cn-chengdu.aliyuncs.com"
var OssPath = "zero-cloud-disk.oss-cn-chengdu.aliyuncs.com"

// 文件列表默认参数
var PageSize = 20

var DateTime = "2006-01-02 15:04:05"
