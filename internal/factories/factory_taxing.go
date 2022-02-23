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

type IFactoryTaxing interface {
	DefaultProvider() providers.IProviderTaxing
	EmptyPeriodProps() props.IPropsTaxing

	GetProps(period types.IPeriod) (props.IPropsTaxing, bool)
}

type factoryTaxing struct {
	defaultProvider   providers.IProviderTaxing
	emptyPeriodProps props.IPropsTaxing

	versions map[providers.VERSION]providers.IProviderTaxing
}

func (f factoryTaxing) DefaultProvider() providers.IProviderTaxing {
	return f.defaultProvider
}

func (f factoryTaxing) EmptyPeriodProps() props.IPropsTaxing {
	return f.emptyPeriodProps
}

func (f factoryTaxing) GetProps(period types.IPeriod) (props.IPropsTaxing, bool) {
	provider, exists := f.getProvider(period, f.defaultProvider)

	if exists == false {
		return f.emptyPeriodProps, true
	}
	return provider.GetProps(period), true
}

func (f factoryTaxing) getProvider(period types.IPeriod, defProvider providers.IProviderTaxing) (providers.IProviderTaxing, bool) {
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

func NewFactoryTaxing() IFactoryTaxing {
	return factoryTaxing{
		defaultProvider:   period2022.NewProviderTaxing2022(),
		emptyPeriodProps: props.EmptyPropsTaxing(),
		versions: map[providers.VERSION]providers.IProviderTaxing{
			providers.VERSION(period2010.TAXING_VERSION_CODE): period2010.NewProviderTaxing2010(),
			providers.VERSION(period2011.TAXING_VERSION_CODE): period2011.NewProviderTaxing2011(),
			providers.VERSION(period2012.TAXING_VERSION_CODE): period2012.NewProviderTaxing2012(),
			providers.VERSION(period2013.TAXING_VERSION_CODE): period2013.NewProviderTaxing2013(),
			providers.VERSION(period2014.TAXING_VERSION_CODE): period2014.NewProviderTaxing2014(),
			providers.VERSION(period2015.TAXING_VERSION_CODE): period2015.NewProviderTaxing2015(),
			providers.VERSION(period2016.TAXING_VERSION_CODE): period2016.NewProviderTaxing2016(),
			providers.VERSION(period2017.TAXING_VERSION_CODE): period2017.NewProviderTaxing2017(),
			providers.VERSION(period2018.TAXING_VERSION_CODE): period2018.NewProviderTaxing2018(),
			providers.VERSION(period2019.TAXING_VERSION_CODE): period2019.NewProviderTaxing2019(),
			providers.VERSION(period2020.TAXING_VERSION_CODE): period2020.NewProviderTaxing2020(),
			providers.VERSION(period2021.TAXING_VERSION_CODE): period2021.NewProviderTaxing2021(),
			providers.VERSION(period2022.TAXING_VERSION_CODE): period2022.NewProviderTaxing2022(),
		},
	}
}
