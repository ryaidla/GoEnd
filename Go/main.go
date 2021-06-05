package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Word struct {
	wordBytes []byte
	count int
}

func wordCopyCheck(words []*Word, word []byte) (bool, int) {
	index := 0
	for _, wordStruct := range words {
		if string(wordStruct.wordBytes) == string(word) {
			return true, index
		}
		index++
	}
	return false, -1
}


func main() {
	start := time.Now()

	file, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		panic(err)
	}

	var words []*Word
	res := make(chan []byte)

	go func() {
		var wordFromBytes []byte
		for _, b := range file {
			if b >= 65 && b <= 90 {
				wordFromBytes = append(wordFromBytes, b + 32)
			} else if b >= 97 && b <= 122 {
				wordFromBytes = append(wordFromBytes, b)
			} else if len(wordFromBytes) > 0 {
				res <- wordFromBytes
				wordFromBytes = []byte{}
			}
		}
		close(res)
	}()

	for word := range res {
		isCopy, index := wordCopyCheck(words, word)
		if !isCopy {
			words = append(words, &Word{word, 1})
		} else {
			if len(words) > 0 {
				words[index].count++
			}
		}
	}


	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})

	for i := 0; i < len(words) && i < 20; i++ {
		fmt.Println(string(words[i].wordBytes), ":", words[i].count)
	}

	fmt.Printf("Process took %s\n", time.Since(start))
}