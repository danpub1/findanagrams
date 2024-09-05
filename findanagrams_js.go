package main

import (
	"syscall/js"
)

func convertResults(results []string) []interface{} {
	rv := make([]interface{}, len(results))
	for ii := range results {
		rv[ii] = results[ii]
	}
	return rv
}

func findAnagramsJS(this js.Value, inputs []js.Value) interface{} {
	defer func() {
		recover()
	}()

	source := inputs[0].String()
	includeWords := inputs[1].String()
	excludeWords := inputs[2].String()
	minLength := inputs[3].Int()
	maxLength := inputs[4].Int()
	minWords := inputs[5].Int()
	maxWords := inputs[6].Int()
	stopAfter := inputs[7].Int()
	locale := uint16(inputs[8].Int())
	part := uint8(inputs[9].Int())
	dict := uint16(inputs[10].Int())
	repeatWords := inputs[11].Bool()
	targetWords := inputs[12].Bool()
	linkToWiktionary := inputs[13].Bool()

	results := findAnagrams(source, includeWords, excludeWords, minLength, maxLength, minWords, maxWords, stopAfter, locale, part, dict, repeatWords, targetWords, linkToWiktionary)

	return js.ValueOf(convertResults(results))
}

// GOOS=js GOARCH=wasm go build -o findanagrams.wasm
func main() {
	c := make(chan bool)
	js.Global().Set("findAnagramsJS", js.FuncOf(findAnagramsJS))
	<-c
}
