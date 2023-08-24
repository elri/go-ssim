package main

import (
	"flag"
	"fmt"

	"github.com/elri/go-ssim"
)

var (
	imgFlag1 = flag.String("img1", "", "Image filepath #1")
	imgFlag2 = flag.String("img2", "", "Image filepath #2")
)

func main() {
	flag.Parse()

	if *imgFlag1 == "" || *imgFlag2 == "" {
		fmt.Println("MISSING ARGUMENTS...")
		flag.Usage()
		return
	}

	index, err := ssim.CalculateSSIM(*imgFlag1, *imgFlag2)
	if err == nil {
		fmt.Printf("SSIM = %f\n", index)
	} else {
		fmt.Println(err)
	}
}
