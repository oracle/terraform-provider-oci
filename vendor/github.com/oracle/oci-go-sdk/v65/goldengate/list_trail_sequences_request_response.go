// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTrailSequencesRequest wrapper for the ListTrailSequences operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListTrailSequences.go.html to see an example of how to use ListTrailSequencesRequest.
type ListTrailSequencesRequest struct {

	// A unique Deployment identifier.
	DeploymentId *string `mandatory:"true" contributesTo:"query" name:"deploymentId"`

	// A Trail File identifier
	TrailFileId *string `mandatory:"true" contributesTo:"query" name:"trailFileId"`

	// A Trail Sequence identifier
	TrailSequenceId *string `mandatory:"false" contributesTo:"query" name:"trailSequenceId"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeLastUpdated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// displayName is the default.
	SortBy ListTrailSequencesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTrailSequencesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTrailSequencesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTrailSequencesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTrailSequencesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTrailSequencesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTrailSequencesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTrailSequencesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTrailSequencesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTrailSequencesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTrailSequencesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTrailSequencesResponse wrapper for the ListTrailSequences operation
type ListTrailSequencesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TrailSequenceCollection instances
	TrailSequenceCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTrailSequencesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTrailSequencesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTrailSequencesSortByEnum Enum with underlying type: string
type ListTrailSequencesSortByEnum string

// Set of constants representing the allowable values for ListTrailSequencesSortByEnum
const (
	ListTrailSequencesSortByTimelastupdated ListTrailSequencesSortByEnum = "timeLastUpdated"
	ListTrailSequencesSortByDisplayname     ListTrailSequencesSortByEnum = "displayName"
)

var mappingListTrailSequencesSortByEnum = map[string]ListTrailSequencesSortByEnum{
	"timeLastUpdated": ListTrailSequencesSortByTimelastupdated,
	"displayName":     ListTrailSequencesSortByDisplayname,
}

var mappingListTrailSequencesSortByEnumLowerCase = map[string]ListTrailSequencesSortByEnum{
	"timelastupdated": ListTrailSequencesSortByTimelastupdated,
	"displayname":     ListTrailSequencesSortByDisplayname,
}

// GetListTrailSequencesSortByEnumValues Enumerates the set of values for ListTrailSequencesSortByEnum
func GetListTrailSequencesSortByEnumValues() []ListTrailSequencesSortByEnum {
	values := make([]ListTrailSequencesSortByEnum, 0)
	for _, v := range mappingListTrailSequencesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTrailSequencesSortByEnumStringValues Enumerates the set of values in String for ListTrailSequencesSortByEnum
func GetListTrailSequencesSortByEnumStringValues() []string {
	return []string{
		"timeLastUpdated",
		"displayName",
	}
}

// GetMappingListTrailSequencesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTrailSequencesSortByEnum(val string) (ListTrailSequencesSortByEnum, bool) {
	enum, ok := mappingListTrailSequencesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTrailSequencesSortOrderEnum Enum with underlying type: string
type ListTrailSequencesSortOrderEnum string

// Set of constants representing the allowable values for ListTrailSequencesSortOrderEnum
const (
	ListTrailSequencesSortOrderAsc  ListTrailSequencesSortOrderEnum = "ASC"
	ListTrailSequencesSortOrderDesc ListTrailSequencesSortOrderEnum = "DESC"
)

var mappingListTrailSequencesSortOrderEnum = map[string]ListTrailSequencesSortOrderEnum{
	"ASC":  ListTrailSequencesSortOrderAsc,
	"DESC": ListTrailSequencesSortOrderDesc,
}

var mappingListTrailSequencesSortOrderEnumLowerCase = map[string]ListTrailSequencesSortOrderEnum{
	"asc":  ListTrailSequencesSortOrderAsc,
	"desc": ListTrailSequencesSortOrderDesc,
}

// GetListTrailSequencesSortOrderEnumValues Enumerates the set of values for ListTrailSequencesSortOrderEnum
func GetListTrailSequencesSortOrderEnumValues() []ListTrailSequencesSortOrderEnum {
	values := make([]ListTrailSequencesSortOrderEnum, 0)
	for _, v := range mappingListTrailSequencesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTrailSequencesSortOrderEnumStringValues Enumerates the set of values in String for ListTrailSequencesSortOrderEnum
func GetListTrailSequencesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTrailSequencesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTrailSequencesSortOrderEnum(val string) (ListTrailSequencesSortOrderEnum, bool) {
	enum, ok := mappingListTrailSequencesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
