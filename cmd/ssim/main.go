package main

import (
	"fmt"
	"log"

	"github.com/elri/go-ssim"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img := ssim.ConvertToGray(ssim.ReadImage("images/testImage0.jpg"))
	img2 := ssim.ConvertToGray(ssim.ReadImage("images/testImage3.jpg"))

	c, err := ssim.Covar(img, img2)
	handleError(err)

	index := ssim.CalculateSSIM(img, img2)

	// fmt.Printf("AVG   %f\n", mean(img))
	// fmt.Printf("STDEV %f\n", stdev(img))
	// fmt.Printf("COV   %f\n", c)
	_ = c
	fmt.Printf("\nSSIM = %f\n", index)
}
