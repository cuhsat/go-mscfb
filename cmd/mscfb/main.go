package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cuhsat/go-mscfb"
)

func main() {
	files, err := filepath.Glob("*.msi")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		rdr, err := os.OpenFile(file, os.O_RDONLY, 0)
		if err != nil {
			panic(err)
		}

		msi, err := mscfb.Open(rdr, mscfb.ValidationPermissive)
		if err != nil {
			panic(err)
		}

		fmt.Println(msi)
	}
}
