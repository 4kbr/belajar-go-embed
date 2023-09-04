package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed kucing.jpg
var kucing []byte

//go:embed files/*txt
var path embed.FS

func main() {
	fmt.Println("version", version)

	err := os.WriteFile("file_gambar_baru.jpg", kucing, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntry, _ := path.ReadDir("files")

	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fmt.Println(entry.Name()) //print setiap nama file di directory ini
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file)) //print isi setiap file yang teks
		}
	}

}
