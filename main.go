package main

import (
	"encoding/json"
	"live/assets"
	"live/pdfproc"
	"log"
)

func main() {
	asts, err := assets.New("assets")
	if err != nil {
		log.Fatalf("Error assets: %v", err)
	}
	tmplDatamatrixJson, err := asts.Json("datamatrix")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	tmplBarJson, err := asts.Json("bar")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	tmplDatamatrix := &pdfproc.MarkTemplate{}
	err = json.Unmarshal(tmplDatamatrixJson, tmplDatamatrix)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v", err)
	}
	tmplBar := &pdfproc.MarkTemplate{}
	err = json.Unmarshal(tmplBarJson, tmplBar)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v", err)
	}
	pdf, err := pdfproc.New(tmplDatamatrix, tmplBar, asts)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v", err)
	}
	err = pdf.PdfDocument()
	if err != nil {
		log.Fatalf("Error unmarshal file: %v", err)
	}
}
