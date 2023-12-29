package goxl

import (
	"github.com/xuri/excelize/v2"
	"io"
)

// ParseFile opens and returns the Excel file from the given path
func ParseFile(filePath string) (*File, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	return &File{
		f: f,
	}, nil
}

// ParseReader opens and returns the Excel file from the given reader
func ParseReader(r io.Reader) (*File, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}

	return &File{
		f: f,
	}, nil
}
