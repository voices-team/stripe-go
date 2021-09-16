//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// FilePurpose is the purpose of a particular file.
type FilePurpose string

// List of values that FilePurpose can take.
const (
	FilePurposeAccountRequirement               FilePurpose = "account_requirement"
	FilePurposeAdditionalVerification           FilePurpose = "additional_verification"
	FilePurposeBusinessIcon                     FilePurpose = "business_icon"
	FilePurposeBusinessLogo                     FilePurpose = "business_logo"
	FilePurposeCustomerSignature                FilePurpose = "customer_signature"
	FilePurposeDisputeEvidence                  FilePurpose = "dispute_evidence"
	FilePurposeDocumentProviderIdentityDocument FilePurpose = "document_provider_identity_document"
	FilePurposeFinanceReportRun                 FilePurpose = "finance_report_run"
	FilePurposeIdentityDocument                 FilePurpose = "identity_document"
	FilePurposeIdentityDocumentDownloadable     FilePurpose = "identity_document_downloadable"
	FilePurposePciDocument                      FilePurpose = "pci_document"
	FilePurposeSelfie                           FilePurpose = "selfie"
	FilePurposeSigmaScheduledQuery              FilePurpose = "sigma_scheduled_query"
	FilePurposeTaxDocumentUserUpload            FilePurpose = "tax_document_user_upload"
)

// FileParams is the set of parameters that can be used when creating a file.
// For more details see https://stripe.com/docs/api#create_file.
type FileParams struct {
	Params `form:"*"`
}

// FileListParams is the set of parameters that can be used when listing
// files. For more details see https://stripe.com/docs/api#list_files.
type FileListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Purpose      *string           `form:"purpose"`
}

// File is the resource representing a Stripe file.
// For more details see https://stripe.com/docs/api#file_object.
type File struct {
	APIResource
	Created   int64         `json:"created"`
	ExpiresAt int64         `json:"expires_at"`
	Filename  string        `json:"filename"`
	ID        string        `json:"id"`
	Links     *FileLinkList `json:"links"`
	Object    string        `json:"object"`
	Purpose   FilePurpose   `json:"purpose"`
	Size      int64         `json:"size"`
	Title     string        `json:"title"`
	Type      string        `json:"type"`
	URL       string        `json:"url"`
}

// FileList is a list of files as retrieved from a list endpoint.
type FileList struct {
	APIResource
	ListMeta
	Data []*File `json:"data"`
}

// UnmarshalJSON handles deserialization of a File.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *File) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type file File
	var v file
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = File(v)
	return nil
}
