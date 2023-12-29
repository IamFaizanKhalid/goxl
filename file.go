package goxl

import (
	"github.com/xuri/excelize/v2"
)

// File represents an Excel file.
type File struct {
	f *excelize.File
}

// OpenDefaultSheet opens and returns the first sheet
func (f *File) OpenDefaultSheet() (*Sheet, error) {
	return f.OpenSheetByIndex(0)
}

// OpenSheet opens and returns the sheet with the given name
func (f *File) OpenSheet(name string) (*Sheet, error) {
	r, err := f.f.Rows(name)
	if err != nil {
		return nil, err
	}

	return &Sheet{
		r: r,
	}, nil
}

// OpenSheetByIndex opens and returns the sheet with the given index
func (f *File) OpenSheetByIndex(index int) (*Sheet, error) {
	name := f.f.GetSheetName(index)
	if name == "" {
		return nil, excelize.ErrSheetNotExist{}
	}
	return f.OpenSheet(name)
}

// Close closes the Excel file
func (f *File) Close() error {
	return f.f.Close()
}
