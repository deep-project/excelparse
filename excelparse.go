package excelparse

import (
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Excelparse struct {
	Sheets    []Sheet
	options   *Options
	excelFile *excelize.File
}

func New(opt *Options) *Excelparse {
	return &Excelparse{options: opt}
}

func (e *Excelparse) Run() (err error) {
	if e.options.Filepath == "" {
		return errors.New("file path not defined")
	}
	excelFile, err := excelize.OpenFile(e.options.Filepath, e.options.ExcelizeOptions)
	if err != nil {
		return
	}
	defer excelFile.Close()
	e.excelFile = excelFile
	if err = e.parseSheets(); err != nil {
		return
	}
	return
}

func (e *Excelparse) parseSheets() (err error) {
	for i, name := range e.excelFile.GetSheetList() {

		// hook
		if e.options.ParseSheetsLoopStartHook != nil {
			continueFlag, breakFlag, err := e.options.ParseSheetsLoopStartHook(e, i, name)
			if err != nil {
				return err
			}
			if breakFlag {
				break
			}
			if continueFlag {
				continue
			}
		}
		if err = e.parseSheet(name); err != nil {
			return fmt.Errorf("failed to parse sheet \"%s\". %s", name, err.Error())
		}
	}
	return
}

func (e *Excelparse) parseSheet(name string) (err error) {
	sheet, err := newSheet(e.options, e.excelFile, name)
	e.Sheets = append(e.Sheets, *sheet)
	return
}
