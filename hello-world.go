package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)



func readFileInfo(fileName string) (int, []string) {    
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineCount := 0
    lines:= make([]string,1)
    for scanner.Scan() {
        lineCount++
        lines = append(lines, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    return lineCount, lines
}


func main(){
    fmt.Println("SonarKube")
    
    lineCount, lines := readFileInfo("amk.txt")
    
    _ = lines // placeholder 
    
    output := fmt.Sprintf("Number of lines: %d", lineCount)
    fmt.Println(output)
}

