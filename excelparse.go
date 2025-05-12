package excelparse

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Excelparse struct {
	options   *Options
	excelFile *excelize.File
	sheets    []Sheet
}

func New(opt *Options) *Excelparse {
	return &Excelparse{options: opt}
}

func (w *Excelparse) Run() (err error) {
	excelFile, err := excelize.OpenFile(w.options.Filepath, w.options.ExcelizeOptions)
	if err != nil {
		return
	}
	defer excelFile.Close()
	w.excelFile = excelFile
	if err = w.parseSheets(); err != nil {
		return
	}
	return
}

func (w *Excelparse) parseSheets() (err error) {
	for _, name := range w.excelFile.GetSheetList() {
		if err = w.parseSheet(name); err != nil {
			return fmt.Errorf("failed to parse sheet \"%s\". %s", name, err.Error())
		}
		if w.options.OnlyParseFirstSheet {
			break
		}
	}
	return
}

func (w *Excelparse) parseSheet(name string) (err error) {
	sheet, err := newSheet(w.options, w.excelFile, name)
	w.sheets = append(w.sheets, *sheet)
	return
}
