package doctron

const responseCodeOK = 0

//RequestI RequestI
type RequestI interface {
	NeedDoctronUpload() bool
}

//CommonRequestDTO CommonRequestDTO
type CommonRequestDTO struct {
	ConvertURL string `url:"url" validate:"required,url"`
	UploadKey  string `url:"uploadKey" validate:"omitempty"`
}

//CommonResponse CommonResponse
type CommonResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	UploadedURL string `json:"data"`
}

//ConvertResponse ConvertResponse
type ConvertResponse struct {
	Data []byte
}

//UploadResponse UploadResponse
type UploadResponse struct {
	CommonResponse
}
