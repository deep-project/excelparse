package hooks

import (
	"errors"
	"fmt"
	"slices"

	"github.com/deep-project/excelparse"
)

// 直接通过数字定义某个
func HeaderRowNumberHookByNumber(n int) excelparse.HeaderRowNumberHookType {
	return func(*excelparse.Sheet) (int, error) {
		return n, nil
	}
}

// 通过关键词定位所在行(只要包含一个就可以)
func HeaderRowNumberHookByAnyKeys(keys []string, findMax int) excelparse.HeaderRowNumberHookType {
	return func(sheet *excelparse.Sheet) (_ int, err error) {
		if len(keys) == 0 {
			return 0, errors.New("keys is empty")
		}
		rows, err := sheet.ExcelFile.Rows(sheet.Name)
		if err != nil {
			return
		}
		if findMax <= 0 {
			findMax = 999
		}
		n := 0
		for rows.Next() {
			n++
			if n > findMax {
				break
			}
			row, err := rows.Columns()
			if err != nil {
				break
			}
			for _, colCell := range row {
				if colCell == "" {
					continue
				}
				if slices.Contains(keys, colCell) {
					return n, nil
				}
			}
		}
		return 0, fmt.Errorf("error locating the header row using the keywords. %s", keys)
	}
}

// 通过关键词定位所在行(必须包含全部关键词)
func HeaderRowNumberHookByAllKeys(keys []string, findMax int) excelparse.HeaderRowNumberHookType {
	return func(sheet *excelparse.Sheet) (_ int, err error) {
		if len(keys) == 0 {
			return 0, errors.New("keys is empty")
		}
		rows, err := sheet.ExcelFile.Rows(sheet.Name)
		if err != nil {
			return
		}
		if findMax <= 0 {
			findMax = 999
		}
		n := 0
		for rows.Next() {
			n++
			if n > findMax {
				break
			}
			row, err := rows.Columns()
			if err != nil {
				break
			}
			found := true
			for _, key := range keys {
				if !slices.Contains(row, key) {
					found = false
					break
				}
			}
			if found {
				return n, nil
			}
		}
		return 0, fmt.Errorf("error locating the header row using the keywords. %s", keys)
	}
}
