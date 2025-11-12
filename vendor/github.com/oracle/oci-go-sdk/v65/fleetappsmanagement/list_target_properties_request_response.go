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

// ListTargetPropertiesRequest wrapper for the ListTargetProperties operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTargetProperties.go.html to see an example of how to use ListTargetPropertiesRequest.
type ListTargetPropertiesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Target Id.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Target name.
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Patch severity.
	Severity ListTargetPropertiesSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTargetPropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListTargetPropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetPropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetPropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetPropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetPropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetPropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetPropertiesSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListTargetPropertiesSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetPropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetPropertiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetPropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetPropertiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetPropertiesResponse wrapper for the ListTargetProperties operation
type ListTargetPropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetPropertyCollection instances
	TargetPropertyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetPropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetPropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetPropertiesSeverityEnum Enum with underlying type: string
type ListTargetPropertiesSeverityEnum string

// Set of constants representing the allowable values for ListTargetPropertiesSeverityEnum
const (
	ListTargetPropertiesSeverityCritical ListTargetPropertiesSeverityEnum = "CRITICAL"
	ListTargetPropertiesSeverityHigh     ListTargetPropertiesSeverityEnum = "HIGH"
	ListTargetPropertiesSeverityMedium   ListTargetPropertiesSeverityEnum = "MEDIUM"
	ListTargetPropertiesSeverityLow      ListTargetPropertiesSeverityEnum = "LOW"
)

var mappingListTargetPropertiesSeverityEnum = map[string]ListTargetPropertiesSeverityEnum{
	"CRITICAL": ListTargetPropertiesSeverityCritical,
	"HIGH":     ListTargetPropertiesSeverityHigh,
	"MEDIUM":   ListTargetPropertiesSeverityMedium,
	"LOW":      ListTargetPropertiesSeverityLow,
}

var mappingListTargetPropertiesSeverityEnumLowerCase = map[string]ListTargetPropertiesSeverityEnum{
	"critical": ListTargetPropertiesSeverityCritical,
	"high":     ListTargetPropertiesSeverityHigh,
	"medium":   ListTargetPropertiesSeverityMedium,
	"low":      ListTargetPropertiesSeverityLow,
}

// GetListTargetPropertiesSeverityEnumValues Enumerates the set of values for ListTargetPropertiesSeverityEnum
func GetListTargetPropertiesSeverityEnumValues() []ListTargetPropertiesSeverityEnum {
	values := make([]ListTargetPropertiesSeverityEnum, 0)
	for _, v := range mappingListTargetPropertiesSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetPropertiesSeverityEnumStringValues Enumerates the set of values in String for ListTargetPropertiesSeverityEnum
func GetListTargetPropertiesSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingListTargetPropertiesSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetPropertiesSeverityEnum(val string) (ListTargetPropertiesSeverityEnum, bool) {
	enum, ok := mappingListTargetPropertiesSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetPropertiesSortOrderEnum Enum with underlying type: string
type ListTargetPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetPropertiesSortOrderEnum
const (
	ListTargetPropertiesSortOrderAsc  ListTargetPropertiesSortOrderEnum = "ASC"
	ListTargetPropertiesSortOrderDesc ListTargetPropertiesSortOrderEnum = "DESC"
)

var mappingListTargetPropertiesSortOrderEnum = map[string]ListTargetPropertiesSortOrderEnum{
	"ASC":  ListTargetPropertiesSortOrderAsc,
	"DESC": ListTargetPropertiesSortOrderDesc,
}

var mappingListTargetPropertiesSortOrderEnumLowerCase = map[string]ListTargetPropertiesSortOrderEnum{
	"asc":  ListTargetPropertiesSortOrderAsc,
	"desc": ListTargetPropertiesSortOrderDesc,
}

// GetListTargetPropertiesSortOrderEnumValues Enumerates the set of values for ListTargetPropertiesSortOrderEnum
func GetListTargetPropertiesSortOrderEnumValues() []ListTargetPropertiesSortOrderEnum {
	values := make([]ListTargetPropertiesSortOrderEnum, 0)
	for _, v := range mappingListTargetPropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetPropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetPropertiesSortOrderEnum
func GetListTargetPropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetPropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetPropertiesSortOrderEnum(val string) (ListTargetPropertiesSortOrderEnum, bool) {
	enum, ok := mappingListTargetPropertiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetPropertiesSortByEnum Enum with underlying type: string
type ListTargetPropertiesSortByEnum string

// Set of constants representing the allowable values for ListTargetPropertiesSortByEnum
const (
	ListTargetPropertiesSortByPropertyname ListTargetPropertiesSortByEnum = "propertyName"
)

var mappingListTargetPropertiesSortByEnum = map[string]ListTargetPropertiesSortByEnum{
	"propertyName": ListTargetPropertiesSortByPropertyname,
}

var mappingListTargetPropertiesSortByEnumLowerCase = map[string]ListTargetPropertiesSortByEnum{
	"propertyname": ListTargetPropertiesSortByPropertyname,
}

// GetListTargetPropertiesSortByEnumValues Enumerates the set of values for ListTargetPropertiesSortByEnum
func GetListTargetPropertiesSortByEnumValues() []ListTargetPropertiesSortByEnum {
	values := make([]ListTargetPropertiesSortByEnum, 0)
	for _, v := range mappingListTargetPropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetPropertiesSortByEnumStringValues Enumerates the set of values in String for ListTargetPropertiesSortByEnum
func GetListTargetPropertiesSortByEnumStringValues() []string {
	return []string{
		"propertyName",
	}
}

// GetMappingListTargetPropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetPropertiesSortByEnum(val string) (ListTargetPropertiesSortByEnum, bool) {
	enum, ok := mappingListTargetPropertiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
