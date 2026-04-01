package core

import (
  "strings"
)

func ColorLogic(inputString, banner, inputWord, color string)(string, string){

 var result string
 var message string


 
colorMap :=make(map[string]string) 

 inputColors := strings.Split(color, ",");
 inputWords:= strings.Split(inputWord, ",");

 minLength := 0;

 if len(inputColors) < len(inputWords){
	minLength = len(inputColors)
 }else{
	minLength = len(inputWords)
 }
 
 for i:= 0; i<minLength; i++{
	colorMap[inputWord[i]] = inputColors[i]
 }


 


  
  AsciiArt := Generate(inputString, banner)

  if inputString == "" {
     result = ""
     message = `<span style="color:red">No text provided</span>`

} else if inputWord == "" && color != "none" {

    result = `<span style="color:` + color + `">` + AsciiArt  + `</span>`
    message = `<span style="color:red">input color text to see colour magic</span>`

}else if inputWord == "" && color == "none" {
    result = AsciiArt
     message = `<span style="color:red">input color text to see colour magic</span>`
}else if inputWord == "" || color == "none" || !strings.Contains(inputString, inputWord) {
   
    result = AsciiArt
    message = `<span style="color:red">select color to see colour magic</span>`

}else if color == "none"{
    result = AsciiArt
    message = `<span style="color:red">input color text to see colour magic</span>`
} else {
  
    type segment struct {
        text    string
        colored bool
    }

    var segments []segment

    remaining := inputString

	for remaining > 0{

		minIdx := -1;
        var stringWord string
		var stringColor string

		for word, color := range colorMap{
        
		idx := strings.Index(remaining, word)
         
		if idx != -1 && (min ==-1 || minIdx < idx){
			minIdx = idx;
			stringWord = word;
			stringColor = color
		}

	    if minIdx > -1{
			segments = append(segments, segment{})
		}
	}


	}

    









    // for {
    //     idx := strings.Index(remaining, inputWord)
    //     if idx == -1 {
    //         if remaining != "" {
    //             segments = append(segments, segment{remaining, false})
    //         }
    //         break
    //     }
    //     if idx > 0 {
    //         segments = append(segments, segment{remaining[:idx], false})
    //     }
    //     segments = append(segments, segment{inputWord, true})
    //     remaining = remaining[idx+len(inputWord):]
    // }

    type asciiSeg struct {
        rows []string
        colored bool
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
            if seg.colored {
                line += `<span style="color:` + color + `">` + row + `</span>`
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