package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dop251/goja"
)

func readFile(file *zip.File) {
	// Open the file inside the zip
	rc, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	// Read the contents of the file
	contents, err := ioutil.ReadAll(rc) // deprecated :/
	if err != nil {
		log.Fatal(err)
	}

	// Regular expressions to replace specific patterns
	reConfig := regexp.MustCompile(`window\.\w+\s*=\s*{`)
	reArray := regexp.MustCompile(`window\.\w+\.\w+\.\w+\s*=\s*\[`)

	// Replace patterns in the content
	processedContents := reConfig.ReplaceAllStringFunc(string(contents), func(s string) string {
		return "var data = {"
	})
	processedContents = reArray.ReplaceAllStringFunc(processedContents, func(s string) string {
		return "var data = ["
	})

	// Parse the JavaScript file using goja
	vm := goja.New()
	_, err = vm.RunString(processedContents)
	if err != nil {
		log.Fatalf("Error parsing JS file: %v", err)
	}

	// Retrieve the value of the 'data' variable from the JavaScript context
	value := vm.Get("data")
	if value == nil {
		log.Fatalf("No data variable found in the JS file")
	}

	// Convert the data to a Go-native type
	data := value.Export()

	// Marshal the Go-native type to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling data to JSON: %v", err)
	}

	// Output the JSON data
	fmt.Println(string(jsonData))
}

func run(path string) {
	// Open the zip file
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the zip archive
	for _, f := range r.File {
		// Check if the file is in the /data directory and has a .js extension
		if strings.HasPrefix(f.Name, "data/") && strings.ToLower(filepath.Ext(f.Name)) == ".js" {
			readFile(f)
			// return // Exit after processing the first .js file
		}
	}
}

func main() {
	// Example usage
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the zip file as an argument.")
	}

	path := os.Args[1]
	run(path)
}
