package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"fmt"
	"regexp"
)

func getWords(text string) []string {
	words := regexp.MustCompile(`\pL+('\pL+)*`)
	return words.FindAllString(text, -1)
}

func countWords(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

//func sortData(data map[string]int) map[string]int {
//	sorted := make(map[string]int)
//	values := make([]int, 0, len(data))
//	for _, v := range data {
//		values = append(values, v)
//	}
//	sort.Ints(values)
//	for _, v := range values {
//		fmt.Println(v)
//	}
//	return sorted
//}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}
	//var result []string
	data := countWords(getWords(s))
	for i := 0; i < 10; i++ {
	}
	fmt.Println(data)
	return nil
}
