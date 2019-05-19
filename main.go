package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const anagram = "poultry outwits ants"
const easiestHash = "e4820b45d2277f3844eac66c903e84be"
const difficultHash = "23170acc097c24edb98fc5488ab033fe"
const hardHash = "665e5bcb0c20062fe8abaaf4628bb154"

var cleanedWordList []string
var prevWord string
var wg = sync.WaitGroup{}

//var msgChan = make(chan string, 50)

func main() {

	anagramLetterCount := len(strings.Replace(anagram, " ", "", -1))
	fmt.Printf("Anagram has %d letters\n", anagramLetterCount)

	anagramLetters := UniqueElementsCount(strings.Split(anagram, ""))
	delete(anagramLetters, " ")

	start := time.Now()

	words := ReadFile("wordlist.dms")

	for i := range words {
		if CheckMatchingLetters(words[i], anagramLetters) && prevWord != words[i] {
			prevWord = words[i]
			cleanedWordList = append(cleanedWordList, words[i])
		}
	}

	fmt.Printf("Found %d possible words in %s. Moving on...\n", len(cleanedWordList), time.Since(start))

	destroy := make(chan struct{})

		for i := range cleanedWordList {
			for j := range cleanedWordList {
				wg.Add(1)
				go func(i int, j int, anagramLetters map[string]int, anagramLetterCount int, start time.Time, destroy <-chan struct{}, finishThis chan<- struct{}, wg *sync.WaitGroup) {
					select {
					case <-destroy: // triggered when the stop channel is closed
						fmt.Println("Broken Check for Triplets")
						break
					default:

						if CheckIfPossibleHashMatch(cleanedWordList[i] + " " + cleanedWordList[j], anagramLetters) {
							for k := range cleanedWordList {

								if len(cleanedWordList[k]) == anagramLetterCount-len(cleanedWordList[i] + " " + cleanedWordList[j]) + 1 {

									if CheckIfActualHashMatch(cleanedWordList[i] + " " + cleanedWordList[j] + " " + cleanedWordList[k], anagramLetters) {

										if GetMD5Hash(cleanedWordList[i] + " " + cleanedWordList[j] + " " + cleanedWordList[k], start) {
											finishThis <- struct{}{}
											close(finishThis)
										}
									}
								}
							}
						}
					}
					wg.Done()
				}(i, j, anagramLetters, anagramLetterCount, start, destroy, destroy, &wg)
			}
		}


		//go func(msgChan <- chan string) {
		//	for m := range msgChan {
		//		fmt.Println(m)
		//	}
		//	wg.Done()
		//}(msgChan)

	wg.Wait()
	fmt.Printf("This took %s\n", time.Since(start))

}


func ReadFile(fileName string) []string {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Open file error: %v", err)

		return nil
	}

	defer f.Close()
	words := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		words = append(words, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)

		return nil
	}

	return words
}

func CheckMatchingLetters(phrase string, anagramLetters map[string]int) bool {


	phraseMap := UniqueElementsCount(strings.Split(phrase, ""))
	//Delete spaces
	delete(phraseMap, " ")
	for phraseLetter := range phraseMap {

		if _, ok := anagramLetters[phraseLetter]; !ok {

			return false
		}
	}

	for anagramLetter, anagramLetterCount := range anagramLetters {
		if anagramLetterCount < phraseMap[anagramLetter] {
			return false
		}
	}

	return true
}

func UniqueElementsCount(s []string) map[string]int {
	unique := make(map[string]int, len(s))
	for _, elem := range s {
		unique[elem] = unique[elem] + 1
	}

	return unique

}

func CheckIfPossibleHashMatch(phrase string, anagramLetters map[string]int) bool {

	phraseMap := UniqueElementsCount(strings.Split(phrase, ""))
	for anagramLetter, anagramLetterCount := range anagramLetters {
		if anagramLetterCount < phraseMap[anagramLetter] {

			return false
		}
	}

	return true
}

func CheckIfActualHashMatch(phrase string, anagramLetters map[string]int) bool {

	phraseMap := UniqueElementsCount(strings.Split(phrase, ""))
	for anagramLetter, count := range anagramLetters {
		if count != phraseMap[anagramLetter] {
			return false
		}
	}

	return true
}

func GetMD5Hash(phrase string, start time.Time) bool{
	h := md5.New()
	h.Write([]byte(phrase))
	hexString := hex.EncodeToString(h.Sum(nil))

	switch hexString {
	case easiestHash:
		fmt.Printf("Found EASY hash: '%s' in %s\n", phrase, time.Since(start))
	case difficultHash:
		fmt.Printf("Found DIFFICULT hash: %s in %s\n", phrase, time.Since(start))
	case hardHash:
		fmt.Printf("Found HARD hash: %s in %s\n", phrase, time.Since(start))
		return true
	}
	return false
}

