package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, _ := qr.Encode("http://blog.csdn.net/wangshubo1989", qr.H, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create("qr2.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
