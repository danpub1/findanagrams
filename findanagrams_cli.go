//go:build !web && !js
package main

import (
	"flag"
	"fmt"
)

// go build
func main() {
	var american bool
	var australian bool
	var british bool
	var canadian bool
	var variants bool
	var contractions bool
	var properNames bool
	var abbreviations bool
	var small bool
	var medium bool
	var large bool
	var huge bool
	var insane bool
	var repeatsOk bool
	var targetsOnly bool
	var linkToWiktionary bool
	var anagramWords string
	var excludeWords string
	var includeWords string
	var minWordLength int
	var maxWordLength int
	var minWordCount int
	var maxWordCount int
	var maxMatches int

    flag.StringVar(&anagramWords, "anagram", "", "Words to Anagram")
    flag.StringVar(&excludeWords, "exclude", "", "Words to Exclude from Anagram")
    flag.StringVar(&includeWords, "include", "", "Words to Include in Anagram")
    flag.BoolVar(&american, "american", true, "Include American Words")
    flag.BoolVar(&australian, "australian", false, "Include Australian Words")
    flag.BoolVar(&british, "british", false, "Include British Words")
    flag.BoolVar(&canadian, "canadian", false, "Include Canadian Words")
    flag.BoolVar(&variants, "variants", false, "Include Spelling Variants")
    flag.BoolVar(&contractions, "contractions", false, "Include Contractions")
    flag.BoolVar(&properNames, "proper-names", false, "Include Proper Names")
    flag.BoolVar(&abbreviations, "abbreviations", false, "Include Abbbreviations")
    flag.BoolVar(&small, "small", true, "Use Small Dictionary")
    flag.BoolVar(&medium, "medium", false, "Use Medium Dictionary")
    flag.BoolVar(&large, "large", false, "Use Large Dictionary")
    flag.BoolVar(&huge, "huge", false, "Use Huge Dictionary")
    flag.BoolVar(&insane, "insane", false, "Use Insane Dictionary")
    flag.BoolVar(&repeatsOk, "repeats-ok", false, "OK to repeat words")
    flag.BoolVar(&targetsOnly, "targets-only", false, "Only Show Target Words")
    flag.BoolVar(&linkToWiktionary, "link-to-wiktionary", false, "Link to Wiktionary")
    flag.IntVar(&minWordLength, "min-word-length", 1, "Minimum Word Length")
    flag.IntVar(&maxWordLength, "max-word-length", 99, "Maximum Word Length")
    flag.IntVar(&minWordCount, "min-word-count", 1, "Minimum Word Count")
    flag.IntVar(&maxWordCount, "max-word-count", 99, "Maximum Word Count")
    flag.IntVar(&maxMatches, "matches", 500, "Maximum Number of Matches")
    flag.Parse()

	var locale uint16
	var part uint8
	var dict uint16

	locale = 1
	if (variants) {
		locale |= uint16(2 + 4 + 8)
	}
	if (american) {
		locale |= uint16(16)
	}
	if (australian) {
		locale |= uint16(32)
		if (variants) {
			locale |= uint16(64 + 128)
		}
	}
	if (british) {
		locale |= uint16(256)
		if (variants) {
			locale |= uint16(512 + 1024 + 2048)
		}
	}
	if (canadian) {
		locale |= uint16(4096)
		if (variants) {
			locale |= uint16(8192 + 16384)
		}
	}
	part = 1
	if (contractions) {
		part |= uint8(4)
	}
	if (properNames) {
		part |= uint8(2 + 16)
	}
	if (abbreviations) {
		part |= uint8(8)
	}
	dict = 0
	if (small) {
		dict |= uint16(1 + 2 + 4)
	}
	if (medium) {
		dict |= uint16(8 + 16)
	}
	if (large) {
		dict |= uint16(32 + 64 + 128)
	}
	if (huge) {
		dict |= uint16(256)
	}
	if (insane) {
		dict |= uint16(512)
	}
	if (dict == 0) {
		dict = uint16(1 + 2 + 4)
	}
	results := findAnagrams(anagramWords, includeWords, excludeWords, minWordLength, maxWordLength, minWordCount, maxWordCount, maxMatches, locale, part, dict, repeatsOk, targetsOnly, linkToWiktionary)

	for idx, result := range results {
		fmt.Printf("%d: %s\n", idx+1, result)
	}
}
