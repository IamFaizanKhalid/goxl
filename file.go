package goxl

import (
	"github.com/xuri/excelize/v2"
)

// File represents an Excel file.
type File struct {
	f *excelize.File
}

// GetDefaultSheet opens and returns the first sheet
func (f *File) GetDefaultSheet() (*Sheet, error) {
	return f.GetSheetByIndex(0)
}

// GetSheet opens and returns the sheet with the given name
func (f *File) GetSheet(name string) (*Sheet, error) {
	r, err := f.f.Rows(name)
	if err != nil {
		return nil, err
	}

	return &Sheet{
		r: r,
	}, nil
}

// GetSheetByIndex opens and returns the sheet with the given index
func (f *File) GetSheetByIndex(index int) (*Sheet, error) {
	name := f.f.GetSheetName(index)
	if name == "" {
		return nil, excelize.ErrSheetNotExist{}
	}
	return f.GetSheet(name)
}

// Close closes the Excel file
func (f *File) Close() error {
	return f.f.Close()
}
