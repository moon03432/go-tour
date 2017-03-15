package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

/*
 * type Image interface {
 *     ColorModel() color.Model
 *     Bounds() Rectangle
 *     At(x, y int) color.Color
 * }
 *
 * Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
 * ColorModel should return color.RGBAModel.
 * At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
 */

type Image struct{
	w,h int
}

func (ima Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (ima Image) Bounds() image.Rectangle {
	return image.Rect(0,0,ima.w,ima.h)
}

func (ima Image) At(x,y int) color.Color {
	v := uint8(x^y)
	return color.RGBA{v,v,255,255}
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}
