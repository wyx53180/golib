package goos

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

type File struct {
	file *os.File
	mod  string
	path string
}

func Open(path, mod string) *File {
	var f *os.File
	var err error
	switch mod {
	case "w":
		f, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	case "a":
		f, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0777)
	case "r":
		f, err = os.OpenFile(path, os.O_RDONLY, 0777)
	default:
		f, err = os.OpenFile(path, os.O_RDONLY, 0777)
	}
	checkerr(err)
	return &File{file: f,
		mod:  mod,
		path: path}
}

func (f *File) Close() {
	f.file.Close()
}

// ReadLine function.
func (f *File) ReadLines() chan string {
	buff := bufio.NewReader(f.file)
	ch := make(chan string)
	go func() {
		for {
			data, _, err := buff.ReadLine()
			if err == io.EOF {
				break
			} else {
				checkerr(err)
			}
			ch <- string(data)
		}
		close(ch)
	}()
	return ch
}

func (f *File) Write(data string) {
	f.file.WriteString(data)
}

// Read func.
func (f *File) Read() string {
	// file, err := ioutil.ReadFile(path)
	// checkerr(err)
	// return string(file)

	// file, err := os.Open(path)
	// checkerr(err)
	content, err := ioutil.ReadAll(f.file)
	checkerr(err)
	return string(content)
}

/*--------walk--------------*/
// ObFile struct.
type ObFile struct {
	Path   string
	Folder []string
	File   []string
}

// Walk func.
func Walk(fPath string) chan ObFile {
	_, err := os.Stat(fPath)
	if err != nil {
		panic("err path.")
	}
	ch := make(chan ObFile)
	go walkFunc(fPath, ch)
	return ch
}

func walkFunc(fPath string, ch chan ObFile) {
	filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			var fileStruct ObFile
			path, _ := filepath.Abs(path)
			fileStruct.Path = path
			files, err := ioutil.ReadDir(path)
			checkerr(err)
			for _, file := range files {
				if file.IsDir() {
					fileStruct.Folder = append(fileStruct.Folder, file.Name())
				} else {
					fileStruct.File = append(fileStruct.File, file.Name())
				}
			}
			ch <- fileStruct
		}
		return nil
	})
	close(ch)
}

// func main() {
// 	// w
// 	f := Open("asdf", "w")
// 	defer f.Close()
// 	f.Write("asdf1234\n")
// 	f.Write("asdf1234\n")

// 	// r
// 	f = Open("asdf", "r")
// 	fmt.Print(f.Read())

// 	// readlines
// 	f = Open("asdf", "r")
// 	for i := range f.ReadLines() {
// 		fmt.Println(i)
// 	}

// 	// walk
// 	for i := range Walk("./") {
// 		fmt.Println(i)
// 	}
// }
