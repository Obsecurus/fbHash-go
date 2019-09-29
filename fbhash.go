package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Viking2012/fbHash/hashing"
)

func main() {
	log.SetPrefix("fbHash ")
	absPath, _ := filepath.Abs("data/filetoread.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Println("Got an unresolved error when trying to open a file")
		log.Fatal(err)
		return
	}
	defer file.Close()

	fileCounts := hashing.RollingHashes(file)
	fmt.Println(hashing.Indent(fileCounts))
}
