package app

import (
	"ascii-art-fs/utils"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	fontsPath     = "assets/fonts/"
	fonstStandard = "standard.txt"
	fileSuffix    = "txt"
	countLines    = 856
)

func isNewLineOnly(str string) (bool, string) {
	str = strings.ReplaceAll(str, "\\n", "\n")
	for _, elem := range str {
		if elem != '\n' {
			return false, ""
		}
	}
	return true, str
}

func fileRead(fileName string) (string, error) {
	data, err := os.ReadFile(fontsPath + fileName)
	if err != nil {
		return "", err
	}
	str1 := bytes.Split(data, []byte("\n"))
	if len(str1) != countLines {
		return "", errors.New(utils.ErrMessageInvalidFileName)
	}
	return string(data), nil
}

func final(cord []string, mapa map[rune]string) {
	result := ""
	for _, words := range cord {
		result += answer(words, mapa)
	}
	fmt.Print(result)
}

func checkAscii(cord, data string) (map[rune]string, []string) {
	mapa := make(map[rune]string)
	temp := ""
	count := 0
	keys := rune(32)
	strand := strings.ReplaceAll(cord, "\n", "\\n")
	splitText := strings.Split(strand, "\\n")

	for _, words := range data {
		temp += string(words)
		if words == '\n' {
			count++
		}
		if count == 9 {
			mapa[keys] = temp[1 : len(temp)-1]
			temp = ""
			count = 0
			keys++
		}

	}
	return mapa, splitText
}

func answer(cord string, letters map[rune]string) string {
	var template [8]string
	if len(cord) == 0 {
		return "\n"
	}
	for _, w := range cord {
		for index, value := range strings.Split(letters[w], "\n") {
			template[index] += value
		}
	}
	result := ""
	for _, words := range template {
		result += string(words) + "\n"
	}

	return result
}

func suffixValidation(s string) bool {
	return s == fileSuffix
}

func getFontName(slice ...string) (string, error) {
	if len(slice) == 1 {
		return fonstStandard, nil
	}
	splitText := strings.SplitN(slice[1], ".", 2)
	if len(splitText) == 2 {
		if !suffixValidation(splitText[1]) {
			return "", errors.New(utils.ErrMessageInvalidFileFormat)
		} else {
			return slice[1], nil
		}
	}
	return splitText[0] + ".txt", nil
}

func Start() error {
	arguments := os.Args[1:]
	if len(arguments) > 2 || len(arguments) < 1 {
		fmt.Println(utils.ErrMessageUsage)
		return nil
	}
	fontName, err := getFontName(arguments...)
	if err != nil {
		return err
	}
	data, err := fileRead(fontName)
	if err != nil {
		return err
	}
	words := arguments[0]
	if ok, value := isNewLineOnly(words); ok {
		fmt.Print(value)
		return nil
	}
	mapa, splitText := checkAscii(words, data)
	final(splitText, mapa)
	return err
}
