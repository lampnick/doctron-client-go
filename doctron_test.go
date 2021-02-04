package doctron

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const domain = "http://47.52.25.206:8080"
const defaultUsername = "doctron"
const defaultPassword = "lampnick"

func TestHTML2Pdf(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultHTML2PdfRequestDTO()
	req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
	req.WaitingTime = 1000
	response, err := client.HTML2Pdf(req)
	if err != nil {
		t.Fatalf("HTML2Pdf err: %s", err.Error())
	}
	t.Logf("%d", len(response.Data))
	fn := "./test/HTML2Pdf.pdf"
	os.Remove(fn)
	ioutil.WriteFile(fn, response.Data, os.ModePerm)
}

func TestHTML2PdfAndUpload(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultHTML2PdfRequestDTO()
	req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
	req.UploadKey = "test.pdf"
	response, err := client.HTML2PdfAndUpload(req)
	if err != nil {
		t.Fatalf("HTML2PdfAndUpload err: %s", err.Error())
	}
	t.Logf("UploadedURL: %s", response.UploadedURL)
}

func TestHTML2Image(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultHTML2ImageRequestDTO()
	req.WaitingTime = 1000
	req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
	response, err := client.HTML2Image(req)
	if err != nil {
		t.Fatalf("HTML2Image err: %s", err.Error())
	}
	t.Logf("%d", len(response.Data))
	fn := "./test/HTML2Image.png"
	os.Remove(fn)
	ioutil.WriteFile(fn, response.Data, os.ModePerm)
}

func TestHTML2ImageAndUpload(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultHTML2ImageRequestDTO()
	req.ConvertURL = "http://doctron.lampnick.com/doctron.html"
	req.UploadKey = "test.png"
	response, err := client.HTML2ImageAndUpload(req)
	if err != nil {
		t.Fatalf("HTML2ImageAndUpload err: %s", err.Error())
	}
	t.Logf("UploadedURL: %s", response.UploadedURL)
}

func TestPdfAddWatermark(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultPdfWatermarkRequestDTO()
	req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
	req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
	response, err := client.PdfAddWatermark(req)
	if err != nil {
		t.Fatalf("PdfAddWatermark err: %s", err.Error())
	}
	t.Logf("%d", len(response.Data))
	fn := "./test/PdfAddWatermark.pdf"
	os.Remove(fn)
	ioutil.WriteFile(fn, response.Data, os.ModePerm)
}

func TestPdfAddWatermarkAndUpload(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed: %v", time.Since(start))
	}()
	client := NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := NewDefaultPdfWatermarkRequestDTO()
	req.ConvertURL = "https://qjhdqx-prod.oss-cn-zhangjiakou.aliyuncs.com/test.pdf"
	req.ImageURL = "https://www.baidu.com/img/flexible/logo/pc/result.png"
	req.UploadKey = "PdfAddWatermarkAndUpload.pdf"
	response, err := client.PdfAddWatermarkAndUpload(req)
	if err != nil {
		t.Fatalf("PdfAddWatermarkAndUpload err: %s", err.Error())
	}
	t.Logf("UploadedURL: %s", response.UploadedURL)
}
