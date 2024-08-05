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

// ListCryptoAnalysisResultsRequest wrapper for the ListCryptoAnalysisResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListCryptoAnalysisResults.go.html to see an example of how to use ListCryptoAnalysisResultsRequest.
type ListCryptoAnalysisResultsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The aggregation mode of the crypto event analysis result.
	AggregationMode ListCryptoAnalysisResultsAggregationModeEnum `mandatory:"false" contributesTo:"query" name:"aggregationMode" omitEmpty:"true"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The host OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the managed instance.
	HostName *string `mandatory:"false" contributesTo:"query" name:"hostName"`

	// Non Compliant Finding Count of CryptoAnalysis Report.
	NonCompliantFindingCount *int `mandatory:"false" contributesTo:"query" name:"nonCompliantFindingCount"`

	// Non Compliant Finding Count of CryptoAnalysis Report.
	NonCompliantFindingCountGreaterThan *int `mandatory:"false" contributesTo:"query" name:"nonCompliantFindingCountGreaterThan"`

	// FindingCount of CryptoAnalysis Report.
	FindingCount *int `mandatory:"false" contributesTo:"query" name:"findingCount"`

	// FindingCount of CryptoAnalysis Report.
	FindingCountGreaterThan *int `mandatory:"false" contributesTo:"query" name:"findingCountGreaterThan"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListCryptoAnalysisResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort crypto event analysis results. Only one sort order can be provided.
	// Default order for _timeCreated_, and _jreVersion_ is **descending**.
	// Default order for _managedInstanceId_, _jreDistribution_, _jreVendor_ and _osName_ is **ascending**.
	// If no value is specified _timeCreated_ is default.
	SortBy ListCryptoAnalysisResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCryptoAnalysisResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAnalysisResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAnalysisResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAnalysisResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAnalysisResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAnalysisResultsAggregationModeEnum(string(request.AggregationMode)); !ok && request.AggregationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AggregationMode: %s. Supported values are: %s.", request.AggregationMode, strings.Join(GetListCryptoAnalysisResultsAggregationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAnalysisResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAnalysisResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAnalysisResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAnalysisResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAnalysisResultsResponse wrapper for the ListCryptoAnalysisResults operation
type ListCryptoAnalysisResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAnalysisResultCollection instances
	CryptoAnalysisResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCryptoAnalysisResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAnalysisResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAnalysisResultsAggregationModeEnum Enum with underlying type: string
type ListCryptoAnalysisResultsAggregationModeEnum string

// Set of constants representing the allowable values for ListCryptoAnalysisResultsAggregationModeEnum
const (
	ListCryptoAnalysisResultsAggregationModeJfr             ListCryptoAnalysisResultsAggregationModeEnum = "JFR"
	ListCryptoAnalysisResultsAggregationModeManagedInstance ListCryptoAnalysisResultsAggregationModeEnum = "MANAGED_INSTANCE"
)

var mappingListCryptoAnalysisResultsAggregationModeEnum = map[string]ListCryptoAnalysisResultsAggregationModeEnum{
	"JFR":              ListCryptoAnalysisResultsAggregationModeJfr,
	"MANAGED_INSTANCE": ListCryptoAnalysisResultsAggregationModeManagedInstance,
}

var mappingListCryptoAnalysisResultsAggregationModeEnumLowerCase = map[string]ListCryptoAnalysisResultsAggregationModeEnum{
	"jfr":              ListCryptoAnalysisResultsAggregationModeJfr,
	"managed_instance": ListCryptoAnalysisResultsAggregationModeManagedInstance,
}

// GetListCryptoAnalysisResultsAggregationModeEnumValues Enumerates the set of values for ListCryptoAnalysisResultsAggregationModeEnum
func GetListCryptoAnalysisResultsAggregationModeEnumValues() []ListCryptoAnalysisResultsAggregationModeEnum {
	values := make([]ListCryptoAnalysisResultsAggregationModeEnum, 0)
	for _, v := range mappingListCryptoAnalysisResultsAggregationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAnalysisResultsAggregationModeEnumStringValues Enumerates the set of values in String for ListCryptoAnalysisResultsAggregationModeEnum
func GetListCryptoAnalysisResultsAggregationModeEnumStringValues() []string {
	return []string{
		"JFR",
		"MANAGED_INSTANCE",
	}
}

// GetMappingListCryptoAnalysisResultsAggregationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAnalysisResultsAggregationModeEnum(val string) (ListCryptoAnalysisResultsAggregationModeEnum, bool) {
	enum, ok := mappingListCryptoAnalysisResultsAggregationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAnalysisResultsSortOrderEnum Enum with underlying type: string
type ListCryptoAnalysisResultsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAnalysisResultsSortOrderEnum
const (
	ListCryptoAnalysisResultsSortOrderAsc  ListCryptoAnalysisResultsSortOrderEnum = "ASC"
	ListCryptoAnalysisResultsSortOrderDesc ListCryptoAnalysisResultsSortOrderEnum = "DESC"
)

var mappingListCryptoAnalysisResultsSortOrderEnum = map[string]ListCryptoAnalysisResultsSortOrderEnum{
	"ASC":  ListCryptoAnalysisResultsSortOrderAsc,
	"DESC": ListCryptoAnalysisResultsSortOrderDesc,
}

var mappingListCryptoAnalysisResultsSortOrderEnumLowerCase = map[string]ListCryptoAnalysisResultsSortOrderEnum{
	"asc":  ListCryptoAnalysisResultsSortOrderAsc,
	"desc": ListCryptoAnalysisResultsSortOrderDesc,
}

// GetListCryptoAnalysisResultsSortOrderEnumValues Enumerates the set of values for ListCryptoAnalysisResultsSortOrderEnum
func GetListCryptoAnalysisResultsSortOrderEnumValues() []ListCryptoAnalysisResultsSortOrderEnum {
	values := make([]ListCryptoAnalysisResultsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAnalysisResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAnalysisResultsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAnalysisResultsSortOrderEnum
func GetListCryptoAnalysisResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAnalysisResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAnalysisResultsSortOrderEnum(val string) (ListCryptoAnalysisResultsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAnalysisResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAnalysisResultsSortByEnum Enum with underlying type: string
type ListCryptoAnalysisResultsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAnalysisResultsSortByEnum
const (
	ListCryptoAnalysisResultsSortByTimecreated       ListCryptoAnalysisResultsSortByEnum = "timeCreated"
	ListCryptoAnalysisResultsSortByManagedinstanceid ListCryptoAnalysisResultsSortByEnum = "managedInstanceId"
	ListCryptoAnalysisResultsSortByWorkrequestid     ListCryptoAnalysisResultsSortByEnum = "workRequestId"
)

var mappingListCryptoAnalysisResultsSortByEnum = map[string]ListCryptoAnalysisResultsSortByEnum{
	"timeCreated":       ListCryptoAnalysisResultsSortByTimecreated,
	"managedInstanceId": ListCryptoAnalysisResultsSortByManagedinstanceid,
	"workRequestId":     ListCryptoAnalysisResultsSortByWorkrequestid,
}

var mappingListCryptoAnalysisResultsSortByEnumLowerCase = map[string]ListCryptoAnalysisResultsSortByEnum{
	"timecreated":       ListCryptoAnalysisResultsSortByTimecreated,
	"managedinstanceid": ListCryptoAnalysisResultsSortByManagedinstanceid,
	"workrequestid":     ListCryptoAnalysisResultsSortByWorkrequestid,
}

// GetListCryptoAnalysisResultsSortByEnumValues Enumerates the set of values for ListCryptoAnalysisResultsSortByEnum
func GetListCryptoAnalysisResultsSortByEnumValues() []ListCryptoAnalysisResultsSortByEnum {
	values := make([]ListCryptoAnalysisResultsSortByEnum, 0)
	for _, v := range mappingListCryptoAnalysisResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAnalysisResultsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAnalysisResultsSortByEnum
func GetListCryptoAnalysisResultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
	}
}

// GetMappingListCryptoAnalysisResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAnalysisResultsSortByEnum(val string) (ListCryptoAnalysisResultsSortByEnum, bool) {
	enum, ok := mappingListCryptoAnalysisResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
