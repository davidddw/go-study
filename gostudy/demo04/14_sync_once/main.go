package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"sync"
)

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon get icon
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcon(imagePath string) image.Image {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	fmt.Println("load", imagePath)
	return img
}

func main() {
	Icon("left")
	Icon("right")
}
