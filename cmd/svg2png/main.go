/*
svg2png

	svgファイルをpngファイルに変換します。

# REFERENCES
  - https://github.com/typomedia/rasterize/tree/master
*/
package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/devlights/svg2png"
)

type (
	Args struct {
		in      string
		out     string
		size    int
		opacity float64
		debug   bool
	}
)

var (
	args Args
)

func init() {
	flag.StringVar(&args.in, "in", "", "input file")
	flag.StringVar(&args.out, "out", "out.png", "output file")
	flag.IntVar(&args.size, "size", 1000, "size in pixels")
	flag.Float64Var(&args.opacity, "opacity", 1.0, "opacity")
	flag.BoolVar(&args.debug, "debug", false, "debug mode")
}

func main() {
	flag.Parse()

	if args.in == "" || args.size <= 0 || args.opacity < 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if args.out == "" {
		args.out = "out.png"
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		img image.Image
		err error
	)
	if img, err = svg2png.Convert(args.in, args.size, args.opacity); err != nil {
		return err
	}

	var (
		file *os.File
	)
	if file, err = os.Create(args.out); err != nil {
		return err
	}
	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
