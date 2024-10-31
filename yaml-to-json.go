package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	filename := flag.String("f", "", "filename")
	flag.Parse()
	if filename == nil || *filename == "" {
		log.Fatalf("filename is empty")
	}

	ext := filepath.Ext(*filename)
	if !(ext == ".yaml" || ext == ".yml") {
		log.Fatalf("file extension is wrong")
	}

	yamlFile, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	f := map[string]interface{}{}
	err = yaml.Unmarshal(yamlFile, &f)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	jsonBytes, err := json.Marshal(f)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(string(jsonBytes))
	// err = os.WriteFile("test.json", jsonBytes, 0644)
	// if err != nil {
	// 	log.Fatalf("failed to write file: %v", err)
	// }
}
