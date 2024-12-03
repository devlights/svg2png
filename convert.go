package svg2png

import (
	"image"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// Convert は、指定されたsvgファイルのデータを [image.Image] に変換します。
func Convert(svgFile string, size int, opacity float64) (image.Image, error) {
	var (
		svg *oksvg.SvgIcon
		err error
	)
	if svg, err = oksvg.ReadIcon(svgFile, oksvg.WarnErrorMode); err != nil {
		return nil, err
	}

	var (
		width   = int(svg.ViewBox.W)
		height  = int(svg.ViewBox.H)
		maxSize = max(width, height)
	)
	width = (width * size) / maxSize
	height = (height * size) / maxSize

	var (
		img     = image.NewRGBA(image.Rect(0, 0, width, height))
		scanner = rasterx.NewScannerGV(width, height, img, img.Bounds())
		raster  = rasterx.NewDasher(width, height, scanner)
	)
	svg.SetTarget(0, 0, float64(width), float64(height))
	svg.Draw(raster, opacity)

	return img, nil
}
