package main

import "fmt"

func main() {
	img := ssim.convertToGray(readImage("images/testImage0.jpg"))
	img2 := ssim.convertToGray(readImage("images/testImage3.jpg"))

	c, err := ssim.covar(img, img2)
	util.handleError(err)

	index := ssim.calculateSSIM(img, img2)

	// fmt.Printf("AVG   %f\n", mean(img))
	// fmt.Printf("STDEV %f\n", stdev(img))
	// fmt.Printf("COV   %f\n", c)
	_ = c
	fmt.Printf("\nSSIM = %f\n", index)
}
