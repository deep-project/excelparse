package excelparse

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	Options               *Options
	ExcelFile             *excelize.File
	Name                  string     // 名称
	Visible               bool       // 是否是显示状态
	HeaderRowNumber       int        // 表头所在行(默认第一行)
	ContentBeginRowNumber int        // 表内容开始行(默认表头所在行+1)
	HeaderNames           []string   // 表头名称
	ContentRows           []TableRow // 行数据
}

func newSheet(options *Options, excelFile *excelize.File, name string) (s *Sheet, err error) {
	s = &Sheet{Options: options, ExcelFile: excelFile, Name: name}

	if err = s.parseHidden(); err != nil {
		return
	}
	if err = s.parseHeaderRowNumber(); err != nil {
		return
	}
	if err = s.parseContentRowNumber(); err != nil {
		return
	}
	if err = s.parseHeaderNames(); err != nil {
		return
	}
	err = s.parseContentRows()
	return
}

func (s *Sheet) parseHidden() (err error) {
	s.Visible, err = s.ExcelFile.GetSheetVisible(s.Name)
	return
}

// 解析表头所在行号
func (s *Sheet) parseHeaderRowNumber() (err error) {
	if s.Options.HeaderRowNumberHook != nil {
		s.HeaderRowNumber, err = s.Options.HeaderRowNumberHook(s)
	}
	if s.HeaderRowNumber <= 0 {
		s.HeaderRowNumber = 1
	}
	return
}

// 解析内容行所在行号
func (s *Sheet) parseContentRowNumber() (err error) {
	s.ContentBeginRowNumber = s.HeaderRowNumber + 1 // TODO 可以自定义hook
	return
}

// 解析列名
func (s *Sheet) parseHeaderNames() (err error) {
	rows, err := s.ExcelFile.Rows(s.Name)
	if err != nil {
		return
	}
	var i = 0
	for rows.Next() {
		i++
		cols, err := rows.Columns()
		if err != nil {
			return err
		}
		if i == s.HeaderRowNumber {
			s.HeaderNames = cols
			break
		}
	}
	if len(s.HeaderNames) == 0 {
		return errors.New("Header not found in the dataset.")
	}
	return
}

// 解析行内容
func (s *Sheet) parseContentRows() (err error) {

	rows, err := s.ExcelFile.Rows(s.Name)
	if err != nil {
		return
	}

	var rowLineNumber = 0
	for rows.Next() {
		rowLineNumber++
		// hook
		if s.Options.ParseContentRowsLoopStartHook != nil {
			continueFlag, breakFlag, err := s.Options.ParseContentRowsLoopStartHook(s, rowLineNumber)
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

		cols, err := rows.Columns()
		if err != nil {
			if s.Options.IgnoreContentRowError {
				continue
			}
			return err
		}

		// 跳过非数据行
		if rowLineNumber < s.ContentBeginRowNumber {
			continue
		}

		row := TableRow{
			Map:  make(map[string]*TableRowData),
			List: []*TableRowData{},
		}

		// 补齐列数
		diff := len(s.HeaderNames) - len(cols)
		if diff > 0 {
			cols = append(cols, make([]string, diff)...) // 补 diff 个空字符串
		}

		// 解析行内容
		for i, v := range cols {

			colLineNumber := i + 1

			rowData := &TableRowData{
				ColumnIndex: i,
				HeaderName:  s.getHeaderNameByIndex(i),
				Value:       v,
			}

			// 获取单元格类型
			//
			if s.Options.GetCellType {
				cell, err := excelize.CoordinatesToCellName(colLineNumber, rowLineNumber)
				if err == nil {
					rowData.CellType, _ = s.ExcelFile.GetCellType(s.Name, cell)
				}
			}

			// 获取图片
			if s.Options.GetCellPictures {
				cell, err := excelize.CoordinatesToCellName(colLineNumber, rowLineNumber)
				if err == nil {
					rowData.CellPictures, _ = s.ExcelFile.GetPictures(s.Name, cell)
				}
			}

			// 加入map
			if rowData.HeaderName != "" {

				if s.Options.MustExtractLast {
					row.Map[rowData.HeaderName] = rowData
				} else {
					// 这里增加一个判断，当多列名重复时
					// 如果已经存在值，判断要赋予的值是否为空，为空则不赋予
					// 这样可以保证解析的数据不为空
					_, ok := row.Map[rowData.HeaderName]
					if ok {
						if v != "" {
							row.Map[rowData.HeaderName] = rowData
						}
					} else {
						row.Map[rowData.HeaderName] = rowData
					}
				}

			}
			// 加入list
			row.List = append(row.List, rowData)
		}

		// hook
		if s.Options.ParseContentRowsLoopAppendBeforeHook != nil {
			continueFlag, breakFlag, err := s.Options.ParseContentRowsLoopAppendBeforeHook(s, rowLineNumber, &row)
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

		// 添加
		s.ContentRows = append(s.ContentRows, row)
	}

	return
}

func (s *Sheet) getHeaderNameByIndex(index int) string {
	if s.HeaderNames == nil || index < 0 || index >= len(s.HeaderNames) {
		return ""
	}
	return s.HeaderNames[index]
}
