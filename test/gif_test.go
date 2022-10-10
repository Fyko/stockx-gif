package gifutil

import (
	"testing"

	"github.com/Fyko/stockx-gif/internal/gifutil"
)

var urls = []string{
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img01.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img02.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img03.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img04.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img05.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img06.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img07.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img08.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img09.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img10.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img11.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img12.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img13.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img14.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img15.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img16.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img17.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img18.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img19.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img20.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img21.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img22.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img23.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img24.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img25.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img26.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img27.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img28.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img29.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img30.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img31.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img32.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img33.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img34.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img35.jpg?w=1280&fm=jpg",
	"https://stockx-360.imgix.net/adidas-Yeezy-Boost-350-V2-Synth/Images/adidas-Yeezy-Boost-350-V2-Synth/Lv2/img36.jpg?w=1280&fm=jpg",
}

func BenchmarkWriteGIF(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gifutil.WriteGIF(urls)
	}
}
