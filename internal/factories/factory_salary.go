package factories

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2010"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2011"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2012"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2013"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2014"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2015"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2016"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2017"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2018"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2019"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2020"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2021"
	"github.com/mzdyhrave/legaliosgo/internal/providers/period2022"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IFactorySalary interface {
	DefaultProvider() providers.IProviderSalary
	EmptyPeriodProps() props.IPropsSalary

	GetProps(period types.IPeriod) (props.IPropsSalary, bool)
}

type factorySalary struct {
	defaultProvider   providers.IProviderSalary
	emptyPeriodProps props.IPropsSalary

	versions map[providers.VERSION]providers.IProviderSalary
}

func (f factorySalary) DefaultProvider() providers.IProviderSalary {
	return f.defaultProvider
}

func (f factorySalary) EmptyPeriodProps() props.IPropsSalary {
	return f.emptyPeriodProps
}

func (f factorySalary) GetProps(period types.IPeriod) (props.IPropsSalary, bool) {
	provider, exists := f.getProvider(period, f.defaultProvider)

	if exists == false {
		return f.emptyPeriodProps, true
	}
	return provider.GetProps(period), true
}

func (f factorySalary) getProvider(period types.IPeriod, defProvider providers.IProviderSalary) (providers.IProviderSalary, bool) {
	provider, exists := f.versions[providers.VERSION(period.Year())]

	if exists == false {
		if period.Year() > defProvider.GetVersionValue() {
			return defProvider, true
		}
		return nil, false
	}
	if period.Year() != provider.GetVersionValue() {
		return nil, false
	}
	return provider, true
}

func NewFactorySalary() IFactorySalary {
	return factorySalary{
		defaultProvider:   period2022.NewProviderSalary2022(),
		emptyPeriodProps: props.EmptyPropsSalary(),
		versions: map[providers.VERSION]providers.IProviderSalary{
			providers.VERSION(period2010.SALARY_VERSION_CODE): period2010.NewProviderSalary2010(),
			providers.VERSION(period2011.SALARY_VERSION_CODE): period2011.NewProviderSalary2011(),
			providers.VERSION(period2012.SALARY_VERSION_CODE): period2012.NewProviderSalary2012(),
			providers.VERSION(period2013.SALARY_VERSION_CODE): period2013.NewProviderSalary2013(),
			providers.VERSION(period2014.SALARY_VERSION_CODE): period2014.NewProviderSalary2014(),
			providers.VERSION(period2015.SALARY_VERSION_CODE): period2015.NewProviderSalary2015(),
			providers.VERSION(period2016.SALARY_VERSION_CODE): period2016.NewProviderSalary2016(),
			providers.VERSION(period2017.SALARY_VERSION_CODE): period2017.NewProviderSalary2017(),
			providers.VERSION(period2018.SALARY_VERSION_CODE): period2018.NewProviderSalary2018(),
			providers.VERSION(period2019.SALARY_VERSION_CODE): period2019.NewProviderSalary2019(),
			providers.VERSION(period2020.SALARY_VERSION_CODE): period2020.NewProviderSalary2020(),
			providers.VERSION(period2021.SALARY_VERSION_CODE): period2021.NewProviderSalary2021(),
			providers.VERSION(period2022.SALARY_VERSION_CODE): period2022.NewProviderSalary2022(),
		},
	}
}
