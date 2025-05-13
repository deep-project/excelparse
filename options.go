package excelparse

import "github.com/xuri/excelize/v2"

type Options struct {
	Filepath              string
	ExcelizeOptions       excelize.Options // excelize 配置
	IgnoreContentRowError bool             // 忽略内容行解析错误
	GetCellPictures       bool             // 获取单元格图片
	GetCellType           bool             // 获取单元格类型

	ParseSheetsLoopStartHook             ParseSheetsLoopStartHookType             // 解析sheets循环开始hook
	HeaderRowNumberHook                  HeaderRowNumberHookType                  // 定位表头所在行hook
	ParseContentRowsLoopStartHook        ParseContentRowsLoopStartHookType        // 解析内容行循环开始hook
	ParseContentRowsLoopAppendBeforeHook ParseContentRowsLoopAppendBeforeHookType // 解析内容行循环添加数据之前hook

}

type (
	ParseSheetsLoopStartHookType             = func(ep *Excelparse, sjeeetIndex int, sheetName string) (continueFlag, breakFlag bool, err error)
	HeaderRowNumberHookType                  = func(*Sheet) (int, error)
	ParseContentRowsLoopStartHookType        = func(s *Sheet, rowLineNumber int) (continueFlag, breakFlag bool, err error)
	ParseContentRowsLoopAppendBeforeHookType = func(s *Sheet, rowLineNumber int, row *TableRow) (continueFlag, breakFlag bool, err error)
)
