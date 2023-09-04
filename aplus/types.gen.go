package aplus

import (
	"net/http"
	"time"
)

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional information that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

type ErrorList []Error

type SearchContentDocumentsParams struct {
	// A marketplace identifier. Specifies the marketplace for the item.
	MarketplaceId string `json:"marketplaceId"`

	PageToken *string `json:"pageToken,omitempty"`
}

type SearchContentDocumentsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *SearchContentDocumentsResponse
}

type SearchContentDocumentsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors                 *ErrorList               `json:"errors,omitempty"`
	Warnings               []any                    `json:"warnings,omitempty"`
	NextPageToken          *string                  `json:"nextPageToken,omitempty"`
	ContentMetadataRecords []*ContentMetadataRecord `json:"contentMetadataRecords,omitempty"`
}

type ContentMetadataRecord struct {
	ContentReferenceKey string `json:"contentReferenceKey,omitempty"`
	ContentMetadata     struct {
		Name          string    `json:"name,omitempty"`
		MarketplaceID string    `json:"marketplaceId,omitempty"`
		Status        string    `json:"status,omitempty"`
		BadgeSet      []string  `json:"badgeSet,omitempty"`
		UpdateTime    time.Time `json:"updateTime,omitempty"`
	} `json:"contentMetadata,omitempty"`
}

type SearchContentPublishRecordsParams struct {
	// A marketplace identifier. Specifies the marketplace for the item.
	MarketplaceId string  `json:"marketplaceId"`
	PageToken     *string `json:"pageToken,omitempty"`
}

type SearchContentPublishRecordsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *SearchContentPublishRecordsResponse
}

type SearchContentPublishRecordsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors            *ErrorList           `json:"errors,omitempty"`
	Warnings          []any                `json:"warnings,omitempty"`
	NextPageToken     *string              `json:"nextPageToken,omitempty"`
	PublishRecordList []*PublishRecordItem `json:"publishRecordList,omitempty"`
}

type PublishRecordItem struct {
	MarketplaceID       string `json:"marketplaceId,omitempty"`
	Locale              string `json:"locale,omitempty"`
	Asin                string `json:"asin,omitempty"`
	ContentType         string `json:"contentType,omitempty"`
	ContentSubType      string `json:"contentSubType,omitempty"`
	ContentReferenceKey string `json:"contentReferenceKey,omitempty"`
}
