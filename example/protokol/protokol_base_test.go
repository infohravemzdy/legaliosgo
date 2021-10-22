// +build protokolFile

package protokol

import (
	"fmt"
	. "github.com/shopspring/decimal"
	"os"
	"path/filepath"
)

const (
	PROTOKOL_FOLDER_PATH = "../../../../test_values"
	PROTOKOL_FOLDER_NAME = "test_values"
	PARENT_PROTOKOL_FOLDER_NAME = "legaliosgo"
)

func createProtokolFile(baseName string, fileName string) (f *os.File, err error) {
	currPath, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	for  {
		_, lastName := filepath.Split(currPath)
		if lastName == PARENT_PROTOKOL_FOLDER_NAME || currPath == "." || currPath == "/" {
			break
		}
		currPath = filepath.Dir(currPath)
	}
	basePath := filepath.Join(currPath, PROTOKOL_FOLDER_NAME)
	filePath := filepath.Join(basePath, fileName)
	fullPath, err := filepath.Abs(filePath)
	println(fullPath)
	if err != nil {
		return nil, err
	}
	f, err = os.Create(fullPath)
	return f, err
}

func exportPropsIntValue(protokol *os.File, value int32) {
	if  _, err := fmt.Fprintf(protokol,"\t%d", value); err != nil {
		return
	}
}

func exportPropsDecValue(protokol *os.File, value Decimal) {
	var intValue = value.Mul(NewFromInt32(100)).IntPart()
	//valueFloat, _ := value.Float64()
	if  _, err := fmt.Fprintf(protokol,"\t%d", intValue); err != nil {
		return
	}
}

func exportPropsStart(protokol *os.File) {
	if  _, err := fmt.Fprintf(protokol,""); err != nil {
		return
	}
	if  _, err := fmt.Fprintf(protokol,"YEAR"); err != nil {
		return
	}
	for testMonth := 1; testMonth <=12; testMonth++ {
		if  _, err := fmt.Fprintf(protokol,"\t%d", testMonth); err != nil {
			return
		}
	}
	if  _, err := fmt.Fprintf(protokol,"\n"); err != nil {
		return
	}
}

func exportPropsEnd(protokol *os.File) {
	if  _, err := fmt.Fprintf(protokol,"\n"); err != nil {
		return
	}
}

func exportPropsYear(protokol *os.File, value int16) {
	if  _, err := fmt.Fprintf(protokol,"%d", value); err != nil {
		return
	}
}
