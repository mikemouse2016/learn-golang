package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	//var png []byte
	//png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	qrcode.WriteFile("testing!", qrcode.Medium, 256, "qr.png")
}
