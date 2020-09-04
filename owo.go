package owo

import (
	"math/rand"
	"strings"
	"time"
)

// Define a prefixes variable as a slice of strings
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

// Define a suffixes variable as a slice of strings
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

// Translate the given text. You can set withPrefix and withSuffix
// to true or false, if you want them or not.
// It returns the translated text
func Translate(text string, withPrefix bool, withSuffix bool) string {

	// Define a concat variable as a slice of strings
	var concat []string

	// if withPrefix is set to true add it to the concat slice
	if withPrefix {
		concat = append(concat, prefixes[random(len(prefixes))])
	}

	// substitute characters and add the processed string to the concat slice
	concat = append(concat, substitute(text))

	// if withSuffix is set to true add it to the concat slice
	if withSuffix {
		concat = append(concat, suffixes[random(len(suffixes))])
	}

	// Join the concat slice and return it as a string
	return strings.Join(concat, " ")
}

// It replaces characters in the given text
// It returns the processed string
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

// It returns a random number
func random(max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
}
