package model

var (
	fontCache = map[string]*PdfFont{}
)

func CacheTtfFont(fontname string, filepath string) (err error) {
	font, err := NewCompositePdfFontFromTTFFile(filepath)
	if err != nil {
		return
	}
	fontCache[fontname] = font
	return
}

func FontInCache(fontname string) bool {
	if fontCache[fontname] != nil {
		return true
	}
	return false
}

func GetCacheFont(fontname string) (*PdfFont, bool) {
	font, err := fontCache[fontname]
	return font, err
}
