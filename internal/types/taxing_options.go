package types

type TaxDeclSignOption int16
const DECL_TAX_NO_SIGNED = TaxDeclSignOption(0)
const DECL_TAX_DO_SIGNED = TaxDeclSignOption(1)

type TaxNoneSignOption int16
const NOSIGN_TAX_WITHHOLD = TaxNoneSignOption(0)
const NOSIGN_TAX_ADVANCES = TaxNoneSignOption(1)


type TaxDeclBenfOption int16
const DECL_TAX_BENEF0 = TaxDeclBenfOption(0)
const DECL_TAX_BENEF1 = TaxDeclBenfOption(1)

type TaxDeclDisabOption int16
const DISB_TAX_BENEF0 = TaxDeclDisabOption(0)
const DISB_TAX_DISAB1 = TaxDeclDisabOption(1)
const DISB_TAX_DISAB2 = TaxDeclDisabOption(2)
const DISB_TAX_DISAB3 = TaxDeclDisabOption(3)


