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

type propsTaxingIntFunc func(props props.IPropsTaxing) int32
type propsTaxingDecFunc func(props props.IPropsTaxing) Decimal

func exportTaxingPropsIntLine(protokol *os.File, testYear int16, sut factories.IFactoryTaxing, function propsTaxingIntFunc) {
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

func exportTaxingPropsDecLine(protokol *os.File, testYear int16, sut factories.IFactoryTaxing, function propsTaxingDecFunc) {
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

func exportTaxingPropsIntFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactoryTaxing, function propsTaxingIntFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportTaxingPropsIntLine(testProtokol, testYear, sut, function)
	}
}

func exportTaxingPropsDecFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactoryTaxing, function propsTaxingDecFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportTaxingPropsDecLine(testProtokol, testYear, sut, function )
	}
}
