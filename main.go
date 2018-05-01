package main

import
(
	"os"
	"image/jpeg"
	"image"
	"image/draw"
	"image/png"
	"log"
	"encoding/json"
	"fmt"
	"time"
)





func main() {
	j := `
	[
		{"name":"Z_Head_01.png", "X":0, "Y":0},
		{"name":"Chinese-style-hat_01.png", "X":1, "Y":1},
		{"name":"Z_Nose_01.png", "X":88, "Y":88},
		{"name":"Z_HairCut_06_1.png", "X":5, "Y":5},
		{"name":"Z_Mouth_03.png", "X":100, "Y":300},
		{"name":"Z_Glasses_01.png", "X":99, "Y":99},
		{"name":"Z_Eyes_01_L.png", "X":150, "Y":300},
		{"name":"Z_Eyes_01_R.png", "X":200, "Y":50}
	]`

	var imgs []Imgs
	err := json.Unmarshal([]byte(j), &imgs)
	if err != nil {
		panic(err)
	}

	t1 := time.Now() // get current time

	for i := 0; i <= 1; i++ {
		mix(imgs,i)
	}

	elapsed := time.Since(t1)
	fmt.Println("时长: ", elapsed)
}

type Imgs struct {
	Name string
	X int
	Y int
}

func mix(img []Imgs, xx int) {
	father_img,err := os.Open(img[0].Name)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
		return
	}

	father, err := png.Decode(father_img)
	if err != nil {
		log.Fatalf("failed to decode1: %s", err)
	}

	defer father_img.Close()


	b := father.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, father, image.ZP, draw.Src)
	for i := 1; i < len(img); i++ {
		fmt.Println(i)
		son_image,err := os.Open(img[i].Name)
		if err != nil {
			log.Fatalf("failed to open: %s", err)
		}
		son,err := png.Decode(son_image)
		if err != nil {
			log.Fatalf("failed to decode: %s", err)
		}
		defer son_image.Close()
		offset := image.Pt(img[i].X, img[i].Y)
		draw.Draw(image3, son.Bounds().Add(offset), son, image.ZP, draw.Over)
	}

	final,err := os.Create(fmt.Sprintf("result_%d.jpg",xx))
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(final, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer final.Close()
}