//
//
// File generated from our OpenAPI spec
//
//

package stripe

// SigmaScheduledQueryRunParams is the set of parameters that can be used when updating a scheduled query run.
type SigmaScheduledQueryRunParams struct {
	Params `form:"*"`
}

// SigmaScheduledQueryRunListParams is the set of parameters that can be used when listing scheduled query runs.
type SigmaScheduledQueryRunListParams struct {
	ListParams `form:"*"`
}
type SigmaScheduledQueryRunError struct {
	Message string `json:"message"`
}

// SigmaScheduledQueryRun is the resource representing a scheduled query run.
type SigmaScheduledQueryRun struct {
	APIResource
	Created              int64                        `json:"created"`
	DataLoadTime         int64                        `json:"data_load_time"`
	Error                *SigmaScheduledQueryRunError `json:"error"`
	File                 *File                        `json:"file"`
	ID                   string                       `json:"id"`
	Livemode             bool                         `json:"livemode"`
	Object               string                       `json:"object"`
	ResultAvailableUntil int64                        `json:"result_available_until"`
	Sql                  string                       `json:"sql"`
	Status               string                       `json:"status"`
	Title                string                       `json:"title"`
}

// SigmaScheduledQueryRunList is a list of scheduled query runs as retrieved from a list endpoint.
type SigmaScheduledQueryRunList struct {
	APIResource
	ListMeta
	Data []*SigmaScheduledQueryRun `json:"data"`
}
