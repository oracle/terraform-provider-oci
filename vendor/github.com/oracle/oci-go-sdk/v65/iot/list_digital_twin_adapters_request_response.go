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

// ListDigitalTwinAdaptersRequest wrapper for the ListDigitalTwinAdapters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinAdapters.go.html to see an example of how to use ListDigitalTwinAdaptersRequest.
type ListDigitalTwinAdaptersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources.
	IotDomainId *string `mandatory:"true" contributesTo:"query" name:"iotDomainId"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter resources that match the specified URI (DTMI) of the digital twin model.
	DigitalTwinModelSpecUri *string `mandatory:"false" contributesTo:"query" name:"digitalTwinModelSpecUri"`

	// Filter resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model.
	DigitalTwinModelId *string `mandatory:"false" contributesTo:"query" name:"digitalTwinModelId"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState ListDigitalTwinAdaptersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Page representing the requested page of items.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListDigitalTwinAdaptersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDigitalTwinAdaptersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalTwinAdaptersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalTwinAdaptersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalTwinAdaptersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalTwinAdaptersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalTwinAdaptersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalTwinAdaptersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalTwinAdaptersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinAdaptersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalTwinAdaptersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinAdaptersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalTwinAdaptersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalTwinAdaptersResponse wrapper for the ListDigitalTwinAdapters operation
type ListDigitalTwinAdaptersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalTwinAdapterCollection instances
	DigitalTwinAdapterCollection `presentIn:"body"`

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

func (response ListDigitalTwinAdaptersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalTwinAdaptersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalTwinAdaptersLifecycleStateEnum Enum with underlying type: string
type ListDigitalTwinAdaptersLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalTwinAdaptersLifecycleStateEnum
const (
	ListDigitalTwinAdaptersLifecycleStateActive  ListDigitalTwinAdaptersLifecycleStateEnum = "ACTIVE"
	ListDigitalTwinAdaptersLifecycleStateDeleted ListDigitalTwinAdaptersLifecycleStateEnum = "DELETED"
)

var mappingListDigitalTwinAdaptersLifecycleStateEnum = map[string]ListDigitalTwinAdaptersLifecycleStateEnum{
	"ACTIVE":  ListDigitalTwinAdaptersLifecycleStateActive,
	"DELETED": ListDigitalTwinAdaptersLifecycleStateDeleted,
}

var mappingListDigitalTwinAdaptersLifecycleStateEnumLowerCase = map[string]ListDigitalTwinAdaptersLifecycleStateEnum{
	"active":  ListDigitalTwinAdaptersLifecycleStateActive,
	"deleted": ListDigitalTwinAdaptersLifecycleStateDeleted,
}

// GetListDigitalTwinAdaptersLifecycleStateEnumValues Enumerates the set of values for ListDigitalTwinAdaptersLifecycleStateEnum
func GetListDigitalTwinAdaptersLifecycleStateEnumValues() []ListDigitalTwinAdaptersLifecycleStateEnum {
	values := make([]ListDigitalTwinAdaptersLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalTwinAdaptersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinAdaptersLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalTwinAdaptersLifecycleStateEnum
func GetListDigitalTwinAdaptersLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDigitalTwinAdaptersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinAdaptersLifecycleStateEnum(val string) (ListDigitalTwinAdaptersLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalTwinAdaptersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinAdaptersSortOrderEnum Enum with underlying type: string
type ListDigitalTwinAdaptersSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalTwinAdaptersSortOrderEnum
const (
	ListDigitalTwinAdaptersSortOrderAsc  ListDigitalTwinAdaptersSortOrderEnum = "ASC"
	ListDigitalTwinAdaptersSortOrderDesc ListDigitalTwinAdaptersSortOrderEnum = "DESC"
)

var mappingListDigitalTwinAdaptersSortOrderEnum = map[string]ListDigitalTwinAdaptersSortOrderEnum{
	"ASC":  ListDigitalTwinAdaptersSortOrderAsc,
	"DESC": ListDigitalTwinAdaptersSortOrderDesc,
}

var mappingListDigitalTwinAdaptersSortOrderEnumLowerCase = map[string]ListDigitalTwinAdaptersSortOrderEnum{
	"asc":  ListDigitalTwinAdaptersSortOrderAsc,
	"desc": ListDigitalTwinAdaptersSortOrderDesc,
}

// GetListDigitalTwinAdaptersSortOrderEnumValues Enumerates the set of values for ListDigitalTwinAdaptersSortOrderEnum
func GetListDigitalTwinAdaptersSortOrderEnumValues() []ListDigitalTwinAdaptersSortOrderEnum {
	values := make([]ListDigitalTwinAdaptersSortOrderEnum, 0)
	for _, v := range mappingListDigitalTwinAdaptersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinAdaptersSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalTwinAdaptersSortOrderEnum
func GetListDigitalTwinAdaptersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalTwinAdaptersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinAdaptersSortOrderEnum(val string) (ListDigitalTwinAdaptersSortOrderEnum, bool) {
	enum, ok := mappingListDigitalTwinAdaptersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinAdaptersSortByEnum Enum with underlying type: string
type ListDigitalTwinAdaptersSortByEnum string

// Set of constants representing the allowable values for ListDigitalTwinAdaptersSortByEnum
const (
	ListDigitalTwinAdaptersSortByTimecreated ListDigitalTwinAdaptersSortByEnum = "timeCreated"
	ListDigitalTwinAdaptersSortByDisplayname ListDigitalTwinAdaptersSortByEnum = "displayName"
)

var mappingListDigitalTwinAdaptersSortByEnum = map[string]ListDigitalTwinAdaptersSortByEnum{
	"timeCreated": ListDigitalTwinAdaptersSortByTimecreated,
	"displayName": ListDigitalTwinAdaptersSortByDisplayname,
}

var mappingListDigitalTwinAdaptersSortByEnumLowerCase = map[string]ListDigitalTwinAdaptersSortByEnum{
	"timecreated": ListDigitalTwinAdaptersSortByTimecreated,
	"displayname": ListDigitalTwinAdaptersSortByDisplayname,
}

// GetListDigitalTwinAdaptersSortByEnumValues Enumerates the set of values for ListDigitalTwinAdaptersSortByEnum
func GetListDigitalTwinAdaptersSortByEnumValues() []ListDigitalTwinAdaptersSortByEnum {
	values := make([]ListDigitalTwinAdaptersSortByEnum, 0)
	for _, v := range mappingListDigitalTwinAdaptersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinAdaptersSortByEnumStringValues Enumerates the set of values in String for ListDigitalTwinAdaptersSortByEnum
func GetListDigitalTwinAdaptersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDigitalTwinAdaptersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinAdaptersSortByEnum(val string) (ListDigitalTwinAdaptersSortByEnum, bool) {
	enum, ok := mappingListDigitalTwinAdaptersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
