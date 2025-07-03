package hooks

import (
	"fmt"

	"github.com/deep-project/excelparse"
)

func ParseContentRowsLoopStartHookLogger(gap int, logger func(string)) excelparse.ParseContentRowsLoopStartHookType {
	if gap == 0 {
		gap = 1000
	}
	return func(s *excelparse.Sheet, rowLineNumber int) (continueFlag, breakFlag bool, err error) {
		if rowLineNumber%gap == 0 {
			logger(fmt.Sprintf("正在解析 %s 第 %d 行数据...", s.Name, rowLineNumber))
		}
		return false, false, nil
	}
}
