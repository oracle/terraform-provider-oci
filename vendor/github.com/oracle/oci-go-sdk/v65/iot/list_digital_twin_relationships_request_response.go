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

// ListDigitalTwinRelationshipsRequest wrapper for the ListDigitalTwinRelationships operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinRelationships.go.html to see an example of how to use ListDigitalTwinRelationshipsRequest.
type ListDigitalTwinRelationshipsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources.
	IotDomainId *string `mandatory:"true" contributesTo:"query" name:"iotDomainId"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filters resources that match the content path of the digital twin relationship.
	ContentPath *string `mandatory:"false" contributesTo:"query" name:"contentPath"`

	// Filter resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance.
	SourceDigitalTwinInstanceId *string `mandatory:"false" contributesTo:"query" name:"sourceDigitalTwinInstanceId"`

	// Filter resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.
	TargetDigitalTwinInstanceId *string `mandatory:"false" contributesTo:"query" name:"targetDigitalTwinInstanceId"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState ListDigitalTwinRelationshipsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Page representing the requested page of items.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListDigitalTwinRelationshipsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDigitalTwinRelationshipsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalTwinRelationshipsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalTwinRelationshipsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalTwinRelationshipsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalTwinRelationshipsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalTwinRelationshipsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalTwinRelationshipsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalTwinRelationshipsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinRelationshipsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalTwinRelationshipsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalTwinRelationshipsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalTwinRelationshipsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalTwinRelationshipsResponse wrapper for the ListDigitalTwinRelationships operation
type ListDigitalTwinRelationshipsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalTwinRelationshipCollection instances
	DigitalTwinRelationshipCollection `presentIn:"body"`

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

func (response ListDigitalTwinRelationshipsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalTwinRelationshipsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalTwinRelationshipsLifecycleStateEnum Enum with underlying type: string
type ListDigitalTwinRelationshipsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalTwinRelationshipsLifecycleStateEnum
const (
	ListDigitalTwinRelationshipsLifecycleStateActive  ListDigitalTwinRelationshipsLifecycleStateEnum = "ACTIVE"
	ListDigitalTwinRelationshipsLifecycleStateDeleted ListDigitalTwinRelationshipsLifecycleStateEnum = "DELETED"
)

var mappingListDigitalTwinRelationshipsLifecycleStateEnum = map[string]ListDigitalTwinRelationshipsLifecycleStateEnum{
	"ACTIVE":  ListDigitalTwinRelationshipsLifecycleStateActive,
	"DELETED": ListDigitalTwinRelationshipsLifecycleStateDeleted,
}

var mappingListDigitalTwinRelationshipsLifecycleStateEnumLowerCase = map[string]ListDigitalTwinRelationshipsLifecycleStateEnum{
	"active":  ListDigitalTwinRelationshipsLifecycleStateActive,
	"deleted": ListDigitalTwinRelationshipsLifecycleStateDeleted,
}

// GetListDigitalTwinRelationshipsLifecycleStateEnumValues Enumerates the set of values for ListDigitalTwinRelationshipsLifecycleStateEnum
func GetListDigitalTwinRelationshipsLifecycleStateEnumValues() []ListDigitalTwinRelationshipsLifecycleStateEnum {
	values := make([]ListDigitalTwinRelationshipsLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalTwinRelationshipsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinRelationshipsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalTwinRelationshipsLifecycleStateEnum
func GetListDigitalTwinRelationshipsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDigitalTwinRelationshipsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinRelationshipsLifecycleStateEnum(val string) (ListDigitalTwinRelationshipsLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalTwinRelationshipsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinRelationshipsSortOrderEnum Enum with underlying type: string
type ListDigitalTwinRelationshipsSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalTwinRelationshipsSortOrderEnum
const (
	ListDigitalTwinRelationshipsSortOrderAsc  ListDigitalTwinRelationshipsSortOrderEnum = "ASC"
	ListDigitalTwinRelationshipsSortOrderDesc ListDigitalTwinRelationshipsSortOrderEnum = "DESC"
)

var mappingListDigitalTwinRelationshipsSortOrderEnum = map[string]ListDigitalTwinRelationshipsSortOrderEnum{
	"ASC":  ListDigitalTwinRelationshipsSortOrderAsc,
	"DESC": ListDigitalTwinRelationshipsSortOrderDesc,
}

var mappingListDigitalTwinRelationshipsSortOrderEnumLowerCase = map[string]ListDigitalTwinRelationshipsSortOrderEnum{
	"asc":  ListDigitalTwinRelationshipsSortOrderAsc,
	"desc": ListDigitalTwinRelationshipsSortOrderDesc,
}

// GetListDigitalTwinRelationshipsSortOrderEnumValues Enumerates the set of values for ListDigitalTwinRelationshipsSortOrderEnum
func GetListDigitalTwinRelationshipsSortOrderEnumValues() []ListDigitalTwinRelationshipsSortOrderEnum {
	values := make([]ListDigitalTwinRelationshipsSortOrderEnum, 0)
	for _, v := range mappingListDigitalTwinRelationshipsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinRelationshipsSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalTwinRelationshipsSortOrderEnum
func GetListDigitalTwinRelationshipsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalTwinRelationshipsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinRelationshipsSortOrderEnum(val string) (ListDigitalTwinRelationshipsSortOrderEnum, bool) {
	enum, ok := mappingListDigitalTwinRelationshipsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalTwinRelationshipsSortByEnum Enum with underlying type: string
type ListDigitalTwinRelationshipsSortByEnum string

// Set of constants representing the allowable values for ListDigitalTwinRelationshipsSortByEnum
const (
	ListDigitalTwinRelationshipsSortByTimecreated ListDigitalTwinRelationshipsSortByEnum = "timeCreated"
	ListDigitalTwinRelationshipsSortByDisplayname ListDigitalTwinRelationshipsSortByEnum = "displayName"
)

var mappingListDigitalTwinRelationshipsSortByEnum = map[string]ListDigitalTwinRelationshipsSortByEnum{
	"timeCreated": ListDigitalTwinRelationshipsSortByTimecreated,
	"displayName": ListDigitalTwinRelationshipsSortByDisplayname,
}

var mappingListDigitalTwinRelationshipsSortByEnumLowerCase = map[string]ListDigitalTwinRelationshipsSortByEnum{
	"timecreated": ListDigitalTwinRelationshipsSortByTimecreated,
	"displayname": ListDigitalTwinRelationshipsSortByDisplayname,
}

// GetListDigitalTwinRelationshipsSortByEnumValues Enumerates the set of values for ListDigitalTwinRelationshipsSortByEnum
func GetListDigitalTwinRelationshipsSortByEnumValues() []ListDigitalTwinRelationshipsSortByEnum {
	values := make([]ListDigitalTwinRelationshipsSortByEnum, 0)
	for _, v := range mappingListDigitalTwinRelationshipsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalTwinRelationshipsSortByEnumStringValues Enumerates the set of values in String for ListDigitalTwinRelationshipsSortByEnum
func GetListDigitalTwinRelationshipsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDigitalTwinRelationshipsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalTwinRelationshipsSortByEnum(val string) (ListDigitalTwinRelationshipsSortByEnum, bool) {
	enum, ok := mappingListDigitalTwinRelationshipsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
