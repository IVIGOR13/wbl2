package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

var outputFileName = "index.html"

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("missing URL")
	}

	if err := wget(flag.Arg(0)); err != nil {
		log.Fatalln(err)
	}
}

func wget(url string) error {
	// get request to host
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create output file
	file, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// copy data to output file
	if _, err = io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}
