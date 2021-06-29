//
//
// File generated from our OpenAPI spec
//
//

package stripe

type TokenBankAccountParams struct {
	AccountHolderName *string `form:"account_holder_name"`
	AccountHolderType *string `form:"account_holder_type"`
	AccountNumber     *string `form:"account_number"`
	Country           *string `form:"country"`
	Currency          *string `form:"currency"`
	RoutingNumber     *string `form:"routing_number"`
}
type TokenCardParams struct {
	AddressCity    *string `form:"address_city"`
	AddressCountry *string `form:"address_country"`
	AddressLine1   *string `form:"address_line1"`
	AddressLine2   *string `form:"address_line2"`
	AddressState   *string `form:"address_state"`
	AddressZip     *string `form:"address_zip"`
	Currency       *string `form:"currency"`
	CVC            *string `form:"cvc"`
	ExpMonth       *string `form:"exp_month"`
	ExpYear        *string `form:"exp_year"`
	Name           *string `form:"name"`
	Number         *string `form:"number"`
}

// TokenCVCUpdateParams is the set of parameters that can be used when creating a CVC token.
type TokenCVCUpdateParams struct {
	CVC *string `form:"cvc"`
}
type TokenAccountCompanyAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenAccountCompanyAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenAccountCompanyVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}
type TokenAccountCompanyVerificationParams struct {
	Document *TokenAccountCompanyVerificationDocumentParams `form:"document"`
}
type TokenAccountCompanyParams struct {
	Address            *AddressParams                         `form:"address"`
	AddressKana        *TokenAccountCompanyAddressKanaParams  `form:"address_kana"`
	AddressKanji       *TokenAccountCompanyAddressKanjiParams `form:"address_kanji"`
	DirectorsProvided  *bool                                  `form:"directors_provided"`
	ExecutivesProvided *bool                                  `form:"executives_provided"`
	Name               *string                                `form:"name"`
	NameKana           *string                                `form:"name_kana"`
	NameKanji          *string                                `form:"name_kanji"`
	OwnersProvided     *bool                                  `form:"owners_provided"`
	Phone              *string                                `form:"phone"`
	RegistrationNumber *string                                `form:"registration_number"`
	Structure          *string                                `form:"structure"`
	TaxID              *string                                `form:"tax_id"`
	TaxIDRegistrar     *string                                `form:"tax_id_registrar"`
	VATID              *string                                `form:"vat_id"`
	Verification       *TokenAccountCompanyVerificationParams `form:"verification"`
}
type TokenAccountIndividualAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenAccountIndividualAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenAccountIndividualDobParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}
type TokenAccountIndividualVerificationAdditionalDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}
type TokenAccountIndividualVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}
type TokenAccountIndividualVerificationParams struct {
	AdditionalDocument *TokenAccountIndividualVerificationAdditionalDocumentParams `form:"additional_document"`
	Document           *TokenAccountIndividualVerificationDocumentParams           `form:"document"`
}
type TokenAccountIndividualParams struct {
	Address           *AddressParams                            `form:"address"`
	AddressKana       *TokenAccountIndividualAddressKanaParams  `form:"address_kana"`
	AddressKanji      *TokenAccountIndividualAddressKanjiParams `form:"address_kanji"`
	Dob               *TokenAccountIndividualDobParams          `form:"dob"`
	Email             *string                                   `form:"email"`
	FirstName         *string                                   `form:"first_name"`
	FirstNameKana     *string                                   `form:"first_name_kana"`
	FirstNameKanji    *string                                   `form:"first_name_kanji"`
	Gender            *string                                   `form:"gender"`
	IDNumber          *string                                   `form:"id_number"`
	LastName          *string                                   `form:"last_name"`
	LastNameKana      *string                                   `form:"last_name_kana"`
	LastNameKanji     *string                                   `form:"last_name_kanji"`
	MaidenName        *string                                   `form:"maiden_name"`
	Metadata          map[string]string                         `form:"metadata"`
	Phone             *string                                   `form:"phone"`
	PoliticalExposure *string                                   `form:"political_exposure"`
	SsnLast4          *string                                   `form:"ssn_last_4"`
	Verification      *TokenAccountIndividualVerificationParams `form:"verification"`
}

// TokenAccountParams is the set of parameters that can be used when creating a Account token.
type TokenAccountParams struct {
	BusinessType        *string                       `form:"business_type"`
	Company             *TokenAccountCompanyParams    `form:"company"`
	Individual          *TokenAccountIndividualParams `form:"individual"`
	TosShownAndAccepted *bool                         `form:"tos_shown_and_accepted"`
}

// TokenParams is the set of parameters that can be used when creating a token.
// For more details see:
// - https://stripe.com/docs/api#create_card_token
// - https://stripe.com/docs/api#create_bank_account_token
// - https://stripe.com/docs/api/tokens/create_account
// - https://stripe.com/docs/api/tokens/create_person
type TokenParams struct {
	Params      `form:"*"`
	Account     *TokenAccountParams     `form:"account"`
	BankAccount *TokenBankAccountParams `form:"bank_account"`
	// Card *[todo({"shape":"nullable","type":{"shape":"ref","namespaces":[],"ref":"TokenCardParams"}} | {"shape":"primitive","primitive":"string","maxLength":5000})] `form:"card"`
	Customer  *string               `form:"customer"`
	CVCUpdate *TokenCVCUpdateParams `form:"cvc_update"`
	Person    *TokenPersonParams    `form:"person"`
	Pii       *TokenPiiParams       `form:"pii"`
}
type TokenPersonAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenPersonAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}
type TokenPersonDobParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}
type TokenPersonDocumentsCompanyAuthorizationParams struct {
	Files []*string `form:"files"`
}
type TokenPersonDocumentsPassportParams struct {
	Files []*string `form:"files"`
}
type TokenPersonDocumentsVisaParams struct {
	Files []*string `form:"files"`
}
type TokenPersonDocumentsParams struct {
	CompanyAuthorization *TokenPersonDocumentsCompanyAuthorizationParams `form:"company_authorization"`
	Passport             *TokenPersonDocumentsPassportParams             `form:"passport"`
	Visa                 *TokenPersonDocumentsVisaParams                 `form:"visa"`
}
type TokenPersonRelationshipParams struct {
	Director         *bool    `form:"director"`
	Executive        *bool    `form:"executive"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}
type TokenPersonVerificationAdditionalDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}
type TokenPersonVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}
type TokenPersonVerificationParams struct {
	AdditionalDocument *TokenPersonVerificationAdditionalDocumentParams `form:"additional_document"`
	Document           *TokenPersonVerificationDocumentParams           `form:"document"`
}
type TokenPersonParams struct {
	Address           *AddressParams                 `form:"address"`
	AddressKana       *TokenPersonAddressKanaParams  `form:"address_kana"`
	AddressKanji      *TokenPersonAddressKanjiParams `form:"address_kanji"`
	Dob               *TokenPersonDobParams          `form:"dob"`
	Documents         *TokenPersonDocumentsParams    `form:"documents"`
	Email             *string                        `form:"email"`
	FirstName         *string                        `form:"first_name"`
	FirstNameKana     *string                        `form:"first_name_kana"`
	FirstNameKanji    *string                        `form:"first_name_kanji"`
	Gender            *string                        `form:"gender"`
	IDNumber          *string                        `form:"id_number"`
	LastName          *string                        `form:"last_name"`
	LastNameKana      *string                        `form:"last_name_kana"`
	LastNameKanji     *string                        `form:"last_name_kanji"`
	MaidenName        *string                        `form:"maiden_name"`
	Metadata          map[string]string              `form:"metadata"`
	Nationality       *string                        `form:"nationality"`
	Phone             *string                        `form:"phone"`
	PoliticalExposure *string                        `form:"political_exposure"`
	Relationship      *TokenPersonRelationshipParams `form:"relationship"`
	SsnLast4          *string                        `form:"ssn_last_4"`
	Verification      *TokenPersonVerificationParams `form:"verification"`
}
type TokenPiiParams struct {
	IDNumber *string `form:"id_number"`
}

// Token is the resource representing a Stripe token.
// For more details see https://stripe.com/docs/api#tokens.
type Token struct {
	APIResource
	BankAccount *BankAccount `json:"bank_account"`
	Card        *Card        `json:"card"`
	ClientIP    string       `json:"client_ip"`
	Created     int64        `json:"created"`
	ID          string       `json:"id"`
	Livemode    bool         `json:"livemode"`
	Object      string       `json:"object"`
	Type        string       `json:"type"`
	Used        bool         `json:"used"`
}
