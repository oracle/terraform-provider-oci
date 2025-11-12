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

// ListInstalledPatchesRequest wrapper for the ListInstalledPatches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListInstalledPatches.go.html to see an example of how to use ListInstalledPatchesRequest.
type ListInstalledPatchesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Target Id.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Target name.
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Patch severity.
	Severity ListInstalledPatchesSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// Patch level.
	PatchLevel *string `mandatory:"false" contributesTo:"query" name:"patchLevel"`

	// Patch Type.
	PatchType *string `mandatory:"false" contributesTo:"query" name:"patchType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInstalledPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListInstalledPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstalledPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstalledPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstalledPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstalledPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstalledPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstalledPatchesSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListInstalledPatchesSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstalledPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstalledPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstalledPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstalledPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstalledPatchesResponse wrapper for the ListInstalledPatches operation
type ListInstalledPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledPatchCollection instances
	InstalledPatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstalledPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstalledPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstalledPatchesSeverityEnum Enum with underlying type: string
type ListInstalledPatchesSeverityEnum string

// Set of constants representing the allowable values for ListInstalledPatchesSeverityEnum
const (
	ListInstalledPatchesSeverityCritical ListInstalledPatchesSeverityEnum = "CRITICAL"
	ListInstalledPatchesSeverityHigh     ListInstalledPatchesSeverityEnum = "HIGH"
	ListInstalledPatchesSeverityMedium   ListInstalledPatchesSeverityEnum = "MEDIUM"
	ListInstalledPatchesSeverityLow      ListInstalledPatchesSeverityEnum = "LOW"
)

var mappingListInstalledPatchesSeverityEnum = map[string]ListInstalledPatchesSeverityEnum{
	"CRITICAL": ListInstalledPatchesSeverityCritical,
	"HIGH":     ListInstalledPatchesSeverityHigh,
	"MEDIUM":   ListInstalledPatchesSeverityMedium,
	"LOW":      ListInstalledPatchesSeverityLow,
}

var mappingListInstalledPatchesSeverityEnumLowerCase = map[string]ListInstalledPatchesSeverityEnum{
	"critical": ListInstalledPatchesSeverityCritical,
	"high":     ListInstalledPatchesSeverityHigh,
	"medium":   ListInstalledPatchesSeverityMedium,
	"low":      ListInstalledPatchesSeverityLow,
}

// GetListInstalledPatchesSeverityEnumValues Enumerates the set of values for ListInstalledPatchesSeverityEnum
func GetListInstalledPatchesSeverityEnumValues() []ListInstalledPatchesSeverityEnum {
	values := make([]ListInstalledPatchesSeverityEnum, 0)
	for _, v := range mappingListInstalledPatchesSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstalledPatchesSeverityEnumStringValues Enumerates the set of values in String for ListInstalledPatchesSeverityEnum
func GetListInstalledPatchesSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingListInstalledPatchesSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstalledPatchesSeverityEnum(val string) (ListInstalledPatchesSeverityEnum, bool) {
	enum, ok := mappingListInstalledPatchesSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstalledPatchesSortOrderEnum Enum with underlying type: string
type ListInstalledPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListInstalledPatchesSortOrderEnum
const (
	ListInstalledPatchesSortOrderAsc  ListInstalledPatchesSortOrderEnum = "ASC"
	ListInstalledPatchesSortOrderDesc ListInstalledPatchesSortOrderEnum = "DESC"
)

var mappingListInstalledPatchesSortOrderEnum = map[string]ListInstalledPatchesSortOrderEnum{
	"ASC":  ListInstalledPatchesSortOrderAsc,
	"DESC": ListInstalledPatchesSortOrderDesc,
}

var mappingListInstalledPatchesSortOrderEnumLowerCase = map[string]ListInstalledPatchesSortOrderEnum{
	"asc":  ListInstalledPatchesSortOrderAsc,
	"desc": ListInstalledPatchesSortOrderDesc,
}

// GetListInstalledPatchesSortOrderEnumValues Enumerates the set of values for ListInstalledPatchesSortOrderEnum
func GetListInstalledPatchesSortOrderEnumValues() []ListInstalledPatchesSortOrderEnum {
	values := make([]ListInstalledPatchesSortOrderEnum, 0)
	for _, v := range mappingListInstalledPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstalledPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListInstalledPatchesSortOrderEnum
func GetListInstalledPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstalledPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstalledPatchesSortOrderEnum(val string) (ListInstalledPatchesSortOrderEnum, bool) {
	enum, ok := mappingListInstalledPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstalledPatchesSortByEnum Enum with underlying type: string
type ListInstalledPatchesSortByEnum string

// Set of constants representing the allowable values for ListInstalledPatchesSortByEnum
const (
	ListInstalledPatchesSortByPatchname ListInstalledPatchesSortByEnum = "patchName"
)

var mappingListInstalledPatchesSortByEnum = map[string]ListInstalledPatchesSortByEnum{
	"patchName": ListInstalledPatchesSortByPatchname,
}

var mappingListInstalledPatchesSortByEnumLowerCase = map[string]ListInstalledPatchesSortByEnum{
	"patchname": ListInstalledPatchesSortByPatchname,
}

// GetListInstalledPatchesSortByEnumValues Enumerates the set of values for ListInstalledPatchesSortByEnum
func GetListInstalledPatchesSortByEnumValues() []ListInstalledPatchesSortByEnum {
	values := make([]ListInstalledPatchesSortByEnum, 0)
	for _, v := range mappingListInstalledPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstalledPatchesSortByEnumStringValues Enumerates the set of values in String for ListInstalledPatchesSortByEnum
func GetListInstalledPatchesSortByEnumStringValues() []string {
	return []string{
		"patchName",
	}
}

// GetMappingListInstalledPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstalledPatchesSortByEnum(val string) (ListInstalledPatchesSortByEnum, bool) {
	enum, ok := mappingListInstalledPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
