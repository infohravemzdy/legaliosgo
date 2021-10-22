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

type propsSocialIntFunc func(props props.IPropsSocial) int32
type propsSocialDecFunc func(props props.IPropsSocial) Decimal

func exportSocialPropsIntLine(protokol *os.File, testYear int16, sut factories.IFactorySocial, function propsSocialIntFunc) {
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

func exportSocialPropsDecLine(protokol *os.File, testYear int16, sut factories.IFactorySocial, function propsSocialDecFunc) {
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

func exportSocialPropsIntFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactorySocial, function propsSocialIntFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportSocialPropsIntLine(testProtokol, testYear, sut, function)
	}
}

func exportSocialPropsDecFile(t *testing.T, baseName string, fileName string, minYear int16, maxYear int16, sut factories.IFactorySocial, function propsSocialDecFunc) {
	testProtokol, err := createProtokolFile(baseName, fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportSocialPropsDecLine(testProtokol, testYear, sut, function )
	}
}
