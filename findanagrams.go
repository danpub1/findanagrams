package main

import (
    "fmt"
    "sort"
    "strings"
)

/*
var primes = [...]uint64{
    1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
    53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 1, 1, 1, 1, 1,
    1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
    53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 1, 1, 1, 1, 1,
}

func hash(s string) Code {
    var low uint64 = 1
    var high uint64 = 0

    for _, ss := range s {
        h1, l1 := bits.Mul64(low, primes[ss-' '])
        _, l2 := bits.Mul64(high, primes[ss-' '])
        low = l1
        high = l2 + h1
    }

    return Code{low, high}
}

func title(source string) string {
    if len(source) > 1 {
        return strings.ToUpper(string(source[0])) + strings.ToLower(source[1:])
    } else if len(source) > 0 {
        return strings.ToUpper(source)
    } else {
        return source
    }
}
*/

// Code is the hash value
type Code struct {
    low, high uint64
}

func remainder(source, word string) (string, bool) {
    ll := len(source)

    if len(word) > ll {
        return "", false
    }

    for _, ss := range word {
        orgLen := len(source)
        source = strings.Replace(source, string(ss), "", 1)
        if orgLen == len(source) {
            return "", false
        }
    }

    return source, true
}

// ByLen implements sort.Interface for []string based on string length descending
type ByLen []string

func (a ByLen) Len() int      { return len(a) }
func (a ByLen) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool {
    return len(strings.SplitN(a[i], "/", 2)[0]) >= len(strings.SplitN(a[j], "/", 2)[0])
}

func findSome(source string, result string, numWords int, startIdx int, results []string, words []string, minWordLength int, minWords int, maxWords int, repeatWords bool, stopAfter int) []string {
    numWords++
    // fmt.Printf("findSome(\"%s\" \"%s\")\n", source, result)
    for idx := startIdx; idx < len(words); idx++ {
        startWords := strings.SplitN(words[idx], "/", 2)
        if numWords < maxWords || (numWords == maxWords && len(startWords[0]) == len(source)) {
            if rem, exists := remainder(source, startWords[0]); exists {
                var newResult string
                if len(result) > 0 {
                    newResult = result + " " + startWords[1]
                } else {
                    newResult = startWords[1]
                }
                if len(rem) == 0 && numWords >= minWords && numWords <= maxWords {
                    results = append(results, newResult)
                    if stopAfter > 0 && len(results) >= stopAfter {
                        return results
                    }
                } else if len(rem) >= minWordLength && numWords < maxWords {
                    nextIdx := idx
                    if !repeatWords {
                        nextIdx++
                    }
                    results = findSome(rem, newResult, numWords, nextIdx, results, words, minWordLength, minWords, maxWords, repeatWords, stopAfter)
                    if stopAfter > 0 && len(results) >= stopAfter {
                        return results
                    }
                }
            }
        }
    }
    return results
}

func isExcluded(entry string, excludedWords []string) bool {
    for _, excludedWord := range excludedWords {
        if entry == excludedWord {
            return true
        }
    }
    return false
}

func LinkToWiktionary(results []string) []string {
    for ii := range results {
        words := strings.Split(results[ii], " ")
        for jj := range words {
            word := strings.Split(words[jj], "/")
            for kk := range word {
                word[kk] = fmt.Sprintf("<a href=\"https://en.wiktionary.org/wiki/%v\" target=\"_blank\">%v</a>", word[kk], word[kk])
            }
            words[jj] = strings.Join(word, "/")
        }
        results[ii] = strings.Join(words, " ")
    }
    return results
}

func findAnagrams(source string, includeWords string, excludeWords string, minWordLength int, maxWordLength int, minWords int, maxWords int, stopAfter int, locale uint16, part uint8, dict uint16, repeatWords bool, targetWordList bool, linkToWiktionary bool) []string {
    cleanSource := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return r
        case r >= 'a' && r <= 'z':
            return r - 'a' + 'A'
        }
        return ' '
    }

    source = strings.ReplaceAll(strings.Map(cleanSource, source), " ", "")
    var excludedWords []string
    if len(excludeWords) > 0 {
        excludedWords = strings.Split(strings.TrimSpace(excludeWords), " ")
        for idx, excludedWord := range excludedWords {
            excludedWords[idx] = strings.ReplaceAll(strings.Map(cleanSource, excludedWord), " ", "")
        }
    }

    var wordMap map[Code]string = make(map[Code]string)
    for _, aword := range allWords {
        awordCode := Code{aword.low, aword.high}
        wordLength := len(strings.ReplaceAll(aword.entry, string("'"), ""))

        if wordLength >= minWordLength && wordLength <= maxWordLength && (aword.locale&locale) != 0 && (aword.part&part) != 0 && (aword.dictSize&dict) != 0 {
            entry := strings.ToUpper(strings.ReplaceAll(aword.entry, string("'"), ""))
            if !isExcluded(entry, excludedWords) {
                if _, foundInSource := remainder(source, entry); foundInSource {
                    prev, exists := wordMap[awordCode]
                    if exists {
                        wordMap[awordCode] = prev + "/" + aword.entry
                    } else {
                        if targetWordList {
                            wordMap[awordCode] = aword.entry
                        } else {
                            wordMap[awordCode] = entry + "/" + aword.entry
                        }
                    }
                }
            }
        }
    }

    var words []string = make([]string, len(wordMap))
    var ii = 0
    for _, word := range wordMap {
        words[ii] = word
        ii = ii + 1
    }

    wordMap = nil

    if targetWordList {
        sort.Strings(words)
        if linkToWiktionary {
            words = LinkToWiktionary(words)
        }
        return words
    }

    sort.Sort(ByLen(words))

    var results []string = make([]string, 0)
    if len(includeWords) > 0 {
        result := ""
        notFound := false
        for _, includedWord := range strings.Split(strings.TrimSpace(includeWords), " ") {
            includedWord = strings.ReplaceAll(strings.Map(cleanSource, includedWord), " ", "")
            if len(includedWord) > 0 {
                if remainder, exists := remainder(source, includedWord); exists {
                    source = remainder
                    var found []string = make([]string, 0)
                    found = findSome(includedWord, "", 0, 0, found, words, 1, 1, 1, false, 1)
                    if len(result) > 0 {
                        result = result + " " + found[0]
                    } else {
                        result = found[0]
                    }
                } else {
                    notFound = true
                }
            }
        }
        if !notFound {
            results = findSome(source, result, 0, 0, results, words, minWordLength, minWords-2, maxWords-1, repeatWords, stopAfter)
        }
    } else {
        results = findSome(source, "", 0, 0, results, words, minWordLength, minWords, maxWords, repeatWords, stopAfter)
    }

    if linkToWiktionary {
        results = LinkToWiktionary(results)
    }

    return results
}
