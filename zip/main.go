// main.go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ZipFile(fPath string) {
	fmt.Println("creating zip archive")
	// Create zip archive
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
		// this is to catch errors if any
	}
	defer archive.Close()
	fmt.Println("archive file created successfully....")
	//we use the defer key to close it, once we create an archive we need to close it using the defer keyword
	defer archive.Close()
	fmt.Println("archive file created successfully")

	// Create zip writer
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()
	fmt.Println("opening first file")

	// Add files to the zip archive
	f1, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("adding file to archive..")
	w1, err := zipWriter.Create("inzip.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}
	fmt.Println("closing archive")
}

func ZipFolder() {
	zf, err := os.Create("output.zip")
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	w := zip.NewWriter(zf)
	defer w.Close()

	walker := func(path string, fi os.FileInfo, err error) error {
		log.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			path = fmt.Sprintf("%s%c", path, os.PathSeparator)
			_, err := w.Create(path)
			return err
		} else {
			// Ensure that `path` is not absolute; it should not start with "/".
			// This snippet happens to work because I don't use
			// absolute paths, but ensure your real-world code
			// transforms path into a zip-root relative path.
			fInZip, err := w.Create(path)
			if err != nil {
				return err
			}

			fSrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()

			_, err = io.Copy(fInZip, fSrc)
			if err != nil {
				return err
			}
		}

		return nil
	}

	root := "C:\\Users\\JLi21\\Desktop\\MemoGo\\zip\\target"
	err = filepath.Walk(root, walker)
	if err != nil {
		panic(err)
	}
}

func main() {
	fPath := "C:\\Users\\JLi21\\Desktop\\MemoGo\\zip\\target.txt"
	ZipFile(fPath)
	//ZipFolder()
}
