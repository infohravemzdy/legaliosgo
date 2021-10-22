package factories

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IBundleProps interface {
	GetPeriodProps() types.IPeriod

	GetSalaryProps() props.IPropsSalary
	GetHealthProps() props.IPropsHealth
	GetSocialProps() props.IPropsSocial
	GetTaxingProps() props.IPropsTaxing
}

type bundleProps struct {
	periodProps types.IPeriod
	salaryProps props.IPropsSalary
	healthProps props.IPropsHealth
	socialProps props.IPropsSocial
	taxingProps props.IPropsTaxing
}

func (b bundleProps) GetPeriodProps() types.IPeriod {
	return b.periodProps
}

func (b bundleProps) GetSalaryProps() props.IPropsSalary {
	return b.salaryProps
}

func (b bundleProps) GetHealthProps() props.IPropsHealth {
	return b.healthProps
}

func (b bundleProps) GetSocialProps() props.IPropsSocial {
	return b.socialProps
}

func (b bundleProps) GetTaxingProps() props.IPropsTaxing {
	return b.taxingProps
}


func NewBundleProps(period types.IPeriod,
	bundleSalary props.IPropsSalary,
	bundleHealth props.IPropsHealth,
	bundleSocial props.IPropsSocial,
	bundleTaxing props.IPropsTaxing) IBundleProps {
	return bundleProps{ periodProps: period,
		salaryProps: bundleSalary,
		healthProps: bundleHealth,
		socialProps: bundleSocial,
		taxingProps: bundleTaxing }
}

func EmptyBundleProps(period types.IPeriod) IBundleProps {
	return bundleProps{ periodProps: period,
		salaryProps: props.EmptyPropsSalary(),
		healthProps: props.EmptyPropsHealth(),
		socialProps: props.EmptyPropsSocial(),
		taxingProps: props.EmptyPropsTaxing(),
	}
}

