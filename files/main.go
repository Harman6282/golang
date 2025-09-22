package main

import (
	// "fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name: ", fileInfo.Name())
	// fmt.Println("is folder", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())

	//? read file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 18)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++{

	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// ? /////////////////////////////////////////////////

	//    dir , err := os.Open(".")

	//    if err != nil {
	// 	panic(err)
	//    }

	//    defer dir.Close()

	//    fileInfo, err := dir.ReadDir(-1)

	//    for _, fi := range fileInfo {
	// 	   fmt.Println(fi.Name())
	// 	}

	//? create file

	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// f.WriteString("Hi Harman")
	// f.WriteString("this is ubuntu linux")

	// bytes := []byte("Hello world")

	// f.Write(bytes)


}
