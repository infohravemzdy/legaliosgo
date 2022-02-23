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
		{"2009",
			[]testParams{
				{ "2009-1", 2009, 1 },
				{ "2009-2", 2009, 2 },
				{ "2009-3", 2009, 3 },
				{ "2009-4", 2009, 4 },
				{ "2009-5", 2009, 5 },
				{ "2009-6", 2009, 6 },
				{ "2009-7", 2009, 7 },
				{ "2009-8", 2009, 8 },
				{ "2009-9", 2009, 9 },
				{ "2009-10", 2009, 10 },
				{ "2009-11", 2009, 11 },
				{ "2009-12", 2009, 12 },
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
