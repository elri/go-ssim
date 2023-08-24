package main

import (
	"fmt"

	"github.com/elri/go-ssim"
	"github.com/elri/go-ssim/util"
)

func main() {
	img := ssim.ConvertToGray(ssim.ReadImage("images/testImage0.jpg"))
	img2 := ssim.ConvertToGray(ssim.ReadImage("images/testImage3.jpg"))

	c, err := ssim.Covar(img, img2)
	util.HandleError(err)

	index := ssim.CalculateSSIM(img, img2)

	// fmt.Printf("AVG   %f\n", mean(img))
	// fmt.Printf("STDEV %f\n", stdev(img))
	// fmt.Printf("COV   %f\n", c)
	_ = c
	fmt.Printf("\nSSIM = %f\n", index)
}
