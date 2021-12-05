package simple

import "fmt"

type File struct {
	Name string `json:"name"`
}

func (f *File) Close() {
	fmt.Printf("Close File", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}

	return file, func() {
		file.Close()
	}
}
