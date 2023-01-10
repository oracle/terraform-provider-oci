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

// ListDbHomesRequest wrapper for the ListDbHomes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbHomes.go.html to see an example of how to use ListDbHomesRequest.
type ListDbHomesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The DB system OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" contributesTo:"query" name:"vmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup. Specify a backupId to list only the DB systems or DB homes that support creating a database using this backup in this compartment.
	BackupId *string `mandatory:"false" contributesTo:"query" name:"backupId"`

	// A filter to return only DB Homes that match the specified dbVersion.
	DbVersion *string `mandatory:"false" contributesTo:"query" name:"dbVersion"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListDbHomesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbHomesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DbHomeSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbHomesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbHomesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbHomesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbHomesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbHomesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbHomesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbHomesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbHomesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbHomesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbHomeSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbHomeSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbHomesResponse wrapper for the ListDbHomes operation
type ListDbHomesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbHomeSummary instances
	Items []DbHomeSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbHomesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbHomesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbHomesSortByEnum Enum with underlying type: string
type ListDbHomesSortByEnum string

// Set of constants representing the allowable values for ListDbHomesSortByEnum
const (
	ListDbHomesSortByTimecreated ListDbHomesSortByEnum = "TIMECREATED"
	ListDbHomesSortByDisplayname ListDbHomesSortByEnum = "DISPLAYNAME"
)

var mappingListDbHomesSortByEnum = map[string]ListDbHomesSortByEnum{
	"TIMECREATED": ListDbHomesSortByTimecreated,
	"DISPLAYNAME": ListDbHomesSortByDisplayname,
}

var mappingListDbHomesSortByEnumLowerCase = map[string]ListDbHomesSortByEnum{
	"timecreated": ListDbHomesSortByTimecreated,
	"displayname": ListDbHomesSortByDisplayname,
}

// GetListDbHomesSortByEnumValues Enumerates the set of values for ListDbHomesSortByEnum
func GetListDbHomesSortByEnumValues() []ListDbHomesSortByEnum {
	values := make([]ListDbHomesSortByEnum, 0)
	for _, v := range mappingListDbHomesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbHomesSortByEnumStringValues Enumerates the set of values in String for ListDbHomesSortByEnum
func GetListDbHomesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDbHomesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbHomesSortByEnum(val string) (ListDbHomesSortByEnum, bool) {
	enum, ok := mappingListDbHomesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbHomesSortOrderEnum Enum with underlying type: string
type ListDbHomesSortOrderEnum string

// Set of constants representing the allowable values for ListDbHomesSortOrderEnum
const (
	ListDbHomesSortOrderAsc  ListDbHomesSortOrderEnum = "ASC"
	ListDbHomesSortOrderDesc ListDbHomesSortOrderEnum = "DESC"
)

var mappingListDbHomesSortOrderEnum = map[string]ListDbHomesSortOrderEnum{
	"ASC":  ListDbHomesSortOrderAsc,
	"DESC": ListDbHomesSortOrderDesc,
}

var mappingListDbHomesSortOrderEnumLowerCase = map[string]ListDbHomesSortOrderEnum{
	"asc":  ListDbHomesSortOrderAsc,
	"desc": ListDbHomesSortOrderDesc,
}

// GetListDbHomesSortOrderEnumValues Enumerates the set of values for ListDbHomesSortOrderEnum
func GetListDbHomesSortOrderEnumValues() []ListDbHomesSortOrderEnum {
	values := make([]ListDbHomesSortOrderEnum, 0)
	for _, v := range mappingListDbHomesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbHomesSortOrderEnumStringValues Enumerates the set of values in String for ListDbHomesSortOrderEnum
func GetListDbHomesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbHomesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbHomesSortOrderEnum(val string) (ListDbHomesSortOrderEnum, bool) {
	enum, ok := mappingListDbHomesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
