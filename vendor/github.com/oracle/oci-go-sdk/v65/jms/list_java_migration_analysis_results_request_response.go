// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaMigrationAnalysisResultsRequest wrapper for the ListJavaMigrationAnalysisResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaMigrationAnalysisResults.go.html to see an example of how to use ListJavaMigrationAnalysisResultsRequest.
type ListJavaMigrationAnalysisResultsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The host OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the managed instance.
	HostName *string `mandatory:"false" contributesTo:"query" name:"hostName"`

	// The name of the application.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaMigrationAnalysisResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field that sorts the Java migration analysis results. Only one sort order can be provided.
	// The default order for _timeCreated_, _managedInstanceId_ and _workRequestId_ is **descending**.
	// If no value is specified, then _timeCreated_ is default.
	SortBy ListJavaMigrationAnalysisResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaMigrationAnalysisResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaMigrationAnalysisResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaMigrationAnalysisResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaMigrationAnalysisResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaMigrationAnalysisResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaMigrationAnalysisResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaMigrationAnalysisResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaMigrationAnalysisResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaMigrationAnalysisResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaMigrationAnalysisResultsResponse wrapper for the ListJavaMigrationAnalysisResults operation
type ListJavaMigrationAnalysisResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaMigrationAnalysisResultCollection instances
	JavaMigrationAnalysisResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaMigrationAnalysisResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaMigrationAnalysisResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaMigrationAnalysisResultsSortOrderEnum Enum with underlying type: string
type ListJavaMigrationAnalysisResultsSortOrderEnum string

// Set of constants representing the allowable values for ListJavaMigrationAnalysisResultsSortOrderEnum
const (
	ListJavaMigrationAnalysisResultsSortOrderAsc  ListJavaMigrationAnalysisResultsSortOrderEnum = "ASC"
	ListJavaMigrationAnalysisResultsSortOrderDesc ListJavaMigrationAnalysisResultsSortOrderEnum = "DESC"
)

var mappingListJavaMigrationAnalysisResultsSortOrderEnum = map[string]ListJavaMigrationAnalysisResultsSortOrderEnum{
	"ASC":  ListJavaMigrationAnalysisResultsSortOrderAsc,
	"DESC": ListJavaMigrationAnalysisResultsSortOrderDesc,
}

var mappingListJavaMigrationAnalysisResultsSortOrderEnumLowerCase = map[string]ListJavaMigrationAnalysisResultsSortOrderEnum{
	"asc":  ListJavaMigrationAnalysisResultsSortOrderAsc,
	"desc": ListJavaMigrationAnalysisResultsSortOrderDesc,
}

// GetListJavaMigrationAnalysisResultsSortOrderEnumValues Enumerates the set of values for ListJavaMigrationAnalysisResultsSortOrderEnum
func GetListJavaMigrationAnalysisResultsSortOrderEnumValues() []ListJavaMigrationAnalysisResultsSortOrderEnum {
	values := make([]ListJavaMigrationAnalysisResultsSortOrderEnum, 0)
	for _, v := range mappingListJavaMigrationAnalysisResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaMigrationAnalysisResultsSortOrderEnumStringValues Enumerates the set of values in String for ListJavaMigrationAnalysisResultsSortOrderEnum
func GetListJavaMigrationAnalysisResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaMigrationAnalysisResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaMigrationAnalysisResultsSortOrderEnum(val string) (ListJavaMigrationAnalysisResultsSortOrderEnum, bool) {
	enum, ok := mappingListJavaMigrationAnalysisResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaMigrationAnalysisResultsSortByEnum Enum with underlying type: string
type ListJavaMigrationAnalysisResultsSortByEnum string

// Set of constants representing the allowable values for ListJavaMigrationAnalysisResultsSortByEnum
const (
	ListJavaMigrationAnalysisResultsSortByTimecreated       ListJavaMigrationAnalysisResultsSortByEnum = "timeCreated"
	ListJavaMigrationAnalysisResultsSortByManagedinstanceid ListJavaMigrationAnalysisResultsSortByEnum = "managedInstanceId"
	ListJavaMigrationAnalysisResultsSortByWorkrequestid     ListJavaMigrationAnalysisResultsSortByEnum = "workRequestId"
)

var mappingListJavaMigrationAnalysisResultsSortByEnum = map[string]ListJavaMigrationAnalysisResultsSortByEnum{
	"timeCreated":       ListJavaMigrationAnalysisResultsSortByTimecreated,
	"managedInstanceId": ListJavaMigrationAnalysisResultsSortByManagedinstanceid,
	"workRequestId":     ListJavaMigrationAnalysisResultsSortByWorkrequestid,
}

var mappingListJavaMigrationAnalysisResultsSortByEnumLowerCase = map[string]ListJavaMigrationAnalysisResultsSortByEnum{
	"timecreated":       ListJavaMigrationAnalysisResultsSortByTimecreated,
	"managedinstanceid": ListJavaMigrationAnalysisResultsSortByManagedinstanceid,
	"workrequestid":     ListJavaMigrationAnalysisResultsSortByWorkrequestid,
}

// GetListJavaMigrationAnalysisResultsSortByEnumValues Enumerates the set of values for ListJavaMigrationAnalysisResultsSortByEnum
func GetListJavaMigrationAnalysisResultsSortByEnumValues() []ListJavaMigrationAnalysisResultsSortByEnum {
	values := make([]ListJavaMigrationAnalysisResultsSortByEnum, 0)
	for _, v := range mappingListJavaMigrationAnalysisResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaMigrationAnalysisResultsSortByEnumStringValues Enumerates the set of values in String for ListJavaMigrationAnalysisResultsSortByEnum
func GetListJavaMigrationAnalysisResultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
	}
}

// GetMappingListJavaMigrationAnalysisResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaMigrationAnalysisResultsSortByEnum(val string) (ListJavaMigrationAnalysisResultsSortByEnum, bool) {
	enum, ok := mappingListJavaMigrationAnalysisResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
