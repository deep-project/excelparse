# Excel数据解析器 / Excel Data Parser

## 介绍 / Introduction

工具的作用是自动把excel文件解析成结构化的数据

The tool automatically parses Excel file into structured data.


## 特点 / Features

- 自动识别表头所在行 / Automatically detects the header row
- 可获取表格中的图片 / Supports extracting images from the sheet
- 可解析多个sheet页的数据 / Parses data across multiple sheets
- 埋点多个hooks / Provides multiple hooks for customization



## 使用 / Usage
```go

import (
    "github.com/deep-project/excelparse"
)

func main(){
  ex := excelparse.New(&excelparse.Options{
		Filepath: "xlsx file path",
	})

	if err = ex.Run(); err != nil {
		return
	}

  for _, sheet := range ex.Sheets {
		for _, row := range sheet.ContentRows {
      .....
		}
	}
}
```

## 感谢 / Acknowledgements

- [excelize](https://github.com/xuri/excelize/v2) 
