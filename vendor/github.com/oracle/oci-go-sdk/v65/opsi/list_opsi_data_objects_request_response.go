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

// ListOpsiDataObjectsRequest wrapper for the ListOpsiDataObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOpsiDataObjects.go.html to see an example of how to use ListOpsiDataObjectsRequest.
type ListOpsiDataObjectsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OPSI data object types.
	DataObjectType []OpsiDataObjectTypeEnum `contributesTo:"query" name:"dataObjectType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

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
	SortOrder ListOpsiDataObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// OPSI data object list sort options.
	SortBy ListOpsiDataObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only data objects that belongs to the group of the given group name. By default, no filtering will be applied on group name.
	GroupName *string `mandatory:"false" contributesTo:"query" name:"groupName"`

	// A filter to return only data objects that match the entire data object name. By default, no filtering will be applied on data object name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpsiDataObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpsiDataObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpsiDataObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpsiDataObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpsiDataObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DataObjectType {
		if _, ok := GetMappingOpsiDataObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataObjectType: %s. Supported values are: %s.", val, strings.Join(GetOpsiDataObjectTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListOpsiDataObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpsiDataObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpsiDataObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpsiDataObjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpsiDataObjectsResponse wrapper for the ListOpsiDataObjects operation
type ListOpsiDataObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpsiDataObjectsCollection instances
	OpsiDataObjectsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpsiDataObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpsiDataObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpsiDataObjectsSortOrderEnum Enum with underlying type: string
type ListOpsiDataObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListOpsiDataObjectsSortOrderEnum
const (
	ListOpsiDataObjectsSortOrderAsc  ListOpsiDataObjectsSortOrderEnum = "ASC"
	ListOpsiDataObjectsSortOrderDesc ListOpsiDataObjectsSortOrderEnum = "DESC"
)

var mappingListOpsiDataObjectsSortOrderEnum = map[string]ListOpsiDataObjectsSortOrderEnum{
	"ASC":  ListOpsiDataObjectsSortOrderAsc,
	"DESC": ListOpsiDataObjectsSortOrderDesc,
}

var mappingListOpsiDataObjectsSortOrderEnumLowerCase = map[string]ListOpsiDataObjectsSortOrderEnum{
	"asc":  ListOpsiDataObjectsSortOrderAsc,
	"desc": ListOpsiDataObjectsSortOrderDesc,
}

// GetListOpsiDataObjectsSortOrderEnumValues Enumerates the set of values for ListOpsiDataObjectsSortOrderEnum
func GetListOpsiDataObjectsSortOrderEnumValues() []ListOpsiDataObjectsSortOrderEnum {
	values := make([]ListOpsiDataObjectsSortOrderEnum, 0)
	for _, v := range mappingListOpsiDataObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiDataObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListOpsiDataObjectsSortOrderEnum
func GetListOpsiDataObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpsiDataObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiDataObjectsSortOrderEnum(val string) (ListOpsiDataObjectsSortOrderEnum, bool) {
	enum, ok := mappingListOpsiDataObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpsiDataObjectsSortByEnum Enum with underlying type: string
type ListOpsiDataObjectsSortByEnum string

// Set of constants representing the allowable values for ListOpsiDataObjectsSortByEnum
const (
	ListOpsiDataObjectsSortByDisplayname    ListOpsiDataObjectsSortByEnum = "displayName"
	ListOpsiDataObjectsSortByDataobjecttype ListOpsiDataObjectsSortByEnum = "dataObjectType"
	ListOpsiDataObjectsSortByName           ListOpsiDataObjectsSortByEnum = "name"
)

var mappingListOpsiDataObjectsSortByEnum = map[string]ListOpsiDataObjectsSortByEnum{
	"displayName":    ListOpsiDataObjectsSortByDisplayname,
	"dataObjectType": ListOpsiDataObjectsSortByDataobjecttype,
	"name":           ListOpsiDataObjectsSortByName,
}

var mappingListOpsiDataObjectsSortByEnumLowerCase = map[string]ListOpsiDataObjectsSortByEnum{
	"displayname":    ListOpsiDataObjectsSortByDisplayname,
	"dataobjecttype": ListOpsiDataObjectsSortByDataobjecttype,
	"name":           ListOpsiDataObjectsSortByName,
}

// GetListOpsiDataObjectsSortByEnumValues Enumerates the set of values for ListOpsiDataObjectsSortByEnum
func GetListOpsiDataObjectsSortByEnumValues() []ListOpsiDataObjectsSortByEnum {
	values := make([]ListOpsiDataObjectsSortByEnum, 0)
	for _, v := range mappingListOpsiDataObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiDataObjectsSortByEnumStringValues Enumerates the set of values in String for ListOpsiDataObjectsSortByEnum
func GetListOpsiDataObjectsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"dataObjectType",
		"name",
	}
}

// GetMappingListOpsiDataObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiDataObjectsSortByEnum(val string) (ListOpsiDataObjectsSortByEnum, bool) {
	enum, ok := mappingListOpsiDataObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
