//
//
// File generated from our OpenAPI spec
//
//

package stripe

type CapabilityRequirementsErrorCode string

const (
	CapabilityRequirementsErrorCodeInvalidAddressCityStatePostalCode                      CapabilityRequirementsErrorCode = "invalid_address_city_state_postal_code"
	CapabilityRequirementsErrorCodeInvalidStreetAddress                                   CapabilityRequirementsErrorCode = "invalid_street_address"
	CapabilityRequirementsErrorCodeInvalidValueOther                                      CapabilityRequirementsErrorCode = "invalid_value_other"
	CapabilityRequirementsErrorCodeVerificationDocumentAddressMismatch                    CapabilityRequirementsErrorCode = "verification_document_address_mismatch"
	CapabilityRequirementsErrorCodeVerificationDocumentAddressMissing                     CapabilityRequirementsErrorCode = "verification_document_address_missing"
	CapabilityRequirementsErrorCodeVerificationDocumentCorrupt                            CapabilityRequirementsErrorCode = "verification_document_corrupt"
	CapabilityRequirementsErrorCodeVerificationDocumentCountryNotSupported                CapabilityRequirementsErrorCode = "verification_document_country_not_supported"
	CapabilityRequirementsErrorCodeVerificationDocumentDobMismatch                        CapabilityRequirementsErrorCode = "verification_document_dob_mismatch"
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
type CapabilityRequirementsError struct {
	Code        CapabilityRequirementsErrorCode `json:"code"`
	Reason      string                          `json:"reason"`
	Requirement string                          `json:"requirement"`
}

// CapabilityRequirements represents information that needs to be collected for a capability.
type CapabilityRequirements struct {
	CurrentDeadline     int64                          `json:"current_deadline"`
	CurrentlyDue        []string                       `json:"currently_due"`
	DisabledReason      string                         `json:"disabled_reason"`
	Errors              []*CapabilityRequirementsError `json:"errors"`
	EventuallyDue       []string                       `json:"eventually_due"`
	PastDue             []string                       `json:"past_due"`
	PendingVerification []string                       `json:"pending_verification"`
}

// Capability is the resource representing a Stripe capability.
// For more details see https://stripe.com/docs/api/capabilities
type Capability struct {
	APIResource
	Account      *Account                `json:"account"`
	ID           string                  `json:"id"`
	Object       string                  `json:"object"`
	Requested    bool                    `json:"requested"`
	RequestedAt  int64                   `json:"requested_at"`
	Requirements *CapabilityRequirements `json:"requirements"`
	Status       CapabilityStatus        `json:"status"`
}

// CapabilityList is a list of capabilities as retrieved from a list endpoint.
type CapabilityList struct {
	APIResource
	ListMeta
	Data []*Capability `json:"data"`
}
