package hooks

import "github.com/deep-project/excelparse"

// 只处理第一个sheet页
func ParseSheetsLoopStartHookOnlyFirstSheet() excelparse.ParseSheetsLoopStartHookType {
	return func(ep *excelparse.Excelparse, index int, sheetName string) (continueFlag, breakFlag bool, err error) {
		if index > 0 {
			return false, true, nil
		}
		return false, false, nil
	}
}
