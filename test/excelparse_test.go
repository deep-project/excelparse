package excelparse

import (
	"fmt"
	"testing"

	"github.com/deep-project/excelparse"
	"github.com/deep-project/excelparse/hooks"
)

func TestMain(t *testing.T) {
	ex := excelparse.New(&excelparse.Options{
		Filepath:                      "path",
		HeaderRowNumberHook:           hooks.HeaderRowNumberHookByAllKeys([]string{"存货名称", "货号"}, 20),
		ParseSheetsLoopStartHook:      hooks.ParseSheetsLoopStartHookOnlyFirstSheet(),
		ParseContentRowsLoopStartHook: hooks.ParseContentRowsLoopStartHookLogger(1000, func(val string) { fmt.Println(val) }),
	})

	if err := ex.Run(); err != nil {
		t.Error(err)
		return
	}
	for _, sheet := range ex.Sheets {
		for _, row := range sheet.ContentRows {
			fmt.Printf("row:%#v", row)
			for K, V := range row.Map {
				fmt.Println(K, "----", V)
			}
		}
	}

}
