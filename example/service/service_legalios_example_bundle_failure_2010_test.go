package service

import (
	service "github.com/mzdyhrave/legaliosgo"
	"testing"
)

func TestBundleFailure_ForYear2010(t *testing.T) {
	type testParams struct {
		testName string
		testYear int16
		testMonth int16
	}
	type testScenario struct {
		title string
		tests []testParams
	}
	testExamples := []testScenario {
		{"2010",
			[]testParams{
				{ "2010-1", 2010, 1 },
				{ "2010-2", 2010, 2 },
				{ "2010-3", 2010, 3 },
				{ "2010-4", 2010, 4 },
				{ "2010-5", 2010, 5 },
				{ "2010-6", 2010, 6 },
				{ "2010-7", 2010, 7 },
				{ "2010-8", 2010, 8 },
				{ "2010-9", 2010, 9 },
				{ "2010-10", 2010, 10 },
				{ "2010-11", 2010, 11 },
				{ "2010-12", 2010, 12 },
			},
		},
	}
	for _, tx := range testExamples {
		for _, tt := range tx.tests {
			t.Run(tt.testName, func(t *testing.T) {
				var period = service.NewPeriodWithYearMonth(tt.testYear, tt.testMonth)
				var service service.IServiceLegalios = service.NewLegaliosService()
				bundle, error := service.GetBundle(period)
				if error == nil {
					t.Errorf("Error getting error from service for period = %d/%d",
						tt.testYear, tt.testMonth)
					return
				}
				if bundle != nil {
					var resultMonth = bundle.GetPeriodProps().Month()
					var resultYear = bundle.GetPeriodProps().Year()
					t.Errorf("Error getting props from service for period = %d/%d; got %d/%d",
						tt.testYear, tt.testMonth, resultYear, resultMonth)
					return
				}
			})
		}
	}
}
