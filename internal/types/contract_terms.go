package types

type WorkContractTerms int16
const WORKTERM_EMPLOYMENT_1 = WorkContractTerms(0)
const WORKTERM_CONTRACTER_A = WorkContractTerms(1)
const WORKTERM_CONTRACTER_T = WorkContractTerms(2)
const WORKTERM_PARTNER_STAT = WorkContractTerms(3)


type WorkTaxingTerms int16
const TAXING_TERM_BY_CONTRACT = WorkTaxingTerms(0)
const TAXING_TERM_EMPLOYMENTS = WorkTaxingTerms(1)
const TAXING_TERM_AGREEM_TASK = WorkTaxingTerms(2)
const TAXING_TERM_STATUT_PART = WorkTaxingTerms(3)


type WorkHealthTerms int16
const HEALTH_TERM_BY_CONTRACT = WorkHealthTerms(0)
const HEALTH_TERM_EMPLOYMENTS = WorkHealthTerms(1)
const HEALTH_TERM_AGREEM_WORK = WorkHealthTerms(2)
const HEALTH_TERM_AGREEM_TASK = WorkHealthTerms(3)


type WorkSocialTerms int16
const SOCIAL_TERM_BY_CONTRACT = WorkSocialTerms(0)
const SOCIAL_TERM_EMPLOYMENTS = WorkSocialTerms(1)
const SOCIAL_TERM_SMALLS_EMPL = WorkSocialTerms(2)
const SOCIAL_TERM_SHORTS_MEET = WorkSocialTerms(3)
const SOCIAL_TERM_SHORTS_DENY = WorkSocialTerms(4)
const SOCIAL_TERM_AGREEM_TASK = WorkSocialTerms(5)
