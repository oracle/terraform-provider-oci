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

// ListRecommendedPatchesRequest wrapper for the ListRecommendedPatches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRecommendedPatches.go.html to see an example of how to use ListRecommendedPatchesRequest.
type ListRecommendedPatchesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Target identifier.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Unique target name
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Patch level.
	PatchLevel ListRecommendedPatchesPatchLevelEnum `mandatory:"false" contributesTo:"query" name:"patchLevel" omitEmpty:"true"`

	// Patch severity.
	Severity ListRecommendedPatchesSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// Patch Type.
	PatchType *string `mandatory:"false" contributesTo:"query" name:"patchType"`

	// Patch identifier.
	PatchId *string `mandatory:"false" contributesTo:"query" name:"patchId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRecommendedPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListRecommendedPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendedPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendedPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendedPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendedPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendedPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendedPatchesPatchLevelEnum(string(request.PatchLevel)); !ok && request.PatchLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchLevel: %s. Supported values are: %s.", request.PatchLevel, strings.Join(GetListRecommendedPatchesPatchLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendedPatchesSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListRecommendedPatchesSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendedPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecommendedPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendedPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecommendedPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendedPatchesResponse wrapper for the ListRecommendedPatches operation
type ListRecommendedPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendedPatchCollection instances
	RecommendedPatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecommendedPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendedPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendedPatchesPatchLevelEnum Enum with underlying type: string
type ListRecommendedPatchesPatchLevelEnum string

// Set of constants representing the allowable values for ListRecommendedPatchesPatchLevelEnum
const (
	ListRecommendedPatchesPatchLevelLatest         ListRecommendedPatchesPatchLevelEnum = "LATEST"
	ListRecommendedPatchesPatchLevelLatestMinusOne ListRecommendedPatchesPatchLevelEnum = "LATEST_MINUS_ONE"
	ListRecommendedPatchesPatchLevelLatestMinusTwo ListRecommendedPatchesPatchLevelEnum = "LATEST_MINUS_TWO"
)

var mappingListRecommendedPatchesPatchLevelEnum = map[string]ListRecommendedPatchesPatchLevelEnum{
	"LATEST":           ListRecommendedPatchesPatchLevelLatest,
	"LATEST_MINUS_ONE": ListRecommendedPatchesPatchLevelLatestMinusOne,
	"LATEST_MINUS_TWO": ListRecommendedPatchesPatchLevelLatestMinusTwo,
}

var mappingListRecommendedPatchesPatchLevelEnumLowerCase = map[string]ListRecommendedPatchesPatchLevelEnum{
	"latest":           ListRecommendedPatchesPatchLevelLatest,
	"latest_minus_one": ListRecommendedPatchesPatchLevelLatestMinusOne,
	"latest_minus_two": ListRecommendedPatchesPatchLevelLatestMinusTwo,
}

// GetListRecommendedPatchesPatchLevelEnumValues Enumerates the set of values for ListRecommendedPatchesPatchLevelEnum
func GetListRecommendedPatchesPatchLevelEnumValues() []ListRecommendedPatchesPatchLevelEnum {
	values := make([]ListRecommendedPatchesPatchLevelEnum, 0)
	for _, v := range mappingListRecommendedPatchesPatchLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendedPatchesPatchLevelEnumStringValues Enumerates the set of values in String for ListRecommendedPatchesPatchLevelEnum
func GetListRecommendedPatchesPatchLevelEnumStringValues() []string {
	return []string{
		"LATEST",
		"LATEST_MINUS_ONE",
		"LATEST_MINUS_TWO",
	}
}

// GetMappingListRecommendedPatchesPatchLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendedPatchesPatchLevelEnum(val string) (ListRecommendedPatchesPatchLevelEnum, bool) {
	enum, ok := mappingListRecommendedPatchesPatchLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendedPatchesSeverityEnum Enum with underlying type: string
type ListRecommendedPatchesSeverityEnum string

// Set of constants representing the allowable values for ListRecommendedPatchesSeverityEnum
const (
	ListRecommendedPatchesSeverityCritical ListRecommendedPatchesSeverityEnum = "CRITICAL"
	ListRecommendedPatchesSeverityHigh     ListRecommendedPatchesSeverityEnum = "HIGH"
	ListRecommendedPatchesSeverityMedium   ListRecommendedPatchesSeverityEnum = "MEDIUM"
	ListRecommendedPatchesSeverityLow      ListRecommendedPatchesSeverityEnum = "LOW"
)

var mappingListRecommendedPatchesSeverityEnum = map[string]ListRecommendedPatchesSeverityEnum{
	"CRITICAL": ListRecommendedPatchesSeverityCritical,
	"HIGH":     ListRecommendedPatchesSeverityHigh,
	"MEDIUM":   ListRecommendedPatchesSeverityMedium,
	"LOW":      ListRecommendedPatchesSeverityLow,
}

var mappingListRecommendedPatchesSeverityEnumLowerCase = map[string]ListRecommendedPatchesSeverityEnum{
	"critical": ListRecommendedPatchesSeverityCritical,
	"high":     ListRecommendedPatchesSeverityHigh,
	"medium":   ListRecommendedPatchesSeverityMedium,
	"low":      ListRecommendedPatchesSeverityLow,
}

// GetListRecommendedPatchesSeverityEnumValues Enumerates the set of values for ListRecommendedPatchesSeverityEnum
func GetListRecommendedPatchesSeverityEnumValues() []ListRecommendedPatchesSeverityEnum {
	values := make([]ListRecommendedPatchesSeverityEnum, 0)
	for _, v := range mappingListRecommendedPatchesSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendedPatchesSeverityEnumStringValues Enumerates the set of values in String for ListRecommendedPatchesSeverityEnum
func GetListRecommendedPatchesSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingListRecommendedPatchesSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendedPatchesSeverityEnum(val string) (ListRecommendedPatchesSeverityEnum, bool) {
	enum, ok := mappingListRecommendedPatchesSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendedPatchesSortOrderEnum Enum with underlying type: string
type ListRecommendedPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListRecommendedPatchesSortOrderEnum
const (
	ListRecommendedPatchesSortOrderAsc  ListRecommendedPatchesSortOrderEnum = "ASC"
	ListRecommendedPatchesSortOrderDesc ListRecommendedPatchesSortOrderEnum = "DESC"
)

var mappingListRecommendedPatchesSortOrderEnum = map[string]ListRecommendedPatchesSortOrderEnum{
	"ASC":  ListRecommendedPatchesSortOrderAsc,
	"DESC": ListRecommendedPatchesSortOrderDesc,
}

var mappingListRecommendedPatchesSortOrderEnumLowerCase = map[string]ListRecommendedPatchesSortOrderEnum{
	"asc":  ListRecommendedPatchesSortOrderAsc,
	"desc": ListRecommendedPatchesSortOrderDesc,
}

// GetListRecommendedPatchesSortOrderEnumValues Enumerates the set of values for ListRecommendedPatchesSortOrderEnum
func GetListRecommendedPatchesSortOrderEnumValues() []ListRecommendedPatchesSortOrderEnum {
	values := make([]ListRecommendedPatchesSortOrderEnum, 0)
	for _, v := range mappingListRecommendedPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendedPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListRecommendedPatchesSortOrderEnum
func GetListRecommendedPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecommendedPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendedPatchesSortOrderEnum(val string) (ListRecommendedPatchesSortOrderEnum, bool) {
	enum, ok := mappingListRecommendedPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendedPatchesSortByEnum Enum with underlying type: string
type ListRecommendedPatchesSortByEnum string

// Set of constants representing the allowable values for ListRecommendedPatchesSortByEnum
const (
	ListRecommendedPatchesSortByPatchname ListRecommendedPatchesSortByEnum = "patchName"
)

var mappingListRecommendedPatchesSortByEnum = map[string]ListRecommendedPatchesSortByEnum{
	"patchName": ListRecommendedPatchesSortByPatchname,
}

var mappingListRecommendedPatchesSortByEnumLowerCase = map[string]ListRecommendedPatchesSortByEnum{
	"patchname": ListRecommendedPatchesSortByPatchname,
}

// GetListRecommendedPatchesSortByEnumValues Enumerates the set of values for ListRecommendedPatchesSortByEnum
func GetListRecommendedPatchesSortByEnumValues() []ListRecommendedPatchesSortByEnum {
	values := make([]ListRecommendedPatchesSortByEnum, 0)
	for _, v := range mappingListRecommendedPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendedPatchesSortByEnumStringValues Enumerates the set of values in String for ListRecommendedPatchesSortByEnum
func GetListRecommendedPatchesSortByEnumStringValues() []string {
	return []string{
		"patchName",
	}
}

// GetMappingListRecommendedPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendedPatchesSortByEnum(val string) (ListRecommendedPatchesSortByEnum, bool) {
	enum, ok := mappingListRecommendedPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
