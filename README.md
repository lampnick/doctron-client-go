# doctron-client-go
### install
```
go get -u github.com/lampnick/doctron-client-go
```
### html2pdf
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultHTML2PdfRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
response, err := client.HTML2Pdf(req)
...
```
### html2pdf and upload
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultHTML2PdfRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
req.UploadKey = "test.pdf" // add this
response, err := client.HTML2PdfAndUpload(req)
```
### html2image
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultHTML2ImageRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
response, err := client.HTML2Image(req)
```
### html2image and upload
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultHTML2ImageRequestDTO()
req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
req.UploadKey = "test.png"
response, err := client.HTML2ImageAndUpload(req)
```
### pdf add watermark
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultPdfWatermarkRequestDTO()
req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
response, err := client.PdfAddWatermark(req)
```
### pdf add watermark and upload
```
client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
req := NewDefaultPdfWatermarkRequestDTO()
req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
req.UploadKey = "PdfAddWatermarkAndUpload.pdf"
response, err := client.PdfAddWatermarkAndUpload(req)
```