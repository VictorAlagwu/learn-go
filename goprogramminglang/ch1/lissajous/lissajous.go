package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 128, 0, 1}, color.RGBA{221, 74, 104, 1}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
	redIndex = 3
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8439", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	r.ParseForm()
	_, hasParam := r.Form["cycles"]
	var cycles float64;
	if hasParam {
		i, err := strconv.Atoi(r.FormValue("cycles"))
		if err != nil {
			cycles = 5
		} else {
			cycles = float64(i)
		}
	} else {
		cycles = 5
	}
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x:= math.Sin(t)
			y := math.Sin(t*freq + phase)
			if i % 2 == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			} else if i % 3 == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			}
			
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}