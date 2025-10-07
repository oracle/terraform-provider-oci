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

// ListDigitalTwinInstancesRequest wrapper for the ListDigitalTwinInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinInstances.go.html to see an example of how to use ListDigitalTwinInstancesRequest.
type ListDigitalTwinInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources.
	IotDomainId *string `mandatory:"true" contributesTo:"query" name:"iotDomainId"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Page representing the requested page of items.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState ListDigitalTwinInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListDigitalTwinInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDigitalTwinInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model.
	DigitalTwinModelId *string `mandatory:"false" contributesTo:"query" name:"digitalTwinModelId"`

	// Filter resources that match the specified URI (DTMI) of the digital twin model.
	DigitalTwinModelSpecUri *string `mandatory:"false" contributesTo:"query" name:"digitalTwinModelSpecUri"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalTwinInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalTwinInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalTwinInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalTwinInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalTwinInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalTwinInstancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalTwinInstancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalTwinInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalTwinInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalTwinInstancesResponse wrapper for the ListDigitalTwinInstances operation
type ListDigitalTwinInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalTwinInstanceCollection instances
	DigitalTwinInstanceCollection `presentIn:"body"`

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

func (response ListDigitalTwinInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalTwinInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalTwinInstancesLifecycleStateEnum Enum with underlying type: string
type ListDigitalTwinInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalTwinInstancesLifecycleStateEnum
const (
	ListDigitalTwinInstancesLifecycleStateActive  ListDigitalTwinInstancesLifecycleStateEnum = "ACTIVE"
	ListDigitalTwinInstancesLifecycleStateDeleted ListDigitalTwinInstancesLifecycleStateEnum = "DELETED"
)

var mappingListDigitalTwinInstancesLifecycleStateEnum = map[string]ListDigitalTwinInstancesLifecycleStateEnum{
	"ACTIVE":  ListDigitalTwinInstancesLifecycleStateActive,
	"DELETED": ListDigitalTwinInstancesLifecycleStateDeleted,
}

var mappingListDigitalTwinInstancesLifecycleStateEnumLowerCase = map[string]ListDigitalTwinInstancesLifecycleStateEnum{
	"active":  ListDigitalTwinInstancesLifecycleStateActive,
	"deleted": ListDigitalTwinInstancesLifecycleStateDeleted,
}

// GetListDigitalTwinInstancesLifecycleStateEnumValues Enumerates the set of values for ListDigitalTwinInstancesLifecycleStateEnum
func GetListDigitalTwinInstancesLifecycleStateEnumValues() []ListDigitalTwinInstancesLifecycleStateEnum {
	values := make([]ListDigitalTwinInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalTwinInstancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinInstancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalTwinInstancesLifecycleStateEnum
func GetListDigitalTwinInstancesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDigitalTwinInstancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinInstancesLifecycleStateEnum(val string) (ListDigitalTwinInstancesLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalTwinInstancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinInstancesSortOrderEnum Enum with underlying type: string
type ListDigitalTwinInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalTwinInstancesSortOrderEnum
const (
	ListDigitalTwinInstancesSortOrderAsc  ListDigitalTwinInstancesSortOrderEnum = "ASC"
	ListDigitalTwinInstancesSortOrderDesc ListDigitalTwinInstancesSortOrderEnum = "DESC"
)

var mappingListDigitalTwinInstancesSortOrderEnum = map[string]ListDigitalTwinInstancesSortOrderEnum{
	"ASC":  ListDigitalTwinInstancesSortOrderAsc,
	"DESC": ListDigitalTwinInstancesSortOrderDesc,
}

var mappingListDigitalTwinInstancesSortOrderEnumLowerCase = map[string]ListDigitalTwinInstancesSortOrderEnum{
	"asc":  ListDigitalTwinInstancesSortOrderAsc,
	"desc": ListDigitalTwinInstancesSortOrderDesc,
}

// GetListDigitalTwinInstancesSortOrderEnumValues Enumerates the set of values for ListDigitalTwinInstancesSortOrderEnum
func GetListDigitalTwinInstancesSortOrderEnumValues() []ListDigitalTwinInstancesSortOrderEnum {
	values := make([]ListDigitalTwinInstancesSortOrderEnum, 0)
	for _, v := range mappingListDigitalTwinInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalTwinInstancesSortOrderEnum
func GetListDigitalTwinInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalTwinInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinInstancesSortOrderEnum(val string) (ListDigitalTwinInstancesSortOrderEnum, bool) {
	enum, ok := mappingListDigitalTwinInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinInstancesSortByEnum Enum with underlying type: string
type ListDigitalTwinInstancesSortByEnum string

// Set of constants representing the allowable values for ListDigitalTwinInstancesSortByEnum
const (
	ListDigitalTwinInstancesSortByTimecreated ListDigitalTwinInstancesSortByEnum = "timeCreated"
	ListDigitalTwinInstancesSortByDisplayname ListDigitalTwinInstancesSortByEnum = "displayName"
)

var mappingListDigitalTwinInstancesSortByEnum = map[string]ListDigitalTwinInstancesSortByEnum{
	"timeCreated": ListDigitalTwinInstancesSortByTimecreated,
	"displayName": ListDigitalTwinInstancesSortByDisplayname,
}

var mappingListDigitalTwinInstancesSortByEnumLowerCase = map[string]ListDigitalTwinInstancesSortByEnum{
	"timecreated": ListDigitalTwinInstancesSortByTimecreated,
	"displayname": ListDigitalTwinInstancesSortByDisplayname,
}

// GetListDigitalTwinInstancesSortByEnumValues Enumerates the set of values for ListDigitalTwinInstancesSortByEnum
func GetListDigitalTwinInstancesSortByEnumValues() []ListDigitalTwinInstancesSortByEnum {
	values := make([]ListDigitalTwinInstancesSortByEnum, 0)
	for _, v := range mappingListDigitalTwinInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinInstancesSortByEnumStringValues Enumerates the set of values in String for ListDigitalTwinInstancesSortByEnum
func GetListDigitalTwinInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDigitalTwinInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinInstancesSortByEnum(val string) (ListDigitalTwinInstancesSortByEnum, bool) {
	enum, ok := mappingListDigitalTwinInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
