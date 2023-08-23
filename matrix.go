package main

import (
	"math/rand"
	"strings"
	"time"

	color "github.com/fatih/color"
)

const DotCount = 3
const Message = "Initializing"

func main() {
	dots := "."
	font := color.New(color.FgGreen, color.Bold)
	ticker := time.NewTicker(1 * time.Second)
	clearspaces := strings.Repeat(" ", DotCount)

	go func() {
		for range ticker.C {
			if len(dots) >= DotCount {
				dots = "."
			} else {
				dots += "."
			}
		}
	}()

	for {
		symbol := RandomSymbol()
		font.Printf("%v %v%v%v\r \r", symbol, Message, dots, clearspaces)
	}
}

func Symbols() []string {
	return []string{
		"゠", "ァ", "ア", "ィ", "イ", "ゥ", "ウ", "ェ", "エ", "ォ", "オ", "カ", "ガ", "キ", "ギ", "ク",
		"グ", "ケ", "ゲ", "コ", "ゴ", "サ", "ザ", "シ", "ジ", "ス", "ズ", "セ", "ゼ", "ソ", "ゾ", "タ",
		"ダ", "チ", "ヂ", "ッ", "ツ", "ヅ", "テ", "デ", "ト", "ド", "ナ", "ニ", "ヌ", "ネ", "ノ", "ハ",
		"バ", "パ", "ヒ", "ビ", "ピ", "フ", "ブ", "プ", "ヘ", "ベ", "ペ", "ホ", "ボ", "ポ", "マ", "ミ",
		"ム", "メ", "モ", "ャ", "ヤ", "ュ", "ユ", "ョ", "ヨ", "ラ", "リ", "ル", "レ", "ロ", "ヮ", "ワ",
		"ヰ", "ヱ", "ヲ", "ン", "ヴ", "ヵ", "ヶ", "ヷ", "ヸ", "ヹ", "ヺ", "ー", "ヽ", "ヾ", "ヿ",
	}
}

func RandomSymbol() string {
	symbols := Symbols()
	index := rand.Intn(len(symbols))
	return symbols[index]
}
