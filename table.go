package excelparse

type TableRow struct {
	Map  map[string]*TableRowData // 表头名称作为key索引
	List []*TableRowData          // 切片数据
}

// 通过列表获取单元格数据
func (t *TableRow) Get(columnName string) *TableRowData {
	return &TableRowData{}
}

type TableRowData struct {
	ColumnIndex int    // 列索引
	HeaderName  string // 表头名称
	Value       string // 结果
}
