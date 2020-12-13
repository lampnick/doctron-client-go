package doctron

// Image compression format (defaults to png).
const FormatPng = "png"
const FormatJpeg = "jpeg"

// Compression quality from range [0..100] (jpeg only).
const DefaultQuality = 0

// Capture the screenshot from the surface, rather than the view. Defaults to true.
const DefaultFromSurface = true

// Capture the screenshot of a given region only.
// X offset in device independent pixels (dip).
const DefaultViewportX = 0

// Y offset in device independent pixels (dip).
const DefaultViewportY = 0

// Rectangle width in device independent pixels (dip).
const DefaultViewportWidth = 996

// Rectangle height in device independent pixels (dip).
const DefaultViewportHeight = 996

// Page scale factor.
const DefaultViewportScale = 1

//NeedDoctronUpload HTML2ImageRequestDTO NeedDoctronUpload
func (r *HTML2ImageRequestDTO) NeedDoctronUpload() bool {
	return r.UploadKey != ""
}

//HTML2ImageRequestDTO HTML2ImageRequestDTO
type HTML2ImageRequestDTO struct {
	CommonRequestDTO
	Format      string  `url:"format" validate:"omitempty"`      // Image compression format (defaults to png).
	Quality     int64   `url:"quality" validate:"omitempty"`     // Compression quality from range [0..100] (jpeg only).
	CustomClip  bool    `url:"customClip" validate:"omitempty"`  //if set this value, the below clip will work,otherwise not work!
	ClipX       float64 `url:"clipX" validate:"omitempty"`       // Capture the screenshot of a given region only.X offset in device independent pixels (dip).
	ClipY       float64 `url:"clipY" validate:"omitempty"`       // Capture the screenshot of a given region only.Y offset in device independent pixels (dip).
	ClipWidth   float64 `url:"clipWidth" validate:"omitempty"`   // Capture the screenshot of a given region only.Rectangle width in device independent pixels (dip).
	ClipHeight  float64 `url:"clipHeight" validate:"omitempty"`  // Capture the screenshot of a given region only.Rectangle height in device independent pixels (dip).
	ClipScale   float64 `url:"clipScale" validate:"omitempty"`   // Capture the screenshot of a given region only.Page scale factor.
	FromSurface bool    `url:"fromSurface" validate:"omitempty"` // Capture the screenshot from the surface, rather than the view. Defaults to true.
}

//NewDefaultHTML2ImageRequestDTO NewDefaultHTML2ImageRequestDTO
func NewDefaultHTML2ImageRequestDTO() *HTML2ImageRequestDTO {
	return &HTML2ImageRequestDTO{
		Format:      FormatPng,
		Quality:     DefaultQuality,
		CustomClip:  false,
		ClipX:       DefaultViewportX,
		ClipY:       DefaultViewportY,
		ClipWidth:   DefaultViewportWidth,
		ClipHeight:  DefaultViewportHeight,
		ClipScale:   DefaultViewportScale,
		FromSurface: DefaultFromSurface,
	}
}
