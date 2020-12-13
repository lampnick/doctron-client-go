package doctron

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/lampnick/doctron-client-go/utils"
	"golang.org/x/net/context/ctxhttp"
	"gopkg.in/go-playground/validator.v9"
)

const separator = "/"
const (
	relativePathHTML2Pdf        = "/convert/html2pdf"
	relativePathHTML2Image      = "/convert/html2image"
	relativePathPdfAddWatermark = "/convert/pdfAddWatermark"
)

var debug utils.Debug

func init() {
	debug = utils.Init("doctron")
}

//Client Client
type Client struct {
	ctx      context.Context
	domain   string
	username string `url:"u" validate:"required"`
	password string `url:"p" validate:"required"`
	hc       *http.Client
}

//NewClient NewClient
func NewClient(ctx context.Context, domain, username, password string) *Client {
	return &Client{
		ctx:      ctx,
		domain:   domain,
		username: username,
		password: password,
		hc: &http.Client{
			Timeout: time.Second * 60,
		},
	}
}

//HTML2Pdf html convert to pdf,return pdf byte data
func (client *Client) HTML2Pdf(req *HTML2PdfRequestDTO) (*ConvertResponse, error) {
	return client.convert(relativePathHTML2Pdf, req)
}

//HTML2PdfAndUpload  html convert to pdf,return uploaded pdf url
func (client *Client) HTML2PdfAndUpload(req *HTML2PdfRequestDTO) (*UploadResponse, error) {
	return client.convertAndUpload(relativePathHTML2Pdf, req)
}

//HTML2Image html convert to image,return image byte data
func (client *Client) HTML2Image(req *HTML2ImageRequestDTO) (*ConvertResponse, error) {
	return client.convert(relativePathHTML2Image, req)
}

//HTML2ImageAndUpload  html convert to image,return uploaded image url
func (client *Client) HTML2ImageAndUpload(req *HTML2ImageRequestDTO) (*UploadResponse, error) {
	return client.convertAndUpload(relativePathHTML2Image, req)
}

//PdfAddWatermark pdf add watermark,return pdf byte data
func (client *Client) PdfAddWatermark(req *PdfWatermarkRequestDTO) (*ConvertResponse, error) {
	return client.convert(relativePathPdfAddWatermark, req)
}

//PdfAddWatermarkAndUpload   pdf add watermark,return uploaded pdf url
func (client *Client) PdfAddWatermarkAndUpload(req *PdfWatermarkRequestDTO) (*UploadResponse, error) {
	return client.convertAndUpload(relativePathPdfAddWatermark, req)
}

func (client *Client) getURIWithAuth(relativePath string) string {
	domain := strings.Trim(client.domain, separator)
	if !strings.HasPrefix(relativePath, separator) {
		relativePath = separator + relativePath
	}
	relativePath = strings.TrimRight(relativePath, separator)
	return fmt.Sprintf("%s%s?u=%s&p=%s", domain, relativePath, client.username, client.password)
}

//convert support Doctron convert to byte
func (client *Client) convert(relativePath string, req RequestI) (*ConvertResponse, error) {
	response := &ConvertResponse{}
	if req.NeedDoctronUpload() {
		return nil, errors.New("convert can't pass uploadKey")
	}
	body, err := client.get(relativePath, req)
	if err != nil {
		return response, err
	}
	response.Data = body
	return response, nil
}

//convertAndUpload support Doctron convert to byte and upload
func (client *Client) convertAndUpload(relativePath string, req RequestI) (*UploadResponse, error) {
	response := &UploadResponse{}
	if !req.NeedDoctronUpload() {
		return nil, errors.New("convert and upload must pass uploadKey")
	}
	body, err := client.get(relativePath, req)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.New("Client.convertAndUpload.json.Unmarshal err:" + err.Error())
	}
	return response, nil
}

func (client *Client) get(relativePath string, req RequestI) (body []byte, err error) {
	// validate
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, errors.New("params err:" + err.Error())
	}
	v, err := query.Values(req)
	if err != nil {
		return nil, errors.New("Client.get.query err:" + err.Error())
	}

	url := client.getURIWithAuth(relativePath) + "&" + v.Encode()
	debug("%s", url)
	resp, err := ctxhttp.Get(client.ctx, client.hc, url)
	if err != nil {
		return nil, errors.New("Client.get.http.Get err:" + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Client.get.http.StatusCode:" + strconv.Itoa(resp.StatusCode))
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Client.get.ioutil.ReadAll err:" + err.Error())
	}
	contentType := resp.Header.Get("Content-Type")
	//only support json,pdf format
	if !strings.Contains(contentType, "application/json") && !strings.Contains(contentType, "application/pdf") && !strings.Contains(contentType, "image") {
		return nil, errors.New("Client.get.contentType err:" + contentType)
	}
	//if json format, judge response code
	if strings.Contains(contentType, "application/json") {
		response := &CommonResponse{}
		err := json.Unmarshal(body, response)
		if err != nil {
			return nil, errors.New("Client.get.json.Unmarshal err:" + err.Error())
		}
		if response.Code != responseCodeOK {
			msg := fmt.Sprintf("Client.get.response.Code:%d,err message:%s", response.Code, response.Message)
			return nil, errors.New(msg)
		}
	}
	return

}
