//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
type PaymentMethodEPSBank string

// List of values that PaymentMethodEPSBank can take
const (
	PaymentMethodEPSBankArzteUndApothekerBank                PaymentMethodEPSBank = "arzte_und_apotheker_bank"
	PaymentMethodEPSBankAustrianAnadiBankAg                  PaymentMethodEPSBank = "austrian_anadi_bank_ag"
	PaymentMethodEPSBankBankAustria                          PaymentMethodEPSBank = "bank_austria"
	PaymentMethodEPSBankBankhausCarlSpangler                 PaymentMethodEPSBank = "bankhaus_carl_spangler"
	PaymentMethodEPSBankBankhausSchelhammerUndSchatteraAg    PaymentMethodEPSBank = "bankhaus_schelhammer_und_schattera_ag"
	PaymentMethodEPSBankBawagPskAg                           PaymentMethodEPSBank = "bawag_psk_ag"
	PaymentMethodEPSBankBksBankAg                            PaymentMethodEPSBank = "bks_bank_ag"
	PaymentMethodEPSBankBrullKallmusBankAg                   PaymentMethodEPSBank = "brull_kallmus_bank_ag"
	PaymentMethodEPSBankBtvVierLanderBank                    PaymentMethodEPSBank = "btv_vier_lander_bank"
	PaymentMethodEPSBankCapitalBankGraweGruppeAg             PaymentMethodEPSBank = "capital_bank_grawe_gruppe_ag"
	PaymentMethodEPSBankDolomitenbank                        PaymentMethodEPSBank = "dolomitenbank"
	PaymentMethodEPSBankEasybankAg                           PaymentMethodEPSBank = "easybank_ag"
	PaymentMethodEPSBankErsteBankUndSparkassen               PaymentMethodEPSBank = "erste_bank_und_sparkassen"
	PaymentMethodEPSBankHypoAlpeadriabankInternationalAg     PaymentMethodEPSBank = "hypo_alpeadriabank_international_ag"
	PaymentMethodEPSBankHypoBankBurgenlandAktiengesellschaft PaymentMethodEPSBank = "hypo_bank_burgenland_aktiengesellschaft"
	PaymentMethodEPSBankHypoNoeLbFurNiederosterreichUWien    PaymentMethodEPSBank = "hypo_noe_lb_fur_niederosterreich_u_wien"
	PaymentMethodEPSBankHypoOberosterreichSalzburgSteiermark PaymentMethodEPSBank = "hypo_oberosterreich_salzburg_steiermark"
	PaymentMethodEPSBankHypoTirolBankAg                      PaymentMethodEPSBank = "hypo_tirol_bank_ag"
	PaymentMethodEPSBankHypoVorarlbergBankAg                 PaymentMethodEPSBank = "hypo_vorarlberg_bank_ag"
	PaymentMethodEPSBankMarchfelderBank                      PaymentMethodEPSBank = "marchfelder_bank"
	PaymentMethodEPSBankOberbankAg                           PaymentMethodEPSBank = "oberbank_ag"
	PaymentMethodEPSBankRaiffeisenBankengruppeOsterreich     PaymentMethodEPSBank = "raiffeisen_bankengruppe_osterreich"
	PaymentMethodEPSBankSchoellerbankAg                      PaymentMethodEPSBank = "schoellerbank_ag"
	PaymentMethodEPSBankSpardaBankWien                       PaymentMethodEPSBank = "sparda_bank_wien"
	PaymentMethodEPSBankVolksbankGruppe                      PaymentMethodEPSBank = "volksbank_gruppe"
	PaymentMethodEPSBankVolkskreditbankAg                    PaymentMethodEPSBank = "volkskreditbank_ag"
	PaymentMethodEPSBankVrBankBraunau                        PaymentMethodEPSBank = "vr_bank_braunau"
)

// PaymentMethodFPXAccountHolderType is a list of string values that FPX AccountHolderType accepts.
type PaymentMethodFPXAccountHolderType string

// List of values that PaymentMethodFPXAccountHolderType can take
const (
	PaymentMethodFPXAccountHolderTypeCompany    PaymentMethodFPXAccountHolderType = "company"
	PaymentMethodFPXAccountHolderTypeIndividual PaymentMethodFPXAccountHolderType = "individual"
)

// The customer's bank, if provided. Can be one of `affin_bank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, or `pb_enterprise`.
type PaymentMethodFPXBank string

// List of values that PaymentMethodFPXBank can take
const (
	PaymentMethodFPXBankAffinBank         PaymentMethodFPXBank = "affin_bank"
	PaymentMethodFPXBankAllianceBank      PaymentMethodFPXBank = "alliance_bank"
	PaymentMethodFPXBankAmbank            PaymentMethodFPXBank = "ambank"
	PaymentMethodFPXBankBankIslam         PaymentMethodFPXBank = "bank_islam"
	PaymentMethodFPXBankBankMuamalat      PaymentMethodFPXBank = "bank_muamalat"
	PaymentMethodFPXBankBankRakyat        PaymentMethodFPXBank = "bank_rakyat"
	PaymentMethodFPXBankBsn               PaymentMethodFPXBank = "bsn"
	PaymentMethodFPXBankCimb              PaymentMethodFPXBank = "cimb"
	PaymentMethodFPXBankDeutscheBank      PaymentMethodFPXBank = "deutsche_bank"
	PaymentMethodFPXBankHongLeongBank     PaymentMethodFPXBank = "hong_leong_bank"
	PaymentMethodFPXBankHsbc              PaymentMethodFPXBank = "hsbc"
	PaymentMethodFPXBankKfh               PaymentMethodFPXBank = "kfh"
	PaymentMethodFPXBankMaybank2e         PaymentMethodFPXBank = "maybank2e"
	PaymentMethodFPXBankMaybank2u         PaymentMethodFPXBank = "maybank2u"
	PaymentMethodFPXBankOcbc              PaymentMethodFPXBank = "ocbc"
	PaymentMethodFPXBankPbEnterprise      PaymentMethodFPXBank = "pb_enterprise"
	PaymentMethodFPXBankPublicBank        PaymentMethodFPXBank = "public_bank"
	PaymentMethodFPXBankRhb               PaymentMethodFPXBank = "rhb"
	PaymentMethodFPXBankStandardChartered PaymentMethodFPXBank = "standard_chartered"
	PaymentMethodFPXBankUob               PaymentMethodFPXBank = "uob"
)

// The customer's bank, if provided. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
type PaymentMethodIdealBank string

// List of values that PaymentMethodIdealBank can take
const (
	PaymentMethodIdealBankAbnAmro       PaymentMethodIdealBank = "abn_amro"
	PaymentMethodIdealBankAsnBank       PaymentMethodIdealBank = "asn_bank"
	PaymentMethodIdealBankBunq          PaymentMethodIdealBank = "bunq"
	PaymentMethodIdealBankHandelsbanken PaymentMethodIdealBank = "handelsbanken"
	PaymentMethodIdealBankIng           PaymentMethodIdealBank = "ing"
	PaymentMethodIdealBankKnab          PaymentMethodIdealBank = "knab"
	PaymentMethodIdealBankMoneyou       PaymentMethodIdealBank = "moneyou"
	PaymentMethodIdealBankRabobank      PaymentMethodIdealBank = "rabobank"
	PaymentMethodIdealBankRegiobank     PaymentMethodIdealBank = "regiobank"
	PaymentMethodIdealBankRevolut       PaymentMethodIdealBank = "revolut"
	PaymentMethodIdealBankSnsBank       PaymentMethodIdealBank = "sns_bank"
	PaymentMethodIdealBankTriodosBank   PaymentMethodIdealBank = "triodos_bank"
	PaymentMethodIdealBankVanLanschot   PaymentMethodIdealBank = "van_lanschot"
)

// The Bank Identifier Code of the customer's bank, if the bank was provided.
type PaymentMethodIdealBic string

// List of values that PaymentMethodIdealBic can take
const (
	PaymentMethodIdealBicABNANL2A PaymentMethodIdealBic = "ABNANL2A"
	PaymentMethodIdealBicASNBNL21 PaymentMethodIdealBic = "ASNBNL21"
	PaymentMethodIdealBicBUNQNL2A PaymentMethodIdealBic = "BUNQNL2A"
	PaymentMethodIdealBicFVLBNL22 PaymentMethodIdealBic = "FVLBNL22"
	PaymentMethodIdealBicHANDNL2A PaymentMethodIdealBic = "HANDNL2A"
	PaymentMethodIdealBicINGBNL2A PaymentMethodIdealBic = "INGBNL2A"
	PaymentMethodIdealBicKNABNL2H PaymentMethodIdealBic = "KNABNL2H"
	PaymentMethodIdealBicMOYONL21 PaymentMethodIdealBic = "MOYONL21"
	PaymentMethodIdealBicRABONL2U PaymentMethodIdealBic = "RABONL2U"
	PaymentMethodIdealBicRBRBNL21 PaymentMethodIdealBic = "RBRBNL21"
	PaymentMethodIdealBicREVOLT21 PaymentMethodIdealBic = "REVOLT21"
	PaymentMethodIdealBicSNSBNL2A PaymentMethodIdealBic = "SNSBNL2A"
	PaymentMethodIdealBicTRIONL2U PaymentMethodIdealBic = "TRIONL2U"
)

// The customer's bank, if provided.
type PaymentMethodP24Bank string

// List of values that PaymentMethodP24Bank can take
const (
	PaymentMethodP24BankAliorBank            PaymentMethodP24Bank = "alior_bank"
	PaymentMethodP24BankBankMillennium       PaymentMethodP24Bank = "bank_millennium"
	PaymentMethodP24BankBankNowyBfgSa        PaymentMethodP24Bank = "bank_nowy_bfg_sa"
	PaymentMethodP24BankBankPekaoSa          PaymentMethodP24Bank = "bank_pekao_sa"
	PaymentMethodP24BankBankiSpbdzielcze     PaymentMethodP24Bank = "banki_spbdzielcze"
	PaymentMethodP24BankBlik                 PaymentMethodP24Bank = "blik"
	PaymentMethodP24BankBnpParibas           PaymentMethodP24Bank = "bnp_paribas"
	PaymentMethodP24BankBoz                  PaymentMethodP24Bank = "boz"
	PaymentMethodP24BankCitiHandlowy         PaymentMethodP24Bank = "citi_handlowy"
	PaymentMethodP24BankCreditAgricole       PaymentMethodP24Bank = "credit_agricole"
	PaymentMethodP24BankEnvelobank           PaymentMethodP24Bank = "envelobank"
	PaymentMethodP24BankEtransferPocztowy24  PaymentMethodP24Bank = "etransfer_pocztowy24"
	PaymentMethodP24BankGetinBank            PaymentMethodP24Bank = "getin_bank"
	PaymentMethodP24BankIdeabank             PaymentMethodP24Bank = "ideabank"
	PaymentMethodP24BankIng                  PaymentMethodP24Bank = "ing"
	PaymentMethodP24BankInteligo             PaymentMethodP24Bank = "inteligo"
	PaymentMethodP24BankMbankMtransfer       PaymentMethodP24Bank = "mbank_mtransfer"
	PaymentMethodP24BankNestPrzelew          PaymentMethodP24Bank = "nest_przelew"
	PaymentMethodP24BankNoblePay             PaymentMethodP24Bank = "noble_pay"
	PaymentMethodP24BankPbacZIpko            PaymentMethodP24Bank = "pbac_z_ipko"
	PaymentMethodP24BankPlusBank             PaymentMethodP24Bank = "plus_bank"
	PaymentMethodP24BankSantanderPrzelew24   PaymentMethodP24Bank = "santander_przelew24"
	PaymentMethodP24BankTmobileUsbugiBankowe PaymentMethodP24Bank = "tmobile_usbugi_bankowe"
	PaymentMethodP24BankToyotaBank           PaymentMethodP24Bank = "toyota_bank"
	PaymentMethodP24BankVolkswagenBank       PaymentMethodP24Bank = "volkswagen_bank"
)

// PaymentMethodType is the list of allowed values for the payment method type.
type PaymentMethodType string

// List of values that PaymentMethodType can take.
const (
	PaymentMethodTypeACSSDebit        PaymentMethodType = "acss_debit"
	PaymentMethodTypeAfterpayClearpay PaymentMethodType = "afterpay_clearpay"
	PaymentMethodTypeAlipay           PaymentMethodType = "alipay"
	PaymentMethodTypeAUBECSDebit      PaymentMethodType = "au_becs_debit"
	PaymentMethodTypeBACSDebit        PaymentMethodType = "bacs_debit"
	PaymentMethodTypeBancontact       PaymentMethodType = "bancontact"
	PaymentMethodTypeBoleto           PaymentMethodType = "boleto"
	PaymentMethodTypeCard             PaymentMethodType = "card"
	PaymentMethodTypeCardPresent      PaymentMethodType = "card_present"
	PaymentMethodTypeEPS              PaymentMethodType = "eps"
	PaymentMethodTypeFPX              PaymentMethodType = "fpx"
	PaymentMethodTypeGiropay          PaymentMethodType = "giropay"
	PaymentMethodTypeGrabpay          PaymentMethodType = "grabpay"
	PaymentMethodTypeIdeal            PaymentMethodType = "ideal"
	PaymentMethodTypeInteracPresent   PaymentMethodType = "interac_present"
	PaymentMethodTypeOXXO             PaymentMethodType = "oxxo"
	PaymentMethodTypeP24              PaymentMethodType = "p24"
	PaymentMethodTypeSepaDebit        PaymentMethodType = "sepa_debit"
	PaymentMethodTypeSofort           PaymentMethodType = "sofort"
	PaymentMethodTypeWechatPay        PaymentMethodType = "wechat_pay"
)

// PaymentMethodCardWalletType is the list of allowed values for the type a wallet can take on
// a Card PaymentMethod.
type PaymentMethodCardWalletType string

// List of values that PaymentMethodCardWalletType can take.
const (
	PaymentMethodCardWalletTypeAmexExpressCheckout PaymentMethodCardWalletType = "amex_express_checkout"
	PaymentMethodCardWalletTypeApplePay            PaymentMethodCardWalletType = "apple_pay"
	PaymentMethodCardWalletTypeGooglePay           PaymentMethodCardWalletType = "google_pay"
	PaymentMethodCardWalletTypeMasterpass          PaymentMethodCardWalletType = "masterpass"
	PaymentMethodCardWalletTypeSamsungPay          PaymentMethodCardWalletType = "samsung_pay"
	PaymentMethodCardWalletTypeVisaCheckout        PaymentMethodCardWalletType = "visa_checkout"
)

// BillingDetailsParams is the set of parameters that can be used as billing details
// when creating or updating a PaymentMethod
type BillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// PaymentMethodACSSDebitParams TODO
type PaymentMethodACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}

// PaymentMethodAfterpayClearpayParams is the set of parameters allowed for the
// `afterpay_clearpay`  hash when creating a PaymentMethod of type AfterpayClearpay.
type PaymentMethodAfterpayClearpayParams struct{}

// PaymentMethodAlipayParams is the set of parameters allowed for the `alipay` hash when creating a
// PaymentMethod of type Alipay.
type PaymentMethodAlipayParams struct{}

// PaymentMethodAUBECSDebitParams is the set of parameters allowed for the `AUBECSDebit` hash when creating a
// PaymentMethod of type AUBECSDebit.
type PaymentMethodAUBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BSBNumber     *string `form:"bsb_number"`
}

// PaymentMethodBACSDebitParams is the set of parameters allowed for BACS Debit payment method.
type PaymentMethodBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}

// PaymentMethodBancontactParams is the set of parameters allowed for the `bancontact` hash when creating a
// PaymentMethod of type Bancontact.
type PaymentMethodBancontactParams struct{}

// PaymentMethodBoletoParams is the set of parameters allowed for the `boleto` hash when creating a
// PaymentMethod of type Boleto
type PaymentMethodBoletoParams struct {
	TaxID *string `form:"tax_id"`
}

// PaymentMethodCardParams is the set of parameters allowed for the `card` hash when creating a
// PaymentMethod of type card.
type PaymentMethodCardParams struct {
	CVC      *string `form:"cvc"`
	ExpMonth *int64  `form:"exp_month"`
	ExpYear  *int64  `form:"exp_year"`
	Number   *string `form:"number"`
	Token    *string `form:"token"`
}

// PaymentMethodEPSParams is the set of parameters allowed for the `eps` hash when creating a
// PaymentMethod of type EPS.
type PaymentMethodEPSParams struct {
	Bank *string `form:"bank"`
}

// PaymentMethodFPXParams is the set of parameters allowed for the `fpx` hash when creating a
// PaymentMethod of type fpx.
type PaymentMethodFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// PaymentMethodGiropayParams is the set of parameters allowed for the `giropay` hash when creating a
// PaymentMethod of type Giropay.
type PaymentMethodGiropayParams struct{}

// PaymentMethodGrabpayParams is the set of parameters allowed for the `grabpay` hash when creating a
// PaymentMethod of type Grabpay.
type PaymentMethodGrabpayParams struct{}

// PaymentMethodIdealParams is the set of parameters allowed for the `ideal` hash when creating a
// PaymentMethod of type ideal.
type PaymentMethodIdealParams struct {
	Bank *string `form:"bank"`
}

// PaymentMethodInteracPresentParams is the set of parameters allowed for the `interac_present` hash when creating a
// PaymentMethod of type interac_present.
type PaymentMethodInteracPresentParams struct{}

// PaymentMethodOXXOParams is the set of parameters allowed for the `oxxo` hash when creating a
// PaymentMethod of type OXXO.
type PaymentMethodOXXOParams struct{}

// PaymentMethodP24Params is the set of parameters allowed for the `p24` hash when creating a
// PaymentMethod of type P24.
type PaymentMethodP24Params struct {
	Bank *string `form:"bank"`
}

// PaymentMethodSepaDebitParams is the set of parameters allowed for the `sepa_debit` hash when
// creating a PaymentMethod of type sepa_debit.
type PaymentMethodSepaDebitParams struct {
	Iban *string `form:"iban"`
}

// PaymentMethodSofortParams is the set of parameters allowed for the `sofort` hash when
// creating a PaymentMethod of type sofort.
type PaymentMethodSofortParams struct {
	Country *string `form:"country"`
}

// PaymentMethodWechatPayParams is the set of parameters allowed for the `wechat_pay` hash when
// creating a PaymentMethod of type wechat_pay.
type PaymentMethodWechatPayParams struct{}

// PaymentMethodParams is the set of parameters that can be used when creating or updating a
// PaymentMethod.
type PaymentMethodParams struct {
	Params           `form:"*"`
	ACSSDebit        *PaymentMethodACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentMethodAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentMethodAlipayParams           `form:"alipay"`
	AUBECSDebit      *PaymentMethodAUBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentMethodBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentMethodBancontactParams       `form:"bancontact"`
	BillingDetails   *BillingDetailsParams                `form:"billing_details"`
	Boleto           *PaymentMethodBoletoParams           `form:"boleto"`
	// Card *[todo({"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentMethodCardParams"}} | {"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"PaymentMethodCardParams"}})] `form:"card"`
	Customer       *string                            `form:"customer"`
	EPS            *PaymentMethodEPSParams            `form:"eps"`
	FPX            *PaymentMethodFPXParams            `form:"fpx"`
	Giropay        *PaymentMethodGiropayParams        `form:"giropay"`
	Grabpay        *PaymentMethodGrabpayParams        `form:"grabpay"`
	Ideal          *PaymentMethodIdealParams          `form:"ideal"`
	InteracPresent *PaymentMethodInteracPresentParams `form:"interac_present"`
	OXXO           *PaymentMethodOXXOParams           `form:"oxxo"`
	P24            *PaymentMethodP24Params            `form:"p24"`
	PaymentMethod  *string                            `form:"payment_method"`
	SepaDebit      *PaymentMethodSepaDebitParams      `form:"sepa_debit"`
	Sofort         *PaymentMethodSofortParams         `form:"sofort"`
	Type           *string                            `form:"type"`
	WechatPay      *PaymentMethodWechatPayParams      `form:"wechat_pay"`
}

// PaymentMethodAttachParams is the set of parameters that can be used when attaching a
// PaymentMethod to a Customer.
type PaymentMethodAttachParams struct {
	Params   `form:"*"`
	Customer *string `form:"customer"`
}

// PaymentMethodDetachParams is the set of parameters that can be used when detaching a
// PaymentMethod.
type PaymentMethodDetachParams struct {
	Params `form:"*"`
}

// PaymentMethodListParams is the set of parameters that can be used when listing PaymentMethods.
type PaymentMethodListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Type       *string `form:"type"`
}

// PaymentMethodACSSDebit TODO
type PaymentMethodACSSDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	TransitNumber     string `json:"transit_number"`
}

// PaymentMethodAfterpayClearpay represents the AfterpayClearpay properties.
type PaymentMethodAfterpayClearpay struct{}

// PaymentMethodAlipay represents the Alipay properties.
type PaymentMethodAlipay struct{}

// PaymentMethodAUBECSDebit represents AUBECSDebit-specific properties (Australia Only).
type PaymentMethodAUBECSDebit struct {
	BSBNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
}

// PaymentMethodBACSDebit represents the BACS Debit properties.
type PaymentMethodBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	SortCode    string `json:"sort_code"`
}

// PaymentMethodBancontact represents the Bancontact properties.
type PaymentMethodBancontact struct{}
type PaymentMethodBillingDetails struct {
	Address *Address `json:"address"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
}

// PaymentMethodBoleto represents the Boleto properties.
type PaymentMethodBoleto struct {
	TaxID string `json:"tax_id"`
}

// PaymentMethodCardChecks represents the checks associated with a Card PaymentMethod.
type PaymentMethodCardChecks struct {
	AddressLine1Check      string `json:"address_line1_check"`
	AddressPostalCodeCheck string `json:"address_postal_code_check"`
	CVCCheck               string `json:"cvc_check"`
}

// PaymentMethodCardNetworks represents the card networks that can be used to process the payment.
type PaymentMethodCardNetworks struct {
	Available []string `json:"available"`
	Preferred string   `json:"preferred"`
}

// PaymentMethodCardThreeDSecureUsage represents the 3DS usage for that Card PaymentMethod.
type PaymentMethodCardThreeDSecureUsage struct {
	Supported bool `json:"supported"`
}
type PaymentMethodCardWalletAmexExpressCheckout struct{}
type PaymentMethodCardWalletApplePay struct{}
type PaymentMethodCardWalletGooglePay struct{}
type PaymentMethodCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}
type PaymentMethodCardWalletSamsungPay struct{}
type PaymentMethodCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// PaymentMethodCardWallet represents the details of the card wallet if this Card PaymentMethod
// is part of a card wallet.
type PaymentMethodCardWallet struct {
	AmexExpressCheckout *PaymentMethodCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *PaymentMethodCardWalletApplePay            `json:"apple_pay"`
	DynamicLast4        string                                      `json:"dynamic_last4"`
	GooglePay           *PaymentMethodCardWalletGooglePay           `json:"google_pay"`
	Masterpass          *PaymentMethodCardWalletMasterpass          `json:"masterpass"`
	SamsungPay          *PaymentMethodCardWalletSamsungPay          `json:"samsung_pay"`
	Type                PaymentMethodCardWalletType                 `json:"type"`
	VisaCheckout        *PaymentMethodCardWalletVisaCheckout        `json:"visa_checkout"`
}

// PaymentMethodCard represents the card-specific properties.
type PaymentMethodCard struct {
	Brand             string                              `json:"brand"`
	Checks            *PaymentMethodCardChecks            `json:"checks"`
	Country           string                              `json:"country"`
	Description       string                              `json:"description"`
	ExpMonth          int64                               `json:"exp_month"`
	ExpYear           int64                               `json:"exp_year"`
	Fingerprint       string                              `json:"fingerprint"`
	Funding           string                              `json:"funding"`
	Iin               string                              `json:"iin"`
	Issuer            string                              `json:"issuer"`
	Last4             string                              `json:"last4"`
	Networks          *PaymentMethodCardNetworks          `json:"networks"`
	ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
	Wallet            *PaymentMethodCardWallet            `json:"wallet"`
}

// PaymentMethodCardPresent represents the card-present-specific properties.
type PaymentMethodCardPresent struct{}

// PaymentMethodEPS represents the EPS properties.
type PaymentMethodEPS struct {
	Bank PaymentMethodEPSBank `json:"bank"`
}

// PaymentMethodFPX represents FPX-specific properties (Malaysia Only).
type PaymentMethodFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              PaymentMethodFPXBank              `json:"bank"`
}

// PaymentMethodGiropay represents the Giropay properties.
type PaymentMethodGiropay struct{}

// PaymentMethodGrabpay represents the Grabpay properties.
type PaymentMethodGrabpay struct{}

// PaymentMethodIdeal represents the iDEAL-specific properties.
type PaymentMethodIdeal struct {
	Bank PaymentMethodIdealBank `json:"bank"`
	Bic  PaymentMethodIdealBic  `json:"bic"`
}

// PaymentMethodInteracPresent represents the interac present properties.
type PaymentMethodInteracPresent struct{}

// PaymentMethodOXXO represents the OXXO-specific properties.
type PaymentMethodOXXO struct{}

// PaymentMethodP24 represents the P24 properties.
type PaymentMethodP24 struct {
	Bank PaymentMethodP24Bank `json:"bank"`
}

// PaymentMethodSepaDebitGeneratedFrom represents information about the object
// that generated this PaymentMethod
type PaymentMethodSepaDebitGeneratedFrom struct {
	Charge       *Charge       `json:"charge"`
	SetupAttempt *SetupAttempt `json:"setup_attempt"`
}

// PaymentMethodSepaDebit represents the SEPA-debit-specific properties.
type PaymentMethodSepaDebit struct {
	BankCode      string                               `json:"bank_code"`
	BranchCode    string                               `json:"branch_code"`
	Country       string                               `json:"country"`
	Fingerprint   string                               `json:"fingerprint"`
	GeneratedFrom *PaymentMethodSepaDebitGeneratedFrom `json:"generated_from"`
	Last4         string                               `json:"last4"`
}

// PaymentMethodSofort represents the Sofort-specific properties.
type PaymentMethodSofort struct {
	Country string `json:"country"`
}

// PaymentMethodWechatPay represents the WeChatPay-specific properties.
type PaymentMethodWechatPay struct{}

// PaymentMethod is the resource representing a PaymentMethod.
type PaymentMethod struct {
	APIResource
	ACSSDebit        *PaymentMethodACSSDebit        `json:"acss_debit"`
	AfterpayClearpay *PaymentMethodAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *PaymentMethodAlipay           `json:"alipay"`
	AUBECSDebit      *PaymentMethodAUBECSDebit      `json:"au_becs_debit"`
	BACSDebit        *PaymentMethodBACSDebit        `json:"bacs_debit"`
	Bancontact       *PaymentMethodBancontact       `json:"bancontact"`
	BillingDetails   *PaymentMethodBillingDetails   `json:"billing_details"`
	Boleto           *PaymentMethodBoleto           `json:"boleto"`
	Card             *PaymentMethodCard             `json:"card"`
	CardPresent      *PaymentMethodCardPresent      `json:"card_present"`
	Created          int64                          `json:"created"`
	Customer         *Customer                      `json:"customer"`
	EPS              *PaymentMethodEPS              `json:"eps"`
	FPX              *PaymentMethodFPX              `json:"fpx"`
	Giropay          *PaymentMethodGiropay          `json:"giropay"`
	Grabpay          *PaymentMethodGrabpay          `json:"grabpay"`
	ID               string                         `json:"id"`
	Ideal            *PaymentMethodIdeal            `json:"ideal"`
	InteracPresent   *PaymentMethodInteracPresent   `json:"interac_present"`
	Livemode         bool                           `json:"livemode"`
	Metadata         map[string]string              `json:"metadata"`
	Object           string                         `json:"object"`
	OXXO             *PaymentMethodOXXO             `json:"oxxo"`
	P24              *PaymentMethodP24              `json:"p24"`
	SepaDebit        *PaymentMethodSepaDebit        `json:"sepa_debit"`
	Sofort           *PaymentMethodSofort           `json:"sofort"`
	Type             PaymentMethodType              `json:"type"`
	WechatPay        *PaymentMethodWechatPay        `json:"wechat_pay"`
}

// PaymentMethodList is a list of PaymentMethods as retrieved from a list endpoint.
type PaymentMethodList struct {
	APIResource
	ListMeta
	Data []*PaymentMethod `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentMethod.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentMethod) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentMethod PaymentMethod
	var v paymentMethod
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentMethod(v)
	return nil
}
