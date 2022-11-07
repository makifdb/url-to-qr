package converter

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/makifdb/url-to-qr/color"
	"github.com/skip2/go-qrcode"
)

func CreateQRCode(URL, BGColor, FGColor string) (string, error) {

	// create temp file
	file, err := os.CreateTemp("", "qr-*")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	// parse the bg color
	bgColor, err := color.ParseColor(BGColor)
	if err != nil {
		return "", err
	}

	// parse the fg color
	fgColor, err := color.ParseColor(FGColor)
	if err != nil {
		return "", err
	}

	// generate a QR code for the url
	err = qrcode.WriteColorFile(URL, qrcode.Medium, 512, bgColor, fgColor, file.Name())
	if err != nil {
		return "", err
	}

	// read the file
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	// convert to base64
	b64 := base64.StdEncoding.EncodeToString(data)

	// convert to data url
	image := strings.Join([]string{"data:image/png;base64,", b64}, "")

	return image, nil
}
