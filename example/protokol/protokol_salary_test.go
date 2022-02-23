// +build protokolFile

package protokol

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"testing"
)

func TestProtokolSalary_WorkingShiftWeek(t *testing.T) {
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
			factory := factories.NewFactorySalary()

			// 02_Salary_01_WorkingShiftWeek
			exportSalaryPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"02_Salary_01_WorkingShiftWeek.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSalary) int32 {return prop.WorkingShiftWeek()})
		})
	}
}

func TestProtokolSalary_WorkingShiftTime(t *testing.T) {
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
			factory := factories.NewFactorySalary()

			// 02_Salary_02_WorkingShiftTime
			exportSalaryPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"02_Salary_02_WorkingShiftTime.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSalary) int32 {return prop.WorkingShiftTime()})
		})
	}
}
func TestProtokolSalary_MinMonthlyWage(t *testing.T) {
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
			factory := factories.NewFactorySalary()

			// 02_Salary_03_MinMonthlyWage
			exportSalaryPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"02_Salary_03_MinMonthlyWage.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSalary) int32 {return prop.MinMonthlyWage()})
		})
	}
}
func TestProtokolSalary_MinHourlyWage(t *testing.T)  {
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
			factory := factories.NewFactorySalary()

			// 02_Salary_04_MinHourlyWage
			exportSalaryPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"02_Salary_04_MinHourlyWage.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsSalary) int32 {return prop.MinHourlyWage()})
		})
	}
}

