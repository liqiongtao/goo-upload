# 上传到本地

```
s := goo.NewServer()
s.GET("/upload/local", func(c *gin.Context) {
    filename, _ := upload.Local.Upload(c)
    c.JSON(200, gin.H{"filename": filename})
})
s.Run(":18080")
```

# 上传到OSS

```
var (
	conf = upload.OSSConfig{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Bucket:          "",
		Domain:          "",
	}
)

func main() {
	body, _ := ioutil.ReadFile("1.txt")
	url, _ := upload.OSS.Upload("1.txt", body)
	fmt.Println(url)
}
```