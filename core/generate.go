package core 


import (
	"bufio"
	"os"
	"strings"
	"log"
)

func Generate(text, banner string) string {
    var word []string

    inputText := strings.Split(text, "\n")

    stdFileRep, err := os.Open("./banner/" + banner + ".txt")
    if err != nil {
        log.Printf("Error opening file: %v", err)
        return "Error: Banner file not found"
    }
    defer stdFileRep.Close()

    scanner := bufio.NewScanner(stdFileRep)

    for scanner.Scan() {
        word = append(word, scanner.Text())
    }

    if err := scanner.Err(); err != nil {      
        log.Printf("Scanner error: %v", err)
        return "Error: Could not read banner file"
    }

    var filePrint []string

    for _, char := range inputText {
        if char == "" {
            filePrint = append(filePrint, "\n")
            continue
        }
        for row := 0; row < 8; row++ {
            rows := ""
            for i := 0; i < len(char); i++ {
                asciiValue := int(char[i]-32)*9 + 1
                index := asciiValue + row
                if index >= len(word) {
                    continue
                }
                rows += word[index]
            }
            filePrint = append(filePrint, rows+"\n")
        }
    }

    return strings.Join(filePrint, "")
}