package until

import (
	"fmt"
	"os"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	// 设置上传到的空间
	bucket       = ""
	host         = ""
	callbackbody = `{"key":$(key), "hash":$(etag),"filesize":$(fsize)}`
)

// 构造返回值字段
type PutRet struct {
	Hash     string `json:"hash"`
	Key      string `json:"key"`
	Filesize int    `json:"filesize"`
}

func UpQiniu(path, notify, key string) {
	conf.ACCESS_KEY = ""//七牛配置ACCESS_KEY
	conf.SECRET_KEY = ""//七牛配置SECRET_KEY

	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + key,
		//设置Token过期时间
		Expires:      3600,
		CallbackUrl:  notify + "?path=" + host + key,
		CallbackBody: callbackbody,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	// 设置上传文件的路径
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, path, nil)
	// 打印返回的信息
	fmt.Println("upload success", key)
	os.Remove(path)
	// 打印出错信息
	if res != nil {
		ErrorNotify(notify)
		fmt.Println("io.Put failed:", res)
		return
	}

}
