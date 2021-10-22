// +build protokolFile

package service

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	EXAMPLE_FOLDER_PATH = "../../../../test_expected"
	EXAMPLE_FOLDER_NAME = "test_expected"
	PARENT_EXAMPLE_FOLDER_NAME = "legaliosgo"
)

func createLoggerFile(baseName string, fileName string) (f *os.File, err error) {
	currPath, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	for  {
		_, lastName := filepath.Split(currPath)
		if lastName == PARENT_EXAMPLE_FOLDER_NAME || currPath == "." || currPath == "/" {
			break
		}
		currPath = filepath.Dir(currPath)
	}
	basePath := filepath.Join(currPath, EXAMPLE_FOLDER_NAME)
	filePath := filepath.Join(basePath, fileName)
	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	f, err = os.Create(fullPath)
	return f, err
}

func logTestStart(protokol *os.File) {
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

func logTestEnd(protokol *os.File) {
	if  _, err := fmt.Fprintf(protokol,"\n"); err != nil {
		return
	}
}

func logTestYear(protokol *os.File, value string) {
	if  _, err := fmt.Fprintf(protokol,"%s", value); err != nil {
		return
	}
}

func logExampleIntValue(protokol *os.File, value int32) {
	if  _, err := fmt.Fprintf(protokol,"\t%d", value); err != nil {
		return
	}
}

func logExampleDecValue(protokol *os.File, value float64) {
	var intValue = int(value*100)
	if  _, err := fmt.Fprintf(protokol,"\t%d", intValue); err != nil {
		return
	}
}

func logTestIntExamples(t *testing.T, fileName string, examples []testIntScenario) {
	testProtokol, err := createLoggerFile(EXAMPLE_FOLDER_PATH, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	logTestStart(testProtokol)

	for _, tx := range examples {
		logTestYear(testProtokol, tx.title)

		for _, tt := range tx.tests {
			logExampleIntValue(testProtokol, tt.resultValue)
		}
		logTestEnd(testProtokol)
	}
}

func logTestDecExamples(t *testing.T, fileName string, examples []testDecScenario) {
	testProtokol, err := createLoggerFile(EXAMPLE_FOLDER_PATH, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	logTestStart(testProtokol)

	for _, tx := range examples {
		logTestYear(testProtokol, tx.title)

		for _, tt := range tx.tests {
			logExampleDecValue(testProtokol, tt.resultValue)
		}
		logTestEnd(testProtokol)
	}
}
