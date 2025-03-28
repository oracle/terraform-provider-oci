// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package autoscaling

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutoScalingConfigurationsRequest wrapper for the ListAutoScalingConfigurations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/autoscaling/ListAutoScalingConfigurations.go.html to see an example of how to use ListAutoScalingConfigurationsRequest.
type ListAutoScalingConfigurationsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the
	// resource. Use tenancyId to search in
	// the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return in a paginated "List" call. For important details
	// about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	SortBy ListAutoScalingConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListAutoScalingConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutoScalingConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutoScalingConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutoScalingConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutoScalingConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutoScalingConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutoScalingConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutoScalingConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutoScalingConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutoScalingConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutoScalingConfigurationsResponse wrapper for the ListAutoScalingConfigurations operation
type ListAutoScalingConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutoScalingConfigurationSummary instances
	Items []AutoScalingConfigurationSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAutoScalingConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutoScalingConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutoScalingConfigurationsSortByEnum Enum with underlying type: string
type ListAutoScalingConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListAutoScalingConfigurationsSortByEnum
const (
	ListAutoScalingConfigurationsSortByTimecreated ListAutoScalingConfigurationsSortByEnum = "TIMECREATED"
	ListAutoScalingConfigurationsSortByDisplayname ListAutoScalingConfigurationsSortByEnum = "DISPLAYNAME"
)

var mappingListAutoScalingConfigurationsSortByEnum = map[string]ListAutoScalingConfigurationsSortByEnum{
	"TIMECREATED": ListAutoScalingConfigurationsSortByTimecreated,
	"DISPLAYNAME": ListAutoScalingConfigurationsSortByDisplayname,
}

var mappingListAutoScalingConfigurationsSortByEnumLowerCase = map[string]ListAutoScalingConfigurationsSortByEnum{
	"timecreated": ListAutoScalingConfigurationsSortByTimecreated,
	"displayname": ListAutoScalingConfigurationsSortByDisplayname,
}

// GetListAutoScalingConfigurationsSortByEnumValues Enumerates the set of values for ListAutoScalingConfigurationsSortByEnum
func GetListAutoScalingConfigurationsSortByEnumValues() []ListAutoScalingConfigurationsSortByEnum {
	values := make([]ListAutoScalingConfigurationsSortByEnum, 0)
	for _, v := range mappingListAutoScalingConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoScalingConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListAutoScalingConfigurationsSortByEnum
func GetListAutoScalingConfigurationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutoScalingConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoScalingConfigurationsSortByEnum(val string) (ListAutoScalingConfigurationsSortByEnum, bool) {
	enum, ok := mappingListAutoScalingConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutoScalingConfigurationsSortOrderEnum Enum with underlying type: string
type ListAutoScalingConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListAutoScalingConfigurationsSortOrderEnum
const (
	ListAutoScalingConfigurationsSortOrderAsc  ListAutoScalingConfigurationsSortOrderEnum = "ASC"
	ListAutoScalingConfigurationsSortOrderDesc ListAutoScalingConfigurationsSortOrderEnum = "DESC"
)

var mappingListAutoScalingConfigurationsSortOrderEnum = map[string]ListAutoScalingConfigurationsSortOrderEnum{
	"ASC":  ListAutoScalingConfigurationsSortOrderAsc,
	"DESC": ListAutoScalingConfigurationsSortOrderDesc,
}

var mappingListAutoScalingConfigurationsSortOrderEnumLowerCase = map[string]ListAutoScalingConfigurationsSortOrderEnum{
	"asc":  ListAutoScalingConfigurationsSortOrderAsc,
	"desc": ListAutoScalingConfigurationsSortOrderDesc,
}

// GetListAutoScalingConfigurationsSortOrderEnumValues Enumerates the set of values for ListAutoScalingConfigurationsSortOrderEnum
func GetListAutoScalingConfigurationsSortOrderEnumValues() []ListAutoScalingConfigurationsSortOrderEnum {
	values := make([]ListAutoScalingConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListAutoScalingConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoScalingConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListAutoScalingConfigurationsSortOrderEnum
func GetListAutoScalingConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutoScalingConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoScalingConfigurationsSortOrderEnum(val string) (ListAutoScalingConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListAutoScalingConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
