//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"
import "github.com/stripe/stripe-go/v72/form"

type AccountExternalAccountsType string

// List of values that AccountExternalAccountsType can take
const (
	AccountExternalAccountsTypeBankAccount AccountExternalAccountsType = "bank_account"
	AccountExternalAccountsTypeCard        AccountExternalAccountsType = "card"
)

// The code for the type of error.
type AccountFutureRequirementsErrorCode string

// List of values that AccountFutureRequirementsErrorCode can take
const (
	AccountFutureRequirementsErrorCodeInvalidAddressCityStatePostalCode                      AccountFutureRequirementsErrorCode = "invalid_address_city_state_postal_code"
	AccountFutureRequirementsErrorCodeInvalidStreetAddress                                   AccountFutureRequirementsErrorCode = "invalid_street_address"
	AccountFutureRequirementsErrorCodeInvalidValueOther                                      AccountFutureRequirementsErrorCode = "invalid_value_other"
	AccountFutureRequirementsErrorCodeVerificationDocumentAddressMismatch                    AccountFutureRequirementsErrorCode = "verification_document_address_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentAddressMissing                     AccountFutureRequirementsErrorCode = "verification_document_address_missing"
	AccountFutureRequirementsErrorCodeVerificationDocumentCorrupt                            AccountFutureRequirementsErrorCode = "verification_document_corrupt"
	AccountFutureRequirementsErrorCodeVerificationDocumentCountryNotSupported                AccountFutureRequirementsErrorCode = "verification_document_country_not_supported"
	AccountFutureRequirementsErrorCodeVerificationDocumentDOBMismatch                        AccountFutureRequirementsErrorCode = "verification_document_dob_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentDuplicateType                      AccountFutureRequirementsErrorCode = "verification_document_duplicate_type"
	AccountFutureRequirementsErrorCodeVerificationDocumentExpired                            AccountFutureRequirementsErrorCode = "verification_document_expired"
	AccountFutureRequirementsErrorCodeVerificationDocumentFailedCopy                         AccountFutureRequirementsErrorCode = "verification_document_failed_copy"
	AccountFutureRequirementsErrorCodeVerificationDocumentFailedGreyscale                    AccountFutureRequirementsErrorCode = "verification_document_failed_greyscale"
	AccountFutureRequirementsErrorCodeVerificationDocumentFailedOther                        AccountFutureRequirementsErrorCode = "verification_document_failed_other"
	AccountFutureRequirementsErrorCodeVerificationDocumentFailedTestMode                     AccountFutureRequirementsErrorCode = "verification_document_failed_test_mode"
	AccountFutureRequirementsErrorCodeVerificationDocumentFraudulent                         AccountFutureRequirementsErrorCode = "verification_document_fraudulent"
	AccountFutureRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   AccountFutureRequirementsErrorCode = "verification_document_id_number_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentIDNumberMissing                    AccountFutureRequirementsErrorCode = "verification_document_id_number_missing"
	AccountFutureRequirementsErrorCodeVerificationDocumentIncomplete                         AccountFutureRequirementsErrorCode = "verification_document_incomplete"
	AccountFutureRequirementsErrorCodeVerificationDocumentInvalid                            AccountFutureRequirementsErrorCode = "verification_document_invalid"
	AccountFutureRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           AccountFutureRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	AccountFutureRequirementsErrorCodeVerificationDocumentManipulated                        AccountFutureRequirementsErrorCode = "verification_document_manipulated"
	AccountFutureRequirementsErrorCodeVerificationDocumentMissingBack                        AccountFutureRequirementsErrorCode = "verification_document_missing_back"
	AccountFutureRequirementsErrorCodeVerificationDocumentMissingFront                       AccountFutureRequirementsErrorCode = "verification_document_missing_front"
	AccountFutureRequirementsErrorCodeVerificationDocumentNameMismatch                       AccountFutureRequirementsErrorCode = "verification_document_name_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentNameMissing                        AccountFutureRequirementsErrorCode = "verification_document_name_missing"
	AccountFutureRequirementsErrorCodeVerificationDocumentNationalityMismatch                AccountFutureRequirementsErrorCode = "verification_document_nationality_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentNotReadable                        AccountFutureRequirementsErrorCode = "verification_document_not_readable"
	AccountFutureRequirementsErrorCodeVerificationDocumentNotSigned                          AccountFutureRequirementsErrorCode = "verification_document_not_signed"
	AccountFutureRequirementsErrorCodeVerificationDocumentNotUploaded                        AccountFutureRequirementsErrorCode = "verification_document_not_uploaded"
	AccountFutureRequirementsErrorCodeVerificationDocumentPhotoMismatch                      AccountFutureRequirementsErrorCode = "verification_document_photo_mismatch"
	AccountFutureRequirementsErrorCodeVerificationDocumentTooLarge                           AccountFutureRequirementsErrorCode = "verification_document_too_large"
	AccountFutureRequirementsErrorCodeVerificationDocumentTypeNotSupported                   AccountFutureRequirementsErrorCode = "verification_document_type_not_supported"
	AccountFutureRequirementsErrorCodeVerificationFailedAddressMatch                         AccountFutureRequirementsErrorCode = "verification_failed_address_match"
	AccountFutureRequirementsErrorCodeVerificationFailedBusinessIecNumber                    AccountFutureRequirementsErrorCode = "verification_failed_business_iec_number"
	AccountFutureRequirementsErrorCodeVerificationFailedDocumentMatch                        AccountFutureRequirementsErrorCode = "verification_failed_document_match"
	AccountFutureRequirementsErrorCodeVerificationFailedIDNumberMatch                        AccountFutureRequirementsErrorCode = "verification_failed_id_number_match"
	AccountFutureRequirementsErrorCodeVerificationFailedKeyedIdentity                        AccountFutureRequirementsErrorCode = "verification_failed_keyed_identity"
	AccountFutureRequirementsErrorCodeVerificationFailedKeyedMatch                           AccountFutureRequirementsErrorCode = "verification_failed_keyed_match"
	AccountFutureRequirementsErrorCodeVerificationFailedNameMatch                            AccountFutureRequirementsErrorCode = "verification_failed_name_match"
	AccountFutureRequirementsErrorCodeVerificationFailedOther                                AccountFutureRequirementsErrorCode = "verification_failed_other"
	AccountFutureRequirementsErrorCodeVerificationFailedTaxIDMatch                           AccountFutureRequirementsErrorCode = "verification_failed_tax_id_match"
	AccountFutureRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       AccountFutureRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	AccountFutureRequirementsErrorCodeVerificationMissingExecutives                          AccountFutureRequirementsErrorCode = "verification_missing_executives"
	AccountFutureRequirementsErrorCodeVerificationMissingOwners                              AccountFutureRequirementsErrorCode = "verification_missing_owners"
	AccountFutureRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations AccountFutureRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// The code for the type of error.
type AccountRequirementsErrorCode string

// List of values that AccountRequirementsErrorCode can take
const (
	AccountRequirementsErrorCodeInvalidAddressCityStatePostalCode                      AccountRequirementsErrorCode = "invalid_address_city_state_postal_code"
	AccountRequirementsErrorCodeInvalidStreetAddress                                   AccountRequirementsErrorCode = "invalid_street_address"
	AccountRequirementsErrorCodeInvalidValueOther                                      AccountRequirementsErrorCode = "invalid_value_other"
	AccountRequirementsErrorCodeVerificationDocumentAddressMismatch                    AccountRequirementsErrorCode = "verification_document_address_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentAddressMissing                     AccountRequirementsErrorCode = "verification_document_address_missing"
	AccountRequirementsErrorCodeVerificationDocumentCorrupt                            AccountRequirementsErrorCode = "verification_document_corrupt"
	AccountRequirementsErrorCodeVerificationDocumentCountryNotSupported                AccountRequirementsErrorCode = "verification_document_country_not_supported"
	AccountRequirementsErrorCodeVerificationDocumentDOBMismatch                        AccountRequirementsErrorCode = "verification_document_dob_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentDuplicateType                      AccountRequirementsErrorCode = "verification_document_duplicate_type"
	AccountRequirementsErrorCodeVerificationDocumentExpired                            AccountRequirementsErrorCode = "verification_document_expired"
	AccountRequirementsErrorCodeVerificationDocumentFailedCopy                         AccountRequirementsErrorCode = "verification_document_failed_copy"
	AccountRequirementsErrorCodeVerificationDocumentFailedGreyscale                    AccountRequirementsErrorCode = "verification_document_failed_greyscale"
	AccountRequirementsErrorCodeVerificationDocumentFailedOther                        AccountRequirementsErrorCode = "verification_document_failed_other"
	AccountRequirementsErrorCodeVerificationDocumentFailedTestMode                     AccountRequirementsErrorCode = "verification_document_failed_test_mode"
	AccountRequirementsErrorCodeVerificationDocumentFraudulent                         AccountRequirementsErrorCode = "verification_document_fraudulent"
	AccountRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   AccountRequirementsErrorCode = "verification_document_id_number_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentIDNumberMissing                    AccountRequirementsErrorCode = "verification_document_id_number_missing"
	AccountRequirementsErrorCodeVerificationDocumentIncomplete                         AccountRequirementsErrorCode = "verification_document_incomplete"
	AccountRequirementsErrorCodeVerificationDocumentInvalid                            AccountRequirementsErrorCode = "verification_document_invalid"
	AccountRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           AccountRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	AccountRequirementsErrorCodeVerificationDocumentManipulated                        AccountRequirementsErrorCode = "verification_document_manipulated"
	AccountRequirementsErrorCodeVerificationDocumentMissingBack                        AccountRequirementsErrorCode = "verification_document_missing_back"
	AccountRequirementsErrorCodeVerificationDocumentMissingFront                       AccountRequirementsErrorCode = "verification_document_missing_front"
	AccountRequirementsErrorCodeVerificationDocumentNameMismatch                       AccountRequirementsErrorCode = "verification_document_name_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentNameMissing                        AccountRequirementsErrorCode = "verification_document_name_missing"
	AccountRequirementsErrorCodeVerificationDocumentNationalityMismatch                AccountRequirementsErrorCode = "verification_document_nationality_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentNotReadable                        AccountRequirementsErrorCode = "verification_document_not_readable"
	AccountRequirementsErrorCodeVerificationDocumentNotSigned                          AccountRequirementsErrorCode = "verification_document_not_signed"
	AccountRequirementsErrorCodeVerificationDocumentNotUploaded                        AccountRequirementsErrorCode = "verification_document_not_uploaded"
	AccountRequirementsErrorCodeVerificationDocumentPhotoMismatch                      AccountRequirementsErrorCode = "verification_document_photo_mismatch"
	AccountRequirementsErrorCodeVerificationDocumentTooLarge                           AccountRequirementsErrorCode = "verification_document_too_large"
	AccountRequirementsErrorCodeVerificationDocumentTypeNotSupported                   AccountRequirementsErrorCode = "verification_document_type_not_supported"
	AccountRequirementsErrorCodeVerificationFailedAddressMatch                         AccountRequirementsErrorCode = "verification_failed_address_match"
	AccountRequirementsErrorCodeVerificationFailedBusinessIecNumber                    AccountRequirementsErrorCode = "verification_failed_business_iec_number"
	AccountRequirementsErrorCodeVerificationFailedDocumentMatch                        AccountRequirementsErrorCode = "verification_failed_document_match"
	AccountRequirementsErrorCodeVerificationFailedIDNumberMatch                        AccountRequirementsErrorCode = "verification_failed_id_number_match"
	AccountRequirementsErrorCodeVerificationFailedKeyedIdentity                        AccountRequirementsErrorCode = "verification_failed_keyed_identity"
	AccountRequirementsErrorCodeVerificationFailedKeyedMatch                           AccountRequirementsErrorCode = "verification_failed_keyed_match"
	AccountRequirementsErrorCodeVerificationFailedNameMatch                            AccountRequirementsErrorCode = "verification_failed_name_match"
	AccountRequirementsErrorCodeVerificationFailedOther                                AccountRequirementsErrorCode = "verification_failed_other"
	AccountRequirementsErrorCodeVerificationFailedTaxIDMatch                           AccountRequirementsErrorCode = "verification_failed_tax_id_match"
	AccountRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       AccountRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	AccountRequirementsErrorCodeVerificationMissingExecutives                          AccountRequirementsErrorCode = "verification_missing_executives"
	AccountRequirementsErrorCodeVerificationMissingOwners                              AccountRequirementsErrorCode = "verification_missing_owners"
	AccountRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations AccountRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// AccountType is the type of an account.
type AccountType string

// List of values that AccountType can take.
const (
	AccountTypeCustom   AccountType = "custom"
	AccountTypeExpress  AccountType = "express"
	AccountTypeStandard AccountType = "standard"
)

// AccountBusinessType describes the business type associated with an account.
type AccountBusinessType string

// List of values that AccountBusinessType can take.
const (
	AccountBusinessTypeCompany          AccountBusinessType = "company"
	AccountBusinessTypeGovernmentEntity AccountBusinessType = "government_entity"
	AccountBusinessTypeIndividual       AccountBusinessType = "individual"
	AccountBusinessTypeNonProfit        AccountBusinessType = "non_profit"
)

// The status of the Canadian pre-authorized debits payments capability of the account, or whether the account can directly process Canadian pre-authorized debits charges.
type AccountCapabilitiesACSSDebitPayments string

// List of values that AccountCapabilitiesACSSDebitPayments can take
const (
	AccountCapabilitiesACSSDebitPaymentsActive   AccountCapabilitiesACSSDebitPayments = "active"
	AccountCapabilitiesACSSDebitPaymentsInactive AccountCapabilitiesACSSDebitPayments = "inactive"
	AccountCapabilitiesACSSDebitPaymentsPending  AccountCapabilitiesACSSDebitPayments = "pending"
)

// The status of the Afterpay Clearpay capability of the account, or whether the account can directly process Afterpay Clearpay charges.
type AccountCapabilitiesAfterpayClearpayPayments string

// List of values that AccountCapabilitiesAfterpayClearpayPayments can take
const (
	AccountCapabilitiesAfterpayClearpayPaymentsActive   AccountCapabilitiesAfterpayClearpayPayments = "active"
	AccountCapabilitiesAfterpayClearpayPaymentsInactive AccountCapabilitiesAfterpayClearpayPayments = "inactive"
	AccountCapabilitiesAfterpayClearpayPaymentsPending  AccountCapabilitiesAfterpayClearpayPayments = "pending"
)

// The status of the BECS Direct Debit (AU) payments capability of the account, or whether the account can directly process BECS Direct Debit (AU) charges.
type AccountCapabilitiesAUBECSDebitPayments string

// List of values that AccountCapabilitiesAUBECSDebitPayments can take
const (
	AccountCapabilitiesAUBECSDebitPaymentsActive   AccountCapabilitiesAUBECSDebitPayments = "active"
	AccountCapabilitiesAUBECSDebitPaymentsInactive AccountCapabilitiesAUBECSDebitPayments = "inactive"
	AccountCapabilitiesAUBECSDebitPaymentsPending  AccountCapabilitiesAUBECSDebitPayments = "pending"
)

// The status of the Bacs Direct Debits payments capability of the account, or whether the account can directly process Bacs Direct Debits charges.
type AccountCapabilitiesBACSDebitPayments string

// List of values that AccountCapabilitiesBACSDebitPayments can take
const (
	AccountCapabilitiesBACSDebitPaymentsActive   AccountCapabilitiesBACSDebitPayments = "active"
	AccountCapabilitiesBACSDebitPaymentsInactive AccountCapabilitiesBACSDebitPayments = "inactive"
	AccountCapabilitiesBACSDebitPaymentsPending  AccountCapabilitiesBACSDebitPayments = "pending"
)

// The status of the Bancontact payments capability of the account, or whether the account can directly process Bancontact charges.
type AccountCapabilitiesBancontactPayments string

// List of values that AccountCapabilitiesBancontactPayments can take
const (
	AccountCapabilitiesBancontactPaymentsActive   AccountCapabilitiesBancontactPayments = "active"
	AccountCapabilitiesBancontactPaymentsInactive AccountCapabilitiesBancontactPayments = "inactive"
	AccountCapabilitiesBancontactPaymentsPending  AccountCapabilitiesBancontactPayments = "pending"
)

// The status of the boleto payments capability of the account, or whether the account can directly process boleto charges.
type AccountCapabilitiesBoletoPayments string

// List of values that AccountCapabilitiesBoletoPayments can take
const (
	AccountCapabilitiesBoletoPaymentsActive   AccountCapabilitiesBoletoPayments = "active"
	AccountCapabilitiesBoletoPaymentsInactive AccountCapabilitiesBoletoPayments = "inactive"
	AccountCapabilitiesBoletoPaymentsPending  AccountCapabilitiesBoletoPayments = "pending"
)

// The status of the card issuing capability of the account, or whether you can use Issuing to distribute funds on cards
type AccountCapabilitiesCardIssuing string

// List of values that AccountCapabilitiesCardIssuing can take
const (
	AccountCapabilitiesCardIssuingActive   AccountCapabilitiesCardIssuing = "active"
	AccountCapabilitiesCardIssuingInactive AccountCapabilitiesCardIssuing = "inactive"
	AccountCapabilitiesCardIssuingPending  AccountCapabilitiesCardIssuing = "pending"
)

// The status of the card payments capability of the account, or whether the account can directly process credit and debit card charges.
type AccountCapabilitiesCardPayments string

// List of values that AccountCapabilitiesCardPayments can take
const (
	AccountCapabilitiesCardPaymentsActive   AccountCapabilitiesCardPayments = "active"
	AccountCapabilitiesCardPaymentsInactive AccountCapabilitiesCardPayments = "inactive"
	AccountCapabilitiesCardPaymentsPending  AccountCapabilitiesCardPayments = "pending"
)

// The status of the Cartes Bancaires payments capability of the account, or whether the account can directly process Cartes Bancaires card charges in EUR currency.
type AccountCapabilitiesCartesBancairesPayments string

// List of values that AccountCapabilitiesCartesBancairesPayments can take
const (
	AccountCapabilitiesCartesBancairesPaymentsActive   AccountCapabilitiesCartesBancairesPayments = "active"
	AccountCapabilitiesCartesBancairesPaymentsInactive AccountCapabilitiesCartesBancairesPayments = "inactive"
	AccountCapabilitiesCartesBancairesPaymentsPending  AccountCapabilitiesCartesBancairesPayments = "pending"
)

// The status of the EPS payments capability of the account, or whether the account can directly process EPS charges.
type AccountCapabilitiesEPSPayments string

// List of values that AccountCapabilitiesEPSPayments can take
const (
	AccountCapabilitiesEPSPaymentsActive   AccountCapabilitiesEPSPayments = "active"
	AccountCapabilitiesEPSPaymentsInactive AccountCapabilitiesEPSPayments = "inactive"
	AccountCapabilitiesEPSPaymentsPending  AccountCapabilitiesEPSPayments = "pending"
)

// The status of the FPX payments capability of the account, or whether the account can directly process FPX charges.
type AccountCapabilitiesFPXPayments string

// List of values that AccountCapabilitiesFPXPayments can take
const (
	AccountCapabilitiesFPXPaymentsActive   AccountCapabilitiesFPXPayments = "active"
	AccountCapabilitiesFPXPaymentsInactive AccountCapabilitiesFPXPayments = "inactive"
	AccountCapabilitiesFPXPaymentsPending  AccountCapabilitiesFPXPayments = "pending"
)

// The status of the giropay payments capability of the account, or whether the account can directly process giropay charges.
type AccountCapabilitiesGiropayPayments string

// List of values that AccountCapabilitiesGiropayPayments can take
const (
	AccountCapabilitiesGiropayPaymentsActive   AccountCapabilitiesGiropayPayments = "active"
	AccountCapabilitiesGiropayPaymentsInactive AccountCapabilitiesGiropayPayments = "inactive"
	AccountCapabilitiesGiropayPaymentsPending  AccountCapabilitiesGiropayPayments = "pending"
)

// The status of the GrabPay payments capability of the account, or whether the account can directly process GrabPay charges.
type AccountCapabilitiesGrabpayPayments string

// List of values that AccountCapabilitiesGrabpayPayments can take
const (
	AccountCapabilitiesGrabpayPaymentsActive   AccountCapabilitiesGrabpayPayments = "active"
	AccountCapabilitiesGrabpayPaymentsInactive AccountCapabilitiesGrabpayPayments = "inactive"
	AccountCapabilitiesGrabpayPaymentsPending  AccountCapabilitiesGrabpayPayments = "pending"
)

// The status of the iDEAL payments capability of the account, or whether the account can directly process iDEAL charges.
type AccountCapabilitiesIdealPayments string

// List of values that AccountCapabilitiesIdealPayments can take
const (
	AccountCapabilitiesIdealPaymentsActive   AccountCapabilitiesIdealPayments = "active"
	AccountCapabilitiesIdealPaymentsInactive AccountCapabilitiesIdealPayments = "inactive"
	AccountCapabilitiesIdealPaymentsPending  AccountCapabilitiesIdealPayments = "pending"
)

// The status of the JCB payments capability of the account, or whether the account (Japan only) can directly process JCB credit card charges in JPY currency.
type AccountCapabilitiesJcbPayments string

// List of values that AccountCapabilitiesJcbPayments can take
const (
	AccountCapabilitiesJcbPaymentsActive   AccountCapabilitiesJcbPayments = "active"
	AccountCapabilitiesJcbPaymentsInactive AccountCapabilitiesJcbPayments = "inactive"
	AccountCapabilitiesJcbPaymentsPending  AccountCapabilitiesJcbPayments = "pending"
)

// The status of the legacy payments capability of the account.
type AccountCapabilitiesLegacyPayments string

// List of values that AccountCapabilitiesLegacyPayments can take
const (
	AccountCapabilitiesLegacyPaymentsActive   AccountCapabilitiesLegacyPayments = "active"
	AccountCapabilitiesLegacyPaymentsInactive AccountCapabilitiesLegacyPayments = "inactive"
	AccountCapabilitiesLegacyPaymentsPending  AccountCapabilitiesLegacyPayments = "pending"
)

// The status of the OXXO payments capability of the account, or whether the account can directly process OXXO charges.
type AccountCapabilitiesOXXOPayments string

// List of values that AccountCapabilitiesOXXOPayments can take
const (
	AccountCapabilitiesOXXOPaymentsActive   AccountCapabilitiesOXXOPayments = "active"
	AccountCapabilitiesOXXOPaymentsInactive AccountCapabilitiesOXXOPayments = "inactive"
	AccountCapabilitiesOXXOPaymentsPending  AccountCapabilitiesOXXOPayments = "pending"
)

// The status of the P24 payments capability of the account, or whether the account can directly process P24 charges.
type AccountCapabilitiesP24Payments string

// List of values that AccountCapabilitiesP24Payments can take
const (
	AccountCapabilitiesP24PaymentsActive   AccountCapabilitiesP24Payments = "active"
	AccountCapabilitiesP24PaymentsInactive AccountCapabilitiesP24Payments = "inactive"
	AccountCapabilitiesP24PaymentsPending  AccountCapabilitiesP24Payments = "pending"
)

// The status of the SEPA Direct Debits payments capability of the account, or whether the account can directly process SEPA Direct Debits charges.
type AccountCapabilitiesSepaDebitPayments string

// List of values that AccountCapabilitiesSepaDebitPayments can take
const (
	AccountCapabilitiesSepaDebitPaymentsActive   AccountCapabilitiesSepaDebitPayments = "active"
	AccountCapabilitiesSepaDebitPaymentsInactive AccountCapabilitiesSepaDebitPayments = "inactive"
	AccountCapabilitiesSepaDebitPaymentsPending  AccountCapabilitiesSepaDebitPayments = "pending"
)

// The status of the Sofort payments capability of the account, or whether the account can directly process Sofort charges.
type AccountCapabilitiesSofortPayments string

// List of values that AccountCapabilitiesSofortPayments can take
const (
	AccountCapabilitiesSofortPaymentsActive   AccountCapabilitiesSofortPayments = "active"
	AccountCapabilitiesSofortPaymentsInactive AccountCapabilitiesSofortPayments = "inactive"
	AccountCapabilitiesSofortPaymentsPending  AccountCapabilitiesSofortPayments = "pending"
)

// The status of the tax reporting 1099-K (US) capability of the account.
type AccountCapabilitiesTaxReportingUs1099K string

// List of values that AccountCapabilitiesTaxReportingUs1099K can take
const (
	AccountCapabilitiesTaxReportingUs1099KActive   AccountCapabilitiesTaxReportingUs1099K = "active"
	AccountCapabilitiesTaxReportingUs1099KInactive AccountCapabilitiesTaxReportingUs1099K = "inactive"
	AccountCapabilitiesTaxReportingUs1099KPending  AccountCapabilitiesTaxReportingUs1099K = "pending"
)

// The status of the tax reporting 1099-MISC (US) capability of the account.
type AccountCapabilitiesTaxReportingUs1099Misc string

// List of values that AccountCapabilitiesTaxReportingUs1099Misc can take
const (
	AccountCapabilitiesTaxReportingUs1099MiscActive   AccountCapabilitiesTaxReportingUs1099Misc = "active"
	AccountCapabilitiesTaxReportingUs1099MiscInactive AccountCapabilitiesTaxReportingUs1099Misc = "inactive"
	AccountCapabilitiesTaxReportingUs1099MiscPending  AccountCapabilitiesTaxReportingUs1099Misc = "pending"
)

// The status of the transfers capability of the account, or whether your platform can transfer funds to the account.
type AccountCapabilitiesTransfers string

// List of values that AccountCapabilitiesTransfers can take
const (
	AccountCapabilitiesTransfersActive   AccountCapabilitiesTransfers = "active"
	AccountCapabilitiesTransfersInactive AccountCapabilitiesTransfers = "inactive"
	AccountCapabilitiesTransfersPending  AccountCapabilitiesTransfers = "pending"
)

// AccountCompanyStructure describes the structure associated with a company.
type AccountCompanyStructure string

// List of values that AccountCompanyStructure can take.
const (
	AccountCompanyStructureFreeZoneEstablishment              AccountCompanyStructure = "free_zone_establishment"
	AccountCompanyStructureFreeZoneLlc                        AccountCompanyStructure = "free_zone_llc"
	AccountCompanyStructureGovernmentInstrumentality          AccountCompanyStructure = "government_instrumentality"
	AccountCompanyStructureGovernmentalUnit                   AccountCompanyStructure = "governmental_unit"
	AccountCompanyStructureIncorporatedNonProfit              AccountCompanyStructure = "incorporated_non_profit"
	AccountCompanyStructureLimitedLiabilityPartnership        AccountCompanyStructure = "limited_liability_partnership"
	AccountCompanyStructureLlc                                AccountCompanyStructure = "llc"
	AccountCompanyStructureMultiMemberLlc                     AccountCompanyStructure = "multi_member_llc"
	AccountCompanyStructurePrivateCompany                     AccountCompanyStructure = "private_company"
	AccountCompanyStructurePrivateCorporation                 AccountCompanyStructure = "private_corporation"
	AccountCompanyStructurePrivatePartnership                 AccountCompanyStructure = "private_partnership"
	AccountCompanyStructurePublicCompany                      AccountCompanyStructure = "public_company"
	AccountCompanyStructurePublicCorporation                  AccountCompanyStructure = "public_corporation"
	AccountCompanyStructurePublicPartnership                  AccountCompanyStructure = "public_partnership"
	AccountCompanyStructureSingleMemberLlc                    AccountCompanyStructure = "single_member_llc"
	AccountCompanyStructureSoleEstablishment                  AccountCompanyStructure = "sole_establishment"
	AccountCompanyStructureSoleProprietorship                 AccountCompanyStructure = "sole_proprietorship"
	AccountCompanyStructureTaxExemptGovernmentInstrumentality AccountCompanyStructure = "tax_exempt_government_instrumentality"
	AccountCompanyStructureUnincorporatedAssociation          AccountCompanyStructure = "unincorporated_association"
	AccountCompanyStructureUnincorporatedNonProfit            AccountCompanyStructure = "unincorporated_non_profit"
)

// AccountControllerType describes the controller type of the account.
type AccountControllerType string

// List of values that AccountControllerType can take.
const (
	AccountControllerTypeAccount     AccountControllerType = "account"
	AccountControllerTypeApplication AccountControllerType = "application"
)

// AccountBusinessProfileParams are the parameters allowed for an account's business information
type AccountBusinessProfileParams struct {
	Mcc                *string        `form:"mcc"`
	Name               *string        `form:"name"`
	ProductDescription *string        `form:"product_description"`
	SupportAddress     *AddressParams `form:"support_address"`
	SupportEmail       *string        `form:"support_email"`
	SupportPhone       *string        `form:"support_phone"`
	SupportURL         *string        `form:"support_url"`
	URL                *string        `form:"url"`
}

// AccountCapabilitiesACSSDebitPaymentsParams represent allowed parameters to configure the ACSS Debit capability on an account.
type AccountCapabilitiesACSSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The afterpay_clearpay_payments capability.
type AccountCapabilitiesAfterpayClearpayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesAUBECSDebitPaymentsParams represent allowed parameters to configure the AU BECS Debit capability on an account.
type AccountCapabilitiesAUBECSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesBACSDebitPaymentsParams represent allowed parameters to configure the BACS Debit capability on an account.
type AccountCapabilitiesBACSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesBancontactPaymentsParams represent allowed parameters to configure the bancontact payments capability on an account.
type AccountCapabilitiesBancontactPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesBoletoPaymentsParams represent allowed parameters to configure the boleto payments capability on an account.
type AccountCapabilitiesBoletoPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesCardIssuingParams represent allowed parameters to configure the Issuing capability on an account.
type AccountCapabilitiesCardIssuingParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesCardPaymentsParams represent allowed parameters to configure the card payments capability on an account.
type AccountCapabilitiesCardPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesCartesBancairesPaymentsParams represent allowed parameters to configure the Cartes Bancaires payments capability on an account.
type AccountCapabilitiesCartesBancairesPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesEPSPaymentsParams represent allowed parameters to configure the EPS payments capability on an account.
type AccountCapabilitiesEPSPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesFPXPaymentsParams represent allowed parameters to configure the FPX payments capability on an account.
type AccountCapabilitiesFPXPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesGiropayPaymentsParams represent allowed parameters to configure the giropay payments capability on an account.
type AccountCapabilitiesGiropayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesGrabpayPaymentsParams represent allowed parameters to configure the grabpay payments capability on an account.
type AccountCapabilitiesGrabpayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesIdealPaymentsParams represent allowed parameters to configure the ideal payments capability on an account.
type AccountCapabilitiesIdealPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The jcb_payments capability.
type AccountCapabilitiesJcbPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesLegacyPaymentsParams represent allowed parameters to configure the legacy payments capability on an account.
type AccountCapabilitiesLegacyPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesOXXOPaymentsParams represent allowed parameters to configure the OXXO capability on an account.
type AccountCapabilitiesOXXOPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesP24PaymentsParams represent allowed parameters to configure the P24 payments capability on an account.
type AccountCapabilitiesP24PaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The sepa_debit_payments capability.
type AccountCapabilitiesSepaDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesSofortPaymentsParams represent allowed parameters to configure the sofort payments capability on an account.
type AccountCapabilitiesSofortPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_k capability.
type AccountCapabilitiesTaxReportingUs1099KParams struct {
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_misc capability.
type AccountCapabilitiesTaxReportingUs1099MiscParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesTransfersParams represent allowed parameters to configure the transfers capability on an account.
type AccountCapabilitiesTransfersParams struct {
	Requested *bool `form:"requested"`
}

// AccountCapabilitiesParams represent allowed parameters to configure capabilities on an account.
type AccountCapabilitiesParams struct {
	ListParams               `form:"*"`
	Account                  *string                                            `form:"-"` // Included in URL
	ACSSDebitPayments        *AccountCapabilitiesACSSDebitPaymentsParams        `form:"acss_debit_payments"`
	AfterpayClearpayPayments *AccountCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments"`
	AUBECSDebitPayments      *AccountCapabilitiesAUBECSDebitPaymentsParams      `form:"au_becs_debit_payments"`
	BACSDebitPayments        *AccountCapabilitiesBACSDebitPaymentsParams        `form:"bacs_debit_payments"`
	BancontactPayments       *AccountCapabilitiesBancontactPaymentsParams       `form:"bancontact_payments"`
	BoletoPayments           *AccountCapabilitiesBoletoPaymentsParams           `form:"boleto_payments"`
	CardIssuing              *AccountCapabilitiesCardIssuingParams              `form:"card_issuing"`
	CardPayments             *AccountCapabilitiesCardPaymentsParams             `form:"card_payments"`
	CartesBancairesPayments  *AccountCapabilitiesCartesBancairesPaymentsParams  `form:"cartes_bancaires_payments"`
	EPSPayments              *AccountCapabilitiesEPSPaymentsParams              `form:"eps_payments"`
	FPXPayments              *AccountCapabilitiesFPXPaymentsParams              `form:"fpx_payments"`
	GiropayPayments          *AccountCapabilitiesGiropayPaymentsParams          `form:"giropay_payments"`
	GrabpayPayments          *AccountCapabilitiesGrabpayPaymentsParams          `form:"grabpay_payments"`
	IdealPayments            *AccountCapabilitiesIdealPaymentsParams            `form:"ideal_payments"`
	JcbPayments              *AccountCapabilitiesJcbPaymentsParams              `form:"jcb_payments"`
	LegacyPayments           *AccountCapabilitiesLegacyPaymentsParams           `form:"legacy_payments"`
	OXXOPayments             *AccountCapabilitiesOXXOPaymentsParams             `form:"oxxo_payments"`
	P24Payments              *AccountCapabilitiesP24PaymentsParams              `form:"p24_payments"`
	SepaDebitPayments        *AccountCapabilitiesSepaDebitPaymentsParams        `form:"sepa_debit_payments"`
	SofortPayments           *AccountCapabilitiesSofortPaymentsParams           `form:"sofort_payments"`
	TaxReportingUs1099K      *AccountCapabilitiesTaxReportingUs1099KParams      `form:"tax_reporting_us_1099_k"`
	TaxReportingUs1099Misc   *AccountCapabilitiesTaxReportingUs1099MiscParams   `form:"tax_reporting_us_1099_misc"`
	Transfers                *AccountCapabilitiesTransfersParams                `form:"transfers"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountCompanyAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// The Kanji variation of the company's primary address (Japan only).
type AccountCompanyAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// AccountCompanyVerificationDocumentParams are the parameters allowed to pass for a document
// verifying a company.
type AccountCompanyVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// AccountCompanyVerificationParams are the parameters allowed to verify a company.
type AccountCompanyVerificationParams struct {
	Document *AccountCompanyVerificationDocumentParams `form:"document"`
}

// AccountCompanyParams are the parameters describing the company associated with the account.
type AccountCompanyParams struct {
	Address            *AddressParams                    `form:"address"`
	AddressKana        *AccountCompanyAddressKanaParams  `form:"address_kana"`
	AddressKanji       *AccountCompanyAddressKanjiParams `form:"address_kanji"`
	DirectorsProvided  *bool                             `form:"directors_provided"`
	ExecutivesProvided *bool                             `form:"executives_provided"`
	Name               *string                           `form:"name"`
	NameKana           *string                           `form:"name_kana"`
	NameKanji          *string                           `form:"name_kanji"`
	OwnersProvided     *bool                             `form:"owners_provided"`
	Phone              *string                           `form:"phone"`
	RegistrationNumber *string                           `form:"registration_number"`
	Structure          *string                           `form:"structure"`
	TaxID              *string                           `form:"tax_id"`
	TaxIDRegistrar     *string                           `form:"tax_id_registrar"`
	VATID              *string                           `form:"vat_id"`
	Verification       *AccountCompanyVerificationParams `form:"verification"`
}

// AccountDocumentsBankAccountOwnershipVerificationParams represents the
// parameters allowed for passing bank account ownership verification documents
// on an account.
type AccountDocumentsBankAccountOwnershipVerificationParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsCompanyLicenseParams represents the parameters allowed for
// passing company license verification documents on an account.
type AccountDocumentsCompanyLicenseParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsCompanyMemorandumOfAssociationParams represents the
// parameters allowed for passing company memorandum of association documents
// on an account.
type AccountDocumentsCompanyMemorandumOfAssociationParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsCompanyMinisterialDecreeParams represents the parameters
// allowed for passing company ministerial decree documents on an account.
type AccountDocumentsCompanyMinisterialDecreeParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsCompanyRegistrationVerificationParams represents the
// parameters allowed for passing company registration verification documents.
type AccountDocumentsCompanyRegistrationVerificationParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsCompanyTaxIDVerificationParams represents the parameters
// allowed for passing company tax id verification documents on an account.
type AccountDocumentsCompanyTaxIDVerificationParams struct {
	Files []*string `form:"files"`
}

// AccountDocumentsParams represents the parameters allowed for passing additional documents on an account.
type AccountDocumentsParams struct {
	BankAccountOwnershipVerification *AccountDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification"`
	CompanyLicense                   *AccountDocumentsCompanyLicenseParams                   `form:"company_license"`
	CompanyMemorandumOfAssociation   *AccountDocumentsCompanyMemorandumOfAssociationParams   `form:"company_memorandum_of_association"`
	CompanyMinisterialDecree         *AccountDocumentsCompanyMinisterialDecreeParams         `form:"company_ministerial_decree"`
	CompanyRegistrationVerification  *AccountDocumentsCompanyRegistrationVerificationParams  `form:"company_registration_verification"`
	CompanyTaxIDVerification         *AccountDocumentsCompanyTaxIDVerificationParams         `form:"company_tax_id_verification"`
}

// The Kana variation of the the individual's primary address (Japan only).
type AccountIndividualAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// The Kanji variation of the the individual's primary address (Japan only).
type AccountIndividualAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// The individual's date of birth.
type AccountIndividualDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type AccountIndividualVerificationAdditionalDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// An identifying document, either a passport or local ID card.
type AccountIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// The individual's verification document information.
type AccountIndividualVerificationParams struct {
	AdditionalDocument *AccountIndividualVerificationAdditionalDocumentParams `form:"additional_document"`
	Document           *AccountIndividualVerificationDocumentParams           `form:"document"`
}

// Information about the person represented by the account. This field is null unless `business_type` is set to `individual`.
type AccountIndividualParams struct {
	Address           *AddressParams                       `form:"address"`
	AddressKana       *AccountIndividualAddressKanaParams  `form:"address_kana"`
	AddressKanji      *AccountIndividualAddressKanjiParams `form:"address_kanji"`
	DOB               *AccountIndividualDOBParams          `form:"dob"`
	Email             *string                              `form:"email"`
	FirstName         *string                              `form:"first_name"`
	FirstNameKana     *string                              `form:"first_name_kana"`
	FirstNameKanji    *string                              `form:"first_name_kanji"`
	Gender            *string                              `form:"gender"`
	IDNumber          *string                              `form:"id_number"`
	LastName          *string                              `form:"last_name"`
	LastNameKana      *string                              `form:"last_name_kana"`
	LastNameKanji     *string                              `form:"last_name_kanji"`
	MaidenName        *string                              `form:"maiden_name"`
	Metadata          map[string]string                    `form:"metadata"`
	Phone             *string                              `form:"phone"`
	PoliticalExposure *string                              `form:"political_exposure"`
	SsnLast4          *string                              `form:"ssn_last_4"`
	Verification      *AccountIndividualVerificationParams `form:"verification"`
}

// AccountSettingsBrandingParams represent allowed parameters to configure settings specific to the
// account’s branding.
type AccountSettingsBrandingParams struct {
	Icon           *string `form:"icon"`
	Logo           *string `form:"logo"`
	PrimaryColor   *string `form:"primary_color"`
	SecondaryColor *string `form:"secondary_color"`
}

// Details on the account's acceptance of the [Stripe Issuing Terms and Disclosures](https://stripe.com/docs/issuing/connect/tos_acceptance).
type AccountSettingsCardIssuingTosAcceptanceParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}

// AccountSettingsCardIssuingParams represent allowed parameters relating to the acceptance of the terms of service agreement.
type AccountSettingsCardIssuingParams struct {
	TosAcceptance *AccountSettingsCardIssuingTosAcceptanceParams `form:"tos_acceptance"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type AccountSettingsCardPaymentsDeclineOnParams struct {
	AvsFailure *bool `form:"avs_failure"`
	CVCFailure *bool `form:"cvc_failure"`
}

// AccountSettingsCardPaymentsParams represent allowed parameters to configure settings specific to
// card charging on the account.
type AccountSettingsCardPaymentsParams struct {
	DeclineOn                 *AccountSettingsCardPaymentsDeclineOnParams `form:"decline_on"`
	StatementDescriptorPrefix *string                                     `form:"statement_descriptor_prefix"`
}

// AccountSettingsPaymentsParams represent allowed parameters to configure settings  across payment
// methods for charging on the account.
type AccountSettingsPaymentsParams struct {
	StatementDescriptor      *string `form:"statement_descriptor"`
	StatementDescriptorKana  *string `form:"statement_descriptor_kana"`
	StatementDescriptorKanji *string `form:"statement_descriptor_kanji"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation.
type AccountSettingsPayoutsScheduleParams struct {
	DelayDays        *int64  `form:"delay_days"`
	DelayDaysMinimum *bool   `form:"-"` // See custom AppendTo
	Interval         *string `form:"interval"`
	MonthlyAnchor    *int64  `form:"monthly_anchor"`
	WeeklyAnchor     *string `form:"weekly_anchor"`
}

// AppendTo implements custom encoding logic for AccountSettingsPayoutsScheduleParams.
func (a *AccountSettingsPayoutsScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(a.DelayDaysMinimum) {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// AccountSettingsPayoutsParams represent allowed parameters to configure settings specific to the
// account’s payouts.
type AccountSettingsPayoutsParams struct {
	DebitNegativeBalances *bool                                 `form:"debit_negative_balances"`
	Schedule              *AccountSettingsPayoutsScheduleParams `form:"schedule"`
	StatementDescriptor   *string                               `form:"statement_descriptor"`
}

// AccountSettingsParams are the parameters allowed for the account's settings.
type AccountSettingsParams struct {
	Branding     *AccountSettingsBrandingParams     `form:"branding"`
	CardIssuing  *AccountSettingsCardIssuingParams  `form:"card_issuing"`
	CardPayments *AccountSettingsCardPaymentsParams `form:"card_payments"`
	Payments     *AccountSettingsPaymentsParams     `form:"payments"`
	Payouts      *AccountSettingsPayoutsParams      `form:"payouts"`
}

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params          `form:"*"`
	AccountToken    *string                       `form:"account_token"`
	BusinessProfile *AccountBusinessProfileParams `form:"business_profile"`
	BusinessType    *string                       `form:"business_type"`
	Capabilities    *AccountCapabilitiesParams    `form:"capabilities"`
	Company         *AccountCompanyParams         `form:"company"`
	Country         *string                       `form:"country"`
	DefaultCurrency *string                       `form:"default_currency"`
	Documents       *AccountDocumentsParams       `form:"documents"`
	Email           *string                       `form:"email"`
	ExternalAccount *string                       `form:"external_account"`
	Individual      *AccountIndividualParams      `form:"individual"`
	Settings        *AccountSettingsParams        `form:"settings"`
	TosAcceptance   *AccountTosAcceptanceParams   `form:"tos_acceptance"`
	Type            *string                       `form:"type"`
}

// Details on the account's acceptance of the [Stripe Services Agreement](https://stripe.com/docs/connect/updating-accounts#tos-acceptance).
type AccountTosAcceptanceParams struct {
	Date             *int64  `form:"date"`
	IP               *string `form:"ip"`
	ServiceAgreement *string `form:"service_agreement"`
	UserAgent        *string `form:"user_agent"`
}

// AccountListParams are the parameters allowed during account listing.
type AccountListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// AccountRejectParams is the structure for the Reject function.
type AccountRejectParams struct {
	Params `form:"*"`
	Reason *string `form:"reason"`
}

// AccountBusinessProfile represents optional information related to the business.
type AccountBusinessProfile struct {
	Mcc                string   `json:"mcc"`
	Name               string   `json:"name"`
	ProductDescription string   `json:"product_description"`
	SupportAddress     *Address `json:"support_address"`
	SupportEmail       string   `json:"support_email"`
	SupportPhone       string   `json:"support_phone"`
	SupportURL         string   `json:"support_url"`
	URL                string   `json:"url"`
}

// AccountCapabilities is the resource representing the capabilities enabled on that account.
type AccountCapabilities struct {
	ACSSDebitPayments        AccountCapabilitiesACSSDebitPayments        `json:"acss_debit_payments"`
	AfterpayClearpayPayments AccountCapabilitiesAfterpayClearpayPayments `json:"afterpay_clearpay_payments"`
	AUBECSDebitPayments      AccountCapabilitiesAUBECSDebitPayments      `json:"au_becs_debit_payments"`
	BACSDebitPayments        AccountCapabilitiesBACSDebitPayments        `json:"bacs_debit_payments"`
	BancontactPayments       AccountCapabilitiesBancontactPayments       `json:"bancontact_payments"`
	BoletoPayments           AccountCapabilitiesBoletoPayments           `json:"boleto_payments"`
	CardIssuing              AccountCapabilitiesCardIssuing              `json:"card_issuing"`
	CardPayments             AccountCapabilitiesCardPayments             `json:"card_payments"`
	CartesBancairesPayments  AccountCapabilitiesCartesBancairesPayments  `json:"cartes_bancaires_payments"`
	EPSPayments              AccountCapabilitiesEPSPayments              `json:"eps_payments"`
	FPXPayments              AccountCapabilitiesFPXPayments              `json:"fpx_payments"`
	GiropayPayments          AccountCapabilitiesGiropayPayments          `json:"giropay_payments"`
	GrabpayPayments          AccountCapabilitiesGrabpayPayments          `json:"grabpay_payments"`
	IdealPayments            AccountCapabilitiesIdealPayments            `json:"ideal_payments"`
	JcbPayments              AccountCapabilitiesJcbPayments              `json:"jcb_payments"`
	LegacyPayments           AccountCapabilitiesLegacyPayments           `json:"legacy_payments"`
	OXXOPayments             AccountCapabilitiesOXXOPayments             `json:"oxxo_payments"`
	P24Payments              AccountCapabilitiesP24Payments              `json:"p24_payments"`
	SepaDebitPayments        AccountCapabilitiesSepaDebitPayments        `json:"sepa_debit_payments"`
	SofortPayments           AccountCapabilitiesSofortPayments           `json:"sofort_payments"`
	TaxReportingUs1099K      AccountCapabilitiesTaxReportingUs1099K      `json:"tax_reporting_us_1099_k"`
	TaxReportingUs1099Misc   AccountCapabilitiesTaxReportingUs1099Misc   `json:"tax_reporting_us_1099_misc"`
	Transfers                AccountCapabilitiesTransfers                `json:"transfers"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountCompanyAddressKana struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Town       string `json:"town"`
}

// The Kanji variation of the company's primary address (Japan only).
type AccountCompanyAddressKanji struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Town       string `json:"town"`
}

// AccountCompanyVerificationDocument represents details about a company's verification state.
type AccountCompanyVerificationDocument struct {
	Back        *File  `json:"back"`
	Details     string `json:"details"`
	DetailsCode string `json:"details_code"`
	Front       *File  `json:"front"`
}

// AccountCompanyVerification represents details about a company's verification state.
type AccountCompanyVerification struct {
	Document *AccountCompanyVerificationDocument `json:"document"`
}

// AccountCompany represents details about the company or business associated with the account.
type AccountCompany struct {
	Address            *Address                    `json:"address"`
	AddressKana        *AccountCompanyAddressKana  `json:"address_kana"`
	AddressKanji       *AccountCompanyAddressKanji `json:"address_kanji"`
	DirectorsProvided  bool                        `json:"directors_provided"`
	ExecutivesProvided bool                        `json:"executives_provided"`
	Name               string                      `json:"name"`
	NameKana           string                      `json:"name_kana"`
	NameKanji          string                      `json:"name_kanji"`
	OwnersProvided     bool                        `json:"owners_provided"`
	Phone              string                      `json:"phone"`
	Structure          AccountCompanyStructure     `json:"structure"`
	TaxIDProvided      bool                        `json:"tax_id_provided"`
	TaxIDRegistrar     string                      `json:"tax_id_registrar"`
	VATIDProvided      bool                        `json:"vat_id_provided"`
	Verification       *AccountCompanyVerification `json:"verification"`
}

// AccountController contains information about the control of the account.
type AccountController struct {
	IsController bool                  `json:"is_controller"`
	Type         AccountControllerType `json:"type"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type AccountFutureRequirementsError struct {
	Code        AccountFutureRequirementsErrorCode `json:"code"`
	Reason      string                             `json:"reason"`
	Requirement string                             `json:"requirement"`
}
type AccountFutureRequirements struct {
	Alternatives        []*AccountFutureRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                   `json:"current_deadline"`
	CurrentlyDue        []string                                `json:"currently_due"`
	DisabledReason      string                                  `json:"disabled_reason"`
	Errors              []*AccountFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                                `json:"eventually_due"`
	PastDue             []string                                `json:"past_due"`
	PendingVerification []string                                `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// AccountRequirementsError represents details about an error with a requirement.
type AccountRequirementsError struct {
	Code        AccountRequirementsErrorCode `json:"code"`
	Reason      string                       `json:"reason"`
	Requirement string                       `json:"requirement"`
}

// AccountRequirements represents information that needs to be collected for an account.
type AccountRequirements struct {
	Alternatives        []*AccountRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                             `json:"current_deadline"`
	CurrentlyDue        []string                          `json:"currently_due"`
	DisabledReason      string                            `json:"disabled_reason"`
	Errors              []*AccountRequirementsError       `json:"errors"`
	EventuallyDue       []string                          `json:"eventually_due"`
	PastDue             []string                          `json:"past_due"`
	PendingVerification []string                          `json:"pending_verification"`
}

// AccountSettingsBACSDebitPayments represents settings specific to the account’s charging
// via BACS Debit.
type AccountSettingsBACSDebitPayments struct {
	DisplayName string `json:"display_name"`
}

// AccountSettingsBranding represents settings specific to the account's branding.
type AccountSettingsBranding struct {
	Icon           *File  `json:"icon"`
	Logo           *File  `json:"logo"`
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
}
type AccountSettingsCardIssuingTosAcceptance struct {
	Date      int64  `json:"date"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

// AccountSettingsCardIssuing represents settings specific to card issuing on the account.
type AccountSettingsCardIssuing struct {
	TosAcceptance *AccountSettingsCardIssuingTosAcceptance `json:"tos_acceptance"`
}
type AccountSettingsCardPaymentsDeclineOn struct {
	AvsFailure bool `json:"avs_failure"`
	CVCFailure bool `json:"cvc_failure"`
}

// AccountSettingsCardPayments represents settings specific to card charging on the account.
type AccountSettingsCardPayments struct {
	DeclineOn                 *AccountSettingsCardPaymentsDeclineOn `json:"decline_on"`
	StatementDescriptorPrefix string                                `json:"statement_descriptor_prefix"`
}

// AccountSettingsDashboard represents settings specific to the account's Dashboard.
type AccountSettingsDashboard struct {
	DisplayName string `json:"display_name"`
	Timezone    string `json:"timezone"`
}

// AccountSettingsPayments represents settings that apply across payment methods for charging on
// the account.
type AccountSettingsPayments struct {
	StatementDescriptor      string `json:"statement_descriptor"`
	StatementDescriptorKana  string `json:"statement_descriptor_kana"`
	StatementDescriptorKanji string `json:"statement_descriptor_kanji"`
}
type AccountSettingsPayoutsSchedule struct {
	DelayDays     int64  `json:"delay_days"`
	Interval      string `json:"interval"`
	MonthlyAnchor int64  `json:"monthly_anchor"`
	WeeklyAnchor  string `json:"weekly_anchor"`
}

// AccountSettingsPayouts represents settings specific to the account’s payouts.
type AccountSettingsPayouts struct {
	DebitNegativeBalances bool                            `json:"debit_negative_balances"`
	Schedule              *AccountSettingsPayoutsSchedule `json:"schedule"`
	StatementDescriptor   string                          `json:"statement_descriptor"`
}
type AccountSettingsSepaDebitPayments struct {
	CreditorID string `json:"creditor_id"`
}

// AccountSettings represents options for customizing how the account functions within Stripe.
type AccountSettings struct {
	BACSDebitPayments *AccountSettingsBACSDebitPayments `json:"bacs_debit_payments"`
	Branding          *AccountSettingsBranding          `json:"branding"`
	CardIssuing       *AccountSettingsCardIssuing       `json:"card_issuing"`
	CardPayments      *AccountSettingsCardPayments      `json:"card_payments"`
	Dashboard         *AccountSettingsDashboard         `json:"dashboard"`
	Payments          *AccountSettingsPayments          `json:"payments"`
	Payouts           *AccountSettingsPayouts           `json:"payouts"`
	SepaDebitPayments *AccountSettingsSepaDebitPayments `json:"sepa_debit_payments"`
}
type AccountTosAcceptance struct {
	Date             int64  `json:"date"`
	IP               string `json:"ip"`
	ServiceAgreement string `json:"service_agreement"`
	UserAgent        string `json:"user_agent"`
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	APIResource
	BusinessProfile    *AccountBusinessProfile      `json:"business_profile"`
	BusinessType       AccountBusinessType          `json:"business_type"`
	Capabilities       *AccountCapabilities         `json:"capabilities"`
	ChargesEnabled     bool                         `json:"charges_enabled"`
	Company            *AccountCompany              `json:"company"`
	Controller         *AccountController           `json:"controller"`
	Country            string                       `json:"country"`
	Created            int64                        `json:"created"`
	DefaultCurrency    Currency                     `json:"default_currency"`
	Deleted            bool                         `json:"deleted"`
	DetailsSubmitted   bool                         `json:"details_submitted"`
	Email              string                       `json:"email"`
	ExternalAccounts   *AccountExternalAccountsList `json:"external_accounts"`
	FutureRequirements *AccountFutureRequirements   `json:"future_requirements"`
	ID                 string                       `json:"id"`
	Individual         *Person                      `json:"individual"`
	Metadata           map[string]string            `json:"metadata"`
	Object             string                       `json:"object"`
	PayoutsEnabled     bool                         `json:"payouts_enabled"`
	Requirements       *AccountRequirements         `json:"requirements"`
	Settings           *AccountSettings             `json:"settings"`
	TosAcceptance      *AccountTosAcceptance        `json:"tos_acceptance"`
	Type               AccountType                  `json:"type"`
}
type AccountExternalAccounts struct {
	ID   string                      `json:"id"`
	Type AccountExternalAccountsType `json:"object"`

	BankAccount *BankAccount `json:"-"`
	Card        *Card        `json:"-"`
}

// AccountList is a list of accounts as returned from a list endpoint.
type AccountList struct {
	APIResource
	ListMeta
	Data []*Account `json:"data"`
}

// UnmarshalJSON handles deserialization of an Account.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (a *Account) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		a.ID = id
		return nil
	}

	type account Account
	var v account
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*a = Account(v)
	return nil
}

// UnmarshalJSON handles deserialization of an AccountExternalAccounts.
// This custom unmarshaling is needed because the specific type of
// AccountExternalAccounts it refers to is specified in the JSON
func (a *AccountExternalAccounts) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		a.ID = id
		return nil
	}

	type accountExternalAccounts AccountExternalAccounts
	var v accountExternalAccounts
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*a = AccountExternalAccounts(v)
	var err error

	switch a.Type {
	case AccountExternalAccountsTypeBankAccount:
		err = json.Unmarshal(data, &a.BankAccount)
	case AccountExternalAccountsTypeCard:
		err = json.Unmarshal(data, &a.Card)
	}
	return err
}
