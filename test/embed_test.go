package test

// import (
// 	"embed"
// 	_ "embed"
// 	"fmt"
// 	"io/fs"
// 	"os"
// 	"testing"
// )

// //cara pakai go embed, pakai komen kaya ini:

// //go:embed version.txt
// var version string

// // tanpa komen di go:embed version.txt
// var version2 string

// func TestEmbedTXT(t *testing.T) {
// 	fmt.Println("ini dari file version:", version)
// 	fmt.Println("ini dari file version ke 2: ", version2)
// }

// //embed file image

// //go:embed kucing.jpg
// var kucing []byte

// func TestByte(t *testing.T) {
// 	err := os.WriteFile("file_gambar_baru.jpg", kucing, fs.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Sukses buat file gambar")

// }

// //embed multiple files

// //go:embed files/a.txt
// //go:embed files/b.txt
// //go:embed files/c.txt
// var files embed.FS

// func TestMultipleFiles(t *testing.T) {
// 	a, _ := files.ReadFile("files/a.txt")
// 	fmt.Println("a:", string(a)) //aaaaa
// 	b, _ := files.ReadFile("files/b.txt")
// 	fmt.Println("a:", string(b)) //bbbbb
// 	c, _ := files.ReadFile("files/c.txt")
// 	fmt.Println("a:", string(c)) //ccccc
// }

// // bisa juga :

// //go:embed files/*.txt
// var path embed.FS

// func TestPathMatcher(t *testing.T) {
// 	dirEntry, _ := path.ReadDir("files")

// 	for _, entry := range dirEntry {
// 		if !entry.IsDir() {
// 			fmt.Println(entry.Name()) //print setiap nama file di directory ini
// 			file, _ := path.ReadFile("files/" + entry.Name())
// 			fmt.Println(string(file)) //print isi setiap file yang teks
// 		}
// 	}
// }
