package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.jpg")
	if err != nil {
		log.Fatalf("os.Open() failed with %s\n", err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode() failed with %s\n", err)
	}

	f, err = os.Create("jpg_to_png.png")
	if err != nil {
		log.Fatalf("os.Create() failed with %s\n", err)
	}

	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode() failed with %s\n", err)
	}
	fmt.Println("Image converted Successfully!")

}
