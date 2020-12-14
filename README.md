# doctron-client-go
### install
```
go get -u github.com/lampnick/doctron-client-go
```
### full demo 
```
package main

import (
	"context"
	"log"
	"github.com/lampnick/doctron-client-go"
)

const domain = "http://47.52.25.206:8080"
const defaultUsername = "doctron"
const defaultPassword = "lampnick"

func main() {
	client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := doctron.NewDefaultHTML2PdfRequestDTO()
	req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
	response, err := client.HTML2Pdf(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(response.Data))
}

```
### html2pdf
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultHTML2PdfRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
response, err := client.HTML2Pdf(req)
...
```
### html2pdf and upload
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultHTML2PdfRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
req.UploadKey = "test.pdf" // add this
response, err := client.HTML2PdfAndUpload(req)
```
### html2image
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultHTML2ImageRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
response, err := client.HTML2Image(req)
```
### html2image and upload
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultHTML2ImageRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
req.UploadKey = "test.png"
response, err := client.HTML2ImageAndUpload(req)
```
### pdf add watermark
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultPdfWatermarkRequestDTO()
req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
response, err := client.PdfAddWatermark(req)
```
### pdf add watermark and upload
```
client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := doctron.NewDefaultPdfWatermarkRequestDTO()
req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
req.UploadKey = "PdfAddWatermarkAndUpload.pdf"
response, err := client.PdfAddWatermarkAndUpload(req)
```