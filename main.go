package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func sortingValues(word string) string {

	sortingWord := strings.Split(word, "")
	sort.Strings(sortingWord)
	return strings.Join(sortingWord, "")

}

func hashedValues(value string) int32 {

	sortedValues := sortingValues(value)

	hash := fnv.New32()
	hash.Write([]byte(sortedValues))

	data := hash.Sum32()

	return int32(data)

}

func timerToCheck(startTime time.Time, name string) {

	timeDuration := time.Since(startTime)
	log.Printf("%s program took time: %s", name, timeDuration)

}

func main() {

	defer timerToCheck(time.Now(), "anagrams")

	file, err := os.Open("wordList.txt")

	if err != nil {

		log.Fatal(err)

	}

	defer file.Close()

	wordsMap := make(map[int]map[int32][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		words := scanner.Text()
		wordLength := len(words)
		wordsHashed := hashedValues(words)
		wordList := wordsMap[wordLength][wordsHashed]
		wordList = append(wordList, words)

		if wordsMap[wordLength] == nil {

			temporalHash := make(map[int32][]string)
			temporalHash[wordsHashed] = []string{words}
			wordsMap[wordLength] = temporalHash

		} else {

			wordsMap[wordLength][wordsHashed] = wordList

		}

	}

	file, err = os.Open("wordList.txt")

	if err != nil {

		log.Fatal(err)

	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	setsQuantity := 0

	for scanner.Scan() {

		word := scanner.Text()
		wordsLenght := len(word)
		hashedWords := hashedValues(word)

		if len(wordsMap[wordsLenght][hashedWords]) > 1 {

			fmt.Println(wordsMap[wordsLenght][hashedWords])
			wordsMap[wordsLenght][hashedWords] = nil
			setsQuantity++

		}

	}

	fmt.Println("Te quantity of the sets are: ", setsQuantity)

}
