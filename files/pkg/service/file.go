package service

import (
	"os"
	"path/filepath"
)

type File struct {
	FilePath   string
	fileName   string
	OutputFile *os.File
}

func NewFile() *File {
	return &File{}
}

func (f *File) SetFile(fileName, path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	f.fileName = fileName
	f.FilePath = filepath.Join(path, fileName)

	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}

	f.OutputFile = file

	return nil

}

func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}
	if _, err := f.OutputFile.Write(chunk); err != nil {
		return err
	}

	return nil
}

func (f *File) Close() error {
	return f.OutputFile.Close()
}
