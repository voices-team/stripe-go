//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The code for the type of error.
type CapabilityFutureRequirementsErrorCode string

// List of values that CapabilityFutureRequirementsErrorCode can take
const (
	CapabilityFutureRequirementsErrorCodeInvalidAddressCityStatePostalCode                      CapabilityFutureRequirementsErrorCode = "invalid_address_city_state_postal_code"
	CapabilityFutureRequirementsErrorCodeInvalidStreetAddress                                   CapabilityFutureRequirementsErrorCode = "invalid_street_address"
	CapabilityFutureRequirementsErrorCodeInvalidValueOther                                      CapabilityFutureRequirementsErrorCode = "invalid_value_other"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentAddressMismatch                    CapabilityFutureRequirementsErrorCode = "verification_document_address_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentAddressMissing                     CapabilityFutureRequirementsErrorCode = "verification_document_address_missing"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentCorrupt                            CapabilityFutureRequirementsErrorCode = "verification_document_corrupt"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentCountryNotSupported                CapabilityFutureRequirementsErrorCode = "verification_document_country_not_supported"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentDOBMismatch                        CapabilityFutureRequirementsErrorCode = "verification_document_dob_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentDuplicateType                      CapabilityFutureRequirementsErrorCode = "verification_document_duplicate_type"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentExpired                            CapabilityFutureRequirementsErrorCode = "verification_document_expired"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentFailedCopy                         CapabilityFutureRequirementsErrorCode = "verification_document_failed_copy"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentFailedGreyscale                    CapabilityFutureRequirementsErrorCode = "verification_document_failed_greyscale"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentFailedOther                        CapabilityFutureRequirementsErrorCode = "verification_document_failed_other"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentFailedTestMode                     CapabilityFutureRequirementsErrorCode = "verification_document_failed_test_mode"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentFraudulent                         CapabilityFutureRequirementsErrorCode = "verification_document_fraudulent"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   CapabilityFutureRequirementsErrorCode = "verification_document_id_number_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentIDNumberMissing                    CapabilityFutureRequirementsErrorCode = "verification_document_id_number_missing"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentIncomplete                         CapabilityFutureRequirementsErrorCode = "verification_document_incomplete"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentInvalid                            CapabilityFutureRequirementsErrorCode = "verification_document_invalid"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           CapabilityFutureRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentManipulated                        CapabilityFutureRequirementsErrorCode = "verification_document_manipulated"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentMissingBack                        CapabilityFutureRequirementsErrorCode = "verification_document_missing_back"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentMissingFront                       CapabilityFutureRequirementsErrorCode = "verification_document_missing_front"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNameMismatch                       CapabilityFutureRequirementsErrorCode = "verification_document_name_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNameMissing                        CapabilityFutureRequirementsErrorCode = "verification_document_name_missing"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNationalityMismatch                CapabilityFutureRequirementsErrorCode = "verification_document_nationality_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNotReadable                        CapabilityFutureRequirementsErrorCode = "verification_document_not_readable"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNotSigned                          CapabilityFutureRequirementsErrorCode = "verification_document_not_signed"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentNotUploaded                        CapabilityFutureRequirementsErrorCode = "verification_document_not_uploaded"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentPhotoMismatch                      CapabilityFutureRequirementsErrorCode = "verification_document_photo_mismatch"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentTooLarge                           CapabilityFutureRequirementsErrorCode = "verification_document_too_large"
	CapabilityFutureRequirementsErrorCodeVerificationDocumentTypeNotSupported                   CapabilityFutureRequirementsErrorCode = "verification_document_type_not_supported"
	CapabilityFutureRequirementsErrorCodeVerificationFailedAddressMatch                         CapabilityFutureRequirementsErrorCode = "verification_failed_address_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedBusinessIecNumber                    CapabilityFutureRequirementsErrorCode = "verification_failed_business_iec_number"
	CapabilityFutureRequirementsErrorCodeVerificationFailedDocumentMatch                        CapabilityFutureRequirementsErrorCode = "verification_failed_document_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedIDNumberMatch                        CapabilityFutureRequirementsErrorCode = "verification_failed_id_number_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedKeyedIdentity                        CapabilityFutureRequirementsErrorCode = "verification_failed_keyed_identity"
	CapabilityFutureRequirementsErrorCodeVerificationFailedKeyedMatch                           CapabilityFutureRequirementsErrorCode = "verification_failed_keyed_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedNameMatch                            CapabilityFutureRequirementsErrorCode = "verification_failed_name_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedOther                                CapabilityFutureRequirementsErrorCode = "verification_failed_other"
	CapabilityFutureRequirementsErrorCodeVerificationFailedTaxIDMatch                           CapabilityFutureRequirementsErrorCode = "verification_failed_tax_id_match"
	CapabilityFutureRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       CapabilityFutureRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	CapabilityFutureRequirementsErrorCodeVerificationMissingExecutives                          CapabilityFutureRequirementsErrorCode = "verification_missing_executives"
	CapabilityFutureRequirementsErrorCodeVerificationMissingOwners                              CapabilityFutureRequirementsErrorCode = "verification_missing_owners"
	CapabilityFutureRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations CapabilityFutureRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// The code for the type of error.
type CapabilityRequirementsErrorCode string

// List of values that CapabilityRequirementsErrorCode can take
const (
	CapabilityRequirementsErrorCodeInvalidAddressCityStatePostalCode                      CapabilityRequirementsErrorCode = "invalid_address_city_state_postal_code"
	CapabilityRequirementsErrorCodeInvalidStreetAddress                                   CapabilityRequirementsErrorCode = "invalid_street_address"
	CapabilityRequirementsErrorCodeInvalidValueOther                                      CapabilityRequirementsErrorCode = "invalid_value_other"
	CapabilityRequirementsErrorCodeVerificationDocumentAddressMismatch                    CapabilityRequirementsErrorCode = "verification_document_address_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentAddressMissing                     CapabilityRequirementsErrorCode = "verification_document_address_missing"
	CapabilityRequirementsErrorCodeVerificationDocumentCorrupt                            CapabilityRequirementsErrorCode = "verification_document_corrupt"
	CapabilityRequirementsErrorCodeVerificationDocumentCountryNotSupported                CapabilityRequirementsErrorCode = "verification_document_country_not_supported"
	CapabilityRequirementsErrorCodeVerificationDocumentDOBMismatch                        CapabilityRequirementsErrorCode = "verification_document_dob_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentDuplicateType                      CapabilityRequirementsErrorCode = "verification_document_duplicate_type"
	CapabilityRequirementsErrorCodeVerificationDocumentExpired                            CapabilityRequirementsErrorCode = "verification_document_expired"
	CapabilityRequirementsErrorCodeVerificationDocumentFailedCopy                         CapabilityRequirementsErrorCode = "verification_document_failed_copy"
	CapabilityRequirementsErrorCodeVerificationDocumentFailedGreyscale                    CapabilityRequirementsErrorCode = "verification_document_failed_greyscale"
	CapabilityRequirementsErrorCodeVerificationDocumentFailedOther                        CapabilityRequirementsErrorCode = "verification_document_failed_other"
	CapabilityRequirementsErrorCodeVerificationDocumentFailedTestMode                     CapabilityRequirementsErrorCode = "verification_document_failed_test_mode"
	CapabilityRequirementsErrorCodeVerificationDocumentFraudulent                         CapabilityRequirementsErrorCode = "verification_document_fraudulent"
	CapabilityRequirementsErrorCodeVerificationDocumentIDNumberMismatch                   CapabilityRequirementsErrorCode = "verification_document_id_number_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentIDNumberMissing                    CapabilityRequirementsErrorCode = "verification_document_id_number_missing"
	CapabilityRequirementsErrorCodeVerificationDocumentIncomplete                         CapabilityRequirementsErrorCode = "verification_document_incomplete"
	CapabilityRequirementsErrorCodeVerificationDocumentInvalid                            CapabilityRequirementsErrorCode = "verification_document_invalid"
	CapabilityRequirementsErrorCodeVerificationDocumentIssueOrExpiryDateMissing           CapabilityRequirementsErrorCode = "verification_document_issue_or_expiry_date_missing"
	CapabilityRequirementsErrorCodeVerificationDocumentManipulated                        CapabilityRequirementsErrorCode = "verification_document_manipulated"
	CapabilityRequirementsErrorCodeVerificationDocumentMissingBack                        CapabilityRequirementsErrorCode = "verification_document_missing_back"
	CapabilityRequirementsErrorCodeVerificationDocumentMissingFront                       CapabilityRequirementsErrorCode = "verification_document_missing_front"
	CapabilityRequirementsErrorCodeVerificationDocumentNameMismatch                       CapabilityRequirementsErrorCode = "verification_document_name_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentNameMissing                        CapabilityRequirementsErrorCode = "verification_document_name_missing"
	CapabilityRequirementsErrorCodeVerificationDocumentNationalityMismatch                CapabilityRequirementsErrorCode = "verification_document_nationality_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentNotReadable                        CapabilityRequirementsErrorCode = "verification_document_not_readable"
	CapabilityRequirementsErrorCodeVerificationDocumentNotSigned                          CapabilityRequirementsErrorCode = "verification_document_not_signed"
	CapabilityRequirementsErrorCodeVerificationDocumentNotUploaded                        CapabilityRequirementsErrorCode = "verification_document_not_uploaded"
	CapabilityRequirementsErrorCodeVerificationDocumentPhotoMismatch                      CapabilityRequirementsErrorCode = "verification_document_photo_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentTooLarge                           CapabilityRequirementsErrorCode = "verification_document_too_large"
	CapabilityRequirementsErrorCodeVerificationDocumentTypeNotSupported                   CapabilityRequirementsErrorCode = "verification_document_type_not_supported"
	CapabilityRequirementsErrorCodeVerificationFailedAddressMatch                         CapabilityRequirementsErrorCode = "verification_failed_address_match"
	CapabilityRequirementsErrorCodeVerificationFailedBusinessIecNumber                    CapabilityRequirementsErrorCode = "verification_failed_business_iec_number"
	CapabilityRequirementsErrorCodeVerificationFailedDocumentMatch                        CapabilityRequirementsErrorCode = "verification_failed_document_match"
	CapabilityRequirementsErrorCodeVerificationFailedIDNumberMatch                        CapabilityRequirementsErrorCode = "verification_failed_id_number_match"
	CapabilityRequirementsErrorCodeVerificationFailedKeyedIdentity                        CapabilityRequirementsErrorCode = "verification_failed_keyed_identity"
	CapabilityRequirementsErrorCodeVerificationFailedKeyedMatch                           CapabilityRequirementsErrorCode = "verification_failed_keyed_match"
	CapabilityRequirementsErrorCodeVerificationFailedNameMatch                            CapabilityRequirementsErrorCode = "verification_failed_name_match"
	CapabilityRequirementsErrorCodeVerificationFailedOther                                CapabilityRequirementsErrorCode = "verification_failed_other"
	CapabilityRequirementsErrorCodeVerificationFailedTaxIDMatch                           CapabilityRequirementsErrorCode = "verification_failed_tax_id_match"
	CapabilityRequirementsErrorCodeVerificationFailedTaxIDNotIssued                       CapabilityRequirementsErrorCode = "verification_failed_tax_id_not_issued"
	CapabilityRequirementsErrorCodeVerificationMissingExecutives                          CapabilityRequirementsErrorCode = "verification_missing_executives"
	CapabilityRequirementsErrorCodeVerificationMissingOwners                              CapabilityRequirementsErrorCode = "verification_missing_owners"
	CapabilityRequirementsErrorCodeVerificationRequiresAdditionalMemorandumOfAssociations CapabilityRequirementsErrorCode = "verification_requires_additional_memorandum_of_associations"
)

// CapabilityStatus describes the different statuses for a capability's status.
type CapabilityStatus string

// List of values that CapabilityStatus can take.
const (
	CapabilityStatusActive      CapabilityStatus = "active"
	CapabilityStatusDisabled    CapabilityStatus = "disabled"
	CapabilityStatusInactive    CapabilityStatus = "inactive"
	CapabilityStatusPending     CapabilityStatus = "pending"
	CapabilityStatusUnrequested CapabilityStatus = "unrequested"
)

// CapabilityParams is the set of parameters that can be used when updating a capability.
// For more details see https://stripe.com/docs/api/capabilities/update
type CapabilityParams struct {
	Params    `form:"*"`
	Account   *string `form:"-"` // Included in URL
	Requested *bool   `form:"requested"`
}

// CapabilityListParams is the set of parameters that can be used when listing capabilities.
// For more detail see https://stripe.com/docs/api/capabilities/list
type CapabilityListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type CapabilityFutureRequirementsError struct {
	Code        CapabilityFutureRequirementsErrorCode `json:"code"`
	Reason      string                                `json:"reason"`
	Requirement string                                `json:"requirement"`
}
type CapabilityFutureRequirements struct {
	Alternatives        []*CapabilityFutureRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                      `json:"current_deadline"`
	CurrentlyDue        []string                                   `json:"currently_due"`
	DisabledReason      string                                     `json:"disabled_reason"`
	Errors              []*CapabilityFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                                   `json:"eventually_due"`
	PastDue             []string                                   `json:"past_due"`
	PendingVerification []string                                   `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type CapabilityRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type CapabilityRequirementsError struct {
	Code        CapabilityRequirementsErrorCode `json:"code"`
	Reason      string                          `json:"reason"`
	Requirement string                          `json:"requirement"`
}

// CapabilityRequirements represents information that needs to be collected for a capability.
type CapabilityRequirements struct {
	Alternatives        []*CapabilityRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                `json:"current_deadline"`
	CurrentlyDue        []string                             `json:"currently_due"`
	DisabledReason      string                               `json:"disabled_reason"`
	Errors              []*CapabilityRequirementsError       `json:"errors"`
	EventuallyDue       []string                             `json:"eventually_due"`
	PastDue             []string                             `json:"past_due"`
	PendingVerification []string                             `json:"pending_verification"`
}

// Capability is the resource representing a Stripe capability.
// For more details see https://stripe.com/docs/api/capabilities
type Capability struct {
	APIResource
	Account            *Account                      `json:"account"`
	FutureRequirements *CapabilityFutureRequirements `json:"future_requirements"`
	ID                 string                        `json:"id"`
	Object             string                        `json:"object"`
	Requested          bool                          `json:"requested"`
	RequestedAt        int64                         `json:"requested_at"`
	Requirements       *CapabilityRequirements       `json:"requirements"`
	Status             CapabilityStatus              `json:"status"`
}

// CapabilityList is a list of capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data"`
}
