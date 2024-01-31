// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAwrDbWaitEventBucketsRequest wrapper for the SummarizeAwrDbWaitEventBuckets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbWaitEventBuckets.go.html to see an example of how to use SummarizeAwrDbWaitEventBucketsRequest.
type SummarizeAwrDbWaitEventBucketsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The required single value query parameter to filter the entity name.
	Name *string `mandatory:"true" contributesTo:"query" name:"name"`

	// The optional single value query parameter to filter the database instance number.
	InstNum *string `mandatory:"false" contributesTo:"query" name:"instNum"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The number of buckets within the histogram.
	NumBucket *int `mandatory:"false" contributesTo:"query" name:"numBucket"`

	// The minimum value of the histogram.
	MinValue *float64 `mandatory:"false" contributesTo:"query" name:"minValue"`

	// The maximum value of the histogram.
	MaxValue *float64 `mandatory:"false" contributesTo:"query" name:"maxValue"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort distribution data.
	SortBy SummarizeAwrDbWaitEventBucketsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder SummarizeAwrDbWaitEventBucketsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDbWaitEventBucketsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbWaitEventBucketsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbWaitEventBucketsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbWaitEventBucketsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDbWaitEventBucketsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDbWaitEventBucketsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDbWaitEventBucketsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDbWaitEventBucketsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDbWaitEventBucketsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDbWaitEventBucketsResponse wrapper for the SummarizeAwrDbWaitEventBuckets operation
type SummarizeAwrDbWaitEventBucketsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbWaitEventBucketCollection instances
	AwrDbWaitEventBucketCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbWaitEventBucketsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbWaitEventBucketsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbWaitEventBucketsSortByEnum Enum with underlying type: string
type SummarizeAwrDbWaitEventBucketsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbWaitEventBucketsSortByEnum
const (
	SummarizeAwrDbWaitEventBucketsSortByCategory   SummarizeAwrDbWaitEventBucketsSortByEnum = "CATEGORY"
	SummarizeAwrDbWaitEventBucketsSortByPercentage SummarizeAwrDbWaitEventBucketsSortByEnum = "PERCENTAGE"
)

var mappingSummarizeAwrDbWaitEventBucketsSortByEnum = map[string]SummarizeAwrDbWaitEventBucketsSortByEnum{
	"CATEGORY":   SummarizeAwrDbWaitEventBucketsSortByCategory,
	"PERCENTAGE": SummarizeAwrDbWaitEventBucketsSortByPercentage,
}

var mappingSummarizeAwrDbWaitEventBucketsSortByEnumLowerCase = map[string]SummarizeAwrDbWaitEventBucketsSortByEnum{
	"category":   SummarizeAwrDbWaitEventBucketsSortByCategory,
	"percentage": SummarizeAwrDbWaitEventBucketsSortByPercentage,
}

// GetSummarizeAwrDbWaitEventBucketsSortByEnumValues Enumerates the set of values for SummarizeAwrDbWaitEventBucketsSortByEnum
func GetSummarizeAwrDbWaitEventBucketsSortByEnumValues() []SummarizeAwrDbWaitEventBucketsSortByEnum {
	values := make([]SummarizeAwrDbWaitEventBucketsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbWaitEventBucketsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbWaitEventBucketsSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDbWaitEventBucketsSortByEnum
func GetSummarizeAwrDbWaitEventBucketsSortByEnumStringValues() []string {
	return []string{
		"CATEGORY",
		"PERCENTAGE",
	}
}

// GetMappingSummarizeAwrDbWaitEventBucketsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbWaitEventBucketsSortByEnum(val string) (SummarizeAwrDbWaitEventBucketsSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDbWaitEventBucketsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDbWaitEventBucketsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbWaitEventBucketsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbWaitEventBucketsSortOrderEnum
const (
	SummarizeAwrDbWaitEventBucketsSortOrderAsc  SummarizeAwrDbWaitEventBucketsSortOrderEnum = "ASC"
	SummarizeAwrDbWaitEventBucketsSortOrderDesc SummarizeAwrDbWaitEventBucketsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbWaitEventBucketsSortOrderEnum = map[string]SummarizeAwrDbWaitEventBucketsSortOrderEnum{
	"ASC":  SummarizeAwrDbWaitEventBucketsSortOrderAsc,
	"DESC": SummarizeAwrDbWaitEventBucketsSortOrderDesc,
}

var mappingSummarizeAwrDbWaitEventBucketsSortOrderEnumLowerCase = map[string]SummarizeAwrDbWaitEventBucketsSortOrderEnum{
	"asc":  SummarizeAwrDbWaitEventBucketsSortOrderAsc,
	"desc": SummarizeAwrDbWaitEventBucketsSortOrderDesc,
}

// GetSummarizeAwrDbWaitEventBucketsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbWaitEventBucketsSortOrderEnum
func GetSummarizeAwrDbWaitEventBucketsSortOrderEnumValues() []SummarizeAwrDbWaitEventBucketsSortOrderEnum {
	values := make([]SummarizeAwrDbWaitEventBucketsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbWaitEventBucketsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDbWaitEventBucketsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDbWaitEventBucketsSortOrderEnum
func GetSummarizeAwrDbWaitEventBucketsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDbWaitEventBucketsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDbWaitEventBucketsSortOrderEnum(val string) (SummarizeAwrDbWaitEventBucketsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDbWaitEventBucketsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
