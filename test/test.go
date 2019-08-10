package main

import (
	"fmt"
	"github.com/smcduck/xsys/xfileformat"
	"github.com/smcduck/xsys/xfs"
	"github.com/smcduck/xsys/xproc"
	"path/filepath"
)

func main()  {
	myfolder, err := xproc.GetMyFolder()
	if err != nil {
		fmt.Println(err)
		return
	}
	failedCount := 0

	execDir := filepath.Join(myfolder, "testdata-is-executable")
	_, execList, err := xfs.ListDir(execDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range execList {
		is, err := xfileformat.IsExecutable(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !is {
			fmt.Println(v, "is executable file, api says it is NOT")
			failedCount++
		}
	}

	nonExecDir := filepath.Join(myfolder, "testdata-not-executable")
	_, nonExecList, err := xfs.ListDir(nonExecDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range nonExecList {
		is, err := xfileformat.IsExecutable(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		if is {
			fmt.Println(v, "is NOT executable file, api says it is")
			failedCount++
		}
	}

	if failedCount > 0 {
		fmt.Println("test failed!")
	} else {
		fmt.Println("test passed")
	}
}