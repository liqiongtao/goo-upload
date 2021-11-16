package goo_upload

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
	"io/ioutil"
	"os"
	"path"
)

// -----------------------------------------------
// 定义main文件，然后调用ossScript()
// go build -ldflags "-s -w" -o oss
// -----------------------------------------------

var (
	help = flag.Bool("h", false, "help")

	baseDir = flag.String("dir", "", "base dir")

	endpoint        = flag.String("endpoint", "", "endpoint")
	accessKeyId     = flag.String("access-key-id", "", "access key id")
	accessKeySecret = flag.String("access-key-secret", "", "access key secret")
	bucketName      = flag.String("bucket-name", "", "bucket name")
)

func InitScript() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *baseDir == "" {
		*baseDir = "goo"
	}

	// todo:: 配置一下参数
	if *endpoint == "" {
		*endpoint = ""
	}
	if *accessKeyId == "" {
		*accessKeyId = ""
	}
	if *accessKeySecret == "" {
		*accessKeySecret = ""
	}
	if *bucketName == "" {
		*bucketName = ""
	}
}

func main2() {
	if total := len(flag.Args()); total == 0 {
		fmt.Println("请选择上传文件", total, os.Args)
		return
	}

	client, err := oss.New(*endpoint, *accessKeyId, *accessKeySecret)
	if err != nil {
		fmt.Println("[oss-client]", err.Error())
		return
	}

	bucket, err := client.Bucket(*bucketName)
	if err != nil {
		fmt.Println("[oss-bucket]", err.Error())
		return
	}

	filename := flag.Args()[0]

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("[read-file]", err.Error())
		return
	}

	var md5str = goo_utils.MD5(body)
	filename = fmt.Sprintf("%s/%s/%s%s", md5str[0:2], md5str[2:4], md5str[8:24], path.Ext(filename))

	if *baseDir != "" {
		filename = *baseDir + "/" + filename
	}

	if err := bucket.PutObject(filename, bytes.NewReader(body)); err != nil {
		fmt.Println("[oss-upload]", err.Error())
		return
	}

	fmt.Println("https://" + *bucketName + "." + *endpoint + "/" + filename)
}
