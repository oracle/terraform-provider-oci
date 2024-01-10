// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEnterpriseManagerBridgesRequest wrapper for the ListEnterpriseManagerBridges operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListEnterpriseManagerBridges.go.html to see an example of how to use ListEnterpriseManagerBridgesRequest.
type ListEnterpriseManagerBridgesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Enterprise Manager bridge identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEnterpriseManagerBridgesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListEnterpriseManagerBridgesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnterpriseManagerBridgesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnterpriseManagerBridgesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnterpriseManagerBridgesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnterpriseManagerBridgesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEnterpriseManagerBridgesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListEnterpriseManagerBridgesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEnterpriseManagerBridgesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnterpriseManagerBridgesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEnterpriseManagerBridgesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEnterpriseManagerBridgesResponse wrapper for the ListEnterpriseManagerBridges operation
type ListEnterpriseManagerBridgesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnterpriseManagerBridgeCollection instances
	EnterpriseManagerBridgeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEnterpriseManagerBridgesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnterpriseManagerBridgesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnterpriseManagerBridgesSortOrderEnum Enum with underlying type: string
type ListEnterpriseManagerBridgesSortOrderEnum string

// Set of constants representing the allowable values for ListEnterpriseManagerBridgesSortOrderEnum
const (
	ListEnterpriseManagerBridgesSortOrderAsc  ListEnterpriseManagerBridgesSortOrderEnum = "ASC"
	ListEnterpriseManagerBridgesSortOrderDesc ListEnterpriseManagerBridgesSortOrderEnum = "DESC"
)

var mappingListEnterpriseManagerBridgesSortOrderEnum = map[string]ListEnterpriseManagerBridgesSortOrderEnum{
	"ASC":  ListEnterpriseManagerBridgesSortOrderAsc,
	"DESC": ListEnterpriseManagerBridgesSortOrderDesc,
}

var mappingListEnterpriseManagerBridgesSortOrderEnumLowerCase = map[string]ListEnterpriseManagerBridgesSortOrderEnum{
	"asc":  ListEnterpriseManagerBridgesSortOrderAsc,
	"desc": ListEnterpriseManagerBridgesSortOrderDesc,
}

// GetListEnterpriseManagerBridgesSortOrderEnumValues Enumerates the set of values for ListEnterpriseManagerBridgesSortOrderEnum
func GetListEnterpriseManagerBridgesSortOrderEnumValues() []ListEnterpriseManagerBridgesSortOrderEnum {
	values := make([]ListEnterpriseManagerBridgesSortOrderEnum, 0)
	for _, v := range mappingListEnterpriseManagerBridgesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnterpriseManagerBridgesSortOrderEnumStringValues Enumerates the set of values in String for ListEnterpriseManagerBridgesSortOrderEnum
func GetListEnterpriseManagerBridgesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEnterpriseManagerBridgesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnterpriseManagerBridgesSortOrderEnum(val string) (ListEnterpriseManagerBridgesSortOrderEnum, bool) {
	enum, ok := mappingListEnterpriseManagerBridgesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnterpriseManagerBridgesSortByEnum Enum with underlying type: string
type ListEnterpriseManagerBridgesSortByEnum string

// Set of constants representing the allowable values for ListEnterpriseManagerBridgesSortByEnum
const (
	ListEnterpriseManagerBridgesSortByTimecreated ListEnterpriseManagerBridgesSortByEnum = "timeCreated"
	ListEnterpriseManagerBridgesSortByDisplayname ListEnterpriseManagerBridgesSortByEnum = "displayName"
)

var mappingListEnterpriseManagerBridgesSortByEnum = map[string]ListEnterpriseManagerBridgesSortByEnum{
	"timeCreated": ListEnterpriseManagerBridgesSortByTimecreated,
	"displayName": ListEnterpriseManagerBridgesSortByDisplayname,
}

var mappingListEnterpriseManagerBridgesSortByEnumLowerCase = map[string]ListEnterpriseManagerBridgesSortByEnum{
	"timecreated": ListEnterpriseManagerBridgesSortByTimecreated,
	"displayname": ListEnterpriseManagerBridgesSortByDisplayname,
}

// GetListEnterpriseManagerBridgesSortByEnumValues Enumerates the set of values for ListEnterpriseManagerBridgesSortByEnum
func GetListEnterpriseManagerBridgesSortByEnumValues() []ListEnterpriseManagerBridgesSortByEnum {
	values := make([]ListEnterpriseManagerBridgesSortByEnum, 0)
	for _, v := range mappingListEnterpriseManagerBridgesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnterpriseManagerBridgesSortByEnumStringValues Enumerates the set of values in String for ListEnterpriseManagerBridgesSortByEnum
func GetListEnterpriseManagerBridgesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListEnterpriseManagerBridgesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnterpriseManagerBridgesSortByEnum(val string) (ListEnterpriseManagerBridgesSortByEnum, bool) {
	enum, ok := mappingListEnterpriseManagerBridgesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
