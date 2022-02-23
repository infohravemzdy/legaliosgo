// +build protokolFile

package protokol

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	. "github.com/shopspring/decimal"
	"testing"
)

func TestProtokolSocial_MaxAnnualsBasis(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_01_MaxAnnualsBasis
			exportSocialPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_01_MaxAnnualsBasis.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) int32 {return prop.MaxAnnualsBasis()})
		})
	}
}

func TestProtokolSocial_FactorEmployer(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_02_FactorEmployer
			exportSocialPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_02_FactorEmployer.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) Decimal {return prop.FactorEmployer()})
		})
	}
}

func TestProtokolSocial_FactorEmployerHigher(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_03_FactorEmployerHigher
			exportSocialPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_03_FactorEmployerHigher.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) Decimal {return prop.FactorEmployerHigher()})
		})
	}
}

func TestProtokolSocial_FactorEmployee(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_04_FactorEmployee
			exportSocialPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_04_FactorEmployee.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) Decimal {return prop.FactorEmployee()})
		})
	}
}

func TestProtokolSocial_FactorEmployeeGarant(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_05_FactorEmployeeGarant
			exportSocialPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_05_FactorEmployeeGarant.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) Decimal {return prop.FactorEmployeeGarant()})
		})
	}
}

func TestProtokolSocial_FactorEmployeeReduce(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_06_FactorEmployeeReduce
			exportSocialPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_06_FactorEmployeeReduce.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) Decimal {return prop.FactorEmployeeReduce()})
		})
	}
}

func TestProtokolSocial_MarginIncomeEmp(t *testing.T){
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
			factory := factories.NewFactorySocial()

			// 03_Social_07_MarginIncomeEmp
			exportSocialPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_07_MarginIncomeEmp.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) int32 {return prop.MarginIncomeEmp()})
		})
	}
}

func TestProtokolSocial_MarginIncomeAgr(t *testing.T) {
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
			factory := factories.NewFactorySocial()

			// 03_Social_08_MarginIncomeAgr
			exportSocialPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"03_Social_08_MarginIncomeAgr.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSocial) int32 {return prop.MarginIncomeAgr()})
		})
	}
}


