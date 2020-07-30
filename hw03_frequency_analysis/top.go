package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	delimiters  = "?!.;,-:"
	space       = " "
	emptyString = ""
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func isDelimiter(c string) bool {
	return strings.Contains(delimiters, c)
}

func cleanDelimiter(input string) string {
	var result string
	var previousChar string

	for len(input) > 0 {
		r, w := utf8.DecodeRuneInString(input)
		if (string(r) == space && previousChar != space) || !isDelimiter(string(r)) {
			result += string(r)
			previousChar = string(r)
		} else if previousChar != space && isDelimiter(string(r)) {
			result += emptyString
		}
		input = input[w:]
	}
	return result
}

func wordCounter(s string, unique bool) map[string]int {
	strSlice := strings.Fields(cleanDelimiter(s))
	result := make(map[string]int)
	if unique {
		for _, str := range strSlice {
			result[str]++
		}
	} else {
		for _, str := range strSlice {
			result[strings.ToLower(str)]++
		}
	}
	return result
}

func top10words(text map[string]int) []string {
	if len(text) == 0 {
		return nil
	}

	i := 0
	var result []string
	pl := make(PairList, len(text))
	for k, v := range text {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	if len(pl) > 10 {
		for _, v := range pl[:10] {
			result = append(result, v.Key)
		}
	} else {
		for _, v := range pl {
			result = append(result, v.Key)
		}
	}
	return result
}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	return top10words(wordCounter(s, false))
}
