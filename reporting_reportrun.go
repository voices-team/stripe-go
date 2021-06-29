//
//
// File generated from our OpenAPI spec
//
//

package stripe

type ReportingReportRunParams struct {
	Params     `form:"*"`
	Parameters *ReportingReportRunParametersParams `form:"parameters"`
	ReportType *string                             `form:"report_type"`
}
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
type ReportingReportRunList struct {
	APIResource
	ListMeta
	Data []*ReportingReportRun `json:"data"`
}
