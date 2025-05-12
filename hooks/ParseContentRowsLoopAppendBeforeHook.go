package hooks

import "github.com/deep-project/excelparse"

// TODO 根据某个列名数据是否多次为空，判断是否break循环
func ParseContentRowsLoopAppendBeforeHookIsEmptyBreak(headerName string, max int) excelparse.ParseContentRowsLoopAppendBeforeHookType {
	return func(s *excelparse.Sheet, index int, row *excelparse.TableRow) (continueFlag, breakFlag bool, err error) {
		return false, false, nil
	}
}

// TODO 根据某个列名数据是否在黑名单内，continue循环
func ParseContentRowsLoopAppendBeforeHookContainsKeysContinue(headerName string, blacklist []string) excelparse.ParseContentRowsLoopAppendBeforeHookType {
	return func(s *excelparse.Sheet, index int, row *excelparse.TableRow) (continueFlag, breakFlag bool, err error) {
		return false, false, nil
	}
}
