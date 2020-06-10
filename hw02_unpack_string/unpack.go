package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func manipulateWithLetters(letter string, count string) string {
	c, err := strconv.Atoi(count)
	if err != nil {
		log.Fatal("Shit, could be worse")
	}
	r := strings.Repeat(letter, c-1)
	return r
}

func isDigit(n string) bool {
	_, err := strconv.ParseInt(n, 10, 16)
	return err == nil
}

func isZero(z string) bool {
	d, err := strconv.ParseInt(z, 10, 8)
	if err != nil {
		log.Fatal("Shit, could be worse")
	}
	if d == 0 {
		return true
	}
	return false
}

func removeFromStr(str string) string {
	s := str[:len(str)-1]
	return s
}

func startsWithLetterOrDigit(str string) bool {
	reg := regexp.MustCompile("[a-zA-Z0-9]+")
	return reg.MatchString(str)
}

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	if !startsWithLetterOrDigit(s) {
		return "", ErrInvalidString
	}

	result := strings.Builder{}
	previousItem := ""

	for i := 0; i < len(s); i++ {
		if isDigit(string(s[i])) {
			if previousItem == "" || isDigit(previousItem) {
				return "", ErrInvalidString
			}
			if !isZero(string(s[i])) {
				result.WriteString(manipulateWithLetters(previousItem, string(s[i])))
			} else {
				t := result.String()
				result.Reset()
				result.WriteString(removeFromStr(t))
			}
		} else {
			result.WriteByte(s[i])
		}
		previousItem = string(s[i])
	}
	return result.String(), nil
}
