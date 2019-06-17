package texture

import (
	"github.com/emily33901/vtf"
)

// Texture2D is a generic GPU material struct
type Texture2D struct {
	filePath string
	width    int
	height   int
	vtf      *vtf.Vtf
}

// FilePath Get the filepath this data was loaded from
func (tex *Texture2D) FilePath() string {
	return tex.filePath
}

// Width returns materials width
func (tex *Texture2D) Width() int {
	return tex.width
}

// Height returns materials height
func (tex *Texture2D) Height() int {
	return tex.height
}

// Format returns this materials colour format
func (tex *Texture2D) Format() uint32 {
	return tex.vtf.Header().HighResImageFormat
}

// PixelDataForFrame get raw colour data for this frame
func (tex *Texture2D) PixelDataForFrame(frame int) []byte {
	return tex.vtf.HighestResolutionImageForFrame(frame)
}

// Thumbnail returns a small thumbnail image of a material
func (tex *Texture2D) Thumbnail() []byte {
	return tex.vtf.LowResImageData()
}

// Since a Texture2D is always in memory if it is being used
// These shouldnt do anything

func (tex *Texture2D) Reload() error {
	// logger.Warn("You cannot Reload() a Texture2D")
	return nil
}
func (tex *Texture2D) EvictFromMainMemory() {
	// logger.Warn("You cannot Evict() a Texture2D")
}

// NewTexture2D returns a new texture from Vtf
func NewTexture2D(filePath string, vtf *vtf.Vtf, width int, height int) *Texture2D {
	// @TODO: we should be able to load the vtf all by ourselves!
	return &Texture2D{
		filePath: filePath,
		width:    width,
		height:   height,
		vtf:      vtf,
	}
}
