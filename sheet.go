package goxl

import (
	"github.com/xuri/excelize/v2"
)

// Sheet represents an Excel sheet.
type Sheet struct {
	r *excelize.Rows
}

// Next returns true if there are more rows to read.
func (r *Sheet) Next() bool {
	return r.r.Next()
}

// Read reads the next row from the sheet.
func (r *Sheet) Read() (record []string, err error) {
	return r.r.Columns()
}

// ReadAll reads all the remaining rows from the sheet.
func (r *Sheet) ReadAll() (records [][]string, err error) {
	for r.Next() {
		record, err := r.Read()
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return
}

// Close closes the sheet.
func (r *Sheet) Close() error {
	return r.r.Close()
}

// Error returns the error, if any, that was encountered during iteration.
func (r *Sheet) Error() error {
	return r.r.Error()
}
