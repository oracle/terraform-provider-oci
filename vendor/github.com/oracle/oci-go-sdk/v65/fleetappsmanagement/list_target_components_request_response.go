// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetComponentsRequest wrapper for the ListTargetComponents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTargetComponents.go.html to see an example of how to use ListTargetComponentsRequest.
type ListTargetComponentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Target Id.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Target name.
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Patch severity.
	Severity ListTargetComponentsSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// Target Component Name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTargetComponentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListTargetComponentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetComponentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetComponentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetComponentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetComponentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetComponentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetComponentsSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListTargetComponentsSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetComponentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetComponentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetComponentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetComponentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetComponentsResponse wrapper for the ListTargetComponents operation
type ListTargetComponentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetComponentCollection instances
	TargetComponentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetComponentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetComponentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetComponentsSeverityEnum Enum with underlying type: string
type ListTargetComponentsSeverityEnum string

// Set of constants representing the allowable values for ListTargetComponentsSeverityEnum
const (
	ListTargetComponentsSeverityCritical ListTargetComponentsSeverityEnum = "CRITICAL"
	ListTargetComponentsSeverityHigh     ListTargetComponentsSeverityEnum = "HIGH"
	ListTargetComponentsSeverityMedium   ListTargetComponentsSeverityEnum = "MEDIUM"
	ListTargetComponentsSeverityLow      ListTargetComponentsSeverityEnum = "LOW"
)

var mappingListTargetComponentsSeverityEnum = map[string]ListTargetComponentsSeverityEnum{
	"CRITICAL": ListTargetComponentsSeverityCritical,
	"HIGH":     ListTargetComponentsSeverityHigh,
	"MEDIUM":   ListTargetComponentsSeverityMedium,
	"LOW":      ListTargetComponentsSeverityLow,
}

var mappingListTargetComponentsSeverityEnumLowerCase = map[string]ListTargetComponentsSeverityEnum{
	"critical": ListTargetComponentsSeverityCritical,
	"high":     ListTargetComponentsSeverityHigh,
	"medium":   ListTargetComponentsSeverityMedium,
	"low":      ListTargetComponentsSeverityLow,
}

// GetListTargetComponentsSeverityEnumValues Enumerates the set of values for ListTargetComponentsSeverityEnum
func GetListTargetComponentsSeverityEnumValues() []ListTargetComponentsSeverityEnum {
	values := make([]ListTargetComponentsSeverityEnum, 0)
	for _, v := range mappingListTargetComponentsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetComponentsSeverityEnumStringValues Enumerates the set of values in String for ListTargetComponentsSeverityEnum
func GetListTargetComponentsSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingListTargetComponentsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetComponentsSeverityEnum(val string) (ListTargetComponentsSeverityEnum, bool) {
	enum, ok := mappingListTargetComponentsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetComponentsSortOrderEnum Enum with underlying type: string
type ListTargetComponentsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetComponentsSortOrderEnum
const (
	ListTargetComponentsSortOrderAsc  ListTargetComponentsSortOrderEnum = "ASC"
	ListTargetComponentsSortOrderDesc ListTargetComponentsSortOrderEnum = "DESC"
)

var mappingListTargetComponentsSortOrderEnum = map[string]ListTargetComponentsSortOrderEnum{
	"ASC":  ListTargetComponentsSortOrderAsc,
	"DESC": ListTargetComponentsSortOrderDesc,
}

var mappingListTargetComponentsSortOrderEnumLowerCase = map[string]ListTargetComponentsSortOrderEnum{
	"asc":  ListTargetComponentsSortOrderAsc,
	"desc": ListTargetComponentsSortOrderDesc,
}

// GetListTargetComponentsSortOrderEnumValues Enumerates the set of values for ListTargetComponentsSortOrderEnum
func GetListTargetComponentsSortOrderEnumValues() []ListTargetComponentsSortOrderEnum {
	values := make([]ListTargetComponentsSortOrderEnum, 0)
	for _, v := range mappingListTargetComponentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetComponentsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetComponentsSortOrderEnum
func GetListTargetComponentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetComponentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetComponentsSortOrderEnum(val string) (ListTargetComponentsSortOrderEnum, bool) {
	enum, ok := mappingListTargetComponentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetComponentsSortByEnum Enum with underlying type: string
type ListTargetComponentsSortByEnum string

// Set of constants representing the allowable values for ListTargetComponentsSortByEnum
const (
	ListTargetComponentsSortByComponentname ListTargetComponentsSortByEnum = "componentName"
)

var mappingListTargetComponentsSortByEnum = map[string]ListTargetComponentsSortByEnum{
	"componentName": ListTargetComponentsSortByComponentname,
}

var mappingListTargetComponentsSortByEnumLowerCase = map[string]ListTargetComponentsSortByEnum{
	"componentname": ListTargetComponentsSortByComponentname,
}

// GetListTargetComponentsSortByEnumValues Enumerates the set of values for ListTargetComponentsSortByEnum
func GetListTargetComponentsSortByEnumValues() []ListTargetComponentsSortByEnum {
	values := make([]ListTargetComponentsSortByEnum, 0)
	for _, v := range mappingListTargetComponentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetComponentsSortByEnumStringValues Enumerates the set of values in String for ListTargetComponentsSortByEnum
func GetListTargetComponentsSortByEnumStringValues() []string {
	return []string{
		"componentName",
	}
}

// GetMappingListTargetComponentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetComponentsSortByEnum(val string) (ListTargetComponentsSortByEnum, bool) {
	enum, ok := mappingListTargetComponentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
