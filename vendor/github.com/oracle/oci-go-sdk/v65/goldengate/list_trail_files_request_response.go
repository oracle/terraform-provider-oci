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

// ListTrailFilesRequest wrapper for the ListTrailFiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListTrailFiles.go.html to see an example of how to use ListTrailFilesRequest.
type ListTrailFilesRequest struct {

	// A unique Deployment identifier.
	DeploymentId *string `mandatory:"true" contributesTo:"query" name:"deploymentId"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A Trail File identifier
	TrailFileId *string `mandatory:"false" contributesTo:"query" name:"trailFileId"`

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
	SortBy ListTrailFilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTrailFilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTrailFilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTrailFilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTrailFilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTrailFilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTrailFilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTrailFilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTrailFilesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTrailFilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTrailFilesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTrailFilesResponse wrapper for the ListTrailFiles operation
type ListTrailFilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TrailFileCollection instances
	TrailFileCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTrailFilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTrailFilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTrailFilesSortByEnum Enum with underlying type: string
type ListTrailFilesSortByEnum string

// Set of constants representing the allowable values for ListTrailFilesSortByEnum
const (
	ListTrailFilesSortByTimelastupdated ListTrailFilesSortByEnum = "timeLastUpdated"
	ListTrailFilesSortByDisplayname     ListTrailFilesSortByEnum = "displayName"
)

var mappingListTrailFilesSortByEnum = map[string]ListTrailFilesSortByEnum{
	"timeLastUpdated": ListTrailFilesSortByTimelastupdated,
	"displayName":     ListTrailFilesSortByDisplayname,
}

var mappingListTrailFilesSortByEnumLowerCase = map[string]ListTrailFilesSortByEnum{
	"timelastupdated": ListTrailFilesSortByTimelastupdated,
	"displayname":     ListTrailFilesSortByDisplayname,
}

// GetListTrailFilesSortByEnumValues Enumerates the set of values for ListTrailFilesSortByEnum
func GetListTrailFilesSortByEnumValues() []ListTrailFilesSortByEnum {
	values := make([]ListTrailFilesSortByEnum, 0)
	for _, v := range mappingListTrailFilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTrailFilesSortByEnumStringValues Enumerates the set of values in String for ListTrailFilesSortByEnum
func GetListTrailFilesSortByEnumStringValues() []string {
	return []string{
		"timeLastUpdated",
		"displayName",
	}
}

// GetMappingListTrailFilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTrailFilesSortByEnum(val string) (ListTrailFilesSortByEnum, bool) {
	enum, ok := mappingListTrailFilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTrailFilesSortOrderEnum Enum with underlying type: string
type ListTrailFilesSortOrderEnum string

// Set of constants representing the allowable values for ListTrailFilesSortOrderEnum
const (
	ListTrailFilesSortOrderAsc  ListTrailFilesSortOrderEnum = "ASC"
	ListTrailFilesSortOrderDesc ListTrailFilesSortOrderEnum = "DESC"
)

var mappingListTrailFilesSortOrderEnum = map[string]ListTrailFilesSortOrderEnum{
	"ASC":  ListTrailFilesSortOrderAsc,
	"DESC": ListTrailFilesSortOrderDesc,
}

var mappingListTrailFilesSortOrderEnumLowerCase = map[string]ListTrailFilesSortOrderEnum{
	"asc":  ListTrailFilesSortOrderAsc,
	"desc": ListTrailFilesSortOrderDesc,
}

// GetListTrailFilesSortOrderEnumValues Enumerates the set of values for ListTrailFilesSortOrderEnum
func GetListTrailFilesSortOrderEnumValues() []ListTrailFilesSortOrderEnum {
	values := make([]ListTrailFilesSortOrderEnum, 0)
	for _, v := range mappingListTrailFilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTrailFilesSortOrderEnumStringValues Enumerates the set of values in String for ListTrailFilesSortOrderEnum
func GetListTrailFilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTrailFilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTrailFilesSortOrderEnum(val string) (ListTrailFilesSortOrderEnum, bool) {
	enum, ok := mappingListTrailFilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
