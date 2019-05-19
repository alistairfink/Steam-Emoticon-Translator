package main

import (
	"strings"
    "bufio"
    "os"
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "github.com/alistairfink/Steam-Emoticon-Translator/JsonObjects"
)

var letters map[string][]string

func main() {
	letters = make(map[string][]string)
	
	// Open Json File
	jsonFile, err := os.Open("alphabet.json")
	defer jsonFile.Close()
	if err != nil {
		println("Error:", err)
		return
	}

	// Read Json File as Bytes
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		println("Error:", err)
		return
	}

	// Unmarshal json file and Read Into Map
	var alphabet JsonObjects.Alphabet
	err = json.Unmarshal(jsonBytes, &alphabet)

	for _, letter := range alphabet.Letters {
		letters[letter.Letter] = letter.Emoticons
	}

	reader := bufio.NewReader(os.Stdin)
	for ;; {
		print("Enter Text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			println("Error: ", err)
		} else {
			for text[len(text)-1:len(text)] == "\r" || text[len(text)-1:len(text)] == "\n" {
				text = text[:len(text)-1]
			}
			if text == ":q" {
				println("\nClosing...")
				return
			}

			println(translateText(text))
		}
	}
}

func translateText(text string) string {
	split := strings.Split(text, "")
	translation := ""
	invalid := false
	for i, letter := range split {
		letter = strings.ToLower(letter)
		if letter == " " {
			translation += " "
			continue
		}

		if _, exists := letters[letter]; !exists {
			println("Invalid Character:", letter, " at index", i)
			invalid = true
			continue
		}

		if len(letters[letter]) == 0 {
			translation += letter
		} else {
			translation += letters[letter][rand.Intn(len(letters[letter]))]
		}
	}

	if !invalid {
		return translation
	}

	return ""
}