package history

import (
	"fmt"
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	"testing"
)

func TestFactoriesHistory(t *testing.T) {
	var _sutSalary factories.IFactorySalary = factories.NewFactorySalary()
	var _sutHealth factories.IFactoryHealth = factories.NewFactoryHealth()
	var _sutSocial factories.IFactorySocial = factories.NewFactorySocial()
	var _sutTaxing factories.IFactoryTaxing = factories.NewFactoryTaxing()

	minYear := 2010
	maxYear := 2022

	testProtokol, err := createHistoryFile(HISTORY_FOLDER_PATH, fmt.Sprintf("history_%d_%d.xls", minYear, maxYear))

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	var headerData []tupleChanges

	for testYear := minYear; testYear <= maxYear; testYear++ {
		var yearWithChanges = false

		testPeriod := types.GetPeriodWithYearMonth(int16(testYear), 1)

		var testSalaryProp, _ = _sutSalary.GetProps(testPeriod)
		var testHealthProp, _ = _sutHealth.GetProps(testPeriod)
		var testSocialProp, _ = _sutSocial.GetProps(testPeriod)
		var testTaxingProp, _ = _sutTaxing.GetProps(testPeriod)

		for testMonth := 2; testMonth <= 12; testMonth++ {
			nextPeriod := types.GetPeriodWithYearMonth(int16(testYear), int16(testMonth))

			testSalaryNext, _ := _sutSalary.GetProps(nextPeriod)
			testHealthNext, _ := _sutHealth.GetProps(nextPeriod)
			testSocialNext, _ := _sutSocial.GetProps(nextPeriod)
			testTaxingNext, _ := _sutTaxing.GetProps(nextPeriod)

			if testSalaryNext.ValueEquals(testSalaryProp) == false {
				yearWithChanges = true
			}
			if testHealthNext.ValueEquals(testHealthProp) == false {
				yearWithChanges = true
			}
			if testSocialNext.ValueEquals(testSocialProp) == false {
				yearWithChanges = true
			}
			if testTaxingNext.ValueEquals(testTaxingProp) == false {
				yearWithChanges = true
			}
			testSalaryProp = testSalaryNext
			testHealthProp = testHealthNext
			testSocialProp = testSocialNext
			testTaxingProp = testTaxingNext
		}
		headerData = append(headerData, tupleChanges { a: int16(testYear), b: yearWithChanges })
	}
	exportHistoryStart(testProtokol, headerData)

	var VECT_HEALTH_MIN_MONTHLY_BASIS []tupleValues
	var VECT_HEALTH_MAX_ANNUALS_BASIS []tupleValues
	var VECT_HEALTH_LIM_MONTHLY_STATE []tupleValues
	var VECT_HEALTH_LIM_MONTHLY_DIS50 []tupleValues
	var VECT_HEALTH_FACTOR_COMPOUND []tupleValues
	var VECT_HEALTH_FACTOR_EMPLOYEE []tupleValues
	var VECT_HEALTH_MARGIN_INCOME_EMP []tupleValues
	var VECT_HEALTH_MARGIN_INCOME_AGR []tupleValues
	var VECT_SALARY_WORKING_SHIFT_WEEK []tupleValues
	var VECT_SALARY_WORKING_SHIFT_TIME []tupleValues
	var VECT_SALARY_MIN_MONTHLY_WAGE []tupleValues
	var VECT_SALARY_MIN_HOURLY_WAGE []tupleValues
	var VECT_SOCIAL_MAX_ANNUALS_BASIS []tupleValues
	var VECT_SOCIAL_FACTOR_EMPLOYER []tupleValues
	var VECT_SOCIAL_FACTOR_EMPLOYER_HIGHER []tupleValues
	var VECT_SOCIAL_FACTOR_EMPLOYEE []tupleValues
	var VECT_SOCIAL_FACTOR_EMPLOYEE_GARANT []tupleValues
	var VECT_SOCIAL_FACTOR_EMPLOYEE_REDUCE []tupleValues
	var VECT_SOCIAL_MARGIN_INCOME_EMP []tupleValues
	var VECT_SOCIAL_MARGIN_INCOME_AGR []tupleValues
	var VECT_TAXING_ALLOWANCE_PAYER []tupleValues
	var VECT_TAXING_ALLOWANCE_DISAB_1ST []tupleValues
	var VECT_TAXING_ALLOWANCE_DISAB_2ND []tupleValues
	var VECT_TAXING_ALLOWANCE_DISAB_3RD []tupleValues
	var VECT_TAXING_ALLOWANCE_STUDY []tupleValues
	var VECT_TAXING_ALLOWANCE_CHILD_1ST []tupleValues
	var VECT_TAXING_ALLOWANCE_CHILD_2ND []tupleValues
	var VECT_TAXING_ALLOWANCE_CHILD_3RD []tupleValues
	var VECT_TAXING_FACTOR_ADVANCES []tupleValues
	var VECT_TAXING_FACTOR_WITHHOLD []tupleValues
	var VECT_TAXING_FACTOR_SOLIDARY []tupleValues
	var VECT_TAXING_FACTOR_TAXRATE2 []tupleValues
	var VECT_TAXING_MIN_AMOUNT_OF_TAXBONUS []tupleValues
	var VECT_TAXING_MAX_AMOUNT_OF_TAXBONUS []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_TAXBONUS []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_ROUNDING []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_WITHHOLD []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_SOLIDARY []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_TAXRATE2 []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_WHT_EMP []tupleValues
	var VECT_TAXING_MARGIN_INCOME_OF_WHT_AGR []tupleValues

	for testYear := minYear; testYear <= maxYear; testYear++ {
		var testIntYear int16 = int16(testYear)

		var MES_HEALTH_MIN_MONTHLY_BASIS         int16 = 0
		var MES_HEALTH_MAX_ANNUALS_BASIS         int16 = 0
		var MES_HEALTH_LIM_MONTHLY_STATE         int16 = 0
		var MES_HEALTH_LIM_MONTHLY_DIS50         int16 = 0
		var MES_HEALTH_FACTOR_COMPOUND           int16 = 0
		var MES_HEALTH_FACTOR_EMPLOYEE           int16 = 0
		var MES_HEALTH_MARGIN_INCOME_EMP         int16 = 0
		var MES_HEALTH_MARGIN_INCOME_AGR         int16 = 0
		var MES_SALARY_WORKING_SHIFT_WEEK        int16 = 0
		var MES_SALARY_WORKING_SHIFT_TIME        int16 = 0
		var MES_SALARY_MIN_MONTHLY_WAGE          int16 = 0
		var MES_SALARY_MIN_HOURLY_WAGE           int16 = 0
		var MES_SOCIAL_MAX_ANNUALS_BASIS         int16 = 0
		var MES_SOCIAL_FACTOR_EMPLOYER           int16 = 0
		var MES_SOCIAL_FACTOR_EMPLOYER_HIGHER    int16 = 0
		var MES_SOCIAL_FACTOR_EMPLOYEE           int16 = 0
		var MES_SOCIAL_FACTOR_EMPLOYEE_GARANT    int16 = 0
		var MES_SOCIAL_FACTOR_EMPLOYEE_REDUCE    int16 = 0
		var MES_SOCIAL_MARGIN_INCOME_EMP         int16 = 0
		var MES_SOCIAL_MARGIN_INCOME_AGR         int16 = 0
		var MES_TAXING_ALLOWANCE_PAYER           int16 = 0
		var MES_TAXING_ALLOWANCE_DISAB_1ST       int16 = 0
		var MES_TAXING_ALLOWANCE_DISAB_2ND       int16 = 0
		var MES_TAXING_ALLOWANCE_DISAB_3RD       int16 = 0
		var MES_TAXING_ALLOWANCE_STUDY           int16 = 0
		var MES_TAXING_ALLOWANCE_CHILD_1ST       int16 = 0
		var MES_TAXING_ALLOWANCE_CHILD_2ND       int16 = 0
		var MES_TAXING_ALLOWANCE_CHILD_3RD       int16 = 0
		var MES_TAXING_FACTOR_ADVANCES           int16 = 0
		var MES_TAXING_FACTOR_WITHHOLD           int16 = 0
		var MES_TAXING_FACTOR_SOLIDARY           int16 = 0
		var MES_TAXING_FACTOR_TAXRATE2           int16 = 0
		var MES_TAXING_MIN_AMOUNT_OF_TAXBONUS    int16 = 0
		var MES_TAXING_MAX_AMOUNT_OF_TAXBONUS    int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_TAXBONUS int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_ROUNDING int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_WITHHOLD int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_SOLIDARY int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_TAXRATE2 int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_WHT_EMP  int16 = 0
		var MES_TAXING_MARGIN_INCOME_OF_WHT_AGR  int16 = 0

		testPeriod := types.GetPeriodWithYearMonth(testIntYear, 1)

		var testSalaryProp, _ = _sutSalary.GetProps(testPeriod)
		var testHealthProp, _ = _sutHealth.GetProps(testPeriod)
		var testSocialProp, _ = _sutSocial.GetProps(testPeriod)
		var testTaxingProp, _ = _sutTaxing.GetProps(testPeriod)

		var JAN_HEALTH_MIN_MONTHLY_BASIS         = propsIntValueToString(testHealthProp.MinMonthlyBasis())
		var JAN_HEALTH_MAX_ANNUALS_BASIS         = propsIntValueToString(testHealthProp.MaxAnnualsBasis())
		var JAN_HEALTH_LIM_MONTHLY_STATE         = propsIntValueToString(testHealthProp.LimMonthlyState())
		var JAN_HEALTH_LIM_MONTHLY_DIS50         = propsIntValueToString(testHealthProp.LimMonthlyDis50())
		var JAN_HEALTH_FACTOR_COMPOUND           = propsDecValueToString(testHealthProp.FactorCompound() )
		var JAN_HEALTH_FACTOR_EMPLOYEE           = propsDecValueToString(testHealthProp.FactorEmployee() )
		var JAN_HEALTH_MARGIN_INCOME_EMP         = propsIntValueToString(testHealthProp.MarginIncomeEmp())
		var JAN_HEALTH_MARGIN_INCOME_AGR         = propsIntValueToString(testHealthProp.MarginIncomeAgr())
		var JAN_SALARY_WORKING_SHIFT_WEEK        = propsIntValueToString(testSalaryProp.WorkingShiftWeek())
		var JAN_SALARY_WORKING_SHIFT_TIME        = propsIntValueToString(testSalaryProp.WorkingShiftTime())
		var JAN_SALARY_MIN_MONTHLY_WAGE          = propsIntValueToString(testSalaryProp.MinMonthlyWage())
		var JAN_SALARY_MIN_HOURLY_WAGE           = propsIntValueToString(testSalaryProp.MinHourlyWage()  )
		var JAN_SOCIAL_MAX_ANNUALS_BASIS         = propsIntValueToString(testSocialProp.MaxAnnualsBasis())
		var JAN_SOCIAL_FACTOR_EMPLOYER           = propsDecValueToString(testSocialProp.FactorEmployer())
		var JAN_SOCIAL_FACTOR_EMPLOYER_HIGHER    = propsDecValueToString(testSocialProp.FactorEmployerHigher())
		var JAN_SOCIAL_FACTOR_EMPLOYEE           = propsDecValueToString(testSocialProp.FactorEmployee())
		var JAN_SOCIAL_FACTOR_EMPLOYEE_GARANT    = propsDecValueToString(testSocialProp.FactorEmployeeGarant())
		var JAN_SOCIAL_FACTOR_EMPLOYEE_REDUCE    = propsDecValueToString(testSocialProp.FactorEmployeeReduce())
		var JAN_SOCIAL_MARGIN_INCOME_EMP         = propsIntValueToString(testSocialProp.MarginIncomeEmp())
		var JAN_SOCIAL_MARGIN_INCOME_AGR         = propsIntValueToString(testSocialProp.MarginIncomeAgr())
		var JAN_TAXING_ALLOWANCE_PAYER           = propsIntValueToString(testTaxingProp.AllowancePayer())
		var JAN_TAXING_ALLOWANCE_DISAB_1ST       = propsIntValueToString(testTaxingProp.AllowanceDisab1st() )
		var JAN_TAXING_ALLOWANCE_DISAB_2ND       = propsIntValueToString(testTaxingProp.AllowanceDisab2nd() )
		var JAN_TAXING_ALLOWANCE_DISAB_3RD       = propsIntValueToString(testTaxingProp.AllowanceDisab3rd() )
		var JAN_TAXING_ALLOWANCE_STUDY           = propsIntValueToString(testTaxingProp.AllowanceStudy() )
		var JAN_TAXING_ALLOWANCE_CHILD_1ST       = propsIntValueToString(testTaxingProp.AllowanceChild1st() )
		var JAN_TAXING_ALLOWANCE_CHILD_2ND       = propsIntValueToString(testTaxingProp.AllowanceChild2nd() )
		var JAN_TAXING_ALLOWANCE_CHILD_3RD       = propsIntValueToString(testTaxingProp.AllowanceChild3rd() )
		var JAN_TAXING_FACTOR_ADVANCES           = propsDecValueToString(testTaxingProp.FactorAdvances() )
		var JAN_TAXING_FACTOR_WITHHOLD           = propsDecValueToString(testTaxingProp.FactorWithhold() )
		var JAN_TAXING_FACTOR_SOLIDARY           = propsDecValueToString(testTaxingProp.FactorSolidary() )
		var JAN_TAXING_FACTOR_TAXRATE2           = propsDecValueToString(testTaxingProp.FactorTaxRate2() )
		var JAN_TAXING_MIN_AMOUNT_OF_TAXBONUS    = propsIntValueToString(testTaxingProp.MinAmountOfTaxBonus() )
		var JAN_TAXING_MAX_AMOUNT_OF_TAXBONUS    = propsIntValueToString(testTaxingProp.MaxAmountOfTaxBonus() )
		var JAN_TAXING_MARGIN_INCOME_OF_TAXBONUS = propsIntValueToString(testTaxingProp.MarginIncomeOfTaxBonus() )
		var JAN_TAXING_MARGIN_INCOME_OF_ROUNDING = propsIntValueToString(testTaxingProp.MarginIncomeOfRounding() )
		var JAN_TAXING_MARGIN_INCOME_OF_WITHHOLD = propsIntValueToString(testTaxingProp.MarginIncomeOfWithhold() )
		var JAN_TAXING_MARGIN_INCOME_OF_SOLIDARY = propsIntValueToString(testTaxingProp.MarginIncomeOfSolidary() )
		var JAN_TAXING_MARGIN_INCOME_OF_TAXRATE2 = propsIntValueToString(testTaxingProp.MarginIncomeOfTaxRate2() )
		var JAN_TAXING_MARGIN_INCOME_OF_WHT_EMP  = propsIntValueToString(testTaxingProp.MarginIncomeOfWthEmp() )
		var JAN_TAXING_MARGIN_INCOME_OF_WHT_AGR  = propsIntValueToString(testTaxingProp.MarginIncomeOfWthAgr() )

		for testMonth := 2; testMonth <= 12; testMonth++ {
			nextPeriod := types.GetPeriodWithYearMonth(testIntYear, int16(testMonth))

			testSalaryNext, _ := _sutSalary.GetProps(nextPeriod)
			testHealthNext, _ := _sutHealth.GetProps(nextPeriod)
			testSocialNext, _ := _sutSocial.GetProps(nextPeriod)
			testTaxingNext, _ := _sutTaxing.GetProps(nextPeriod)

			var testIntMonth int16 = int16(testMonth)

			if testHealthNext.MinMonthlyBasis() != testHealthProp.MinMonthlyBasis() { MES_HEALTH_MIN_MONTHLY_BASIS = testIntMonth }
			if testHealthNext.MaxAnnualsBasis() != testHealthProp.MaxAnnualsBasis() { MES_HEALTH_MAX_ANNUALS_BASIS = testIntMonth }
			if testHealthNext.LimMonthlyState() != testHealthProp.LimMonthlyState() { MES_HEALTH_LIM_MONTHLY_STATE = testIntMonth }
			if testHealthNext.LimMonthlyDis50() != testHealthProp.LimMonthlyDis50() { MES_HEALTH_LIM_MONTHLY_DIS50 = testIntMonth }
			if testHealthNext.FactorCompound().Equal(testHealthProp.FactorCompound())==false { MES_HEALTH_FACTOR_COMPOUND = testIntMonth }
			if testHealthNext.FactorEmployee().Equal(testHealthProp.FactorEmployee())==false { MES_HEALTH_FACTOR_EMPLOYEE = testIntMonth }
			if testHealthNext.MarginIncomeEmp() != testHealthProp.MarginIncomeEmp() { MES_HEALTH_MARGIN_INCOME_EMP = testIntMonth }
			if testHealthNext.MarginIncomeAgr() != testHealthProp.MarginIncomeAgr() { MES_HEALTH_MARGIN_INCOME_AGR = testIntMonth }
			if testSalaryNext.WorkingShiftWeek() != testSalaryProp.WorkingShiftWeek() { MES_SALARY_WORKING_SHIFT_WEEK = testIntMonth }
			if testSalaryNext.WorkingShiftTime() != testSalaryProp.WorkingShiftTime() { MES_SALARY_WORKING_SHIFT_TIME = testIntMonth }
			if testSalaryNext.MinMonthlyWage() != testSalaryProp.MinMonthlyWage() { MES_SALARY_MIN_MONTHLY_WAGE = testIntMonth }
			if testSalaryNext.MinHourlyWage() != testSalaryProp.MinHourlyWage() { MES_SALARY_MIN_HOURLY_WAGE = testIntMonth }
			if testSocialNext.MaxAnnualsBasis() != testSocialProp.MaxAnnualsBasis() { MES_SOCIAL_MAX_ANNUALS_BASIS = testIntMonth }
			if testSocialNext.FactorEmployer().Equal(testSocialProp.FactorEmployer())==false { MES_SOCIAL_FACTOR_EMPLOYER = testIntMonth }
			if testSocialNext.FactorEmployerHigher().Equal(testSocialProp.FactorEmployerHigher())==false { MES_SOCIAL_FACTOR_EMPLOYER_HIGHER = testIntMonth }
			if testSocialNext.FactorEmployee().Equal(testSocialProp.FactorEmployee())==false { MES_SOCIAL_FACTOR_EMPLOYEE = testIntMonth }
			if testSocialNext.FactorEmployeeReduce().Equal(testSocialProp.FactorEmployeeReduce())==false { MES_SOCIAL_FACTOR_EMPLOYEE_REDUCE = testIntMonth }
			if testSocialNext.FactorEmployeeGarant().Equal(testSocialProp.FactorEmployeeGarant())==false { MES_SOCIAL_FACTOR_EMPLOYEE_GARANT = testIntMonth }
			if testSocialNext.MarginIncomeEmp() != testSocialProp.MarginIncomeEmp() { MES_SOCIAL_MARGIN_INCOME_EMP = testIntMonth }
			if testSocialNext.MarginIncomeAgr() != testSocialProp.MarginIncomeAgr() { MES_SOCIAL_MARGIN_INCOME_AGR = testIntMonth }
			if testTaxingNext.AllowancePayer() != testTaxingProp.AllowancePayer() { MES_TAXING_ALLOWANCE_PAYER = testIntMonth }
			if testTaxingNext.AllowanceDisab1st() != testTaxingProp.AllowanceDisab1st() { MES_TAXING_ALLOWANCE_DISAB_1ST = testIntMonth }
			if testTaxingNext.AllowanceDisab2nd() != testTaxingProp.AllowanceDisab2nd() { MES_TAXING_ALLOWANCE_DISAB_2ND = testIntMonth }
			if testTaxingNext.AllowanceDisab3rd() != testTaxingProp.AllowanceDisab3rd() { MES_TAXING_ALLOWANCE_DISAB_3RD = testIntMonth }
			if testTaxingNext.AllowanceStudy() != testTaxingProp.AllowanceStudy() { MES_TAXING_ALLOWANCE_STUDY = testIntMonth }
			if testTaxingNext.AllowanceChild1st() != testTaxingProp.AllowanceChild1st() { MES_TAXING_ALLOWANCE_CHILD_1ST = testIntMonth }
			if testTaxingNext.AllowanceChild2nd() != testTaxingProp.AllowanceChild2nd() { MES_TAXING_ALLOWANCE_CHILD_2ND = testIntMonth }
			if testTaxingNext.AllowanceChild3rd() != testTaxingProp.AllowanceChild3rd() { MES_TAXING_ALLOWANCE_CHILD_3RD = testIntMonth }
			if testTaxingNext.FactorAdvances().Equal(testTaxingProp.FactorAdvances())==false { MES_TAXING_FACTOR_ADVANCES = testIntMonth }
			if testTaxingNext.FactorWithhold().Equal(testTaxingProp.FactorWithhold())==false { MES_TAXING_FACTOR_WITHHOLD = testIntMonth }
			if testTaxingNext.FactorSolidary().Equal(testTaxingProp.FactorSolidary())==false { MES_TAXING_FACTOR_SOLIDARY = testIntMonth }
			if testTaxingNext.FactorTaxRate2().Equal(testTaxingProp.FactorTaxRate2())==false { MES_TAXING_FACTOR_TAXRATE2 = testIntMonth }
			if testTaxingNext.MinAmountOfTaxBonus() != testTaxingProp.MinAmountOfTaxBonus() { MES_TAXING_MIN_AMOUNT_OF_TAXBONUS = testIntMonth }
			if testTaxingNext.MaxAmountOfTaxBonus() != testTaxingProp.MaxAmountOfTaxBonus() { MES_TAXING_MAX_AMOUNT_OF_TAXBONUS = testIntMonth }
			if testTaxingNext.MarginIncomeOfTaxBonus() != testTaxingProp.MarginIncomeOfTaxBonus() { MES_TAXING_MARGIN_INCOME_OF_TAXBONUS = testIntMonth }
			if testTaxingNext.MarginIncomeOfRounding() != testTaxingProp.MarginIncomeOfRounding() { MES_TAXING_MARGIN_INCOME_OF_ROUNDING = testIntMonth }
			if testTaxingNext.MarginIncomeOfWithhold() != testTaxingProp.MarginIncomeOfWithhold() { MES_TAXING_MARGIN_INCOME_OF_WITHHOLD = testIntMonth }
			if testTaxingNext.MarginIncomeOfSolidary() != testTaxingProp.MarginIncomeOfSolidary() { MES_TAXING_MARGIN_INCOME_OF_SOLIDARY = testIntMonth }
			if testTaxingNext.MarginIncomeOfTaxRate2() != testTaxingProp.MarginIncomeOfTaxRate2() { MES_TAXING_MARGIN_INCOME_OF_TAXRATE2 = testIntMonth }
			if testTaxingNext.MarginIncomeOfWthEmp() != testTaxingProp.MarginIncomeOfWthEmp() { MES_TAXING_MARGIN_INCOME_OF_WHT_EMP = testIntMonth }
			if testTaxingNext.MarginIncomeOfWthAgr() != testTaxingProp.MarginIncomeOfWthAgr() { MES_TAXING_MARGIN_INCOME_OF_WHT_AGR = testIntMonth }

			testSalaryProp = testSalaryNext
			testHealthProp = testHealthNext
			testSocialProp = testSocialNext
			testTaxingProp = testTaxingNext
		}
		VECT_HEALTH_MIN_MONTHLY_BASIS = append(VECT_HEALTH_MIN_MONTHLY_BASIS, tupleValues {a: testIntYear, b: MES_HEALTH_MIN_MONTHLY_BASIS, c: JAN_HEALTH_MIN_MONTHLY_BASIS, d: propsIntValueToString(testHealthProp.MinMonthlyBasis()) } )
		VECT_HEALTH_MAX_ANNUALS_BASIS = append(VECT_SOCIAL_MAX_ANNUALS_BASIS, tupleValues {a: testIntYear, b: MES_HEALTH_MAX_ANNUALS_BASIS, c: JAN_HEALTH_MAX_ANNUALS_BASIS, d: propsIntValueToString(testHealthProp.MaxAnnualsBasis()) } )
		VECT_HEALTH_LIM_MONTHLY_STATE = append(VECT_HEALTH_LIM_MONTHLY_STATE, tupleValues {a: testIntYear, b: MES_HEALTH_LIM_MONTHLY_STATE, c: JAN_HEALTH_LIM_MONTHLY_STATE, d: propsIntValueToString(testHealthProp.LimMonthlyState()) } )
		VECT_HEALTH_LIM_MONTHLY_DIS50 = append(VECT_HEALTH_LIM_MONTHLY_DIS50, tupleValues {a: testIntYear, b: MES_HEALTH_LIM_MONTHLY_DIS50, c: JAN_HEALTH_LIM_MONTHLY_DIS50, d: propsIntValueToString(testHealthProp.LimMonthlyDis50()) } )
		VECT_HEALTH_FACTOR_COMPOUND = append(VECT_HEALTH_FACTOR_COMPOUND, tupleValues {a: testIntYear, b: MES_HEALTH_FACTOR_COMPOUND, c: JAN_HEALTH_FACTOR_COMPOUND, d: propsDecValueToString(testHealthProp.FactorCompound()) } )
		VECT_HEALTH_FACTOR_EMPLOYEE = append(VECT_HEALTH_FACTOR_EMPLOYEE, tupleValues {a: testIntYear, b: MES_HEALTH_FACTOR_EMPLOYEE, c: JAN_HEALTH_FACTOR_EMPLOYEE, d: propsDecValueToString(testHealthProp.FactorEmployee()) } )
		VECT_HEALTH_MARGIN_INCOME_EMP = append(VECT_HEALTH_MARGIN_INCOME_EMP, tupleValues {a: testIntYear, b: MES_HEALTH_MARGIN_INCOME_EMP, c: JAN_HEALTH_MARGIN_INCOME_EMP, d: propsIntValueToString(testHealthProp.MarginIncomeEmp()) } )
		VECT_HEALTH_MARGIN_INCOME_AGR = append(VECT_HEALTH_MARGIN_INCOME_AGR, tupleValues {a: testIntYear, b: MES_HEALTH_MARGIN_INCOME_AGR, c: JAN_HEALTH_MARGIN_INCOME_AGR, d: propsIntValueToString(testHealthProp.MarginIncomeAgr()) } )
		VECT_SALARY_WORKING_SHIFT_WEEK = append(VECT_SALARY_WORKING_SHIFT_WEEK, tupleValues {a: testIntYear, b: MES_SALARY_WORKING_SHIFT_WEEK, c: JAN_SALARY_WORKING_SHIFT_WEEK, d: propsIntValueToString(testSalaryProp.WorkingShiftWeek()) } )
		VECT_SALARY_WORKING_SHIFT_TIME = append(VECT_SALARY_WORKING_SHIFT_TIME, tupleValues {a: testIntYear, b: MES_SALARY_WORKING_SHIFT_TIME, c: JAN_SALARY_WORKING_SHIFT_TIME, d: propsIntValueToString(testSalaryProp.WorkingShiftTime()) } )
		VECT_SALARY_MIN_MONTHLY_WAGE = append(VECT_SALARY_MIN_MONTHLY_WAGE, tupleValues {a: testIntYear, b: MES_SALARY_MIN_MONTHLY_WAGE, c: JAN_SALARY_MIN_MONTHLY_WAGE, d: propsIntValueToString(testSalaryProp.MinMonthlyWage()) } )
		VECT_SALARY_MIN_HOURLY_WAGE = append(VECT_SALARY_MIN_HOURLY_WAGE, tupleValues {a: testIntYear, b: MES_SALARY_MIN_HOURLY_WAGE, c: JAN_SALARY_MIN_HOURLY_WAGE, d: propsIntValueToString(testSalaryProp.MinHourlyWage()) } )
		VECT_SOCIAL_MAX_ANNUALS_BASIS = append(VECT_SOCIAL_MAX_ANNUALS_BASIS, tupleValues {a: testIntYear, b: MES_SOCIAL_MAX_ANNUALS_BASIS, c: JAN_SOCIAL_MAX_ANNUALS_BASIS, d: propsIntValueToString(testSocialProp.MaxAnnualsBasis()) } )
		VECT_SOCIAL_FACTOR_EMPLOYER = append(VECT_SOCIAL_FACTOR_EMPLOYER, tupleValues {a: testIntYear, b: MES_SOCIAL_FACTOR_EMPLOYER, c: JAN_SOCIAL_FACTOR_EMPLOYER, d: propsDecValueToString(testSocialProp.FactorEmployer()) } )
		VECT_SOCIAL_FACTOR_EMPLOYER_HIGHER = append(VECT_SOCIAL_FACTOR_EMPLOYER_HIGHER, tupleValues {a: testIntYear, b: MES_SOCIAL_FACTOR_EMPLOYER_HIGHER, c: JAN_SOCIAL_FACTOR_EMPLOYER_HIGHER, d: propsDecValueToString(testSocialProp.FactorEmployerHigher()) } )
		VECT_SOCIAL_FACTOR_EMPLOYEE = append(VECT_SOCIAL_FACTOR_EMPLOYEE, tupleValues {a: testIntYear, b: MES_SOCIAL_FACTOR_EMPLOYEE, c: JAN_SOCIAL_FACTOR_EMPLOYEE, d: propsDecValueToString(testSocialProp.FactorEmployee()) } )
		VECT_SOCIAL_FACTOR_EMPLOYEE_GARANT = append(VECT_SOCIAL_FACTOR_EMPLOYEE_GARANT, tupleValues {a: testIntYear, b: MES_SOCIAL_FACTOR_EMPLOYEE_GARANT, c: JAN_SOCIAL_FACTOR_EMPLOYEE_GARANT, d: propsDecValueToString(testSocialProp.FactorEmployeeGarant()) } )
		VECT_SOCIAL_FACTOR_EMPLOYEE_REDUCE = append(VECT_SOCIAL_FACTOR_EMPLOYEE_REDUCE, tupleValues {a: testIntYear, b: MES_SOCIAL_FACTOR_EMPLOYEE_REDUCE, c: JAN_SOCIAL_FACTOR_EMPLOYEE_REDUCE, d: propsDecValueToString(testSocialProp.FactorEmployeeReduce()) } )
		VECT_SOCIAL_MARGIN_INCOME_EMP = append(VECT_SOCIAL_MARGIN_INCOME_EMP, tupleValues {a: testIntYear, b: MES_SOCIAL_MARGIN_INCOME_EMP, c: JAN_SOCIAL_MARGIN_INCOME_EMP, d: propsIntValueToString(testSocialProp.MarginIncomeEmp()) } )
		VECT_SOCIAL_MARGIN_INCOME_AGR = append(VECT_SOCIAL_MARGIN_INCOME_AGR, tupleValues {a: testIntYear, b: MES_SOCIAL_MARGIN_INCOME_AGR, c: JAN_SOCIAL_MARGIN_INCOME_AGR, d: propsIntValueToString(testSocialProp.MarginIncomeAgr()) } )
		VECT_TAXING_ALLOWANCE_PAYER = append(VECT_TAXING_ALLOWANCE_PAYER, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_PAYER, c: JAN_TAXING_ALLOWANCE_PAYER, d: propsIntValueToString(testTaxingProp.AllowancePayer()) } )
		VECT_TAXING_ALLOWANCE_DISAB_1ST = append(VECT_TAXING_ALLOWANCE_DISAB_1ST, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_DISAB_1ST, c: JAN_TAXING_ALLOWANCE_DISAB_1ST, d: propsIntValueToString(testTaxingProp.AllowanceDisab1st()) } )
		VECT_TAXING_ALLOWANCE_DISAB_2ND = append(VECT_TAXING_ALLOWANCE_DISAB_2ND, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_DISAB_2ND, c: JAN_TAXING_ALLOWANCE_DISAB_2ND, d: propsIntValueToString(testTaxingProp.AllowanceDisab2nd()) } )
		VECT_TAXING_ALLOWANCE_DISAB_3RD = append(VECT_TAXING_ALLOWANCE_DISAB_3RD, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_DISAB_3RD, c: JAN_TAXING_ALLOWANCE_DISAB_3RD, d: propsIntValueToString(testTaxingProp.AllowanceDisab3rd()) } )
		VECT_TAXING_ALLOWANCE_STUDY = append(VECT_TAXING_ALLOWANCE_STUDY, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_STUDY, c: JAN_TAXING_ALLOWANCE_STUDY, d: propsIntValueToString(testTaxingProp.AllowanceStudy()) } )
		VECT_TAXING_ALLOWANCE_CHILD_1ST = append(VECT_TAXING_ALLOWANCE_CHILD_1ST, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_CHILD_1ST, c: JAN_TAXING_ALLOWANCE_CHILD_1ST, d: propsIntValueToString(testTaxingProp.AllowanceChild1st()) } )
		VECT_TAXING_ALLOWANCE_CHILD_2ND = append(VECT_TAXING_ALLOWANCE_CHILD_2ND, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_CHILD_2ND, c: JAN_TAXING_ALLOWANCE_CHILD_2ND, d: propsIntValueToString(testTaxingProp.AllowanceChild2nd()) } )
		VECT_TAXING_ALLOWANCE_CHILD_3RD = append(VECT_TAXING_ALLOWANCE_CHILD_3RD, tupleValues {a: testIntYear, b: MES_TAXING_ALLOWANCE_CHILD_3RD, c: JAN_TAXING_ALLOWANCE_CHILD_3RD, d: propsIntValueToString(testTaxingProp.AllowanceChild3rd()) } )
		VECT_TAXING_FACTOR_ADVANCES = append(VECT_TAXING_FACTOR_ADVANCES, tupleValues {a: testIntYear, b: MES_TAXING_FACTOR_ADVANCES, c: JAN_TAXING_FACTOR_ADVANCES, d: propsDecValueToString(testTaxingProp.FactorAdvances()) } )
		VECT_TAXING_FACTOR_WITHHOLD = append(VECT_TAXING_FACTOR_WITHHOLD, tupleValues {a: testIntYear, b: MES_TAXING_FACTOR_WITHHOLD, c: JAN_TAXING_FACTOR_WITHHOLD, d: propsDecValueToString(testTaxingProp.FactorWithhold()) } )
		VECT_TAXING_FACTOR_SOLIDARY = append(VECT_TAXING_FACTOR_SOLIDARY, tupleValues {a: testIntYear, b: MES_TAXING_FACTOR_SOLIDARY, c: JAN_TAXING_FACTOR_SOLIDARY, d: propsDecValueToString(testTaxingProp.FactorSolidary()) } )
		VECT_TAXING_FACTOR_TAXRATE2 = append(VECT_TAXING_FACTOR_TAXRATE2, tupleValues {a: testIntYear, b: MES_TAXING_FACTOR_TAXRATE2, c: JAN_TAXING_FACTOR_TAXRATE2, d: propsDecValueToString(testTaxingProp.FactorTaxRate2()) } )
		VECT_TAXING_MIN_AMOUNT_OF_TAXBONUS = append(VECT_TAXING_MIN_AMOUNT_OF_TAXBONUS, tupleValues {a: testIntYear, b: MES_TAXING_MIN_AMOUNT_OF_TAXBONUS, c: JAN_TAXING_MIN_AMOUNT_OF_TAXBONUS, d: propsIntValueToString(testTaxingProp.MinAmountOfTaxBonus()) } )
		VECT_TAXING_MAX_AMOUNT_OF_TAXBONUS = append(VECT_TAXING_MAX_AMOUNT_OF_TAXBONUS, tupleValues {a: testIntYear, b: MES_TAXING_MAX_AMOUNT_OF_TAXBONUS, c: JAN_TAXING_MAX_AMOUNT_OF_TAXBONUS, d: propsIntValueToString(testTaxingProp.MaxAmountOfTaxBonus()) } )
		VECT_TAXING_MARGIN_INCOME_OF_TAXBONUS = append(VECT_TAXING_MARGIN_INCOME_OF_TAXBONUS, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_TAXBONUS, c: JAN_TAXING_MARGIN_INCOME_OF_TAXBONUS, d: propsIntValueToString(testTaxingProp.MarginIncomeOfTaxBonus()) } )
		VECT_TAXING_MARGIN_INCOME_OF_ROUNDING = append(VECT_TAXING_MARGIN_INCOME_OF_ROUNDING, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_ROUNDING, c: JAN_TAXING_MARGIN_INCOME_OF_ROUNDING, d: propsIntValueToString(testTaxingProp.MarginIncomeOfRounding()) } )
		VECT_TAXING_MARGIN_INCOME_OF_WITHHOLD = append(VECT_TAXING_MARGIN_INCOME_OF_WITHHOLD, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_WITHHOLD, c: JAN_TAXING_MARGIN_INCOME_OF_WITHHOLD, d: propsIntValueToString(testTaxingProp.MarginIncomeOfWithhold()) } )
		VECT_TAXING_MARGIN_INCOME_OF_SOLIDARY = append(VECT_TAXING_MARGIN_INCOME_OF_SOLIDARY, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_SOLIDARY, c: JAN_TAXING_MARGIN_INCOME_OF_SOLIDARY, d: propsIntValueToString(testTaxingProp.MarginIncomeOfSolidary()) } )
		VECT_TAXING_MARGIN_INCOME_OF_TAXRATE2 = append(VECT_TAXING_MARGIN_INCOME_OF_TAXRATE2, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_TAXRATE2, c: JAN_TAXING_MARGIN_INCOME_OF_TAXRATE2, d: propsIntValueToString(testTaxingProp.MarginIncomeOfTaxRate2()) } )
		VECT_TAXING_MARGIN_INCOME_OF_WHT_EMP = append(VECT_TAXING_MARGIN_INCOME_OF_WHT_EMP, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_WHT_EMP, c: JAN_TAXING_MARGIN_INCOME_OF_WHT_EMP, d: propsIntValueToString(testTaxingProp.MarginIncomeOfWthEmp()) } )
		VECT_TAXING_MARGIN_INCOME_OF_WHT_AGR = append(VECT_TAXING_MARGIN_INCOME_OF_WHT_AGR, tupleValues {a: testIntYear, b: MES_TAXING_MARGIN_INCOME_OF_WHT_AGR, c: JAN_TAXING_MARGIN_INCOME_OF_WHT_AGR, d: propsIntValueToString(testTaxingProp.MarginIncomeOfWthAgr()) } )
	}
	var tableData = []tupleListValues {
		tupleListValues { a: HEALTH_MIN_MONTHLY_BASIS, b: VECT_HEALTH_MIN_MONTHLY_BASIS },
		tupleListValues { a: HEALTH_MAX_ANNUALS_BASIS, b: VECT_HEALTH_MAX_ANNUALS_BASIS },
		tupleListValues { a: HEALTH_LIM_MONTHLY_STATE, b: VECT_HEALTH_LIM_MONTHLY_STATE },
		tupleListValues { a: HEALTH_LIM_MONTHLY_DIS50, b: VECT_HEALTH_LIM_MONTHLY_DIS50 },
		tupleListValues { a: HEALTH_FACTOR_COMPOUND, b: VECT_HEALTH_FACTOR_COMPOUND },
		tupleListValues { a: HEALTH_FACTOR_EMPLOYEE, b: VECT_HEALTH_FACTOR_EMPLOYEE },
		tupleListValues { a: HEALTH_MARGIN_INCOME_EMP, b: VECT_HEALTH_MARGIN_INCOME_EMP },
		tupleListValues { a: HEALTH_MARGIN_INCOME_AGR, b: VECT_HEALTH_MARGIN_INCOME_AGR },
		tupleListValues { a: SALARY_WORKING_SHIFT_WEEK, b: VECT_SALARY_WORKING_SHIFT_WEEK },
		tupleListValues { a: SALARY_WORKING_SHIFT_TIME, b: VECT_SALARY_WORKING_SHIFT_TIME },
		tupleListValues { a: SALARY_MIN_MONTHLY_WAGE, b: VECT_SALARY_MIN_MONTHLY_WAGE },
		tupleListValues { a: SALARY_MIN_HOURLY_WAGE, b: VECT_SALARY_MIN_HOURLY_WAGE },
		tupleListValues { a: SOCIAL_MAX_ANNUALS_BASIS, b: VECT_SOCIAL_MAX_ANNUALS_BASIS },
		tupleListValues { a: SOCIAL_FACTOR_EMPLOYER, b: VECT_SOCIAL_FACTOR_EMPLOYER },
		tupleListValues { a: SOCIAL_FACTOR_EMPLOYER_HIGHER, b: VECT_SOCIAL_FACTOR_EMPLOYER_HIGHER },
		tupleListValues { a: SOCIAL_FACTOR_EMPLOYEE, b: VECT_SOCIAL_FACTOR_EMPLOYEE },
		tupleListValues { a: SOCIAL_FACTOR_EMPLOYEE_GARANT, b: VECT_SOCIAL_FACTOR_EMPLOYEE_GARANT },
		tupleListValues { a: SOCIAL_FACTOR_EMPLOYEE_REDUCE, b: VECT_SOCIAL_FACTOR_EMPLOYEE_REDUCE },
		tupleListValues { a: SOCIAL_MARGIN_INCOME_EMP, b: VECT_SOCIAL_MARGIN_INCOME_EMP },
		tupleListValues { a: SOCIAL_MARGIN_INCOME_AGR, b: VECT_SOCIAL_MARGIN_INCOME_AGR },
		tupleListValues { a: TAXING_ALLOWANCE_PAYER, b: VECT_TAXING_ALLOWANCE_PAYER },
		tupleListValues { a: TAXING_ALLOWANCE_DISAB_1ST, b: VECT_TAXING_ALLOWANCE_DISAB_1ST },
		tupleListValues { a: TAXING_ALLOWANCE_DISAB_2ND, b: VECT_TAXING_ALLOWANCE_DISAB_2ND },
		tupleListValues { a: TAXING_ALLOWANCE_DISAB_3RD, b: VECT_TAXING_ALLOWANCE_DISAB_3RD },
		tupleListValues { a: TAXING_ALLOWANCE_STUDY, b: VECT_TAXING_ALLOWANCE_STUDY },
		tupleListValues { a: TAXING_ALLOWANCE_CHILD_1ST, b: VECT_TAXING_ALLOWANCE_CHILD_1ST },
		tupleListValues { a: TAXING_ALLOWANCE_CHILD_2ND, b: VECT_TAXING_ALLOWANCE_CHILD_2ND },
		tupleListValues { a: TAXING_ALLOWANCE_CHILD_3RD, b: VECT_TAXING_ALLOWANCE_CHILD_3RD },
		tupleListValues { a: TAXING_FACTOR_ADVANCES, b: VECT_TAXING_FACTOR_ADVANCES },
		tupleListValues { a: TAXING_FACTOR_WITHHOLD, b: VECT_TAXING_FACTOR_WITHHOLD },
		tupleListValues { a: TAXING_FACTOR_SOLIDARY, b: VECT_TAXING_FACTOR_SOLIDARY },
		tupleListValues { a: TAXING_FACTOR_TAXRATE2, b: VECT_TAXING_FACTOR_TAXRATE2 },
		tupleListValues { a: TAXING_MIN_AMOUNT_OF_TAXBONUS, b: VECT_TAXING_MIN_AMOUNT_OF_TAXBONUS },
		tupleListValues { a: TAXING_MAX_AMOUNT_OF_TAXBONUS, b: VECT_TAXING_MAX_AMOUNT_OF_TAXBONUS },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_TAXBONUS, b: VECT_TAXING_MARGIN_INCOME_OF_TAXBONUS },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_ROUNDING, b: VECT_TAXING_MARGIN_INCOME_OF_ROUNDING },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_WITHHOLD, b: VECT_TAXING_MARGIN_INCOME_OF_WITHHOLD },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_SOLIDARY, b: VECT_TAXING_MARGIN_INCOME_OF_SOLIDARY },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_TAXRATE2, b: VECT_TAXING_MARGIN_INCOME_OF_TAXRATE2 },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_WHT_EMP, b: VECT_TAXING_MARGIN_INCOME_OF_WHT_EMP },
		tupleListValues { a: TAXING_MARGIN_INCOME_OF_WHT_AGR, b: VECT_TAXING_MARGIN_INCOME_OF_WHT_AGR },
	}

	for _, data := range tableData {
		exportHistoryTerm(testProtokol, headerData, &data)
	}
}
