package main

import "golang.org/x/tour/pic"

/*
 * Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
 * When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

 * The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
 */

func Pic(dx, dy int) [][]uint8 {
	var p [][]uint8
	for j:=0; j<dy; j++ {
		s := make([]uint8, dx)
		for i:=0; i<dx ;i++  {
			s[i] = uint8((i+j)/2)
		}
		p = append(p, s)
	}
	return p
}

func main() {
	pic.Show(Pic)
}
