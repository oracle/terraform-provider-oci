// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaDownloadRecordsRequest wrapper for the ListJavaDownloadRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadRecords.go.html to see an example of how to use ListJavaDownloadRecordsRequest.
type ListJavaDownloadRecordsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Java family version identifier.
	FamilyVersion *string `mandatory:"false" contributesTo:"query" name:"familyVersion"`

	// Unique Java release version identifier.
	ReleaseVersion *string `mandatory:"false" contributesTo:"query" name:"releaseVersion"`

	// Target Operating System family of the artifact.
	OsFamily *string `mandatory:"false" contributesTo:"query" name:"osFamily"`

	// Target Operating System architecture of the artifact.
	Architecture *string `mandatory:"false" contributesTo:"query" name:"architecture"`

	// Packaging type detail of the artifact.
	PackageTypeDetail *string `mandatory:"false" contributesTo:"query" name:"packageTypeDetail"`

	// The start of the time period for which reports are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period for which reports are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaDownloadRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If no value is specified _timeDownloaded_ is default.
	SortBy ListJavaDownloadRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaDownloadRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaDownloadRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaDownloadRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaDownloadRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaDownloadRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaDownloadRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaDownloadRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaDownloadRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaDownloadRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaDownloadRecordsResponse wrapper for the ListJavaDownloadRecords operation
type ListJavaDownloadRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaDownloadRecordCollection instances
	JavaDownloadRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaDownloadRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaDownloadRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaDownloadRecordsSortOrderEnum Enum with underlying type: string
type ListJavaDownloadRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListJavaDownloadRecordsSortOrderEnum
const (
	ListJavaDownloadRecordsSortOrderAsc  ListJavaDownloadRecordsSortOrderEnum = "ASC"
	ListJavaDownloadRecordsSortOrderDesc ListJavaDownloadRecordsSortOrderEnum = "DESC"
)

var mappingListJavaDownloadRecordsSortOrderEnum = map[string]ListJavaDownloadRecordsSortOrderEnum{
	"ASC":  ListJavaDownloadRecordsSortOrderAsc,
	"DESC": ListJavaDownloadRecordsSortOrderDesc,
}

var mappingListJavaDownloadRecordsSortOrderEnumLowerCase = map[string]ListJavaDownloadRecordsSortOrderEnum{
	"asc":  ListJavaDownloadRecordsSortOrderAsc,
	"desc": ListJavaDownloadRecordsSortOrderDesc,
}

// GetListJavaDownloadRecordsSortOrderEnumValues Enumerates the set of values for ListJavaDownloadRecordsSortOrderEnum
func GetListJavaDownloadRecordsSortOrderEnumValues() []ListJavaDownloadRecordsSortOrderEnum {
	values := make([]ListJavaDownloadRecordsSortOrderEnum, 0)
	for _, v := range mappingListJavaDownloadRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListJavaDownloadRecordsSortOrderEnum
func GetListJavaDownloadRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaDownloadRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadRecordsSortOrderEnum(val string) (ListJavaDownloadRecordsSortOrderEnum, bool) {
	enum, ok := mappingListJavaDownloadRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaDownloadRecordsSortByEnum Enum with underlying type: string
type ListJavaDownloadRecordsSortByEnum string

// Set of constants representing the allowable values for ListJavaDownloadRecordsSortByEnum
const (
	ListJavaDownloadRecordsSortByTimedownloaded   ListJavaDownloadRecordsSortByEnum = "timeDownloaded"
	ListJavaDownloadRecordsSortByDownloadsourceid ListJavaDownloadRecordsSortByEnum = "downloadSourceId"
	ListJavaDownloadRecordsSortByDownloadtype     ListJavaDownloadRecordsSortByEnum = "downloadType"
)

var mappingListJavaDownloadRecordsSortByEnum = map[string]ListJavaDownloadRecordsSortByEnum{
	"timeDownloaded":   ListJavaDownloadRecordsSortByTimedownloaded,
	"downloadSourceId": ListJavaDownloadRecordsSortByDownloadsourceid,
	"downloadType":     ListJavaDownloadRecordsSortByDownloadtype,
}

var mappingListJavaDownloadRecordsSortByEnumLowerCase = map[string]ListJavaDownloadRecordsSortByEnum{
	"timedownloaded":   ListJavaDownloadRecordsSortByTimedownloaded,
	"downloadsourceid": ListJavaDownloadRecordsSortByDownloadsourceid,
	"downloadtype":     ListJavaDownloadRecordsSortByDownloadtype,
}

// GetListJavaDownloadRecordsSortByEnumValues Enumerates the set of values for ListJavaDownloadRecordsSortByEnum
func GetListJavaDownloadRecordsSortByEnumValues() []ListJavaDownloadRecordsSortByEnum {
	values := make([]ListJavaDownloadRecordsSortByEnum, 0)
	for _, v := range mappingListJavaDownloadRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadRecordsSortByEnumStringValues Enumerates the set of values in String for ListJavaDownloadRecordsSortByEnum
func GetListJavaDownloadRecordsSortByEnumStringValues() []string {
	return []string{
		"timeDownloaded",
		"downloadSourceId",
		"downloadType",
	}
}

// GetMappingListJavaDownloadRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadRecordsSortByEnum(val string) (ListJavaDownloadRecordsSortByEnum, bool) {
	enum, ok := mappingListJavaDownloadRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
