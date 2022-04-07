// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbServersRequest wrapper for the ListDbServers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbServers.go.html to see an example of how to use ListDbServersRequest.
type ListDbServersRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ExadataInfrastructure.
	ExadataInfrastructureId *string `mandatory:"true" contributesTo:"query" name:"exadataInfrastructureId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbServersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Sort by TIMECREATED.  Default order for TIMECREATED is descending.
	SortBy ListDbServersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DbServerSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbServersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbServersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbServersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbServersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbServersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbServersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbServersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbServersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbServersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbServerSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbServerSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbServersResponse wrapper for the ListDbServers operation
type ListDbServersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbServerSummary instances
	Items []DbServerSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbServersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbServersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbServersSortOrderEnum Enum with underlying type: string
type ListDbServersSortOrderEnum string

// Set of constants representing the allowable values for ListDbServersSortOrderEnum
const (
	ListDbServersSortOrderAsc  ListDbServersSortOrderEnum = "ASC"
	ListDbServersSortOrderDesc ListDbServersSortOrderEnum = "DESC"
)

var mappingListDbServersSortOrderEnum = map[string]ListDbServersSortOrderEnum{
	"ASC":  ListDbServersSortOrderAsc,
	"DESC": ListDbServersSortOrderDesc,
}

var mappingListDbServersSortOrderEnumLowerCase = map[string]ListDbServersSortOrderEnum{
	"asc":  ListDbServersSortOrderAsc,
	"desc": ListDbServersSortOrderDesc,
}

// GetListDbServersSortOrderEnumValues Enumerates the set of values for ListDbServersSortOrderEnum
func GetListDbServersSortOrderEnumValues() []ListDbServersSortOrderEnum {
	values := make([]ListDbServersSortOrderEnum, 0)
	for _, v := range mappingListDbServersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbServersSortOrderEnumStringValues Enumerates the set of values in String for ListDbServersSortOrderEnum
func GetListDbServersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbServersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbServersSortOrderEnum(val string) (ListDbServersSortOrderEnum, bool) {
	enum, ok := mappingListDbServersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbServersSortByEnum Enum with underlying type: string
type ListDbServersSortByEnum string

// Set of constants representing the allowable values for ListDbServersSortByEnum
const (
	ListDbServersSortByTimecreated ListDbServersSortByEnum = "TIMECREATED"
)

var mappingListDbServersSortByEnum = map[string]ListDbServersSortByEnum{
	"TIMECREATED": ListDbServersSortByTimecreated,
}

var mappingListDbServersSortByEnumLowerCase = map[string]ListDbServersSortByEnum{
	"timecreated": ListDbServersSortByTimecreated,
}

// GetListDbServersSortByEnumValues Enumerates the set of values for ListDbServersSortByEnum
func GetListDbServersSortByEnumValues() []ListDbServersSortByEnum {
	values := make([]ListDbServersSortByEnum, 0)
	for _, v := range mappingListDbServersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbServersSortByEnumStringValues Enumerates the set of values in String for ListDbServersSortByEnum
func GetListDbServersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListDbServersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbServersSortByEnum(val string) (ListDbServersSortByEnum, bool) {
	enum, ok := mappingListDbServersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
