package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/bon-ami/eztools"

	"github.com/unidoc/unioffice/spreadsheet"
)

func wrH(name string) *os.File {
	eztools.ShowStrln("Processing sheet " + name)
	file := "values-" + name
	if err := os.Mkdir(file, os.ModeDir); err != nil {
		eztools.LogErrPrint(err)
		return nil
	}
	fp, err := os.Create(filepath.Join(file, "strings.xml"))
	if err != nil {
		eztools.LogErrPrint(err)
		return nil
	}
	fp.WriteString("<?xml version='1.0' encoding='utf-8'?>\n")
	fp.WriteString("<resources xmlns:android=\"http://schemas.android.com/apk/res/android\" xmlns:tools=\"http://schemas.android.com/tools\" xmlns:xliff=\"urn:oasis:names:tc:xliff:document:1.2\">\n")
	return fp
}

func wrB(fp *os.File, nm, vl string) {
	//eztools.ShowStrln("\t" + nm + "=" + vl)
	if fp != nil {
		fp.WriteString("<string name=\"" + nm + "\">" + vl + "</string>\n")
	}
}

func wrT(fp *os.File) {
	if fp != nil {
		fp.WriteString("</resources>")
		fp.Close()
	}
}

func rd(file string) (err error) {
	wb, err := spreadsheet.Open(file)
	if err != nil {
		eztools.LogErrPrint(err)
		return
	}
	defer wb.Close()
	for _, sh := range wb.Sheets() {
		nm := sh.Name()
		fp := wrH(nm)
		if fp == nil {
			break
		}
		for i, ro := range sh.Rows() {
			if i == 0 {
				continue
			}
			it, err := strconv.Atoi(*(ro.Cell("F").X().V))
			if err != nil {
				eztools.LogErrPrint(err)
				continue
			}
			if it < 0 || it >= len(wb.SharedStrings.X().Si) {
				eztools.LogPrint("invalid index " + *(ro.Cell("F").X().V))
				continue
			}
			for _, vr := range wb.SharedStrings.X().Si[it].R {
				if vr != nil {
					wrB(fp, ro.Cell("B").GetString(), vr.T)
				} else {
					eztools.LogPrint("no value for " + nm + "," + ro.Cell("B").GetString())
				}
			}
			/*str, err := wb.SharedStrings.GetString(it)
			//str, err := ro.Cell("F").GetRawValue()
			if err == nil {
				if len(str) > 0 {
					fmt.Printf("%s\t%x\n", str, str)
				} else {
					fmt.Println(*(ro.Cell("F").X().RAttr))
				}
			} else {
				eztools.LogErrPrint(err)
			}*/
		}
		wrT(fp)
	}
	return
}
