package excelparse

import (
	"regexp"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type TableRow struct {
	Map  map[string]*TableRowData // 表头名称作为key索引
	List []*TableRowData          // 切片数据
}

// 通过列表获取单元格数据
func (t *TableRow) Get(columnName string) (res *TableRowData, exist bool) {
	res, exist = t.Map[columnName]
	return
}

type TableRowData struct {
	ColumnIndex  int                // 列索引
	HeaderName   string             // 表头名称
	Value        string             // 单元格内容
	CellPictures []excelize.Picture // TODO 单元格图片
	CellType     excelize.CellType  // TODO 单元格类型
}

func (d *TableRowData) String() string {
	return d.Value
}

// 获取浮点数
func (d *TableRowData) Float64() (res float64) {
	res, _ = strconv.ParseFloat(d.FormatNumberStr(), 64)
	return
}

// 获取整数
func (d *TableRowData) Int() (res int) {
	return int(d.Float64())
}

// 格式化数字字符串
func (d *TableRowData) FormatNumberStr() string {
	return regexp.MustCompile(`[^0-9.-]`).ReplaceAllString(d.Value, "")
}
