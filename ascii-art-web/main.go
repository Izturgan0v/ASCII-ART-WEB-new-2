package main

import (
	asciiart "ascii-art-web/ascii-art"
	"fmt"
	"net/http"
	"os"
)

//---------------------------------------------------------------------------------------|

var errFile string

func main() {
	content, err := os.ReadFile("templates/error.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	errFile = string(content)

	ip := "localhost"
	port := 45674

	// Обработчик для статических файлов (CSS, JS, изображения)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/sabik", handlerSabik)
	http.HandleFunc("/ascii-art", handlerAsciiart)

	addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("server runing http://%s\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

//---------------------------------------------------------------------------------------|

func handlerHome(w http.ResponseWriter, r *http.Request) {
	indexPage, err := os.ReadFile("templates/index.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}
	// нужен обработчик неизвестных путей

	if r.URL.Path != "/" {
		handlerError(w, http.StatusNotFound)
		return
	}

	fmt.Fprint(w, string(indexPage))
}

//---------------------------------------------------------------------------------------|

func handlerSabik(w http.ResponseWriter, r *http.Request) {
	sabikPage, err := os.ReadFile("templates/sabik.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}

	// нужен обработчик неизвестных путей

	fmt.Fprint(w, string(sabikPage))

}

//---------------------------------------------------------------------------------------|

func handlerAsciiart(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// проверить что text и банер не пустые
	// проверить что банер существует
	asciiPage, err := os.ReadFile("templates/asciiart.html")
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}
	result, err := asciiGenerator(text, banner)
	if err != nil {
		handlerError(w, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(asciiPage), result)
}

//---------------------------------------------------------------------------------------|

func handlerError(w http.ResponseWriter, errCode int) {
	fmt.Fprintf(w, errFile, errCode)
}

//---------------------------------------------------------------------------------------|

func asciiGenerator(text, banner string) (string, error) {

	asciiArt, err := asciiart.Generate(text, banner)
	if err != nil {
		return "", err
	}

	return asciiArt, nil
}
