package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var prefixes = []string{
	"<3 ",
	"0w0 ",
	"H-hewwo?? ",
	"HIIII! ",
	"Haiiii! ",
	"Huohhhh. ",
	"OWO ",
	"OwO ",
	"UwU ",
}

var suffixes = []string{
	" ( ͡° ᴥ ͡°)",
	" (இωஇ )",
	" (๑•́ ₃ •̀๑)",
	" (• o •)",
	" (●´ω｀●)",
	" (◠‿◠✿)",
	" (✿ ♡‿♡)",
	" (　\"◟ \")",
	" (人◕ω◕)",
	" (；ω；)",
	" (｀へ´)",
	" ._.",
	" :3",
	" :D",
	" :P",
	" ;-;",
	" ;3",
	" ;_;",
	" >_<",
	" >_>",
	" UwU",
	" XDDD",
	" ^-^",
	" ^_^",
	" x3",
	" x3",
	" xD",
	" ÙωÙ",
	" ʕʘ‿ʘʔ",
	" ʕ•̫͡•ʔ",
	" ㅇㅅㅇ",
	", fwendo",
	"（＾ｖ＾）",
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	text := strings.Join(os.Args[1:], " ")
	concat := []string{
		prefixes[random(len(prefixes))],
		substitute(text),
		suffixes[random(len(suffixes))],
	}
	fmt.Printf("%s\n", strings.Join(concat, " "))
}

func substitute(text string) string {
	return strings.NewReplacer(
		"r", "w",
		"l", "w",
		"R", "W",
		"L", "W",
		"no", "nu",
		"has", "haz",
		"have", "haz",
		"you", "uu",
		"the ", "da ",
		"The ", "Da ").Replace(text)
}

func random(max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
}
