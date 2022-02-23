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

type IFactoryHealth interface {
	DefaultProvider() providers.IProviderHealth
	EmptyPeriodProps() props.IPropsHealth

	GetProps(period types.IPeriod) (props.IPropsHealth, bool)
}

type factoryHealth struct {
	defaultProvider   providers.IProviderHealth
	emptyPeriodProps props.IPropsHealth

	versions map[providers.VERSION]providers.IProviderHealth
}

func (f factoryHealth) DefaultProvider() providers.IProviderHealth {
	return f.defaultProvider
}

func (f factoryHealth) EmptyPeriodProps() props.IPropsHealth {
	return f.emptyPeriodProps
}

func (f factoryHealth) GetProps(period types.IPeriod) (props.IPropsHealth, bool) {
	provider, exists := f.getProvider(period, f.defaultProvider)

	if exists == false {
		return f.emptyPeriodProps, true
	}
	return provider.GetProps(period), true
}

func (f factoryHealth) getProvider(period types.IPeriod, defProvider providers.IProviderHealth) (providers.IProviderHealth, bool) {
	provider, exists := f.versions[providers.VERSION(period.Year())]

	if exists == false {
		if period.Year() >= defProvider.GetVersionValue() {
			return defProvider, true
		}
		return nil, false
	}
	if period.Year() != provider.GetVersionValue() {
		return nil, false
	}
	return provider, true
}

func NewFactoryHealth() IFactoryHealth {
	return factoryHealth{
		defaultProvider:   period2022.NewProviderHealth2022(),
		emptyPeriodProps: props.EmptyPropsHealth(),
		versions: map[providers.VERSION]providers.IProviderHealth{
			providers.VERSION(period2010.HEALTH_VERSION_CODE): period2010.NewProviderHealth2010(),
			providers.VERSION(period2011.HEALTH_VERSION_CODE): period2011.NewProviderHealth2011(),
			providers.VERSION(period2012.HEALTH_VERSION_CODE): period2012.NewProviderHealth2012(),
			providers.VERSION(period2013.HEALTH_VERSION_CODE): period2013.NewProviderHealth2013(),
			providers.VERSION(period2014.HEALTH_VERSION_CODE): period2014.NewProviderHealth2014(),
			providers.VERSION(period2015.HEALTH_VERSION_CODE): period2015.NewProviderHealth2015(),
			providers.VERSION(period2016.HEALTH_VERSION_CODE): period2016.NewProviderHealth2016(),
			providers.VERSION(period2017.HEALTH_VERSION_CODE): period2017.NewProviderHealth2017(),
			providers.VERSION(period2018.HEALTH_VERSION_CODE): period2018.NewProviderHealth2018(),
			providers.VERSION(period2019.HEALTH_VERSION_CODE): period2019.NewProviderHealth2019(),
			providers.VERSION(period2020.HEALTH_VERSION_CODE): period2020.NewProviderHealth2020(),
			providers.VERSION(period2021.HEALTH_VERSION_CODE): period2021.NewProviderHealth2021(),
			providers.VERSION(period2022.HEALTH_VERSION_CODE): period2022.NewProviderHealth2022(),
		},
	}
}

