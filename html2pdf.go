package doctron

//NeedDoctronUpload HTML2PdfRequestDTO NeedDoctronUpload
func (r *HTML2PdfRequestDTO) NeedDoctronUpload() bool {
	return r.UploadKey != ""
}

// Paper orientation. Defaults to false.
const DefaultLandscape = false

// Display header and footer. Defaults to false.
const DefaultDisplayHeaderFooter = false

// Print background graphics. Defaults to true.
const DefaultPrintBackground = true

// Scale of the webpage rendering. Defaults to 1.
const DefaultScale = 1

// Paper width in inches. Defaults to 8.5 inches.
const DefaultPaperWidth = 8.5

// Paper height in inches. Defaults to 11 inches.
const DefaultPaperHeight = 11

// Top margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginTop = 0.4

// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginBottom = 0.4

// Left margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginLeft = 0.4

// Right margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginRight = 0.4

// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
const DefaultPageRanges = ""

// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
const DefaultIgnoreInvalidPageRanges = false

// HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
const DefaultHeaderTemplate = ""

// HTML template for the print footer. Should use the same format as the headerTemplate.
const DefaultFooterTemplate = ""

// Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
const DefaultPreferCSSPageSize = false

//DefaultWaitingTime Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
const DefaultWaitingTime = 0

//HTML2PdfRequestDTO HTML2PdfRequestDTO
type HTML2PdfRequestDTO struct {
	CommonRequestDTO
	Landscape               bool    `url:"landscape" validate:"omitempty"`                // Paper orientation. core.Defaults to false.
	DisplayHeaderFooter     bool    `url:"displayHeaderFooter" validate:"omitempty"`      // Display header and footer. core.Defaults to false.
	PrintBackground         bool    `url:"printBackground" validate:"omitempty"`          // Print background graphics. core.Defaults to false.
	Scale                   float64 `url:"scale" validate:"omitempty"`                    // Scale of the webpage rendering. core.Defaults to 1.
	PaperWidth              float64 `url:"paperWidth" validate:"gt=0"`                    // Paper width in inches. core.Defaults to 8.5 inches.
	PaperHeight             float64 `url:"paperHeight" validate:"gt=0"`                   // Paper height in inches. core.Defaults to 11 inches.
	MarginTop               float64 `url:"marginTop" validate:"gte=0"`                    // Top margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginBottom            float64 `url:"marginBottom" validate:"gte=0"`                 // Bottom margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginLeft              float64 `url:"marginLeft" validate:"gte=0"`                   // Left margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginRight             float64 `url:"marginRight" validate:"gte=0"`                  // Right margin in inches. core.Defaults to 1cm (~0.4 inches).
	PageRanges              string  `url:"pageRanges" validate:"omitempty"`               // Paper ranges to print, e.g., '1-5, 8, 11-13'. core.Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges bool    `url:"ignoreInvalidPageRanges" validate:"omitempty"`  // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. core.Defaults to false.
	HeaderTemplate          string  `url:"headerTemplate" validate:"omitempty"`           // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
	FooterTemplate          string  `url:"footerTemplate" validate:"omitempty"`           // HTML template for the print footer. Should use the same format as the headerTemplate.
	PreferCSSPageSize       bool    `url:"preferCSSPageSize" validate:"omitempty"`        // Whether or not to prefer page size as defined by css. core.Defaults to false, in which case the content will be scaled to fit the paper size.
	WaitingTime             int     `schema:"waitingTime,omitempty" validate:"omitempty"` // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
}

//NewDefaultHTML2PdfRequestDTO NewDefaultHTML2PdfRequestDTO
func NewDefaultHTML2PdfRequestDTO() *HTML2PdfRequestDTO {
	return &HTML2PdfRequestDTO{
		Landscape:               DefaultLandscape,
		DisplayHeaderFooter:     DefaultDisplayHeaderFooter,
		PrintBackground:         DefaultPrintBackground,
		Scale:                   DefaultScale,
		PaperWidth:              DefaultPaperWidth,
		PaperHeight:             DefaultPaperHeight,
		MarginTop:               DefaultMarginTop,
		MarginBottom:            DefaultMarginBottom,
		MarginLeft:              DefaultMarginLeft,
		MarginRight:             DefaultMarginRight,
		PageRanges:              DefaultPageRanges,
		IgnoreInvalidPageRanges: DefaultIgnoreInvalidPageRanges,
		HeaderTemplate:          DefaultHeaderTemplate,
		FooterTemplate:          DefaultFooterTemplate,
		PreferCSSPageSize:       DefaultPreferCSSPageSize,
		WaitingTime:             DefaultWaitingTime,
	}
}
