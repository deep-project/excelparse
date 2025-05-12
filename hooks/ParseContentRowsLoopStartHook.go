package hooks

import "github.com/deep-project/excelparse"

// TODO 根据循环位置打印日志
func ParseContentRowsLoopStartHookLogger(headerName string, max int) excelparse.ParseContentRowsLoopStartHookType {
	return func(s *excelparse.Sheet, index int) (continueFlag, breakFlag bool, err error) {

		// TODO 使用hook实现
		// if i%1000 == 0 {
		// 	// logger(fmt.Sprintf("正在解析第 %d 行数据...", i))
		// }

		return false, false, nil
	}
}
