package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func TestUniqueElementsCount(t *testing.T) {
	test := []string{"a", "b", "c", "c", "", "xx"}
	expect := map[string]int{
		"a": 1,
		"b": 1,
		"c": 2,
		 "": 1,
		"xx": 1,

	}
	assert.Equal(t, UniqueElementsCount(test), expect)
}

func TestCheckIfPossibleHashMatch(t *testing.T) {
	test := CheckIfPossibleHashMatch("Zttainooprssttuuw", map[string]int{"a":1, "i":1, "l":1, "n":1, "o":2, "p":1, "r":1, "s":2, "t":4, "u":2, "w":1,})
	assert.Equal(t, true, test)

}

func TestCheckIfActualHashMatch(t *testing.T) {
	test := CheckIfActualHashMatch("ailnoopr ssttttu Zuw", map[string]int{"a":1, "i":1, "l":1, "n":1, "o":2, "p":1, "r":1, "s":2, "t":4, "u":2, "w":1,})
	assert.Equal(t, true, test)

}

func TestCheckMatchingLetters(t *testing.T) {

	testAnagram := UniqueElementsCount(strings.Split("abbcd", ""))
	testWord1 := "ab"
	testWord2 := "ab b"
	testWord3 := "abbc z"
	assert.Equal(t, true, CheckMatchingLetters(testWord1, testAnagram))
	assert.Equal(t, true, CheckMatchingLetters(testWord2, testAnagram))
	assert.Equal(t, false, CheckMatchingLetters(testWord3, testAnagram))
}