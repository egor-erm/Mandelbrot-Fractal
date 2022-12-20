package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var (
	name string
)

const (
	WIDTH  = 1000
	HEIGHT = 1000
	MAX    = 1000
)

func main() {
	b := image.Rect(0, 0, WIDTH, HEIGHT)
	img := image.NewRGBA(b)

	flag.StringVar(&name, "f", "fractal.png", "Имя файла.")
	flag.Parse()

	for x := 0; x < WIDTH; x++ { // рекурсия по x
		for y := 0; y < HEIGHT; y++ { // рекурсия по y
			zx := 2.0                                  // значение сдвига x
			zy := 2.0                                  // значение сдвига y
			xf := float64(x)/WIDTH*zx - (zx/2.0 + 0.5) // десятичное значение x по ширине
			yf := float64(y)/HEIGHT*zy - (zy / 2.0)    // десятичное значение y по высоте
			c := complex(xf, yf)                       // заносим как комплексное число

			main_color := int(mandel(c) * 255) //получаем цвет пикселя

			color_value := color.RGBA{uint8(main_color), uint8(2 * main_color % 255), uint8(3 * main_color % 255), 255}
			if main_color == 255 {
				color_value = color.RGBA{0, 100, 100, 255} // Black
			}
			img.Set(x, y, color_value) // ставим пиксель
		}

	}

	file, err := os.Create(name)
	defer file.Close()

	if err != nil || file == nil {
		file, err = os.Open(name)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", err)
			return
		}
	}

	err = png.Encode(file, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding image: %s\n", err)
		return
	}
}

func mandel(c complex128) float64 {
	z := complex(0, 0)
	for i := 0; i < MAX; i++ { // max значение влияет на
		if cmplx.Abs(z) > 2 {
			return float64(i-1) / MAX
		}
		z = z*z + c // формула из проекта
	}

	return 1
}
