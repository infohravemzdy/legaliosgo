// +build protokolFile

package protokol

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"testing"
)

func TestProtokolTaxing_AllowancePayer(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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
		{"2011-2022", 2011, 2022},
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

func TestProtokolTaxing_FactorSolitary(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_11_FactorSolitary
			exportTaxingPropsDecFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_11_FactorSolitary.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) Decimal {return prop.FactorSolitary()})
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_12_MinAmountOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_12_MinAmountOfTaxBonus.txt",
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_13_MaxAmountOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_13_MaxAmountOfTaxBonus.txt",
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_14_MarginIncomeOfTaxBonus
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_14_MarginIncomeOfTaxBonus.txt",
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_15_MarginIncomeOfRounding
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_15_MarginIncomeOfRounding.txt",
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_16_MarginIncomeOfWithhold
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_16_MarginIncomeOfWithhold.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfWithhold()})
		})
	}
}

func TestProtokolTaxing_MarginIncomeOfSolitary(t *testing.T) {
	type testScenario struct {
		title string
		minYear int16
		maxYear int16
	}
	testExamples := []testScenario {
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_17_MarginIncomeOfSolitary
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_17_MarginIncomeOfSolitary.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfSolitary()})
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_18_MarginIncomeOfWthEmp
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_18_MarginIncomeOfWthEmp.txt",
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
		{"2011-2022", 2011, 2022},
	}
	for _, tt := range testExamples {
		t.Run(tt.title, func(t *testing.T) {
			factory := factories.NewFactoryTaxing()

			// 04_Taxing_19_MarginIncomeOfWthAgr
			exportTaxingPropsIntFile(t, PROTOKOL_FOLDER_PATH,
				"04_Taxing_19_MarginIncomeOfWthAgr.txt",
				tt.minYear, tt.maxYear, factory,
				func (prop props.IPropsTaxing) int32 {return prop.MarginIncomeOfWthAgr()})
		})
	}
}

