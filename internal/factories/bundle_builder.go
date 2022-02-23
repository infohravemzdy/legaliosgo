package factories

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

const (
	MIN_VERSION int16 = 2010
)

type IBundleBuilder interface {
	GetBundle(period types.IPeriod) (IBundleProps, bool)
}

type bundleBuilder struct {
	salaryFactory IFactorySalary
	healthFactory IFactoryHealth
	socialFactory IFactorySocial
	taxingFactory IFactoryTaxing
}

func NewBundleBuilder() IBundleBuilder {
	return bundleBuilder{
		salaryFactory: NewFactorySalary(),
		healthFactory: NewFactoryHealth(),
		socialFactory: NewFactorySocial(),
		taxingFactory: NewFactoryTaxing(),
	}
}

func (b bundleBuilder) GetBundle(period types.IPeriod) (IBundleProps, bool) {
	salaryProps, salaryExists := b.salaryFactory.GetProps(period)
	healthProps, healthExists := b.healthFactory.GetProps(period)
	socialProps, socialExists := b.socialFactory.GetProps(period)
	taxingProps, taxingExists := b.taxingFactory.GetProps(period)

	if salaryExists == false || healthExists == false || socialExists  == false || taxingExists == false {
		return nil, false
	}
	if b.isValidBundle(salaryProps, healthProps, socialProps, taxingProps) == false {
		return nil, false
	}
	return NewBundleProps(period, salaryProps, healthProps, socialProps, taxingProps), true
}

func (b bundleBuilder) isNullOrEmpty(props props.IProps) bool {
	return props == nil || props.GetVersionValue() < MIN_VERSION
}

func (b bundleBuilder) isValidBundle(salary props.IPropsSalary, health props.IPropsHealth, social props.IPropsSocial, taxing props.IPropsTaxing) bool {
	if b.isNullOrEmpty(salary) {
		return false
	}
	if b.isNullOrEmpty(health) {
		return false
	}
	if b.isNullOrEmpty(social) {
		return false
	}
	if b.isNullOrEmpty(taxing) {
		return false
	}
	return true
}

func  (b bundleBuilder) getSalaryProps(period types.IPeriod) (props.IPropsSalary, bool) {
	return b.salaryFactory.GetProps(period)
}

func  (b bundleBuilder) getHealthProps(period types.IPeriod) (props.IPropsHealth, bool) {
	return b.healthFactory.GetProps(period)
}

func  (b bundleBuilder) getSocialProps(period types.IPeriod) (props.IPropsSocial, bool) {
	return b.socialFactory.GetProps(period)
}

func  (b bundleBuilder) getTaxingProps(period types.IPeriod) (props.IPropsTaxing, bool) {
	return b.taxingFactory.GetProps(period)
}
