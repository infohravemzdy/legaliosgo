// +build protokolFile

package protokol

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	. "github.com/shopspring/decimal"
	"testing"
)

func TestProtokolHealth_MinMonthlyBasis(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_01_MinMonthlyBasis
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_01_MinMonthlyBasis.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.MinMonthlyBasis()})
		})
	}
}

func TestProtokolHealth_MaxAnnualsBasis(t *testing.T)  {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_02_MaxAnnualsBasis
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_02_MaxAnnualsBasis.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.MaxAnnualsBasis()})
		})
	}
}

func TestProtokolHealth_LimMonthlyState(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_03_LimMonthlyState
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_03_LimMonthlyState.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.LimMonthlyState()})
		})
	}
}

func TestProtokolHealth_LimMonthlyDis50(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_04_LimMonthlyDis50
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_04_LimMonthlyDis50.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.LimMonthlyDis50()})
		})
	}
}

func TestProtokolHealth_FactorCompound(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_05_FactorCompound
			exportHealthPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_05_FactorCompound.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) Decimal {return prop.FactorCompound()})
		})
	}
}

func TestProtokolHealth_FactorEmployee(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_06_FactorEmployee
			exportHealthPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_06_FactorEmployee.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) Decimal {return prop.FactorEmployee()})
		})
	}
}

func TestProtokolHealth_MarginIncomeEmp(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_07_MarginIncomeEmp
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_07_MarginIncomeEmp.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.MarginIncomeEmp()})
		})
	}
}

func TestProtokolHealth_MarginIncomeAgr(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2010-2022", 2010, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryHealth()

			// 01_Health_08_MarginIncomeAgr
			exportHealthPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"01_Health_08_MarginIncomeAgr.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsHealth) int32 {return prop.MarginIncomeAgr()})
		})
	}
}


