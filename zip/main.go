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

func ZipFolder(srcFolder, dstZipFile string) {
	srcParentFolder, srcFile := filepath.Split(srcFolder)
	srcFolder = filepath.Join(srcParentFolder, srcFile) + string(os.PathSeparator)
	startIdx := len(srcFolder)
	log.Printf("=== Start Index: %d\n", startIdx)
	log.Printf("=== src folder: %s\n", srcFolder)
	zf, err := os.Create(dstZipFile)
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	zipWriter := zip.NewWriter(zf)
	defer zipWriter.Close()

	walker := func(path string, fi os.FileInfo, err error) error {
		log.Printf("dst is %s\n", dstZipFile)
		log.Printf("Crawling: %#v\n", path)
		if err != nil {
			log.Printf("xx Crawling error: %v\n", err)
			return err
		}

		if fi.IsDir() {
			path = fmt.Sprintf("%s%c", path, os.PathSeparator)
			log.Printf("------ path: %s\n", path)
			log.Printf("------ path len: %d\n", len(path))
			log.Printf("  Create Folder zipWriter: %s\n", path[startIdx:])
			_, err := zipWriter.Create(path[startIdx:])
			if err != nil {
				log.Printf("xx zipWriter error: %v\n", err)
			}
			return err
		} else {
			// Ensure that `path` is not absolute; it should not start with "/".
			// This snippet happens to work because I don't use
			// absolute paths, but ensure your real-world code
			// transforms path into a zip-root relative path.
			fInZip, err := zipWriter.Create(path[startIdx:])
			log.Printf("    zipwriter create: %s\n", path)
			if err != nil {
				log.Printf("xxx inzip error: %v\n", err)
				return err
			}

			fSrc, err := os.Open(path)
			if err != nil {
				log.Printf("xxx Open src error: %v\n", err)
				return err
			}
			defer fSrc.Close()

			var written int64
			written, err = io.Copy(fInZip, fSrc)
			if err != nil {
				log.Printf("xxx Copy to zip error: %v\n", err)
				return err
			}
			log.Printf("... written = %d\n", written)
		}

		return nil
	}

	err = filepath.Walk(srcFolder, walker)
	if err != nil {
		panic(err)
	}
}

func main() {
	//fPath := "C:\\Users\\JLi21\\Desktop\\MemoGo\\zip\\target.txt"
	//ZipFile(fPath)

	f := filepath.Join("C:\\a", "b", "c", ".zip")
	log.Println(f)

	//ZipFolder("C:\\Users\\jli21\\Desktop\\MemoGo\\zip\\target", "C:\\Users\\Public\\1.zip")
	//ZipFolder("C:\\Users\\jli21\\Desktop\\MemoGo\\zip\\target\\", "C:\\Users\\Public\\2.zip")
	//ZipFolder("C:\\Users\\jli21\\Desktop\\MemoGo\\zip\\target/", "C:\\Users\\Public\\3.zip")
	//ZipFolder("target", "C:\\Users\\Public\\4.zip")
	//ZipFolder("target\\", "C:\\Users\\Public\\5.zip")
}
