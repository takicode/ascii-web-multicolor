package core

import (
  "strings"
)

func ColorLogic(inputString, banner, inputWord, color string)(string, string){

 var result string
 var message string

 colorMap := make(map[string]string) 

 if inputWord != "" && color != "none" {
	inputColors := strings.Split(color, ",")
	inputWords := strings.Split(inputWord, ",")

	minLength := len(inputColors)
	if len(inputWords) < minLength {
		minLength = len(inputWords)
	}
	
	for i := 0; i < minLength; i++ {
		colorMap[inputWords[i]] = inputColors[i]
	}
 }

 AsciiArt := Generate(inputString, banner)

 if inputString == "" {
	result = ""
	message = `<span style="color:red">No text provided</span>`

 } else if inputWord == "" && color != "none" {
	result = `<span style="color:` + color + `">` + AsciiArt + `</span>`
	message = `<span style="color:red">input color text to see colour magic</span>`

 } else if inputWord == "" && color == "none" {
	result = AsciiArt
	message = `<span style="color:red">input color text to see colour magic</span>`

 } else if color == "none" || !strings.Contains(inputString, inputWord) {
	result = AsciiArt
	message = `<span style="color:red">select color to see colour magic</span>`

 } else {
  
	type segment struct {
		text    string
		colored string
	}

	var segments []segment
	remaining := inputString

	for len(remaining) > 0 {
		minIdx := -1
		var stringWord string
		var stringColor string

		for word, wordColor := range colorMap {
			idx := strings.Index(remaining, word)
			if idx != -1 && (minIdx == -1 || idx < minIdx) {
				minIdx = idx
				stringWord = word
				stringColor = wordColor
			}
		}

		if minIdx == -1 {
			segments = append(segments, segment{remaining, ""})
			break
		}

		if minIdx > 0 {
			segments = append(segments, segment{remaining[:minIdx], ""})
		}
		segments = append(segments, segment{stringWord, stringColor})
		remaining = remaining[minIdx+len(stringWord):]
	} 

	type asciiSeg struct {
		rows    []string
		colored string
	}

	var asciiSegments []asciiSeg

	for _, seg := range segments {
		if seg.text == "" {
			continue
		}
		rows := strings.Split(Generate(seg.text, banner), "\n")
		asciiSegments = append(asciiSegments, asciiSeg{rows, seg.colored})
	}

	maxRows := 0
	for _, seg := range asciiSegments {
		if len(seg.rows) > maxRows {
			maxRows = len(seg.rows)
		}
	}

	var lines []string
	for i := 0; i < maxRows; i++ {
		line := ""
		for _, seg := range asciiSegments {
			row := ""
			if i < len(seg.rows) {
				row = seg.rows[i]
			}
			if seg.colored != "" {
				line += `<span style="color:` + seg.colored + `">` + row + `</span>`
			} else {
				line += row
			}
		}
		lines = append(lines, line)
	}

	result = strings.Join(lines, "\n")
	message = `<span style="color:green">Yup!Check the magic below</span>`
 }

 return result, message
}