package history

import (
	"fmt"
	. "github.com/shopspring/decimal"
	"os"
	"path/filepath"
)

const (
	HISTORY_FOLDER_PATH = "../../../../test_history"
	HISTORY_FOLDER_NAME = "test_history"
	PARENT_HISTORY_FOLDER_NAME = "legaliosgo"
	HEALTH_MIN_MONTHLY_BASIS         int16 = 101
	HEALTH_MAX_ANNUALS_BASIS         int16 = 102
	HEALTH_LIM_MONTHLY_STATE         int16 = 103
	HEALTH_LIM_MONTHLY_DIS50         int16 = 104
	HEALTH_FACTOR_COMPOUND           int16 = 105
	HEALTH_FACTOR_EMPLOYEE           int16 = 106
	HEALTH_MARGIN_INCOME_EMP         int16 = 107
	HEALTH_MARGIN_INCOME_AGR         int16 = 108

	SALARY_WORKING_SHIFT_WEEK        int16 = 201
	SALARY_WORKING_SHIFT_TIME        int16 = 202
	SALARY_MIN_MONTHLY_WAGE          int16 = 203
	SALARY_MIN_HOURLY_WAGE           int16 = 204

	SOCIAL_MAX_ANNUALS_BASIS         int16 = 301
	SOCIAL_FACTOR_EMPLOYER           int16 = 302
	SOCIAL_FACTOR_EMPLOYER_HIGHER    int16 = 303
	SOCIAL_FACTOR_EMPLOYEE           int16 = 304
	SOCIAL_FACTOR_EMPLOYEE_GARANT    int16 = 305
	SOCIAL_FACTOR_EMPLOYEE_REDUCE    int16 = 306
	SOCIAL_MARGIN_INCOME_EMP         int16 = 307
	SOCIAL_MARGIN_INCOME_AGR         int16 = 308

	TAXING_ALLOWANCE_PAYER           int16 = 401
	TAXING_ALLOWANCE_DISAB_1ST       int16 = 402
	TAXING_ALLOWANCE_DISAB_2ND       int16 = 403
	TAXING_ALLOWANCE_DISAB_3RD       int16 = 404
	TAXING_ALLOWANCE_STUDY           int16 = 405
	TAXING_ALLOWANCE_CHILD_1ST       int16 = 406
	TAXING_ALLOWANCE_CHILD_2ND       int16 = 407
	TAXING_ALLOWANCE_CHILD_3RD       int16 = 408
	TAXING_FACTOR_ADVANCES           int16 = 409
	TAXING_FACTOR_WITHHOLD           int16 = 410
	TAXING_FACTOR_SOLIDARY           int16 = 411
	TAXING_FACTOR_TAXRATE2           int16 = 412
	TAXING_MIN_AMOUNT_OF_TAXBONUS    int16 = 413
	TAXING_MAX_AMOUNT_OF_TAXBONUS    int16 = 414
	TAXING_MARGIN_INCOME_OF_TAXBONUS int16 = 415
	TAXING_MARGIN_INCOME_OF_ROUNDING int16 = 416
	TAXING_MARGIN_INCOME_OF_WITHHOLD int16 = 417
	TAXING_MARGIN_INCOME_OF_SOLIDARY int16 = 418
	TAXING_MARGIN_INCOME_OF_TAXRATE2 int16 = 419
	TAXING_MARGIN_INCOME_OF_WHT_EMP  int16 = 420
	TAXING_MARGIN_INCOME_OF_WHT_AGR  int16 = 421
)

func createHistoryFile(baseName string, fileName string) (f *os.File, err error) {
	currPath, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	for  {
		_, lastName := filepath.Split(currPath)
		if lastName == PARENT_HISTORY_FOLDER_NAME || currPath == "." || currPath == "/" {
			break
		}
		currPath = filepath.Dir(currPath)
	}
	basePath := filepath.Join(currPath, HISTORY_FOLDER_NAME)
	filePath := filepath.Join(basePath, fileName)
	fullPath, err := filepath.Abs(filePath)
	println(fullPath)
	if err != nil {
		return nil, err
	}
	f, err = os.Create(fullPath)
	return f, err
}

type tupleChanges struct {
	a int16
	b bool
}

type tupleValues struct {
	a int16
	b int16
	c string
	d string
}

type tupleListChanges struct {
	a int16
	b []tupleChanges
}


type tupleListValues struct {
	a int16
	b []tupleValues
}

func exportHistoryStart(protokol *os.File, data []tupleChanges) {
	if  _, err := fmt.Fprintf(protokol,"History Term"); err != nil {
		return
	}
	for _, col := range data {
		if col.b {
			if  _, err := fmt.Fprintf(protokol,"\t%d Begin Value", col.a); err != nil {
				return
			}
			if  _, err := fmt.Fprintf(protokol,"\t%d Change Month", col.a); err != nil {
				return
			}
			if  _, err := fmt.Fprintf(protokol,"\t%d End Value", col.a); err != nil {
				return
			}
		} else {
			if  _, err := fmt.Fprintf(protokol,"\t%d Value", col.a); err != nil {
				return
			}
		}
	}
	if  _, err := fmt.Fprintf(protokol,"\n"); err != nil {
		return
	}
}
type findTermFunc func (termId int16) bool

func firstTermOrNull(source []tupleChanges, findFunc findTermFunc) (*tupleChanges, bool) {
	for _, v := range source {
		if findFunc(v.a) {
			return &v, true
		}
	}
	return nil, false
}

func exportHistoryTerm(protokol *os.File, head []tupleChanges, data *tupleListValues) {
	if  _, err := fmt.Fprintf(protokol,historyTermName(data.a)); err != nil {
		return
	}
	for _, col := range data.b {
		header, _ := firstTermOrNull(head, func (termId int16) bool {
			return termId == col.a
		})
		var yearOfChange bool = false
		if header != nil {
			yearOfChange = header.b
		}
		if  _, err := fmt.Fprintf(protokol,"\t%s", col.c); err != nil {
			return
		}
		if yearOfChange {
			if col.b == 0 {
				if  _, err := fmt.Fprintf(protokol,"\t"); err != nil {
					return
				}
			} else {
				if  _, err := fmt.Fprintf(protokol,"\t%d", col.b); err != nil {
					return
				}
			}
			if  _, err := fmt.Fprintf(protokol,"\t%s", col.d); err != nil {
				return
			}
		}
	}
	if  _, err := fmt.Fprintf(protokol,"\n"); err != nil {
		return
	}
}

func historyTermName(termId int16) string {
	switch termId {
	case HEALTH_MIN_MONTHLY_BASIS : return "01_Health_01_MinMonthlyBasis"
	case HEALTH_MAX_ANNUALS_BASIS : return "01_Health_02_MaxAnnualsBasis"
	case HEALTH_LIM_MONTHLY_STATE : return "01_Health_03_LimMonthlyState"
	case HEALTH_LIM_MONTHLY_DIS50 : return "01_Health_04_LimMonthlyDis50"
	case HEALTH_FACTOR_COMPOUND : return "01_Health_05_FactorCompound"
	case HEALTH_FACTOR_EMPLOYEE : return "01_Health_06_FactorEmployee"
	case HEALTH_MARGIN_INCOME_EMP : return "01_Health_07_MarginIncomeEmp"
	case HEALTH_MARGIN_INCOME_AGR : return "01_Health_08_MarginIncomeAgr"
	case SALARY_WORKING_SHIFT_WEEK : return "02_Salary_01_WorkingShiftWeek"
	case SALARY_WORKING_SHIFT_TIME : return "02_Salary_02_WorkingShiftTime"
	case SALARY_MIN_MONTHLY_WAGE : return "02_Salary_03_MinMonthlyWage"
	case SALARY_MIN_HOURLY_WAGE : return "02_Salary_04_MinHourlyWage"
	case SOCIAL_MAX_ANNUALS_BASIS : return "03_Social_01_MaxAnnualsBasis"
	case SOCIAL_FACTOR_EMPLOYER : return "03_Social_02_FactorEmployer"
	case SOCIAL_FACTOR_EMPLOYER_HIGHER : return "03_Social_03_FactorEmployerHigher"
	case SOCIAL_FACTOR_EMPLOYEE : return "03_Social_04_FactorEmployee"
	case SOCIAL_FACTOR_EMPLOYEE_GARANT : return "03_Social_05_FactorEmployeeGarant"
	case SOCIAL_FACTOR_EMPLOYEE_REDUCE : return "03_Social_06_FactorEmployeeReduce"
	case SOCIAL_MARGIN_INCOME_EMP : return "03_Social_07_MarginIncomeEmp"
	case SOCIAL_MARGIN_INCOME_AGR : return "03_Social_08_MarginIncomeAgr"
	case TAXING_ALLOWANCE_PAYER : return "04_Taxing_01_AllowancePayer"
	case TAXING_ALLOWANCE_DISAB_1ST : return "04_Taxing_02_AllowanceDisab1st"
	case TAXING_ALLOWANCE_DISAB_2ND : return "04_Taxing_03_AllowanceDisab2nd"
	case TAXING_ALLOWANCE_DISAB_3RD : return "04_Taxing_04_AllowanceDisab3rd"
	case TAXING_ALLOWANCE_STUDY : return "04_Taxing_05_AllowanceStudy"
	case TAXING_ALLOWANCE_CHILD_1ST : return "04_Taxing_06_AllowanceChild1st"
	case TAXING_ALLOWANCE_CHILD_2ND : return "04_Taxing_07_AllowanceChild2nd"
	case TAXING_ALLOWANCE_CHILD_3RD : return "04_Taxing_08_AllowanceChild3rd"
	case TAXING_FACTOR_ADVANCES : return "04_Taxing_09_FactorAdvances"
	case TAXING_FACTOR_WITHHOLD : return "04_Taxing_10_FactorWithhold"
	case TAXING_FACTOR_SOLIDARY : return "04_Taxing_11_FactorSolidary"
	case TAXING_FACTOR_TAXRATE2 : return "04_Taxing_12_FactorTaxRate2"
	case TAXING_MIN_AMOUNT_OF_TAXBONUS : return "04_Taxing_13_MinAmountOfTaxBonus"
	case TAXING_MAX_AMOUNT_OF_TAXBONUS : return "04_Taxing_14_MaxAmountOfTaxBonus"
	case TAXING_MARGIN_INCOME_OF_TAXBONUS : return "04_Taxing_15_MarginIncomeOfTaxBonus"
	case TAXING_MARGIN_INCOME_OF_ROUNDING : return "04_Taxing_16_MarginIncomeOfRounding"
	case TAXING_MARGIN_INCOME_OF_WITHHOLD : return "04_Taxing_17_MarginIncomeOfWithhold"
	case TAXING_MARGIN_INCOME_OF_SOLIDARY : return "04_Taxing_18_MarginIncomeOfSolidary"
	case TAXING_MARGIN_INCOME_OF_TAXRATE2 : return "04_Taxing_18_MarginIncomeOfTaxRate2"
	case TAXING_MARGIN_INCOME_OF_WHT_EMP : return "04_Taxing_20_MarginIncomeOfWthEmp"
	case TAXING_MARGIN_INCOME_OF_WHT_AGR : return "04_Taxing_21_MarginIncomeOfWthAgr"
	default : return "Unknown Term"
	}
}

func historyTermNameCZ(termId int16) string {
	switch termId {
	case HEALTH_MIN_MONTHLY_BASIS : return "01_Health_01 Minimální základ zdravotního pojištění na jednoho pracovníka"
	case HEALTH_MAX_ANNUALS_BASIS : return "01_Health_02 Maximální roční vyměřovací základ na jednoho pracovníka (tzv.strop)"
	case HEALTH_LIM_MONTHLY_STATE : return "01_Health_03 Vyměřovací základ ze kterého platí pojistné stát za státní pojištěnce (mateřská, studenti, důchodci)"
	case HEALTH_LIM_MONTHLY_DIS50 : return "01_Health_04 Vyměřovací základ ze kterého platí pojistné stát za státní pojištěnce (mateřská, studenti, důchodci) u zaměstnavatele zaměstnávajícího více než 50% osob OZP"
	case HEALTH_FACTOR_COMPOUND : return "01_Health_05 složená sazba zdravotního pojištění (zaměstnace + zaměstnavatele)"
	case HEALTH_FACTOR_EMPLOYEE : return "01_Health_06 podíl sazby zdravotního pojištění připadajícího na zaměstnace (1/FACTOR_EMPLOYEE)"
	case HEALTH_MARGIN_INCOME_EMP : return "01_Health_07 hranice příjmu pro vznik účasti na pojištění pro zaměstnace v pracovním poměru"
	case HEALTH_MARGIN_INCOME_AGR : return "01_Health_08 hranice příjmu pro vznik účasti na pojištění pro zaměstnace na dohodu"
	case SALARY_WORKING_SHIFT_WEEK : return "02_Salary_01 Počet pracovních dnů v týdnu"
	case SALARY_WORKING_SHIFT_TIME : return "02_Salary_02 Počet pracovních hodin denně"
	case SALARY_MIN_MONTHLY_WAGE : return "02_Salary_03 Minimální mzda měsíční"
	case SALARY_MIN_HOURLY_WAGE : return "02_Salary_04 Minimální mzda hodinová (100*Kč)"
	case SOCIAL_MAX_ANNUALS_BASIS : return "03_Social_01 Maximální roční vyměřovací základ na jednoho pracovníka (tzv.strop)"
	case SOCIAL_FACTOR_EMPLOYER : return "03_Social_02 Sazba - standardní sociálního pojištění - zaměstnavatele"
	case SOCIAL_FACTOR_EMPLOYER_HIGHER : return "03_Social_03 Sazba - vyší sociálního pojištění - zaměstnavatele"
	case SOCIAL_FACTOR_EMPLOYEE : return "03_Social_04 Sazba sociálního pojištění - zaměstnance"
	case SOCIAL_FACTOR_EMPLOYEE_GARANT : return "03_Social_05 Sazba důchodového spoření - zaměstnance - s důchodovým spořením"
	case SOCIAL_FACTOR_EMPLOYEE_REDUCE : return "03_Social_06 Snížení sazby sociálního pojištění - zaměstnance - s důchodovým spořením"
	case SOCIAL_MARGIN_INCOME_EMP : return "03_Social_07 hranice příjmu pro vznik účasti na pojištění pro zaměstnace v pracovním poměru"
	case SOCIAL_MARGIN_INCOME_AGR : return "03_Social_08 hranice příjmu pro vznik účasti na pojištění pro zaměstnace na dohodu"
	case TAXING_ALLOWANCE_PAYER : return "04_Taxing_01 Částka slevy na poplatníka"
	case TAXING_ALLOWANCE_DISAB_1ST : return "04_Taxing_02 Částka slevy na invaliditu 1.stupně poplatníka"
	case TAXING_ALLOWANCE_DISAB_2ND : return "04_Taxing_03 Částka slevy na invaliditu 2.stupně poplatníka"
	case TAXING_ALLOWANCE_DISAB_3RD : return "04_Taxing_04 Částka slevy na invaliditu 3.stupně poplatníka"
	case TAXING_ALLOWANCE_STUDY : return "04_Taxing_05 Částka slevy na poplatníka studenta"
	case TAXING_ALLOWANCE_CHILD_1ST : return "04_Taxing_06 Částka slevy na dítě 1.pořadí"
	case TAXING_ALLOWANCE_CHILD_2ND : return "04_Taxing_07 Částka slevy na dítě 2.pořadí"
	case TAXING_ALLOWANCE_CHILD_3RD : return "04_Taxing_08 Částka slevy na dítě 3.pořadí"
	case TAXING_FACTOR_ADVANCES : return "04_Taxing_09 Sazba daně na zálohový příjem"
	case TAXING_FACTOR_WITHHOLD : return "04_Taxing_10 Sazba daně na srážkový příjem"
	case TAXING_FACTOR_SOLIDARY : return "04_Taxing_11 Sazba daně na solidární zvýšení"
	case TAXING_FACTOR_TAXRATE2 : return "04_Taxing_12 Sazba daně pro druhé pásmo daně"
	case TAXING_MIN_AMOUNT_OF_TAXBONUS : return "04_Taxing_13 Minimální částka pro daňový bonus"
	case TAXING_MAX_AMOUNT_OF_TAXBONUS : return "04_Taxing_14 Maximální částka pro daňový bonus"
	case TAXING_MARGIN_INCOME_OF_TAXBONUS : return "04_Taxing_15 Minimální výše příjmu pro nároku na daňový bonus"
	case TAXING_MARGIN_INCOME_OF_ROUNDING : return "04_Taxing_16 Maximální výše příjmu pro zaokrouhlování"
	case TAXING_MARGIN_INCOME_OF_WITHHOLD : return "04_Taxing_17 Maximální výše příjmu pro srážkový příjem"
	case TAXING_MARGIN_INCOME_OF_SOLIDARY : return "04_Taxing_18 Minimální výše příjmu pro solidární zvýšení daně"
	case TAXING_MARGIN_INCOME_OF_TAXRATE2 : return "04_Taxing_18 Minimální výše příjmu pro druhé pásmo daně"
	case TAXING_MARGIN_INCOME_OF_WHT_EMP : return "04_Taxing_20 hranice příjmu pro srážkovou daň pro zaměstnace v pracovním poměru (nepodepsal prohlášení)"
	case TAXING_MARGIN_INCOME_OF_WHT_AGR : return "04_Taxing_21 hranice příjmu pro srážkovou daň pro zaměstnace na dohodu (nepodepsal prohlášení)"
	default : return "Neznámý Termín"
	}
}

func propsIntValueToString(value int32) string {
	return fmt.Sprintf("%d", value)
}

func propsDecValueToString(value Decimal) string {
	var intValue = value.Mul(NewFromInt32(100)).IntPart()
	return fmt.Sprintf("%d", intValue)
}
