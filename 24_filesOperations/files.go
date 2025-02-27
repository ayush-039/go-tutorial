package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)

	// }
	// FileInfo, err1 := f.Stat()
	// if err1 != nil {
	// 	panic(err1)
	// }
	// fmt.Println("file name: ", FileInfo.Name())
	// fmt.Println("file type: ", FileInfo.IsDir())
	// fmt.Println("file type: ", FileInfo.Size())
	// fmt.Println("file permission: ", FileInfo.Mode())
	// fmt.Println("file modified at: ", FileInfo.ModTime())

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	fmt.Println("Error getting file info:", err)
	// 	return
	// }

	// buf := make([]byte, fileInfo.Size())
	// _, err = f.Read(buf)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// fmt.Println("File contents:")
	// fmt.Println(string(buf))

	// data, err := os.ReadFile("example.txt") // load whole file in the ram once for all
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// Handling files system
	// dir, err := os.Open("./")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfos, err := dir.ReadDir(-1)

	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name())
	// }

	// Create the files and the folder
	// f, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// // f.WriteString("hi, I'm Ayush Barot. ")
	// // f.WriteString("you can call me Ayush Brahmbhatt")

	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	//  read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destiFile, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer destiFile.Close()
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destiFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	er := writer.WriteByte(b)
	// 	if er != nil {
	// 		panic(er)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("written to new line")

	// delete a file.
	err1 := os.Remove("example2.txt")
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("file is deleted")
}
