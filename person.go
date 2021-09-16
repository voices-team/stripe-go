//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The code for the type of error.
type PersonFutureRequirementsErrorCode string

// List of values that PersonFutureRequirementsErrorCode can take
const (
	PersonFutureRequirementsErrorCodeInvalidAddressCityStatePostalCode                      PersonFutureRequirementsErrorCode = "invalid_address_city_state_postal_code"
	PersonFutureRequirementsErrorCodeInvalidStreetAddress                                   PersonFutureRequirementsErrorCode = "invalid_street_address"
	PersonFutureRequirementsErrorCodeInvalidValueOther                                      PersonFutureRequirementsErrorCode = "invalid_value_other"
	PersonFutureRequirementsErrorCodeVerificationDocumentAddressMismatch                    PersonFutureRequirementsErrorCode = "verification_document_address_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentAddressMissing                     PersonFutureRequirementsErrorCode = "verification_document_address_missing"
	PersonFutureRequirementsErrorCodeVerificationDocumentCorrupt                            PersonFutureRequirementsErrorCode = "verification_document_corrupt"
	PersonFutureRequirementsErrorCodeVerificationDocumentCountryNotSupported                PersonFutureRequirementsErrorCode = "verification_document_country_not_supported"
	PersonFutureRequirementsErrorCodeVerificationDocumentDOBMismatch                        PersonFutureRequirementsErrorCode = "verification_document_dob_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentDuplicateType                      PersonFutureRequirementsErrorCode = "verification_document_duplicate_type"
	PersonFutureRequirementsErrorCodeVerificationDocumentExpired                            PersonFutureRequirementsErrorCode = "verification_document_expired"
	PersonFutureRequirementsErrorCodeVerificationDocumentFailedCopy                         PersonFutureRequirementsErrorCode = "verification_document_failed_copy"
	PersonFutureRequirementsErrorCodeVerificationDocumentFailedGreyscale                    PersonFutureRequirementsErrorCode = "verification_document_failed_greyscale"
	PersonFutureRequirementsErrorCodeVerificationDocumentFailedOther                        PersonFutureRequirementsErrorCode = "verification_document_failed_other"
	PersonFutureRequirementsErrorCodeVerificationDocumentFailedTestMode                     PersonFutureRequirementsErrorCode = "verification_document_failed_test_mode"
	PersonFutureRequirementsErrorCodeVerificationDocumentFraudulent                         PersonFutureRequirementsErrorCode = "verification_document_fraudulent"
	PersonFutureRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   PersonFutureRequirementsErrorCode = "verification_document_id_number_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentIDNumberMissing                    PersonFutureRequirementsErrorCode = "verification_document_id_number_missing"
	PersonFutureRequirementsErrorCodeVerificationDocumentIncomplete                         PersonFutureRequirementsErrorCode = "verification_document_incomplete"
	PersonFutureRequirementsErrorCodeVerificationDocumentInvalid                            PersonFutureRequirementsErrorCode = "verification_document_invalid"
	PersonFutureRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           PersonFutureRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	PersonFutureRequirementsErrorCodeVerificationDocumentManipulated                        PersonFutureRequirementsErrorCode = "verification_document_manipulated"
	PersonFutureRequirementsErrorCodeVerificationDocumentMissingBack                        PersonFutureRequirementsErrorCode = "verification_document_missing_back"
	PersonFutureRequirementsErrorCodeVerificationDocumentMissingFront                       PersonFutureRequirementsErrorCode = "verification_document_missing_front"
	PersonFutureRequirementsErrorCodeVerificationDocumentNameMismatch                       PersonFutureRequirementsErrorCode = "verification_document_name_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentNameMissing                        PersonFutureRequirementsErrorCode = "verification_document_name_missing"
	PersonFutureRequirementsErrorCodeVerificationDocumentNationalityMismatch                PersonFutureRequirementsErrorCode = "verification_document_nationality_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentNotReadable                        PersonFutureRequirementsErrorCode = "verification_document_not_readable"
	PersonFutureRequirementsErrorCodeVerificationDocumentNotSigned                          PersonFutureRequirementsErrorCode = "verification_document_not_signed"
	PersonFutureRequirementsErrorCodeVerificationDocumentNotUploaded                        PersonFutureRequirementsErrorCode = "verification_document_not_uploaded"
	PersonFutureRequirementsErrorCodeVerificationDocumentPhotoMismatch                      PersonFutureRequirementsErrorCode = "verification_document_photo_mismatch"
	PersonFutureRequirementsErrorCodeVerificationDocumentTooLarge                           PersonFutureRequirementsErrorCode = "verification_document_too_large"
	PersonFutureRequirementsErrorCodeVerificationDocumentTypeNotSupported                   PersonFutureRequirementsErrorCode = "verification_document_type_not_supported"
	PersonFutureRequirementsErrorCodeVerificationFailedAddressMatch                         PersonFutureRequirementsErrorCode = "verification_failed_address_match"
	PersonFutureRequirementsErrorCodeVerificationFailedBusinessIecNumber                    PersonFutureRequirementsErrorCode = "verification_failed_business_iec_number"
	PersonFutureRequirementsErrorCodeVerificationFailedDocumentMatch                        PersonFutureRequirementsErrorCode = "verification_failed_document_match"
	PersonFutureRequirementsErrorCodeVerificationFailedIDNumberMatch                        PersonFutureRequirementsErrorCode = "verification_failed_id_number_match"
	PersonFutureRequirementsErrorCodeVerificationFailedKeyedIdentity                        PersonFutureRequirementsErrorCode = "verification_failed_keyed_identity"
	PersonFutureRequirementsErrorCodeVerificationFailedKeyedMatch                           PersonFutureRequirementsErrorCode = "verification_failed_keyed_match"
	PersonFutureRequirementsErrorCodeVerificationFailedNameMatch                            PersonFutureRequirementsErrorCode = "verification_failed_name_match"
	PersonFutureRequirementsErrorCodeVerificationFailedOther                                PersonFutureRequirementsErrorCode = "verification_failed_other"
	PersonFutureRequirementsErrorCodeVerificationFailedTaxIDMatch                           PersonFutureRequirementsErrorCode = "verification_failed_tax_id_match"
	PersonFutureRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       PersonFutureRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	PersonFutureRequirementsErrorCodeVerificationMissingExecutives                          PersonFutureRequirementsErrorCode = "verification_missing_executives"
	PersonFutureRequirementsErrorCodeVerificationMissingOwners                              PersonFutureRequirementsErrorCode = "verification_missing_owners"
	PersonFutureRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations PersonFutureRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// PersonPoliticalExposure describes the political exposure of a given person.
type PersonPoliticalExposure string

// List of values that IdentityVerificationStatus can take.
const (
	PersonPoliticalExposureExisting PersonPoliticalExposure = "existing"
	PersonPoliticalExposureNone     PersonPoliticalExposure = "none"
)

// The code for the type of error.
type PersonRequirementsErrorCode string

// List of values that PersonRequirementsErrorCode can take
const (
	PersonRequirementsErrorCodeInvalidAddressCityStatePostalCode                      PersonRequirementsErrorCode = "invalid_address_city_state_postal_code"
	PersonRequirementsErrorCodeInvalidStreetAddress                                   PersonRequirementsErrorCode = "invalid_street_address"
	PersonRequirementsErrorCodeInvalidValueOther                                      PersonRequirementsErrorCode = "invalid_value_other"
	PersonRequirementsErrorCodeVerificationDocumentAddressMismatch                    PersonRequirementsErrorCode = "verification_document_address_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentAddressMissing                     PersonRequirementsErrorCode = "verification_document_address_missing"
	PersonRequirementsErrorCodeVerificationDocumentCorrupt                            PersonRequirementsErrorCode = "verification_document_corrupt"
	PersonRequirementsErrorCodeVerificationDocumentCountryNotSupported                PersonRequirementsErrorCode = "verification_document_country_not_supported"
	PersonRequirementsErrorCodeVerificationDocumentDOBMismatch                        PersonRequirementsErrorCode = "verification_document_dob_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentDuplicateType                      PersonRequirementsErrorCode = "verification_document_duplicate_type"
	PersonRequirementsErrorCodeVerificationDocumentExpired                            PersonRequirementsErrorCode = "verification_document_expired"
	PersonRequirementsErrorCodeVerificationDocumentFailedCopy                         PersonRequirementsErrorCode = "verification_document_failed_copy"
	PersonRequirementsErrorCodeVerificationDocumentFailedGreyscale                    PersonRequirementsErrorCode = "verification_document_failed_greyscale"
	PersonRequirementsErrorCodeVerificationDocumentFailedOther                        PersonRequirementsErrorCode = "verification_document_failed_other"
	PersonRequirementsErrorCodeVerificationDocumentFailedTestMode                     PersonRequirementsErrorCode = "verification_document_failed_test_mode"
	PersonRequirementsErrorCodeVerificationDocumentFraudulent                         PersonRequirementsErrorCode = "verification_document_fraudulent"
	PersonRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   PersonRequirementsErrorCode = "verification_document_id_number_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentIDNumberMissing                    PersonRequirementsErrorCode = "verification_document_id_number_missing"
	PersonRequirementsErrorCodeVerificationDocumentIncomplete                         PersonRequirementsErrorCode = "verification_document_incomplete"
	PersonRequirementsErrorCodeVerificationDocumentInvalid                            PersonRequirementsErrorCode = "verification_document_invalid"
	PersonRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           PersonRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	PersonRequirementsErrorCodeVerificationDocumentManipulated                        PersonRequirementsErrorCode = "verification_document_manipulated"
	PersonRequirementsErrorCodeVerificationDocumentMissingBack                        PersonRequirementsErrorCode = "verification_document_missing_back"
	PersonRequirementsErrorCodeVerificationDocumentMissingFront                       PersonRequirementsErrorCode = "verification_document_missing_front"
	PersonRequirementsErrorCodeVerificationDocumentNameMismatch                       PersonRequirementsErrorCode = "verification_document_name_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentNameMissing                        PersonRequirementsErrorCode = "verification_document_name_missing"
	PersonRequirementsErrorCodeVerificationDocumentNationalityMismatch                PersonRequirementsErrorCode = "verification_document_nationality_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentNotReadable                        PersonRequirementsErrorCode = "verification_document_not_readable"
	PersonRequirementsErrorCodeVerificationDocumentNotSigned                          PersonRequirementsErrorCode = "verification_document_not_signed"
	PersonRequirementsErrorCodeVerificationDocumentNotUploaded                        PersonRequirementsErrorCode = "verification_document_not_uploaded"
	PersonRequirementsErrorCodeVerificationDocumentPhotoMismatch                      PersonRequirementsErrorCode = "verification_document_photo_mismatch"
	PersonRequirementsErrorCodeVerificationDocumentTooLarge                           PersonRequirementsErrorCode = "verification_document_too_large"
	PersonRequirementsErrorCodeVerificationDocumentTypeNotSupported                   PersonRequirementsErrorCode = "verification_document_type_not_supported"
	PersonRequirementsErrorCodeVerificationFailedAddressMatch                         PersonRequirementsErrorCode = "verification_failed_address_match"
	PersonRequirementsErrorCodeVerificationFailedBusinessIecNumber                    PersonRequirementsErrorCode = "verification_failed_business_iec_number"
	PersonRequirementsErrorCodeVerificationFailedDocumentMatch                        PersonRequirementsErrorCode = "verification_failed_document_match"
	PersonRequirementsErrorCodeVerificationFailedIDNumberMatch                        PersonRequirementsErrorCode = "verification_failed_id_number_match"
	PersonRequirementsErrorCodeVerificationFailedKeyedIdentity                        PersonRequirementsErrorCode = "verification_failed_keyed_identity"
	PersonRequirementsErrorCodeVerificationFailedKeyedMatch                           PersonRequirementsErrorCode = "verification_failed_keyed_match"
	PersonRequirementsErrorCodeVerificationFailedNameMatch                            PersonRequirementsErrorCode = "verification_failed_name_match"
	PersonRequirementsErrorCodeVerificationFailedOther                                PersonRequirementsErrorCode = "verification_failed_other"
	PersonRequirementsErrorCodeVerificationFailedTaxIDMatch                           PersonRequirementsErrorCode = "verification_failed_tax_id_match"
	PersonRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       PersonRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	PersonRequirementsErrorCodeVerificationMissingExecutives                          PersonRequirementsErrorCode = "verification_missing_executives"
	PersonRequirementsErrorCodeVerificationMissingOwners                              PersonRequirementsErrorCode = "verification_missing_owners"
	PersonRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations PersonRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// The Kana variation of the person's address (Japan only).
type PersonAddressKanaParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// The Kanji variation of the person's address (Japan only).
type PersonAddressKanjiParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	Town       *string `form:"town"`
}

// The person's date of birth.
type PersonDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type PersonDocumentsCompanyAuthorizationParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's passport page with photo and personal data.
type PersonDocumentsPassportParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type PersonDocumentsVisaParams struct {
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type PersonDocumentsParams struct {
	CompanyAuthorization *PersonDocumentsCompanyAuthorizationParams `form:"company_authorization"`
	Passport             *PersonDocumentsPassportParams             `form:"passport"`
	Visa                 *PersonDocumentsVisaParams                 `form:"visa"`
}

// The relationship that this person has with the account's legal entity.
type PersonRelationshipParams struct {
	Director         *bool    `form:"director"`
	Executive        *bool    `form:"executive"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationAdditionalDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// PersonVerificationDocumentParams represents the parameters available for the document verifying
// a person's identity.
type PersonVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// PersonVerificationParams is used to represent parameters associated with a person's verification
// details.
type PersonVerificationParams struct {
	AdditionalDocument *PersonVerificationAdditionalDocumentParams `form:"additional_document"`
	Document           *PersonVerificationDocumentParams           `form:"document"`
}

// PersonParams is the set of parameters that can be used when creating or updating a person.
// For more details see https://stripe.com/docs/api#create_person.
type PersonParams struct {
	Params            `form:"*"`
	Account           *string                   `form:"-"` // Included in URL
	Address           *AddressParams            `form:"address"`
	AddressKana       *PersonAddressKanaParams  `form:"address_kana"`
	AddressKanji      *PersonAddressKanjiParams `form:"address_kanji"`
	DOB               *PersonDOBParams          `form:"dob"`
	Documents         *PersonDocumentsParams    `form:"documents"`
	Email             *string                   `form:"email"`
	FirstName         *string                   `form:"first_name"`
	FirstNameKana     *string                   `form:"first_name_kana"`
	FirstNameKanji    *string                   `form:"first_name_kanji"`
	Gender            *string                   `form:"gender"`
	IDNumber          *string                   `form:"id_number"`
	LastName          *string                   `form:"last_name"`
	LastNameKana      *string                   `form:"last_name_kana"`
	LastNameKanji     *string                   `form:"last_name_kanji"`
	MaidenName        *string                   `form:"maiden_name"`
	Nationality       *string                   `form:"nationality"`
	PersonToken       *string                   `form:"person_token"`
	Phone             *string                   `form:"phone"`
	PoliticalExposure *string                   `form:"political_exposure"`
	Relationship      *PersonRelationshipParams `form:"relationship"`
	SsnLast4          *string                   `form:"ssn_last_4"`
	Verification      *PersonVerificationParams `form:"verification"`
}

// Filters on the list of people returned based on the person's relationship to the account's company.
type PersonListRelationshipParams struct {
	Director       *bool `form:"director"`
	Executive      *bool `form:"executive"`
	Owner          *bool `form:"owner"`
	Representative *bool `form:"representative"`
}

// PersonListParams is the set of parameters that can be used when listing persons.
// For more detail see https://stripe.com/docs/api#list_persons.
type PersonListParams struct {
	ListParams   `form:"*"`
	Account      *string                       `form:"-"` // Included in URL
	Relationship *PersonListRelationshipParams `form:"relationship"`
}

// The Kana variation of the person's address (Japan only).
type PersonAddressKana struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Town       string `json:"town"`
}

// The Kanji variation of the person's address (Japan only).
type PersonAddressKanji struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	Town       string `json:"town"`
}
type PersonDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PersonFutureRequirementsError struct {
	Code        PersonFutureRequirementsErrorCode `json:"code"`
	Reason      string                            `json:"reason"`
	Requirement string                            `json:"requirement"`
}

// Information about the upcoming new requirements for this person, including what information needs to be collected, and by when.
type PersonFutureRequirements struct {
	Alternatives        []*PersonFutureRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                               `json:"currently_due"`
	Errors              []*PersonFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                               `json:"eventually_due"`
	PastDue             []string                               `json:"past_due"`
	PendingVerification []string                               `json:"pending_verification"`
}
type PersonRelationship struct {
	Director         bool    `json:"director"`
	Executive        bool    `json:"executive"`
	Owner            bool    `json:"owner"`
	PercentOwnership float64 `json:"percent_ownership"`
	Representative   bool    `json:"representative"`
	Title            string  `json:"title"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PersonRequirementsError struct {
	Code        PersonRequirementsErrorCode `json:"code"`
	Reason      string                      `json:"reason"`
	Requirement string                      `json:"requirement"`
}

// Information about the requirements for this person, including what information needs to be collected, and by when.
type PersonRequirements struct {
	Alternatives        []*PersonRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                         `json:"currently_due"`
	Errors              []*PersonRequirementsError       `json:"errors"`
	EventuallyDue       []string                         `json:"eventually_due"`
	PastDue             []string                         `json:"past_due"`
	PendingVerification []string                         `json:"pending_verification"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationAdditionalDocument struct {
	Back        *File  `json:"back"`
	Details     string `json:"details"`
	DetailsCode string `json:"details_code"`
	Front       *File  `json:"front"`
}

// PersonVerificationDocument represents the documents associated with a Person.
type PersonVerificationDocument struct {
	Back        *File  `json:"back"`
	Details     string `json:"details"`
	DetailsCode string `json:"details_code"`
	Front       *File  `json:"front"`
}

// PersonVerification is the structure for a person's verification details.
type PersonVerification struct {
	AdditionalDocument *PersonVerificationAdditionalDocument `json:"additional_document"`
	Details            string                                `json:"details"`
	DetailsCode        string                                `json:"details_code"`
	Document           *PersonVerificationDocument           `json:"document"`
	Status             string                                `json:"status"`
}

// Person is the resource representing a Stripe person.
// For more details see https://stripe.com/docs/api#persons.
type Person struct {
	APIResource
	Account            string                    `json:"account"`
	Address            *Address                  `json:"address"`
	AddressKana        *PersonAddressKana        `json:"address_kana"`
	AddressKanji       *PersonAddressKanji       `json:"address_kanji"`
	Created            int64                     `json:"created"`
	Deleted            bool                      `json:"deleted"`
	DOB                *PersonDOB                `json:"dob"`
	Email              string                    `json:"email"`
	FirstName          string                    `json:"first_name"`
	FirstNameKana      string                    `json:"first_name_kana"`
	FirstNameKanji     string                    `json:"first_name_kanji"`
	FutureRequirements *PersonFutureRequirements `json:"future_requirements"`
	Gender             string                    `json:"gender"`
	ID                 string                    `json:"id"`
	IDNumberProvided   bool                      `json:"id_number_provided"`
	LastName           string                    `json:"last_name"`
	LastNameKana       string                    `json:"last_name_kana"`
	LastNameKanji      string                    `json:"last_name_kanji"`
	MaidenName         string                    `json:"maiden_name"`
	Metadata           map[string]string         `json:"metadata"`
	Nationality        string                    `json:"nationality"`
	Object             string                    `json:"object"`
	Phone              string                    `json:"phone"`
	PoliticalExposure  PersonPoliticalExposure   `json:"political_exposure"`
	Relationship       *PersonRelationship       `json:"relationship"`
	Requirements       *PersonRequirements       `json:"requirements"`
	SsnLast4Provided   bool                      `json:"ssn_last_4_provided"`
	Verification       *PersonVerification       `json:"verification"`
}

// PersonList is a list of persons as retrieved from a list endpoint.
type PersonList struct {
	APIResource
	ListMeta
	Data []*Person `json:"data"`
}
