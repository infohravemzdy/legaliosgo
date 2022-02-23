// +build protokolFile

package protokol

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	. "github.com/shopspring/decimal"
	"testing"
)

func TestProtokolTaxing_AllowancePayer(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_01_AllowancePayer
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_01_AllowancePayer.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowancePayer()})
		})
	}
}

func TestProtokolTaxing_AllowanceDisab1st(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_02_AllowanceDisab1st
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_02_AllowanceDisab1st.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceDisab1st()})
		})
	}
}

func TestProtokolTaxing_AllowanceDisab2nd(t *testing.T){
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_03_AllowanceDisab2nd
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_03_AllowanceDisab2nd.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceDisab2nd()})
		})
	}
}

func TestProtokolTaxing_AllowanceDisab3rd(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_04_AllowanceDisab3rd
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_04_AllowanceDisab3rd.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceDisab3rd()})
		})
	}
}

func TestProtokolTaxing_AllowanceStudy(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_05_AllowanceStudy
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_05_AllowanceStudy.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceStudy()})
		})
	}
}

func TestProtokolTaxing_AllowanceChild1st(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_06_AllowanceChild1st
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_06_AllowanceChild1st.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceChild1st()})
		})
	}
}

func TestProtokolTaxing_AllowanceChild2nd(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_07_AllowanceChild2nd
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_07_AllowanceChild2nd.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceChild2nd()})
		})
	}
}

func TestProtokolTaxing_AllowanceChild3rd(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_08_AllowanceChild3rd
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_08_AllowanceChild3rd.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.AllowanceChild3rd()})
		})
	}
}

func TestProtokolTaxing_FactorAdvances(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_09_FactorAdvances
			exportTaxingPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_09_FactorAdvances.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) Decimal {return prop.FactorAdvances()})
		})
	}
}

func TestProtokolTaxing_FactorWithhold(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_10_FactorWithhold
			exportTaxingPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_10_FactorWithhold.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) Decimal {return prop.FactorWithhold()})
		})
	}
}

func TestProtokolTaxing_FactorSolidary(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_11_FactorSolidary
			exportTaxingPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_11_FactorSolidary.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) Decimal {return prop.FactorSolidary()})
		})
	}
}

func TestProtokolTaxing_FactorTaxRate2(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_12_FactorTaxRate2
			exportTaxingPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_12_FactorTaxRate2.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) Decimal {return prop.FactorTaxRate2()})
		})
	}
}

func TestProtokolTaxing_MinAmountOfTaxBonus(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_13_MinAmountOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_13_MinAmountOfTaxBonus.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MinAmountOfTaxBonus()})
		})
	}
}

func TestProtokolTaxing_MaxAmountOfTaxBonus(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_14_MaxAmountOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_14_MaxAmountOfTaxBonus.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MaxAmountOfTaxBonus()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfTaxBonus(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_15_MarginIncomeOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_15_MarginIncomeOfTaxBonus.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfTaxBonus()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfRounding(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_16_MarginIncomeOfRounding
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_16_MarginIncomeOfRounding.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfRounding()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfWithhold(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_17_MarginIncomeOfWithhold
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_17_MarginIncomeOfWithhold.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfWithhold()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfSolidary(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_18_MarginIncomeOfSolidary
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_18_MarginIncomeOfSolidary.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfSolidary()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfTaxRate2(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_19_MarginIncomeOfTaxRate2
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_19_MarginIncomeOfTaxRate2.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfTaxRate2()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfWthEmp(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_20_MarginIncomeOfWthEmp
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_20_MarginIncomeOfWthEmp.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfWthEmp()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfWthAgr(t *testing.T) {
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
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_21_MarginIncomeOfWthAgr
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_21_MarginIncomeOfWthAgr.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfWthAgr()})
		})
	}
}

