//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves the details of an existing Report Run.
type ReportingReportRunParams struct {
	Params     `form:"*"`
	Parameters *ReportingReportRunParametersParams `form:"parameters"`
	ReportType *string                             `form:"report_type"`
}

// Parameters specifying how the report should be run. Different Report Types have different required and optional parameters, listed in the [API Access to Reports](https://stripe.com/docs/reporting/statements/api) documentation.
type ReportingReportRunParametersParams struct {
	Columns           []*string `form:"columns"`
	ConnectedAccount  *string   `form:"connected_account"`
	Currency          *string   `form:"currency"`
	IntervalEnd       *int64    `form:"interval_end"`
	IntervalStart     *int64    `form:"interval_start"`
	Payout            *string   `form:"payout"`
	ReportingCategory *string   `form:"reporting_category"`
	Timezone          *string   `form:"timezone"`
}

// Returns a list of Report Runs, with the most recent appearing first.
type ReportingReportRunListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}
type ReportingReportRunParameters struct {
	Columns           []string `json:"columns"`
	ConnectedAccount  string   `json:"connected_account"`
	Currency          Currency `json:"currency"`
	IntervalEnd       int64    `json:"interval_end"`
	IntervalStart     int64    `json:"interval_start"`
	Payout            string   `json:"payout"`
	ReportingCategory string   `json:"reporting_category"`
	Timezone          string   `json:"timezone"`
}

// The Report Run object represents an instance of a report type generated with
// specific run parameters. Once the object is created, Stripe begins processing the report.
// When the report has finished running, it will give you a reference to a file
// where you can retrieve your results. For an overview, see
// [API Access to Reports](https://stripe.com/docs/reporting/statements/api).
//
// Note that certain report types can only be run based on your live-mode data (not test-mode
// data), and will error when queried without a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).
type ReportingReportRun struct {
	APIResource
	Created     int64                         `json:"created"`
	Error       string                        `json:"error"`
	ID          string                        `json:"id"`
	Livemode    bool                          `json:"livemode"`
	Object      string                        `json:"object"`
	Parameters  *ReportingReportRunParameters `json:"parameters"`
	ReportType  string                        `json:"report_type"`
	Result      *File                         `json:"result"`
	Status      string                        `json:"status"`
	SucceededAt int64                         `json:"succeeded_at"`
}

// ReportingReportRunList is a list of ReportRuns as retrieved from a list endpoint.
type ReportingReportRunList struct {
	APIResource
	ListMeta
	Data []*ReportingReportRun `json:"data"`
}
