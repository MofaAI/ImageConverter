package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	format string
)

func init() {
	flag.StringVarP(&format, "format", "f", "jpg", "image format, eg: jpg/png/gif")
}

func convert(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	m, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	ext := strings.ToLower(filepath.Ext(name))
	newFileName := strings.TrimSuffix(name, ext) + "." + format
	newFile, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer newFile.Close()
	switch format {
	case "png":
		return png.Encode(newFile, m)
	case "jpg":
		return jpeg.Encode(newFile, m, nil)
	case "jpeg":
		return jpeg.Encode(newFile, m, nil)
	case "gif":
		return gif.Encode(newFile, m, nil)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := convert(flag.Arg(0)); err != nil {
		panic(err)
	}
}
