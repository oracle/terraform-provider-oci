// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbSystemsRequest wrapper for the ListDbSystems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystems.go.html to see an example of how to use ListDbSystemsRequest.
type ListDbSystemsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup. Specify a backupId to list only the DB systems or DB homes that support creating a database using this backup in this compartment.
	BackupId *string `mandatory:"false" contributesTo:"query" name:"backupId"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DbSystemSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbSystemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbSystemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbSystemSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbSystemsResponse wrapper for the ListDbSystems operation
type ListDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbSystemSummary instances
	Items []DbSystemSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemsSortByEnum Enum with underlying type: string
type ListDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListDbSystemsSortByEnum
const (
	ListDbSystemsSortByTimecreated ListDbSystemsSortByEnum = "TIMECREATED"
	ListDbSystemsSortByDisplayname ListDbSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListDbSystemsSortByEnum = map[string]ListDbSystemsSortByEnum{
	"TIMECREATED": ListDbSystemsSortByTimecreated,
	"DISPLAYNAME": ListDbSystemsSortByDisplayname,
}

var mappingListDbSystemsSortByEnumLowerCase = map[string]ListDbSystemsSortByEnum{
	"timecreated": ListDbSystemsSortByTimecreated,
	"displayname": ListDbSystemsSortByDisplayname,
}

// GetListDbSystemsSortByEnumValues Enumerates the set of values for ListDbSystemsSortByEnum
func GetListDbSystemsSortByEnumValues() []ListDbSystemsSortByEnum {
	values := make([]ListDbSystemsSortByEnum, 0)
	for _, v := range mappingListDbSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemsSortByEnumStringValues Enumerates the set of values in String for ListDbSystemsSortByEnum
func GetListDbSystemsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDbSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemsSortByEnum(val string) (ListDbSystemsSortByEnum, bool) {
	enum, ok := mappingListDbSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbSystemsSortOrderEnum Enum with underlying type: string
type ListDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListDbSystemsSortOrderEnum
const (
	ListDbSystemsSortOrderAsc  ListDbSystemsSortOrderEnum = "ASC"
	ListDbSystemsSortOrderDesc ListDbSystemsSortOrderEnum = "DESC"
)

var mappingListDbSystemsSortOrderEnum = map[string]ListDbSystemsSortOrderEnum{
	"ASC":  ListDbSystemsSortOrderAsc,
	"DESC": ListDbSystemsSortOrderDesc,
}

var mappingListDbSystemsSortOrderEnumLowerCase = map[string]ListDbSystemsSortOrderEnum{
	"asc":  ListDbSystemsSortOrderAsc,
	"desc": ListDbSystemsSortOrderDesc,
}

// GetListDbSystemsSortOrderEnumValues Enumerates the set of values for ListDbSystemsSortOrderEnum
func GetListDbSystemsSortOrderEnumValues() []ListDbSystemsSortOrderEnum {
	values := make([]ListDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListDbSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListDbSystemsSortOrderEnum
func GetListDbSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemsSortOrderEnum(val string) (ListDbSystemsSortOrderEnum, bool) {
	enum, ok := mappingListDbSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
