package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Output struct {
	Result string
	Error  string
}

func printOutput(text string, font string) []byte {
	text = strings.ReplaceAll(text, "\r", "")
	textRows := strings.Split(text, "\n")
	fontRows := LoadFont("font/" + font + ".txt")
	var printOutput []byte
	for r := 0; r < len(textRows); r++ {
		if len(textRows[r]) != 0 {
			for k := 0; k < 8; k++ {
				for i := 0; i < len(textRows[r]); i++ {
					if AllowedChar(textRows[r][i]) {
						printOutput = append(printOutput, fontRows[(int(textRows[r][i])-31)*9-8+k]...)
					}
				}
				printOutput = append(printOutput, byte(10))
			}
		} else {
			printOutput = append(printOutput, byte(10))
		}
	}
	return printOutput[:len(printOutput)-1]
}

func LoadFont(name string) []string {
	file, _ := os.ReadFile(name)
	file = []byte(strings.ReplaceAll(string(file), "\r", ""))
	return strings.Split(string(file), "\n")
}

func AllowedChar(letter byte) bool {
	if letter >= 32 && letter <= 126 {
		return true
	}
	return false
}

func main() {
	http.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))
	tmpl := template.Must(template.ParseFiles("front/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method { //get request method
		case "GET": //Web server got requested to show first page
			if r.URL.Path != "/" {
				http.Error(w, "Bad request - 404 Page Not Found", http.StatusNotFound)
				return
			}
			if r.Method != "GET" {
				http.Error(w, "Bad request - 405 Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
			tmpl.Execute(w, nil)
		case "POST": //web server got request to generate something additional from user filled form
			if r.URL.Path != "/" {
				http.Error(w, "Bad request - 404 Page Not Found", http.StatusNotFound)
				return
			}
			if r.Method != "POST" {
				http.Error(w, "Bad request - 405 Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
			text := r.FormValue("textin")
			font := r.FormValue("font")
			
			err := tmpl.Execute(w, Output{Result: string(printOutput(text, font)), Error: ""})
				if err != nil {
			w.WriteHeader(500)
			}
		default:
			http.Error(w, "Bad request - 405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	fmt.Println("Open http://localhost:8080")
	fmt.Println("To terminate server press CTRL + \"c\"")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
