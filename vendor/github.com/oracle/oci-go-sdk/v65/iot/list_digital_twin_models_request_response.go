// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDigitalTwinModelsRequest wrapper for the ListDigitalTwinModels operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinModels.go.html to see an example of how to use ListDigitalTwinModelsRequest.
type ListDigitalTwinModelsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources.
	IotDomainId *string `mandatory:"true" contributesTo:"query" name:"iotDomainId"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filters resources by spec URI prefix. For example, to search all versions of the `dtmi:example:device;1` model, pass the prefix without the version: `dtmi:example:device`.
	SpecUriStartsWith *string `mandatory:"false" contributesTo:"query" name:"specUriStartsWith"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Page representing the requested page of items.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState ListDigitalTwinModelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListDigitalTwinModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDigitalTwinModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalTwinModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalTwinModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalTwinModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalTwinModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalTwinModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalTwinModelsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalTwinModelsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalTwinModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalTwinModelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalTwinModelsResponse wrapper for the ListDigitalTwinModels operation
type ListDigitalTwinModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalTwinModelCollection instances
	DigitalTwinModelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDigitalTwinModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalTwinModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalTwinModelsLifecycleStateEnum Enum with underlying type: string
type ListDigitalTwinModelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalTwinModelsLifecycleStateEnum
const (
	ListDigitalTwinModelsLifecycleStateActive  ListDigitalTwinModelsLifecycleStateEnum = "ACTIVE"
	ListDigitalTwinModelsLifecycleStateDeleted ListDigitalTwinModelsLifecycleStateEnum = "DELETED"
)

var mappingListDigitalTwinModelsLifecycleStateEnum = map[string]ListDigitalTwinModelsLifecycleStateEnum{
	"ACTIVE":  ListDigitalTwinModelsLifecycleStateActive,
	"DELETED": ListDigitalTwinModelsLifecycleStateDeleted,
}

var mappingListDigitalTwinModelsLifecycleStateEnumLowerCase = map[string]ListDigitalTwinModelsLifecycleStateEnum{
	"active":  ListDigitalTwinModelsLifecycleStateActive,
	"deleted": ListDigitalTwinModelsLifecycleStateDeleted,
}

// GetListDigitalTwinModelsLifecycleStateEnumValues Enumerates the set of values for ListDigitalTwinModelsLifecycleStateEnum
func GetListDigitalTwinModelsLifecycleStateEnumValues() []ListDigitalTwinModelsLifecycleStateEnum {
	values := make([]ListDigitalTwinModelsLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalTwinModelsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinModelsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalTwinModelsLifecycleStateEnum
func GetListDigitalTwinModelsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDigitalTwinModelsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinModelsLifecycleStateEnum(val string) (ListDigitalTwinModelsLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalTwinModelsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinModelsSortOrderEnum Enum with underlying type: string
type ListDigitalTwinModelsSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalTwinModelsSortOrderEnum
const (
	ListDigitalTwinModelsSortOrderAsc  ListDigitalTwinModelsSortOrderEnum = "ASC"
	ListDigitalTwinModelsSortOrderDesc ListDigitalTwinModelsSortOrderEnum = "DESC"
)

var mappingListDigitalTwinModelsSortOrderEnum = map[string]ListDigitalTwinModelsSortOrderEnum{
	"ASC":  ListDigitalTwinModelsSortOrderAsc,
	"DESC": ListDigitalTwinModelsSortOrderDesc,
}

var mappingListDigitalTwinModelsSortOrderEnumLowerCase = map[string]ListDigitalTwinModelsSortOrderEnum{
	"asc":  ListDigitalTwinModelsSortOrderAsc,
	"desc": ListDigitalTwinModelsSortOrderDesc,
}

// GetListDigitalTwinModelsSortOrderEnumValues Enumerates the set of values for ListDigitalTwinModelsSortOrderEnum
func GetListDigitalTwinModelsSortOrderEnumValues() []ListDigitalTwinModelsSortOrderEnum {
	values := make([]ListDigitalTwinModelsSortOrderEnum, 0)
	for _, v := range mappingListDigitalTwinModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinModelsSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalTwinModelsSortOrderEnum
func GetListDigitalTwinModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalTwinModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinModelsSortOrderEnum(val string) (ListDigitalTwinModelsSortOrderEnum, bool) {
	enum, ok := mappingListDigitalTwinModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinModelsSortByEnum Enum with underlying type: string
type ListDigitalTwinModelsSortByEnum string

// Set of constants representing the allowable values for ListDigitalTwinModelsSortByEnum
const (
	ListDigitalTwinModelsSortByTimecreated ListDigitalTwinModelsSortByEnum = "timeCreated"
	ListDigitalTwinModelsSortByDisplayname ListDigitalTwinModelsSortByEnum = "displayName"
)

var mappingListDigitalTwinModelsSortByEnum = map[string]ListDigitalTwinModelsSortByEnum{
	"timeCreated": ListDigitalTwinModelsSortByTimecreated,
	"displayName": ListDigitalTwinModelsSortByDisplayname,
}

var mappingListDigitalTwinModelsSortByEnumLowerCase = map[string]ListDigitalTwinModelsSortByEnum{
	"timecreated": ListDigitalTwinModelsSortByTimecreated,
	"displayname": ListDigitalTwinModelsSortByDisplayname,
}

// GetListDigitalTwinModelsSortByEnumValues Enumerates the set of values for ListDigitalTwinModelsSortByEnum
func GetListDigitalTwinModelsSortByEnumValues() []ListDigitalTwinModelsSortByEnum {
	values := make([]ListDigitalTwinModelsSortByEnum, 0)
	for _, v := range mappingListDigitalTwinModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinModelsSortByEnumStringValues Enumerates the set of values in String for ListDigitalTwinModelsSortByEnum
func GetListDigitalTwinModelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDigitalTwinModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinModelsSortByEnum(val string) (ListDigitalTwinModelsSortByEnum, bool) {
	enum, ok := mappingListDigitalTwinModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
