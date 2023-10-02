package ascii_art

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func AsciiArt(input string, fonts string, page http.ResponseWriter, fontFolder string) {
	fontFilePath := filepath.Join(fontFolder, fonts+".txt")
	var data []byte

	data, err := os.ReadFile(fontFilePath)
	if err != nil {
		http.Error(page, fmt.Sprintf("Error reading font file: %v", err), http.StatusInternalServerError)
		return
	}

	input = strings.Replace(input, "\\n", "\n", -1) // change "\\n" to "\n"
	inputText := strings.Split(input, getNewlineSeparator([]byte(input)))
	fontString := strings.Split(string(data), getNewlineSeparator(data))

	for _, text := range inputText {
		if text == "" {
			fmt.Fprintln(page) // Вывести пустую строку
		} else {
			for i := 0; i < 9; i++ {
				for _, char := range text {
					start, _ := getLetter(char)
					fmt.Fprint(page, fontString[start+i])
				}
				fmt.Fprintln(page) // Перейти на новую строку после строки букв
			}
		}
	}
}

func getLetter(c rune) (int, int) {
	charList := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_'abcdefghijklmnopqrstuvwxyz{|}~"
	index := strings.IndexRune(charList, c)
	if index == -1 {
		fmt.Println("Missing input characters")
		return -1, -1
	}
	firstLine := index * 9
	lastLine := firstLine + 8
	return firstLine, lastLine
}

func error(page http.ResponseWriter) {
	http.Error(page, "Invalid input.", http.StatusBadRequest)
}

func getNewlineSeparator(data []byte) string {
	if strings.ContainsRune(string(data), '\r') {
		return "\r\n" // Windows-style
	}
	return "\n" // Unix-style
}
