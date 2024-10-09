package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"unicode"
	"io"


)

/*
Task:

Create a Go program that reads a text file and counts the occurrence of each word. The program should print the words and their respective counts, sorted by the frequency in descending order.
Requirements:

    The program should accept the file path as a command-line argument.
    Ignore case sensitivity (e.g., "Go" and "go" should be counted as the same word).
    Punctuation marks should be ignored (e.g., "hello," and "hello" should be counted as the same word).
    Print the result as word: count, sorted by the count in descending order.
*/


type KeyValue struct {
	Key   string
	Value int
}


func ReadFileToBuffer(path string)  ([]byte, error) {
	buff := make([]byte, 2048)

	file, err := os.Open(path); if err != nil {
		return nil, fmt.Errorf("error openning file %s, %v\n", path, err)
	}
	defer file.Close()

	for {
		n, err := file.Read(buff); if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading to buffer %s, %v\n", path, err)
		}

		if n == 0 {
			break
		}
	}
	return buff, nil

}

func normalizeString(raw string, ) string {
	fmt.Println("Processing word...", raw)
	var normal string
	for _, char := range raw {
		if !unicode.In(char, unicode.Letter) {
			continue
		}
		normal+=string(char)
	}

	return normal
}




func main() {
	fmt.Println("Program has started...")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <relative file path>")
		os.Exit(1)
	} 

	

	bytesRead, err := ReadFileToBuffer(os.Args[1]); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str := string(bytesRead)
	wordSlice := strings.Split(str, " ")



	result := make(map[string]int)


	
	for _, word := range wordSlice {
		word = strings.ToLower(word)
		word =  normalizeString(word)
		
		_, exists := result[word]
		if exists {
			result[word]++
		} else { 
			result[word] = 1
		}

	}

	var kvSlice []KeyValue
	for k, v := range result {
		kvSlice = append(kvSlice, KeyValue{k, v})
	}


	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	fmt.Println("Map sorted by values:")
	for _, kv := range kvSlice {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}


	fmt.Println("Program has finished...")
}