package excelparse

import "github.com/xuri/excelize/v2"

type Options struct {
	Filepath              string
	ExcelizeOptions       excelize.Options // excelize 配置
	IgnoreContentRowError bool             // 忽略内容行解析错误
	OnlyParseFirstSheet   bool             // 直解析第一个sheet页

	HeaderRowNumberHook HeaderRowNumberHookType // 定位表头所在行hook

	ParseContentRowsLoopStartHook        ParseContentRowsLoopStartHookType        // 解析内容行循环开始hook
	ParseContentRowsLoopAppendBeforeHook ParseContentRowsLoopAppendBeforeHookType // 解析内容行循环添加数据之前hook

	// DataStartLine int // 数据起始行
	// DataEndLine   int // 数据结束行

	// IgnoreRowParseErr     bool     // 跳过行解析失败
	// UseColumnNameIndex    bool     // 使用列名做为结果索引
	// MustContainColumnName []string // 必须包含的列名

	// AutoPositioningColumnNameLine bool     // 是否自动定位列名所在行
	// PositioningColumnNameLineKey  []string // 定位时所用的key

	// columnNameIndexList map[int]string // 列名索引列表

	// RunRowsBeforeHook func(*Options, *excelize.Rows)
	// Logger            Logger // 日志接口
}

type (
	HeaderRowNumberHookType                  = func(*Sheet) (int, error)
	ParseContentRowsLoopStartHookType        = func(s *Sheet, index int) (continueFlag, breakFlag bool, err error)
	ParseContentRowsLoopAppendBeforeHookType = func(s *Sheet, index int, row *TableRow) (continueFlag, breakFlag bool, err error)
)
