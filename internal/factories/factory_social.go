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

type IFactorySocial interface {
	DefaultProvider() providers.IProviderSocial
	EmptyPeriodProps() props.IPropsSocial

	GetProps(period types.IPeriod) (props.IPropsSocial, bool)
}

type factorySocial struct {
	defaultProvider   providers.IProviderSocial
	emptyPeriodProps props.IPropsSocial

	versions map[providers.VERSION]providers.IProviderSocial
}

func (f factorySocial) DefaultProvider() providers.IProviderSocial {
	return f.defaultProvider
}

func (f factorySocial) EmptyPeriodProps() props.IPropsSocial {
	return f.emptyPeriodProps
}

func (f factorySocial) GetProps(period types.IPeriod) (props.IPropsSocial, bool) {
	provider, exists := f.getProvider(period, f.defaultProvider)

	if exists == false {
		return f.emptyPeriodProps, true
	}
	return provider.GetProps(period), true
}

func (f factorySocial) getProvider(period types.IPeriod, defProvider providers.IProviderSocial) (providers.IProviderSocial, bool) {
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

func NewFactorySocial() IFactorySocial {
	return factorySocial{
		defaultProvider:   period2022.NewProviderSocial2022(),
		emptyPeriodProps: props.EmptyPropsSocial(),
		versions: map[providers.VERSION]providers.IProviderSocial{
			providers.VERSION(period2010.SOCIAL_VERSION_CODE): period2010.NewProviderSocial2010(),
			providers.VERSION(period2011.SOCIAL_VERSION_CODE): period2011.NewProviderSocial2011(),
			providers.VERSION(period2012.SOCIAL_VERSION_CODE): period2012.NewProviderSocial2012(),
			providers.VERSION(period2013.SOCIAL_VERSION_CODE): period2013.NewProviderSocial2013(),
			providers.VERSION(period2014.SOCIAL_VERSION_CODE): period2014.NewProviderSocial2014(),
			providers.VERSION(period2015.SOCIAL_VERSION_CODE): period2015.NewProviderSocial2015(),
			providers.VERSION(period2016.SOCIAL_VERSION_CODE): period2016.NewProviderSocial2016(),
			providers.VERSION(period2017.SOCIAL_VERSION_CODE): period2017.NewProviderSocial2017(),
			providers.VERSION(period2018.SOCIAL_VERSION_CODE): period2018.NewProviderSocial2018(),
			providers.VERSION(period2019.SOCIAL_VERSION_CODE): period2019.NewProviderSocial2019(),
			providers.VERSION(period2020.SOCIAL_VERSION_CODE): period2020.NewProviderSocial2020(),
			providers.VERSION(period2021.SOCIAL_VERSION_CODE): period2021.NewProviderSocial2021(),
			providers.VERSION(period2022.SOCIAL_VERSION_CODE): period2022.NewProviderSocial2022(),
		},
	}
}

