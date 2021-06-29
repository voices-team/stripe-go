//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type ChargePaymentMethodDetailsCardThreeDSecureVersion string

const (
	ChargePaymentMethodDetailsCardThreeDSecureVersion102 ChargePaymentMethodDetailsCardThreeDSecureVersion = "1.0.2"
	ChargePaymentMethodDetailsCardThreeDSecureVersion210 ChargePaymentMethodDetailsCardThreeDSecureVersion = "2.1.0"
	ChargePaymentMethodDetailsCardThreeDSecureVersion220 ChargePaymentMethodDetailsCardThreeDSecureVersion = "2.2.0"
)

type ChargePaymentMethodDetailsCardWalletType string

const (
	ChargePaymentMethodDetailsCardWalletTypeAmexExpressCheckout ChargePaymentMethodDetailsCardWalletType = "amex_express_checkout"
	ChargePaymentMethodDetailsCardWalletTypeApplePay            ChargePaymentMethodDetailsCardWalletType = "apple_pay"
	ChargePaymentMethodDetailsCardWalletTypeGooglePay           ChargePaymentMethodDetailsCardWalletType = "google_pay"
	ChargePaymentMethodDetailsCardWalletTypeMasterpass          ChargePaymentMethodDetailsCardWalletType = "masterpass"
	ChargePaymentMethodDetailsCardWalletTypeSamsungPay          ChargePaymentMethodDetailsCardWalletType = "samsung_pay"
	ChargePaymentMethodDetailsCardWalletTypeVisaCheckout        ChargePaymentMethodDetailsCardWalletType = "visa_checkout"
)

type ChargePaymentMethodDetailsCardPresentReadMethod string

const (
	ChargePaymentMethodDetailsCardPresentReadMethodContactEmv               ChargePaymentMethodDetailsCardPresentReadMethod = "contact_emv"
	ChargePaymentMethodDetailsCardPresentReadMethodContactlessEmv           ChargePaymentMethodDetailsCardPresentReadMethod = "contactless_emv"
	ChargePaymentMethodDetailsCardPresentReadMethodContactlessMagstripeMode ChargePaymentMethodDetailsCardPresentReadMethod = "contactless_magstripe_mode"
	ChargePaymentMethodDetailsCardPresentReadMethodMagneticStripeFallback   ChargePaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_fallback"
	ChargePaymentMethodDetailsCardPresentReadMethodMagneticStripeTrack2     ChargePaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_track2"
)

// ChargePaymentMethodDetailsCardPresentReceiptAccountType indicates the type of account backing a card present transaction.
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

type ChargePaymentMethodDetailsAchDebitAccountHolderType string

const (
	ChargePaymentMethodDetailsAchDebitAccountHolderTypeCompany    ChargePaymentMethodDetailsAchDebitAccountHolderType = "company"
	ChargePaymentMethodDetailsAchDebitAccountHolderTypeIndividual ChargePaymentMethodDetailsAchDebitAccountHolderType = "individual"
)

type ChargePaymentMethodDetailsBancontactPreferredLanguage string

const (
	ChargePaymentMethodDetailsBancontactPreferredLanguageDe ChargePaymentMethodDetailsBancontactPreferredLanguage = "de"
	ChargePaymentMethodDetailsBancontactPreferredLanguageEn ChargePaymentMethodDetailsBancontactPreferredLanguage = "en"
	ChargePaymentMethodDetailsBancontactPreferredLanguageFr ChargePaymentMethodDetailsBancontactPreferredLanguage = "fr"
	ChargePaymentMethodDetailsBancontactPreferredLanguageNl ChargePaymentMethodDetailsBancontactPreferredLanguage = "nl"
)

// ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow indicates the type of 3D Secure
// authentication performed.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// ChargePaymentMethodDetailsCardThreeDSecureResult indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// ChargePaymentMethodDetailsCardThreeDSecureResultReason represents dditional information about why
// 3D Secure succeeded or failed
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take.
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

type ChargePaymentMethodDetailsEPSBank string

const (
	ChargePaymentMethodDetailsEPSBankArzteUndApothekerBank                ChargePaymentMethodDetailsEPSBank = "arzte_und_apotheker_bank"
	ChargePaymentMethodDetailsEPSBankAustrianAnadiBankAg                  ChargePaymentMethodDetailsEPSBank = "austrian_anadi_bank_ag"
	ChargePaymentMethodDetailsEPSBankBankAustria                          ChargePaymentMethodDetailsEPSBank = "bank_austria"
	ChargePaymentMethodDetailsEPSBankBankhausCarlSpangler                 ChargePaymentMethodDetailsEPSBank = "bankhaus_carl_spangler"
	ChargePaymentMethodDetailsEPSBankBankhausSchelhammerUndSchatteraAg    ChargePaymentMethodDetailsEPSBank = "bankhaus_schelhammer_und_schattera_ag"
	ChargePaymentMethodDetailsEPSBankBawagPskAg                           ChargePaymentMethodDetailsEPSBank = "bawag_psk_ag"
	ChargePaymentMethodDetailsEPSBankBksBankAg                            ChargePaymentMethodDetailsEPSBank = "bks_bank_ag"
	ChargePaymentMethodDetailsEPSBankBrullKallmusBankAg                   ChargePaymentMethodDetailsEPSBank = "brull_kallmus_bank_ag"
	ChargePaymentMethodDetailsEPSBankBtvVierLanderBank                    ChargePaymentMethodDetailsEPSBank = "btv_vier_lander_bank"
	ChargePaymentMethodDetailsEPSBankCapitalBankGraweGruppeAg             ChargePaymentMethodDetailsEPSBank = "capital_bank_grawe_gruppe_ag"
	ChargePaymentMethodDetailsEPSBankDolomitenbank                        ChargePaymentMethodDetailsEPSBank = "dolomitenbank"
	ChargePaymentMethodDetailsEPSBankEasybankAg                           ChargePaymentMethodDetailsEPSBank = "easybank_ag"
	ChargePaymentMethodDetailsEPSBankErsteBankUndSparkassen               ChargePaymentMethodDetailsEPSBank = "erste_bank_und_sparkassen"
	ChargePaymentMethodDetailsEPSBankHypoAlpeadriabankInternationalAg     ChargePaymentMethodDetailsEPSBank = "hypo_alpeadriabank_international_ag"
	ChargePaymentMethodDetailsEPSBankHypoBankBurgenlandAktiengesellschaft ChargePaymentMethodDetailsEPSBank = "hypo_bank_burgenland_aktiengesellschaft"
	ChargePaymentMethodDetailsEPSBankHypoNoeLbFurNiederosterreichUWien    ChargePaymentMethodDetailsEPSBank = "hypo_noe_lb_fur_niederosterreich_u_wien"
	ChargePaymentMethodDetailsEPSBankHypoOberosterreichSalzburgSteiermark ChargePaymentMethodDetailsEPSBank = "hypo_oberosterreich_salzburg_steiermark"
	ChargePaymentMethodDetailsEPSBankHypoTirolBankAg                      ChargePaymentMethodDetailsEPSBank = "hypo_tirol_bank_ag"
	ChargePaymentMethodDetailsEPSBankHypoVorarlbergBankAg                 ChargePaymentMethodDetailsEPSBank = "hypo_vorarlberg_bank_ag"
	ChargePaymentMethodDetailsEPSBankMarchfelderBank                      ChargePaymentMethodDetailsEPSBank = "marchfelder_bank"
	ChargePaymentMethodDetailsEPSBankOberbankAg                           ChargePaymentMethodDetailsEPSBank = "oberbank_ag"
	ChargePaymentMethodDetailsEPSBankRaiffeisenBankengruppeOsterreich     ChargePaymentMethodDetailsEPSBank = "raiffeisen_bankengruppe_osterreich"
	ChargePaymentMethodDetailsEPSBankSchoellerbankAg                      ChargePaymentMethodDetailsEPSBank = "schoellerbank_ag"
	ChargePaymentMethodDetailsEPSBankSpardaBankWien                       ChargePaymentMethodDetailsEPSBank = "sparda_bank_wien"
	ChargePaymentMethodDetailsEPSBankVolksbankGruppe                      ChargePaymentMethodDetailsEPSBank = "volksbank_gruppe"
	ChargePaymentMethodDetailsEPSBankVolkskreditbankAg                    ChargePaymentMethodDetailsEPSBank = "volkskreditbank_ag"
	ChargePaymentMethodDetailsEPSBankVrBankBraunau                        ChargePaymentMethodDetailsEPSBank = "vr_bank_braunau"
)

type ChargePaymentMethodDetailsFPXAccountHolderType string

const (
	ChargePaymentMethodDetailsFPXAccountHolderTypeCompany    ChargePaymentMethodDetailsFPXAccountHolderType = "company"
	ChargePaymentMethodDetailsFPXAccountHolderTypeIndividual ChargePaymentMethodDetailsFPXAccountHolderType = "individual"
)

type ChargePaymentMethodDetailsFPXBank string

const (
	ChargePaymentMethodDetailsFPXBankAffinBank         ChargePaymentMethodDetailsFPXBank = "affin_bank"
	ChargePaymentMethodDetailsFPXBankAllianceBank      ChargePaymentMethodDetailsFPXBank = "alliance_bank"
	ChargePaymentMethodDetailsFPXBankAmbank            ChargePaymentMethodDetailsFPXBank = "ambank"
	ChargePaymentMethodDetailsFPXBankBankIslam         ChargePaymentMethodDetailsFPXBank = "bank_islam"
	ChargePaymentMethodDetailsFPXBankBankMuamalat      ChargePaymentMethodDetailsFPXBank = "bank_muamalat"
	ChargePaymentMethodDetailsFPXBankBankRakyat        ChargePaymentMethodDetailsFPXBank = "bank_rakyat"
	ChargePaymentMethodDetailsFPXBankBsn               ChargePaymentMethodDetailsFPXBank = "bsn"
	ChargePaymentMethodDetailsFPXBankCimb              ChargePaymentMethodDetailsFPXBank = "cimb"
	ChargePaymentMethodDetailsFPXBankDeutscheBank      ChargePaymentMethodDetailsFPXBank = "deutsche_bank"
	ChargePaymentMethodDetailsFPXBankHongLeongBank     ChargePaymentMethodDetailsFPXBank = "hong_leong_bank"
	ChargePaymentMethodDetailsFPXBankHsbc              ChargePaymentMethodDetailsFPXBank = "hsbc"
	ChargePaymentMethodDetailsFPXBankKfh               ChargePaymentMethodDetailsFPXBank = "kfh"
	ChargePaymentMethodDetailsFPXBankMaybank2e         ChargePaymentMethodDetailsFPXBank = "maybank2e"
	ChargePaymentMethodDetailsFPXBankMaybank2u         ChargePaymentMethodDetailsFPXBank = "maybank2u"
	ChargePaymentMethodDetailsFPXBankOcbc              ChargePaymentMethodDetailsFPXBank = "ocbc"
	ChargePaymentMethodDetailsFPXBankPbEnterprise      ChargePaymentMethodDetailsFPXBank = "pb_enterprise"
	ChargePaymentMethodDetailsFPXBankPublicBank        ChargePaymentMethodDetailsFPXBank = "public_bank"
	ChargePaymentMethodDetailsFPXBankRhb               ChargePaymentMethodDetailsFPXBank = "rhb"
	ChargePaymentMethodDetailsFPXBankStandardChartered ChargePaymentMethodDetailsFPXBank = "standard_chartered"
	ChargePaymentMethodDetailsFPXBankUob               ChargePaymentMethodDetailsFPXBank = "uob"
)

type ChargePaymentMethodDetailsIdealBank string

const (
	ChargePaymentMethodDetailsIdealBankAbnAmro       ChargePaymentMethodDetailsIdealBank = "abn_amro"
	ChargePaymentMethodDetailsIdealBankAsnBank       ChargePaymentMethodDetailsIdealBank = "asn_bank"
	ChargePaymentMethodDetailsIdealBankBunq          ChargePaymentMethodDetailsIdealBank = "bunq"
	ChargePaymentMethodDetailsIdealBankHandelsbanken ChargePaymentMethodDetailsIdealBank = "handelsbanken"
	ChargePaymentMethodDetailsIdealBankIng           ChargePaymentMethodDetailsIdealBank = "ing"
	ChargePaymentMethodDetailsIdealBankKnab          ChargePaymentMethodDetailsIdealBank = "knab"
	ChargePaymentMethodDetailsIdealBankMoneyou       ChargePaymentMethodDetailsIdealBank = "moneyou"
	ChargePaymentMethodDetailsIdealBankRabobank      ChargePaymentMethodDetailsIdealBank = "rabobank"
	ChargePaymentMethodDetailsIdealBankRegiobank     ChargePaymentMethodDetailsIdealBank = "regiobank"
	ChargePaymentMethodDetailsIdealBankRevolut       ChargePaymentMethodDetailsIdealBank = "revolut"
	ChargePaymentMethodDetailsIdealBankSnsBank       ChargePaymentMethodDetailsIdealBank = "sns_bank"
	ChargePaymentMethodDetailsIdealBankTriodosBank   ChargePaymentMethodDetailsIdealBank = "triodos_bank"
	ChargePaymentMethodDetailsIdealBankVanLanschot   ChargePaymentMethodDetailsIdealBank = "van_lanschot"
)

type ChargePaymentMethodDetailsIdealBic string

const (
	ChargePaymentMethodDetailsIdealBicABNANL2A ChargePaymentMethodDetailsIdealBic = "ABNANL2A"
	ChargePaymentMethodDetailsIdealBicASNBNL21 ChargePaymentMethodDetailsIdealBic = "ASNBNL21"
	ChargePaymentMethodDetailsIdealBicBUNQNL2A ChargePaymentMethodDetailsIdealBic = "BUNQNL2A"
	ChargePaymentMethodDetailsIdealBicFVLBNL22 ChargePaymentMethodDetailsIdealBic = "FVLBNL22"
	ChargePaymentMethodDetailsIdealBicHANDNL2A ChargePaymentMethodDetailsIdealBic = "HANDNL2A"
	ChargePaymentMethodDetailsIdealBicINGBNL2A ChargePaymentMethodDetailsIdealBic = "INGBNL2A"
	ChargePaymentMethodDetailsIdealBicKNABNL2H ChargePaymentMethodDetailsIdealBic = "KNABNL2H"
	ChargePaymentMethodDetailsIdealBicMOYONL21 ChargePaymentMethodDetailsIdealBic = "MOYONL21"
	ChargePaymentMethodDetailsIdealBicRABONL2U ChargePaymentMethodDetailsIdealBic = "RABONL2U"
	ChargePaymentMethodDetailsIdealBicRBRBNL21 ChargePaymentMethodDetailsIdealBic = "RBRBNL21"
	ChargePaymentMethodDetailsIdealBicREVOLT21 ChargePaymentMethodDetailsIdealBic = "REVOLT21"
	ChargePaymentMethodDetailsIdealBicSNSBNL2A ChargePaymentMethodDetailsIdealBic = "SNSBNL2A"
	ChargePaymentMethodDetailsIdealBicTRIONL2U ChargePaymentMethodDetailsIdealBic = "TRIONL2U"
)

type ChargePaymentMethodDetailsInteracPresentReadMethod string

const (
	ChargePaymentMethodDetailsInteracPresentReadMethodContactEmv               ChargePaymentMethodDetailsInteracPresentReadMethod = "contact_emv"
	ChargePaymentMethodDetailsInteracPresentReadMethodContactlessEmv           ChargePaymentMethodDetailsInteracPresentReadMethod = "contactless_emv"
	ChargePaymentMethodDetailsInteracPresentReadMethodContactlessMagstripeMode ChargePaymentMethodDetailsInteracPresentReadMethod = "contactless_magstripe_mode"
	ChargePaymentMethodDetailsInteracPresentReadMethodMagneticStripeFallback   ChargePaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_fallback"
	ChargePaymentMethodDetailsInteracPresentReadMethodMagneticStripeTrack2     ChargePaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_track2"
)

type ChargePaymentMethodDetailsInteracPresentReceiptAccountType string

const (
	ChargePaymentMethodDetailsInteracPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsInteracPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsInteracPresentReceiptAccountTypeSavings  ChargePaymentMethodDetailsInteracPresentReceiptAccountType = "savings"
	ChargePaymentMethodDetailsInteracPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsInteracPresentReceiptAccountType = "unknown"
)

type ChargePaymentMethodDetailsP24Bank string

const (
	ChargePaymentMethodDetailsP24BankAliorBank            ChargePaymentMethodDetailsP24Bank = "alior_bank"
	ChargePaymentMethodDetailsP24BankBankMillennium       ChargePaymentMethodDetailsP24Bank = "bank_millennium"
	ChargePaymentMethodDetailsP24BankBankNowyBfgSa        ChargePaymentMethodDetailsP24Bank = "bank_nowy_bfg_sa"
	ChargePaymentMethodDetailsP24BankBankPekaoSa          ChargePaymentMethodDetailsP24Bank = "bank_pekao_sa"
	ChargePaymentMethodDetailsP24BankBankiSpbdzielcze     ChargePaymentMethodDetailsP24Bank = "banki_spbdzielcze"
	ChargePaymentMethodDetailsP24BankBlik                 ChargePaymentMethodDetailsP24Bank = "blik"
	ChargePaymentMethodDetailsP24BankBnpParibas           ChargePaymentMethodDetailsP24Bank = "bnp_paribas"
	ChargePaymentMethodDetailsP24BankBoz                  ChargePaymentMethodDetailsP24Bank = "boz"
	ChargePaymentMethodDetailsP24BankCitiHandlowy         ChargePaymentMethodDetailsP24Bank = "citi_handlowy"
	ChargePaymentMethodDetailsP24BankCreditAgricole       ChargePaymentMethodDetailsP24Bank = "credit_agricole"
	ChargePaymentMethodDetailsP24BankEnvelobank           ChargePaymentMethodDetailsP24Bank = "envelobank"
	ChargePaymentMethodDetailsP24BankEtransferPocztowy24  ChargePaymentMethodDetailsP24Bank = "etransfer_pocztowy24"
	ChargePaymentMethodDetailsP24BankGetinBank            ChargePaymentMethodDetailsP24Bank = "getin_bank"
	ChargePaymentMethodDetailsP24BankIdeabank             ChargePaymentMethodDetailsP24Bank = "ideabank"
	ChargePaymentMethodDetailsP24BankIng                  ChargePaymentMethodDetailsP24Bank = "ing"
	ChargePaymentMethodDetailsP24BankInteligo             ChargePaymentMethodDetailsP24Bank = "inteligo"
	ChargePaymentMethodDetailsP24BankMbankMtransfer       ChargePaymentMethodDetailsP24Bank = "mbank_mtransfer"
	ChargePaymentMethodDetailsP24BankNestPrzelew          ChargePaymentMethodDetailsP24Bank = "nest_przelew"
	ChargePaymentMethodDetailsP24BankNoblePay             ChargePaymentMethodDetailsP24Bank = "noble_pay"
	ChargePaymentMethodDetailsP24BankPbacZIpko            ChargePaymentMethodDetailsP24Bank = "pbac_z_ipko"
	ChargePaymentMethodDetailsP24BankPlusBank             ChargePaymentMethodDetailsP24Bank = "plus_bank"
	ChargePaymentMethodDetailsP24BankSantanderPrzelew24   ChargePaymentMethodDetailsP24Bank = "santander_przelew24"
	ChargePaymentMethodDetailsP24BankTmobileUsbugiBankowe ChargePaymentMethodDetailsP24Bank = "tmobile_usbugi_bankowe"
	ChargePaymentMethodDetailsP24BankToyotaBank           ChargePaymentMethodDetailsP24Bank = "toyota_bank"
	ChargePaymentMethodDetailsP24BankVolkswagenBank       ChargePaymentMethodDetailsP24Bank = "volkswagen_bank"
)

type ChargePaymentMethodDetailsSofortPreferredLanguage string

const (
	ChargePaymentMethodDetailsSofortPreferredLanguageDe ChargePaymentMethodDetailsSofortPreferredLanguage = "de"
	ChargePaymentMethodDetailsSofortPreferredLanguageEn ChargePaymentMethodDetailsSofortPreferredLanguage = "en"
	ChargePaymentMethodDetailsSofortPreferredLanguageEs ChargePaymentMethodDetailsSofortPreferredLanguage = "es"
	ChargePaymentMethodDetailsSofortPreferredLanguageFr ChargePaymentMethodDetailsSofortPreferredLanguage = "fr"
	ChargePaymentMethodDetailsSofortPreferredLanguageIt ChargePaymentMethodDetailsSofortPreferredLanguage = "it"
	ChargePaymentMethodDetailsSofortPreferredLanguageNl ChargePaymentMethodDetailsSofortPreferredLanguage = "nl"
	ChargePaymentMethodDetailsSofortPreferredLanguagePl ChargePaymentMethodDetailsSofortPreferredLanguage = "pl"
)

type ChargeDestinationParams struct {
	Account *string `form:"account"`
	Amount  *int64  `form:"amount"`
}

// ChargeTransferDataParams is the set of parameters allowed for the transfer_data hash.
type ChargeTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
type ChargeParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFee            *int64                    `form:"application_fee"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	Capture                   *bool                     `form:"capture"`
	Currency                  *string                   `form:"currency"`
	Customer                  *string                   `form:"customer"`
	Description               *string                   `form:"description"`
	Destination               *ChargeDestinationParams  `form:"destination"`
	FraudDetails              *ChargeFraudDetailsParams `form:"fraud_details"`
	OnBehalfOf                *string                   `form:"on_behalf_of"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	Shipping                  *ShippingDetailsParams    `form:"shipping"`
	Source                    *string                   `form:"source"`
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// ChargeListParams is the set of parameters that can be used when listing charges.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentIntent *string           `form:"payment_intent"`
	TransferGroup *string           `form:"transfer_group"`
}
type ChargeFraudDetailsParams struct {
	UserReport *string `form:"user_report"`
}
type ChargeCaptureTransferDataParams struct {
	Amount *int64 `form:"amount"`
}
type ChargeCaptureParams struct {
	Params                    `form:"*"`
	Amount                    *int64                           `form:"amount"`
	ApplicationFee            *int64                           `form:"application_fee"`
	ApplicationFeeAmount      *int64                           `form:"application_fee_amount"`
	ReceiptEmail              *string                          `form:"receipt_email"`
	StatementDescriptor       *string                          `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                          `form:"statement_descriptor_suffix"`
	TransferData              *ChargeCaptureTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                          `form:"transfer_group"`
}
type ChargeAlternateStatementDescriptors struct {
	Kana  string `json:"kana"`
	Kanji string `json:"kanji"`
}
type ChargeBillingDetails struct {
	Address *Address `json:"address"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
}
type ChargeFraudDetails struct {
	StripeReport string `json:"stripe_report"`
	UserReport   string `json:"user_report"`
}

// ChargeLevel3LineItem represents a line item on level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3LineItem struct {
	DiscountAmount     int64  `json:"discount_amount"`
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	Quantity           int64  `json:"quantity"`
	TaxAmount          int64  `json:"tax_amount"`
	UnitCost           int64  `json:"unit_cost"`
}

// ChargeLevel3 represents the Level III data.
// This is in private beta and would be empty for most integrations
type ChargeLevel3 struct {
	CustomerReference  string                  `json:"customer_reference"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items"`
	MerchantReference  string                  `json:"merchant_reference"`
	ShippingAddressZip string                  `json:"shipping_address_zip"`
	ShippingAmount     int64                   `json:"shipping_amount"`
	ShippingFromZip    string                  `json:"shipping_from_zip"`
}

// ChargePaymentMethodDetailsAchCreditTransfer represents details about the ACH Credit Transfer
// PaymentMethod.
type ChargePaymentMethodDetailsAchCreditTransfer struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	RoutingNumber string `json:"routing_number"`
	SwiftCode     string `json:"swift_code"`
}

// ChargePaymentMethodDetailsAchDebit represents details about the ACH Debit PaymentMethod.
type ChargePaymentMethodDetailsAchDebit struct {
	AccountHolderType ChargePaymentMethodDetailsAchDebitAccountHolderType `json:"account_holder_type"`
	BankName          string                                              `json:"bank_name"`
	Country           string                                              `json:"country"`
	Fingerprint       string                                              `json:"fingerprint"`
	Last4             string                                              `json:"last4"`
	RoutingNumber     string                                              `json:"routing_number"`
}
type ChargePaymentMethodDetailsACSSDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	Mandate           string `json:"mandate"`
	TransitNumber     string `json:"transit_number"`
}

// ChargePaymentMethodDetailsAfterpayClearpay represents details about the AfterpayClearpay PaymentMethod.
type ChargePaymentMethodDetailsAfterpayClearpay struct {
	Reference string `json:"reference"`
}

// ChargePaymentMethodDetailsAlipay represents details about the Alipay PaymentMethod.
type ChargePaymentMethodDetailsAlipay struct {
	Fingerprint   string `json:"fingerprint"`
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAuBECSDebit struct {
	BsbNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

// ChargePaymentMethodDetailsBACSDebit represents details about the BECS Debit PaymentMethod.
type ChargePaymentMethodDetailsBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
	SortCode    string `json:"sort_code"`
}

// ChargePaymentMethodDetailsBancontact represents details about the Bancontact PaymentMethod.
type ChargePaymentMethodDetailsBancontact struct {
	BankCode                  string                                                `json:"bank_code"`
	BankName                  string                                                `json:"bank_name"`
	Bic                       string                                                `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod                                        `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                                              `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                                                `json:"iban_last4"`
	PreferredLanguage         ChargePaymentMethodDetailsBancontactPreferredLanguage `json:"preferred_language"`
	VerifiedName              string                                                `json:"verified_name"`
}

// ChargePaymentMethodDetailsCardChecks represents the checks associated with the charge's Card
// PaymentMethod.
type ChargePaymentMethodDetailsCardChecks struct {
	AddressLine1Check      string `json:"address_line1_check"`
	AddressPostalCodeCheck string `json:"address_postal_code_check"`
	CVCCheck               string `json:"cvc_check"`
}
type ChargePaymentMethodDetailsCardInstallmentsPlan struct {
	Count    int64   `json:"count"`
	Interval *string `json:"interval"`
	Type     string  `json:"type"`
}

// ChargePaymentMethodDetailsCardInstallments represents details about the installment plan chosen
// for this charge.
type ChargePaymentMethodDetailsCardInstallments struct {
	Plan *ChargePaymentMethodDetailsCardInstallmentsPlan `json:"plan"`
}

// ChargePaymentMethodDetailsCardThreeDSecure represents details about 3DS associated with the
// charge's PaymentMethod.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             ChargePaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       ChargePaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            ChargePaymentMethodDetailsCardThreeDSecureVersion            `json:"version"`
}

// ChargePaymentMethodDetailsCardWalletAmexExpressCheckout represents the details of the Amex
// Express Checkout wallet.
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct{}

// ChargePaymentMethodDetailsCardWalletApplePay represents the details of the Apple Pay wallet.
type ChargePaymentMethodDetailsCardWalletApplePay struct{}

// ChargePaymentMethodDetailsCardWalletGooglePay represents the details of the Google Pay wallet.
type ChargePaymentMethodDetailsCardWalletGooglePay struct{}

// ChargePaymentMethodDetailsCardWalletMasterpass represents the details of the Masterpass wallet.
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// ChargePaymentMethodDetailsCardWalletSamsungPay represents the details of the Samsung Pay wallet.
type ChargePaymentMethodDetailsCardWalletSamsungPay struct{}

// ChargePaymentMethodDetailsCardWalletVisaCheckout represents the details of the Visa Checkout
// wallet.
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// ChargePaymentMethodDetailsCardWallet represents the details of the card wallet if this Card
// PaymentMethod is part of a card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay"`
	DynamicLast4        string                                                   `json:"dynamic_last4"`
	GooglePay           *ChargePaymentMethodDetailsCardWalletGooglePay           `json:"google_pay"`
	Masterpass          *ChargePaymentMethodDetailsCardWalletMasterpass          `json:"masterpass"`
	SamsungPay          *ChargePaymentMethodDetailsCardWalletSamsungPay          `json:"samsung_pay"`
	Type                ChargePaymentMethodDetailsCardWalletType                 `json:"type"`
	VisaCheckout        *ChargePaymentMethodDetailsCardWalletVisaCheckout        `json:"visa_checkout"`
}

// ChargePaymentMethodDetailsCard represents details about the Card PaymentMethod.
type ChargePaymentMethodDetailsCard struct {
	Brand        string                                      `json:"brand"`
	Checks       *ChargePaymentMethodDetailsCardChecks       `json:"checks"`
	Country      string                                      `json:"country"`
	Description  string                                      `json:"description"`
	ExpMonth     int64                                       `json:"exp_month"`
	ExpYear      int64                                       `json:"exp_year"`
	Fingerprint  string                                      `json:"fingerprint"`
	Funding      string                                      `json:"funding"`
	Iin          string                                      `json:"iin"`
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments"`
	Issuer       string                                      `json:"issuer"`
	Last4        string                                      `json:"last4"`
	Moto         bool                                        `json:"moto"`
	Network      string                                      `json:"network"`
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	Wallet       *ChargePaymentMethodDetailsCardWallet       `json:"wallet"`
}

// ChargePaymentMethodDetailsCardPresentReceipt represents details about the receipt on a
// Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	AccountType                  ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
	ApplicationCryptogram        string                                                  `json:"application_cryptogram"`
	ApplicationPreferredName     string                                                  `json:"application_preferred_name"`
	AuthorizationCode            string                                                  `json:"authorization_code"`
	AuthorizationResponseCode    string                                                  `json:"authorization_response_code"`
	CardholderVerificationMethod string                                                  `json:"cardholder_verification_method"`
	DedicatedFileName            string                                                  `json:"dedicated_file_name"`
	TerminalVerificationResults  string                                                  `json:"terminal_verification_results"`
	TransactionStatusInformation string                                                  `json:"transaction_status_information"`
}

// ChargePaymentMethodDetailsCardPresent represents details about the Card Present PaymentMethod.
type ChargePaymentMethodDetailsCardPresent struct {
	Brand          string                                          `json:"brand"`
	CardholderName string                                          `json:"cardholder_name"`
	Country        string                                          `json:"country"`
	Description    string                                          `json:"description"`
	EmvAuthData    string                                          `json:"emv_auth_data"`
	ExpMonth       int64                                           `json:"exp_month"`
	ExpYear        int64                                           `json:"exp_year"`
	Fingerprint    string                                          `json:"fingerprint"`
	Funding        string                                          `json:"funding"`
	GeneratedCard  string                                          `json:"generated_card"`
	Iin            string                                          `json:"iin"`
	Issuer         string                                          `json:"issuer"`
	Last4          string                                          `json:"last4"`
	Network        string                                          `json:"network"`
	ReadMethod     ChargePaymentMethodDetailsCardPresentReadMethod `json:"read_method"`
	Receipt        *ChargePaymentMethodDetailsCardPresentReceipt   `json:"receipt"`
}
type ChargePaymentMethodDetailsEPS struct {
	Bank         ChargePaymentMethodDetailsEPSBank `json:"bank"`
	VerifiedName string                            `json:"verified_name"`
}

// ChargePaymentMethodDetailsFPX represents details about the FPX PaymentMethod.
type ChargePaymentMethodDetailsFPX struct {
	AccountHolderType ChargePaymentMethodDetailsFPXAccountHolderType `json:"account_holder_type"`
	Bank              ChargePaymentMethodDetailsFPXBank              `json:"bank"`
	TransactionID     string                                         `json:"transaction_id"`
}

// ChargePaymentMethodDetailsGiropay represents details about the Giropay PaymentMethod.
type ChargePaymentMethodDetailsGiropay struct {
	BankCode     string `json:"bank_code"`
	BankName     string `json:"bank_name"`
	Bic          string `json:"bic"`
	VerifiedName string `json:"verified_name"`
}

// ChargePaymentMethodDetailsGrabpay represents details about the Grabpay PaymentMethod.
type ChargePaymentMethodDetailsGrabpay struct {
	TransactionID string `json:"transaction_id"`
}

// ChargePaymentMethodDetailsIdeal represents details about the Ideal PaymentMethod.
type ChargePaymentMethodDetailsIdeal struct {
	Bank                      ChargePaymentMethodDetailsIdealBank `json:"bank"`
	Bic                       ChargePaymentMethodDetailsIdealBic  `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod                      `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                            `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                              `json:"iban_last4"`
	VerifiedName              string                              `json:"verified_name"`
}

// ChargePaymentMethodDetailsInteracPresent represents details about the InteracPresent PaymentMethod.
type ChargePaymentMethodDetailsInteracPresent struct {
	Brand            string                                             `json:"brand"`
	CardholderName   string                                             `json:"cardholder_name"`
	Country          string                                             `json:"country"`
	Description      string                                             `json:"description"`
	EmvAuthData      string                                             `json:"emv_auth_data"`
	ExpMonth         int64                                              `json:"exp_month"`
	ExpYear          int64                                              `json:"exp_year"`
	Fingerprint      string                                             `json:"fingerprint"`
	Funding          string                                             `json:"funding"`
	GeneratedCard    string                                             `json:"generated_card"`
	Iin              string                                             `json:"iin"`
	Issuer           string                                             `json:"issuer"`
	Last4            string                                             `json:"last4"`
	Network          string                                             `json:"network"`
	PreferredLocales []string                                           `json:"preferred_locales"`
	ReadMethod       ChargePaymentMethodDetailsInteracPresentReadMethod `json:"read_method"`
	Receipt          *ChargePaymentMethodDetailsInteracPresentReceipt   `json:"receipt"`
}

// ChargePaymentMethodDetailsInteracPresentReceipt represents details about the InteracPresent Receipt.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	AccountType                  ChargePaymentMethodDetailsInteracPresentReceiptAccountType `json:"account_type"`
	ApplicationCryptogram        string                                                     `json:"application_cryptogram"`
	ApplicationPreferredName     string                                                     `json:"application_preferred_name"`
	AuthorizationCode            string                                                     `json:"authorization_code"`
	AuthorizationResponseCode    string                                                     `json:"authorization_response_code"`
	CardholderVerificationMethod string                                                     `json:"cardholder_verification_method"`
	DedicatedFileName            string                                                     `json:"dedicated_file_name"`
	TerminalVerificationResults  string                                                     `json:"terminal_verification_results"`
	TransactionStatusInformation string                                                     `json:"transaction_status_information"`
}

// ChargePaymentMethodDetailsKlarna represents details for the Klarna
// PaymentMethod.
type ChargePaymentMethodDetailsKlarna struct{}

// ChargePaymentMethodDetailsMultibanco represents details about the Multibanco PaymentMethod.
type ChargePaymentMethodDetailsMultibanco struct {
	Entity    string `json:"entity"`
	Reference string `json:"reference"`
}

// ChargePaymentMethodDetailsOXXO represents details about the OXXO PaymentMethod.
type ChargePaymentMethodDetailsOXXO struct {
	Number string `json:"number"`
}

// ChargePaymentMethodDetailsP24 represents details about the P24 PaymentMethod.
type ChargePaymentMethodDetailsP24 struct {
	Bank         ChargePaymentMethodDetailsP24Bank `json:"bank"`
	Reference    string                            `json:"reference"`
	VerifiedName string                            `json:"verified_name"`
}
type ChargePaymentMethodDetailsSepaCreditTransfer struct {
	BankName string `json:"bank_name"`
	Bic      string `json:"bic"`
	Iban     string `json:"iban"`
}

// ChargePaymentMethodDetailsSepaDebit represents details about the Sepa Debit PaymentMethod.
type ChargePaymentMethodDetailsSepaDebit struct {
	BankCode    string `json:"bank_code"`
	BranchCode  string `json:"branch_code"`
	Country     string `json:"country"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

// ChargePaymentMethodDetailsSofort represents details about the Sofort PaymentMethod.
type ChargePaymentMethodDetailsSofort struct {
	BankCode                  string                                            `json:"bank_code"`
	BankName                  string                                            `json:"bank_name"`
	Bic                       string                                            `json:"bic"`
	Country                   string                                            `json:"country"`
	GeneratedSepaDebit        *PaymentMethod                                    `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate                                          `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string                                            `json:"iban_last4"`
	PreferredLanguage         ChargePaymentMethodDetailsSofortPreferredLanguage `json:"preferred_language"`
	VerifiedName              string                                            `json:"verified_name"`
}

// ChargePaymentMethodDetailsStripeAccount represents details about the StripeAccount PaymentMethod.
type ChargePaymentMethodDetailsStripeAccount struct{}

// ChargePaymentMethodDetailsWechat represents details about the Wechat PaymentMethod.
type ChargePaymentMethodDetailsWechat struct{}

// ChargePaymentMethodDetails represents the details about the PaymentMethod associated with the
// charge.
type ChargePaymentMethodDetails struct {
	AchCreditTransfer  *ChargePaymentMethodDetailsAchCreditTransfer  `json:"ach_credit_transfer"`
	AchDebit           *ChargePaymentMethodDetailsAchDebit           `json:"ach_debit"`
	ACSSDebit          *ChargePaymentMethodDetailsACSSDebit          `json:"acss_debit"`
	AfterpayClearpay   *ChargePaymentMethodDetailsAfterpayClearpay   `json:"afterpay_clearpay"`
	Alipay             *ChargePaymentMethodDetailsAlipay             `json:"alipay"`
	AuBECSDebit        *ChargePaymentMethodDetailsAuBECSDebit        `json:"au_becs_debit"`
	BACSDebit          *ChargePaymentMethodDetailsBACSDebit          `json:"bacs_debit"`
	Bancontact         *ChargePaymentMethodDetailsBancontact         `json:"bancontact"`
	Card               *ChargePaymentMethodDetailsCard               `json:"card"`
	CardPresent        *ChargePaymentMethodDetailsCardPresent        `json:"card_present"`
	EPS                *ChargePaymentMethodDetailsEPS                `json:"eps"`
	FPX                *ChargePaymentMethodDetailsFPX                `json:"fpx"`
	Giropay            *ChargePaymentMethodDetailsGiropay            `json:"giropay"`
	Grabpay            *ChargePaymentMethodDetailsGrabpay            `json:"grabpay"`
	Ideal              *ChargePaymentMethodDetailsIdeal              `json:"ideal"`
	InteracPresent     *ChargePaymentMethodDetailsInteracPresent     `json:"interac_present"`
	Klarna             *ChargePaymentMethodDetailsKlarna             `json:"klarna"`
	Multibanco         *ChargePaymentMethodDetailsMultibanco         `json:"multibanco"`
	OXXO               *ChargePaymentMethodDetailsOXXO               `json:"oxxo"`
	P24                *ChargePaymentMethodDetailsP24                `json:"p24"`
	SepaCreditTransfer *ChargePaymentMethodDetailsSepaCreditTransfer `json:"sepa_credit_transfer"`
	SepaDebit          *ChargePaymentMethodDetailsSepaDebit          `json:"sepa_debit"`
	Sofort             *ChargePaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *ChargePaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	Type               string                                        `json:"type"`
	Wechat             *ChargePaymentMethodDetailsWechat             `json:"wechat"`
}

// ChargeTransferData represents the information for the transfer_data associated with a charge.
type ChargeTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	APIResource
	AlternateStatementDescriptors *ChargeAlternateStatementDescriptors `json:"alternate_statement_descriptors"`
	Amount                        int64                                `json:"amount"`
	AmountCaptured                int64                                `json:"amount_captured"`
	AmountRefunded                int64                                `json:"amount_refunded"`
	Application                   *Application                         `json:"application"`
	ApplicationFee                *ApplicationFee                      `json:"application_fee"`
	ApplicationFeeAmount          int64                                `json:"application_fee_amount"`
	AuthorizationCode             string                               `json:"authorization_code"`
	BalanceTransaction            *BalanceTransaction                  `json:"balance_transaction"`
	BillingDetails                *ChargeBillingDetails                `json:"billing_details"`
	CalculatedStatementDescriptor string                               `json:"calculated_statement_descriptor"`
	Captured                      bool                                 `json:"captured"`
	Created                       int64                                `json:"created"`
	Currency                      Currency                             `json:"currency"`
	Customer                      *Customer                            `json:"customer"`
	Description                   string                               `json:"description"`
	Destination                   *Account                             `json:"destination"`
	Dispute                       *Dispute                             `json:"dispute"`
	Disputed                      bool                                 `json:"disputed"`
	FailureCode                   string                               `json:"failure_code"`
	FailureMessage                string                               `json:"failure_message"`
	FraudDetails                  *ChargeFraudDetails                  `json:"fraud_details"`
	ID                            string                               `json:"id"`
	Invoice                       *Invoice                             `json:"invoice"`
	Level3                        *ChargeLevel3                        `json:"level3"`
	Livemode                      bool                                 `json:"livemode"`
	Metadata                      map[string]string                    `json:"metadata"`
	Object                        string                               `json:"object"`
	OnBehalfOf                    *Account                             `json:"on_behalf_of"`
	Order                         *Order                               `json:"order"`
	Outcome                       *ChargeOutcome                       `json:"outcome"`
	Paid                          bool                                 `json:"paid"`
	PaymentIntent                 *PaymentIntent                       `json:"payment_intent"`
	PaymentMethod                 string                               `json:"payment_method"`
	PaymentMethodDetails          *ChargePaymentMethodDetails          `json:"payment_method_details"`
	ReceiptEmail                  string                               `json:"receipt_email"`
	ReceiptNumber                 string                               `json:"receipt_number"`
	ReceiptURL                    string                               `json:"receipt_url"`
	Refunded                      bool                                 `json:"refunded"`
	Refunds                       *RefundList                          `json:"refunds"`
	Review                        *Review                              `json:"review"`
	Shipping                      *ShippingDetails                     `json:"shipping"`
	Source                        *PaymentSource                       `json:"source"`
	SourceTransfer                *Transfer                            `json:"source_transfer"`
	StatementDescriptor           string                               `json:"statement_descriptor"`
	StatementDescriptorSuffix     string                               `json:"statement_descriptor_suffix"`
	Status                        string                               `json:"status"`
	Transfer                      *Transfer                            `json:"transfer"`
	TransferData                  *ChargeTransferData                  `json:"transfer_data"`
	TransferGroup                 string                               `json:"transfer_group"`
}

// ChargeOutcomeRule tells you the Radar rule that blocked the charge, if any.
type ChargeOutcomeRule struct {
	Action    string `json:"action"`
	ID        string `json:"id"`
	Predicate string `json:"predicate"`
}

// ChargeOutcome is the charge's outcome that details whether a payment
// was accepted and why.
type ChargeOutcome struct {
	NetworkStatus string             `json:"network_status"`
	Reason        string             `json:"reason"`
	RiskLevel     string             `json:"risk_level"`
	RiskScore     int64              `json:"risk_score"`
	Rule          *ChargeOutcomeRule `json:"rule"`
	SellerMessage string             `json:"seller_message"`
	Type          string             `json:"type"`
}

// ChargeList is a list of charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}
