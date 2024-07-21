package sqsio

import (
	"fmt"
	"strings"

	"github.com/atbuy/squisher/lib"
)

func ToBin(data ...int) string {
	out := []string{}

	for _, value := range data {
		binValue := fmt.Sprintf("%08b", value)
		out = append(out, binValue)
	}

	return strings.Join(out, "")
}

func ImageToBin(sqsImage lib.SquishImage) string {
	out := []string{}

	for y := 0; y < sqsImage.Height; y++ {
		for x := 0; x < sqsImage.Width; x++ {
			pixel := sqsImage.GetPixel(x, y)

			out = append(out, fmt.Sprintf("%08b", pixel.Red))
			out = append(out, fmt.Sprintf("%08b", pixel.Green))
			out = append(out, fmt.Sprintf("%08b", pixel.Blue))
			out = append(out, fmt.Sprintf("%08b", pixel.Alpha))
		}
	}

	return strings.Join(out, "")
}

func StrToBytes(data string) []byte {
	out := []byte{}

	var now byte
	for _, v := range data {
		now = now<<1 + byte(v-'0')
		out = append(out, now)
	}

	return out
}

func BinToBytes(bins string) []byte {
	length := len(bins)/8 + 1
	byteData := make([]byte, length)

	count, i := 0, 0
	var now byte
	for _, v := range bins {
		if count == 8 {
			byteData[i] = now
			i++
			now, count = 0, 0
		}

		now = now<<1 + byte(v-'0')
		count++
	}

	if count != 0 {
		byteData[i] = now << (8 - byte(count))
		i++
	}

	byteData = byteData[:i:i]
	return byteData
}

func JoinStrBin(bins ...string) string {
	out := []string{}
	for _, value := range bins {
		out = append(out, value)
	}

	return strings.Join(out, "")
}
