// +build protokolFile

package protokol

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
	"os"
	"testing"
)

type propsHealthIntFunc func(props props.IPropsHealth) int32
type propsHealthDecFunc func(props props.IPropsHealth) Decimal

func exportHealthPropsIntLine(protokol *os.File, testYear int16, sut factories.IFactoryHealth, function propsHealthIntFunc) {
	exportPropsYear(protokol, testYear)

	var testMonth int16
	for testMonth = 1; testMonth <=12; testMonth++ {
		testPeriod := types.GetPeriodWithYearMonth(testYear, testMonth)
		testResult, exists := sut.GetProps(testPeriod)
		var testValue int32 = 0
		if exists && testResult != nil {
			testValue = function(testResult)
		}
		exportPropsIntValue(protokol, testValue)
	}
	exportPropsEnd(protokol)
}

func exportHealthPropsDecLine(protokol *os.File, testYear int16, sut factories.IFactoryHealth, function propsHealthDecFunc) {
	exportPropsYear(protokol, testYear)

	var testMonth int16
	for testMonth = 1; testMonth <=12; testMonth++ {
		testPeriod := types.GetPeriodWithYearMonth(testYear, testMonth)
		testResult, exists := sut.GetProps(testPeriod)
		var testValue Decimal = NewFromFloat(0)
		if exists && testResult != nil {
			testValue = function(testResult)
		}
		exportPropsDecValue(protokol, testValue)
	}
	exportPropsEnd(protokol)
}

func exportHealthPropsIntFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactoryHealth, function propsHealthIntFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportHealthPropsIntLine(testProtokol, testYear, sut, function)
	}
}

func exportHealthPropsDecFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactoryHealth, function propsHealthDecFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportHealthPropsDecLine(testProtokol, testYear, sut, function )
	}
}

