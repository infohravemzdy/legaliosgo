package service

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	. "github.com/shopspring/decimal"
	"math"
	"testing"
)

type testIntParams struct {
	testName string
	testYear int16
	testMonth int16
	resultYear int16
	resultMonth int16
	resultValue int32
}
type testIntScenario struct {
	title string
	tests []testIntParams
}

type testDecParams struct {
	testName string
	testYear int16
	testMonth int16
	resultYear int16
	resultMonth int16
	resultValue float64
}
type testDecScenario struct {
	title string
	tests []testDecParams
}

func (tt testIntParams) testBasicResult(t *testing.T, bundle factories.IBundleProps, error error) {
	if error != nil {
		t.Errorf("Error getting props from service for period = %d/%d; got error: %s",
			tt.testYear, tt.testMonth, error)
		return
	}
	if bundle == nil {
		t.Errorf("Error getting props from service for period = %d/%d; got nil",
			tt.testYear, tt.testMonth)
		return
	}
	var resultMonth = bundle.GetPeriodProps().Month()
	var resultYear = bundle.GetPeriodProps().Year()
	if resultYear != tt.resultYear || resultMonth != tt.resultMonth {
		t.Errorf("Error getting props from service for period = %d/%d expected: %d/%d; got %d/%d",
			tt.testYear, tt.testMonth, tt.resultYear, tt.resultMonth, resultYear, resultMonth)
	}
}

func (tt testDecParams) testBasicResult(t *testing.T, bundle factories.IBundleProps, error error) {
	if error != nil {
		t.Errorf("Error getting props from service for period = %d/%d; got error: %s",
			tt.testYear, tt.testMonth, error)
		return
	}
	if bundle == nil {
		t.Errorf("Error getting props from service for period = %d/%d; got nil",
			tt.testYear, tt.testMonth)
		return
	}
	var resultMonth = bundle.GetPeriodProps().Month()
	var resultYear = bundle.GetPeriodProps().Year()
	if resultYear != tt.resultYear || resultMonth != tt.resultMonth {
		t.Errorf("Error getting props from service for period = %d/%d expected: %d/%d; got %d/%d",
			tt.testYear, tt.testMonth, tt.resultYear, tt.resultMonth, resultYear, resultMonth)
	}
}

func (tt testDecParams) resultValueDec() Decimal {
	intValue := int64(math.Floor(tt.resultValue*100))
	return NewFromInt(intValue).Div(NewFromInt(100))
}


