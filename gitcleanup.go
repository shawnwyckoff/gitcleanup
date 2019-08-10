package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/smcduck/xsys/xfileformat"
	"github.com/smcduck/xsys/xfs"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("dir needed, like this:\ncef which-dir-to-detect whether-remove")
		return
	}
	dir := os.Args[1]

	_, files, err := xfs.ListDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	mpb := pb.New(len(files))
	mpb.Start()
	execList := []string{}

	type errItem struct{
		filename string
		err error
	}

	errList := []errItem{}

	for _, v := range files {
		if strings.Contains(v, "/.git/") || strings.Contains(v, ".DS_Store") { // ignore .git files
			continue
		}
		isExec, err := xfileformat.IsExecutable(v)
		if err != nil {
			errList = append(errList, errItem{filename:v, err:err})
			continue
		}
		if isExec {
			execList = append(execList, v)
		}
		mpb.Add(1)
	}
	mpb.Finish()

	if len(errList) > 0 {
		fmt.Println("------- errors: -------")
		for _, v := range errList {
			fmt.Println(v.filename, v.err)
		}
		fmt.Println("")
	}

	if len(execList) == 0 {
		fmt.Println("no executable file found")
		return
	} else {
		fmt.Println(len(execList), "executable files found:")
		fmt.Println(strings.Join(execList, "\n"))
	}

	prompt := promptui.Select{
		Label: "Select Whether Remove Detected Executable Files",
		Items: []string{"yes", "no"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %q\n", result)
	if strings.ToLower(result) == "yes" {
		for _, v := range execList {
			if err := os.Remove(v); err != nil {
				fmt.Println(v, err)
			} else {
				fmt.Println(v, "removed!")
			}
		}
	}
}
