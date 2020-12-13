package doctron

//NeedDoctronUpload PdfWatermarkRequestDTO NeedDoctronUpload
func (r *PdfWatermarkRequestDTO) NeedDoctronUpload() bool {
	return r.UploadKey != ""
}

//PdfWatermarkRequestDTO PdfWatermarkRequestDTO
type PdfWatermarkRequestDTO struct {
	CommonRequestDTO
	WatermarkType int    `url:"watermarkType" validate:"omitempty"` // watermark type will support soon
	ImageURL      string `url:"imageUrl" validate:"required,url"`   // watermark image url
}

//NewDefaultPdfWatermarkRequestDTO NewDefaultPdfWatermarkRequestDTO
func NewDefaultPdfWatermarkRequestDTO() *PdfWatermarkRequestDTO {
	return &PdfWatermarkRequestDTO{
		WatermarkType: 0,
		ImageURL:      "",
	}
}
