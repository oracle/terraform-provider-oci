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

// ListPerformanceTuningAnalysisResultsRequest wrapper for the ListPerformanceTuningAnalysisResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListPerformanceTuningAnalysisResults.go.html to see an example of how to use ListPerformanceTuningAnalysisResultsRequest.
type ListPerformanceTuningAnalysisResultsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The host OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the managed instance.
	HostName *string `mandatory:"false" contributesTo:"query" name:"hostName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListPerformanceTuningAnalysisResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort performance tuning analysis results. Only one sort order may be provided.
	// Default order for _timeCreated_, and _jreVersion_ is **descending**.
	// Default order for _managedInstanceId_, _jreDistribution_, _jreVendor_ and _osName_ is **ascending**.
	// If no value is specified _timeCreated_ is default.
	SortBy ListPerformanceTuningAnalysisResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPerformanceTuningAnalysisResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPerformanceTuningAnalysisResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPerformanceTuningAnalysisResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPerformanceTuningAnalysisResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPerformanceTuningAnalysisResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPerformanceTuningAnalysisResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPerformanceTuningAnalysisResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPerformanceTuningAnalysisResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPerformanceTuningAnalysisResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPerformanceTuningAnalysisResultsResponse wrapper for the ListPerformanceTuningAnalysisResults operation
type ListPerformanceTuningAnalysisResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PerformanceTuningAnalysisResultCollection instances
	PerformanceTuningAnalysisResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPerformanceTuningAnalysisResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPerformanceTuningAnalysisResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPerformanceTuningAnalysisResultsSortOrderEnum Enum with underlying type: string
type ListPerformanceTuningAnalysisResultsSortOrderEnum string

// Set of constants representing the allowable values for ListPerformanceTuningAnalysisResultsSortOrderEnum
const (
	ListPerformanceTuningAnalysisResultsSortOrderAsc  ListPerformanceTuningAnalysisResultsSortOrderEnum = "ASC"
	ListPerformanceTuningAnalysisResultsSortOrderDesc ListPerformanceTuningAnalysisResultsSortOrderEnum = "DESC"
)

var mappingListPerformanceTuningAnalysisResultsSortOrderEnum = map[string]ListPerformanceTuningAnalysisResultsSortOrderEnum{
	"ASC":  ListPerformanceTuningAnalysisResultsSortOrderAsc,
	"DESC": ListPerformanceTuningAnalysisResultsSortOrderDesc,
}

var mappingListPerformanceTuningAnalysisResultsSortOrderEnumLowerCase = map[string]ListPerformanceTuningAnalysisResultsSortOrderEnum{
	"asc":  ListPerformanceTuningAnalysisResultsSortOrderAsc,
	"desc": ListPerformanceTuningAnalysisResultsSortOrderDesc,
}

// GetListPerformanceTuningAnalysisResultsSortOrderEnumValues Enumerates the set of values for ListPerformanceTuningAnalysisResultsSortOrderEnum
func GetListPerformanceTuningAnalysisResultsSortOrderEnumValues() []ListPerformanceTuningAnalysisResultsSortOrderEnum {
	values := make([]ListPerformanceTuningAnalysisResultsSortOrderEnum, 0)
	for _, v := range mappingListPerformanceTuningAnalysisResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPerformanceTuningAnalysisResultsSortOrderEnumStringValues Enumerates the set of values in String for ListPerformanceTuningAnalysisResultsSortOrderEnum
func GetListPerformanceTuningAnalysisResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPerformanceTuningAnalysisResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPerformanceTuningAnalysisResultsSortOrderEnum(val string) (ListPerformanceTuningAnalysisResultsSortOrderEnum, bool) {
	enum, ok := mappingListPerformanceTuningAnalysisResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPerformanceTuningAnalysisResultsSortByEnum Enum with underlying type: string
type ListPerformanceTuningAnalysisResultsSortByEnum string

// Set of constants representing the allowable values for ListPerformanceTuningAnalysisResultsSortByEnum
const (
	ListPerformanceTuningAnalysisResultsSortByTimecreated       ListPerformanceTuningAnalysisResultsSortByEnum = "timeCreated"
	ListPerformanceTuningAnalysisResultsSortByManagedinstanceid ListPerformanceTuningAnalysisResultsSortByEnum = "managedInstanceId"
	ListPerformanceTuningAnalysisResultsSortByWorkrequestid     ListPerformanceTuningAnalysisResultsSortByEnum = "workRequestId"
	ListPerformanceTuningAnalysisResultsSortByWarningcount      ListPerformanceTuningAnalysisResultsSortByEnum = "warningCount"
)

var mappingListPerformanceTuningAnalysisResultsSortByEnum = map[string]ListPerformanceTuningAnalysisResultsSortByEnum{
	"timeCreated":       ListPerformanceTuningAnalysisResultsSortByTimecreated,
	"managedInstanceId": ListPerformanceTuningAnalysisResultsSortByManagedinstanceid,
	"workRequestId":     ListPerformanceTuningAnalysisResultsSortByWorkrequestid,
	"warningCount":      ListPerformanceTuningAnalysisResultsSortByWarningcount,
}

var mappingListPerformanceTuningAnalysisResultsSortByEnumLowerCase = map[string]ListPerformanceTuningAnalysisResultsSortByEnum{
	"timecreated":       ListPerformanceTuningAnalysisResultsSortByTimecreated,
	"managedinstanceid": ListPerformanceTuningAnalysisResultsSortByManagedinstanceid,
	"workrequestid":     ListPerformanceTuningAnalysisResultsSortByWorkrequestid,
	"warningcount":      ListPerformanceTuningAnalysisResultsSortByWarningcount,
}

// GetListPerformanceTuningAnalysisResultsSortByEnumValues Enumerates the set of values for ListPerformanceTuningAnalysisResultsSortByEnum
func GetListPerformanceTuningAnalysisResultsSortByEnumValues() []ListPerformanceTuningAnalysisResultsSortByEnum {
	values := make([]ListPerformanceTuningAnalysisResultsSortByEnum, 0)
	for _, v := range mappingListPerformanceTuningAnalysisResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPerformanceTuningAnalysisResultsSortByEnumStringValues Enumerates the set of values in String for ListPerformanceTuningAnalysisResultsSortByEnum
func GetListPerformanceTuningAnalysisResultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
		"warningCount",
	}
}

// GetMappingListPerformanceTuningAnalysisResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPerformanceTuningAnalysisResultsSortByEnum(val string) (ListPerformanceTuningAnalysisResultsSortByEnum, bool) {
	enum, ok := mappingListPerformanceTuningAnalysisResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
