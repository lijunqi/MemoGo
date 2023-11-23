// main.go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func ZipFile() {
	fmt.Println("creating zip archive")
	//Create a new zip archive and named archive.zip
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

	// Traverse directory
	err = filepath.WalkDir("target_folder", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.Name(), "directory?", d.IsDir())
		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}

	//Create a new zip writer
	zipWriter := zip.NewWriter(archive)
	fmt.Println("opening first file")
	//Add files to the zip archive
	f1, err := os.Open("qq/")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("adding file to archive..")
	w1, err := zipWriter.Create("aaa")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}
	fmt.Println("closing archive")
	zipWriter.Close()
}

func main() {
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
			//return nil
		}
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			_, err = io.Copy(f, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	}
	err = filepath.Walk("target", walker)
	if err != nil {
		panic(err)
	}
}
