package ssim

import (
	"errors"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

// Default SSIM constants
var (
	L  = 255.0
	K1 = 0.01
	K2 = 0.03
	C1 = math.Pow((K1 * L), 2.0)
	C2 = math.Pow((K2 * L), 2.0)
)

/*
Calculate the similarity between two given images

Note! The paths must (probably) be absolute, cannot contain '~'
*/
func CalculateSSIM(imgFile1, imgFile2 string) (float64, error) {
	var err error
	var img, img2 image.Image
	index := -1.0

	//Read images
	img, err = readImage(imgFile1)
	if err == nil {

		img2, err = readImage(imgFile2)
		if err == nil {

			// Convert to grayscale
			img = convertToGray(img)
			img2 = convertToGray(img2)

			// Calculate SIM
			index, err = ssim(img, img2)
		}
	}

	return index, err

}

// SSIM algorithm implementation, see https://en.wikipedia.org/wiki/Structural_similarity
func ssim(x, y image.Image) (float64, error) {
	var index float64

	avg_x := mean(x)
	avg_y := mean(y)

	stdev_x := stdev(x)
	stdev_y := stdev(y)

	cov, err := covar(x, y)
	if err == nil {
		numerator := ((2.0 * avg_x * avg_y) + C1) * ((2.0 * cov) + C2)
		denominator := (math.Pow(avg_x, 2.0) + math.Pow(avg_y, 2.0) + C1) *
			(math.Pow(stdev_x, 2.0) + math.Pow(stdev_y, 2.0) + C2)

		index = numerator / denominator
	}

	return index, err
}

// Calculate the covariance of 2 images
func covar(img1, img2 image.Image) (c float64, err error) {
	if !equalDim(img1, img2) {
		err = errors.New("images must have same dimension")
		return
	}
	avg1 := mean(img1)
	avg2 := mean(img2)
	w, h := dim(img1)
	sum := 0.0
	n := float64((w * h) - 1)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pix1 := getPixVal(img1.At(x, y))
			pix2 := getPixVal(img2.At(x, y))
			sum += (pix1 - avg1) * (pix2 - avg2)
		}
	}
	c = sum / n
	return
}

// Given a path to an image file, read and return as
// an image.Image
func readImage(fname string) (image.Image, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

// Convert an Image to grayscale which
// equalize RGB values
func convertToGray(originalImg image.Image) image.Image {
	bounds := originalImg.Bounds()
	w, h := dim(originalImg)

	grayImg := image.NewGray(bounds)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			originalColor := originalImg.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	return grayImg
}

// Convert uint32 R value to a float. The returnng
// float will have a range of 0-255
func getPixVal(c color.Color) float64 {
	r, _, _, _ := c.RGBA()
	return float64(r >> 8)
}

// Helper function that return the dimension of an image
func dim(img image.Image) (w, h int) {
	w, h = img.Bounds().Max.X, img.Bounds().Max.Y
	return
}

// Check if two images have the same dimension
func equalDim(img1, img2 image.Image) bool {
	w1, h1 := dim(img1)
	w2, h2 := dim(img2)
	return (w1 == w2) && (h1 == h2)
}

// Given an Image, calculate the mean of its
// pixel values
func mean(img image.Image) float64 {
	w, h := dim(img)
	n := float64((w * h) - 1)
	sum := 0.0

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			sum += getPixVal(img.At(x, y))
		}
	}
	return sum / n
}

// Compute the standard deviation with pixel values of Image
func stdev(img image.Image) float64 {
	w, h := dim(img)

	n := float64((w * h) - 1)
	sum := 0.0
	avg := mean(img)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pix := getPixVal(img.At(x, y))
			sum += math.Pow((pix - avg), 2.0)
		}
	}
	return math.Sqrt(sum / n)
}
