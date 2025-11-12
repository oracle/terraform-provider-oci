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

// ListInventoryRecordsRequest wrapper for the ListInventoryRecords operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListInventoryRecords.go.html to see an example of how to use ListInventoryRecordsRequest.
type ListInventoryRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// If set to true, resources will be returned for not only the provided compartment, but all compartments which
	// descend from it. Which resources are returned and their field contents depends on the value of accessLevel.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// unique Fleet identifier
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// Resource Identifier
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInventoryRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If set to true, inventory details will be returned.
	IsDetailsRequired *bool `mandatory:"false" contributesTo:"query" name:"isDetailsRequired"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListInventoryRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInventoryRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInventoryRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInventoryRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInventoryRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInventoryRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInventoryRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInventoryRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInventoryRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInventoryRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInventoryRecordsResponse wrapper for the ListInventoryRecords operation
type ListInventoryRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InventoryRecordCollection instances
	InventoryRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInventoryRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInventoryRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInventoryRecordsSortOrderEnum Enum with underlying type: string
type ListInventoryRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListInventoryRecordsSortOrderEnum
const (
	ListInventoryRecordsSortOrderAsc  ListInventoryRecordsSortOrderEnum = "ASC"
	ListInventoryRecordsSortOrderDesc ListInventoryRecordsSortOrderEnum = "DESC"
)

var mappingListInventoryRecordsSortOrderEnum = map[string]ListInventoryRecordsSortOrderEnum{
	"ASC":  ListInventoryRecordsSortOrderAsc,
	"DESC": ListInventoryRecordsSortOrderDesc,
}

var mappingListInventoryRecordsSortOrderEnumLowerCase = map[string]ListInventoryRecordsSortOrderEnum{
	"asc":  ListInventoryRecordsSortOrderAsc,
	"desc": ListInventoryRecordsSortOrderDesc,
}

// GetListInventoryRecordsSortOrderEnumValues Enumerates the set of values for ListInventoryRecordsSortOrderEnum
func GetListInventoryRecordsSortOrderEnumValues() []ListInventoryRecordsSortOrderEnum {
	values := make([]ListInventoryRecordsSortOrderEnum, 0)
	for _, v := range mappingListInventoryRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoryRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListInventoryRecordsSortOrderEnum
func GetListInventoryRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInventoryRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoryRecordsSortOrderEnum(val string) (ListInventoryRecordsSortOrderEnum, bool) {
	enum, ok := mappingListInventoryRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInventoryRecordsSortByEnum Enum with underlying type: string
type ListInventoryRecordsSortByEnum string

// Set of constants representing the allowable values for ListInventoryRecordsSortByEnum
const (
	ListInventoryRecordsSortByTargetname        ListInventoryRecordsSortByEnum = "targetName"
	ListInventoryRecordsSortByTargetproductname ListInventoryRecordsSortByEnum = "targetProductName"
	ListInventoryRecordsSortByTargetresourceid  ListInventoryRecordsSortByEnum = "targetResourceId"
	ListInventoryRecordsSortByOstype            ListInventoryRecordsSortByEnum = "osType"
	ListInventoryRecordsSortByArchitecture      ListInventoryRecordsSortByEnum = "architecture"
)

var mappingListInventoryRecordsSortByEnum = map[string]ListInventoryRecordsSortByEnum{
	"targetName":        ListInventoryRecordsSortByTargetname,
	"targetProductName": ListInventoryRecordsSortByTargetproductname,
	"targetResourceId":  ListInventoryRecordsSortByTargetresourceid,
	"osType":            ListInventoryRecordsSortByOstype,
	"architecture":      ListInventoryRecordsSortByArchitecture,
}

var mappingListInventoryRecordsSortByEnumLowerCase = map[string]ListInventoryRecordsSortByEnum{
	"targetname":        ListInventoryRecordsSortByTargetname,
	"targetproductname": ListInventoryRecordsSortByTargetproductname,
	"targetresourceid":  ListInventoryRecordsSortByTargetresourceid,
	"ostype":            ListInventoryRecordsSortByOstype,
	"architecture":      ListInventoryRecordsSortByArchitecture,
}

// GetListInventoryRecordsSortByEnumValues Enumerates the set of values for ListInventoryRecordsSortByEnum
func GetListInventoryRecordsSortByEnumValues() []ListInventoryRecordsSortByEnum {
	values := make([]ListInventoryRecordsSortByEnum, 0)
	for _, v := range mappingListInventoryRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoryRecordsSortByEnumStringValues Enumerates the set of values in String for ListInventoryRecordsSortByEnum
func GetListInventoryRecordsSortByEnumStringValues() []string {
	return []string{
		"targetName",
		"targetProductName",
		"targetResourceId",
		"osType",
		"architecture",
	}
}

// GetMappingListInventoryRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoryRecordsSortByEnum(val string) (ListInventoryRecordsSortByEnum, bool) {
	enum, ok := mappingListInventoryRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
