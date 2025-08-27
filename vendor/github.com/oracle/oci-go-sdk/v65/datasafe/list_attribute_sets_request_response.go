// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAttributeSetsRequest wrapper for the ListAttributeSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAttributeSets.go.html to see an example of how to use ListAttributeSetsRequest.
type ListAttributeSetsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAttributeSetsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only attribute set resources that matches the specified attribute set OCID query param.
	AttributeSetId *string `mandatory:"false" contributesTo:"query" name:"attributeSetId"`

	// A filter to return only attribute set resources that matches the specified attribute set type query param.
	AttributeSetType AttributeSetAttributeSetTypeEnum `mandatory:"false" contributesTo:"query" name:"attributeSetType" omitEmpty:"true"`

	// The current state of an attribute set.
	LifecycleState AttributeSetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAttributeSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListAttributeSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return user defined or seeded attribute set resources that matches the specified is user defined query param. A true value indicates user defined attribute set.
	IsUserDefined *bool `mandatory:"false" contributesTo:"query" name:"isUserDefined"`

	// A filter to return attribute set resources that are in use by other associated resources.
	InUse ListAttributeSetsInUseEnum `mandatory:"false" contributesTo:"query" name:"inUse" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAttributeSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttributeSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttributeSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttributeSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttributeSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttributeSetsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAttributeSetsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeSetAttributeSetTypeEnum(string(request.AttributeSetType)); !ok && request.AttributeSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeSetType: %s. Supported values are: %s.", request.AttributeSetType, strings.Join(GetAttributeSetAttributeSetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeSetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAttributeSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttributeSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttributeSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttributeSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttributeSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttributeSetsInUseEnum(string(request.InUse)); !ok && request.InUse != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InUse: %s. Supported values are: %s.", request.InUse, strings.Join(GetListAttributeSetsInUseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttributeSetsResponse wrapper for the ListAttributeSets operation
type ListAttributeSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AttributeSetCollection instances
	AttributeSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAttributeSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttributeSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttributeSetsAccessLevelEnum Enum with underlying type: string
type ListAttributeSetsAccessLevelEnum string

// Set of constants representing the allowable values for ListAttributeSetsAccessLevelEnum
const (
	ListAttributeSetsAccessLevelRestricted ListAttributeSetsAccessLevelEnum = "RESTRICTED"
	ListAttributeSetsAccessLevelAccessible ListAttributeSetsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAttributeSetsAccessLevelEnum = map[string]ListAttributeSetsAccessLevelEnum{
	"RESTRICTED": ListAttributeSetsAccessLevelRestricted,
	"ACCESSIBLE": ListAttributeSetsAccessLevelAccessible,
}

var mappingListAttributeSetsAccessLevelEnumLowerCase = map[string]ListAttributeSetsAccessLevelEnum{
	"restricted": ListAttributeSetsAccessLevelRestricted,
	"accessible": ListAttributeSetsAccessLevelAccessible,
}

// GetListAttributeSetsAccessLevelEnumValues Enumerates the set of values for ListAttributeSetsAccessLevelEnum
func GetListAttributeSetsAccessLevelEnumValues() []ListAttributeSetsAccessLevelEnum {
	values := make([]ListAttributeSetsAccessLevelEnum, 0)
	for _, v := range mappingListAttributeSetsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeSetsAccessLevelEnumStringValues Enumerates the set of values in String for ListAttributeSetsAccessLevelEnum
func GetListAttributeSetsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAttributeSetsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeSetsAccessLevelEnum(val string) (ListAttributeSetsAccessLevelEnum, bool) {
	enum, ok := mappingListAttributeSetsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeSetsSortOrderEnum Enum with underlying type: string
type ListAttributeSetsSortOrderEnum string

// Set of constants representing the allowable values for ListAttributeSetsSortOrderEnum
const (
	ListAttributeSetsSortOrderAsc  ListAttributeSetsSortOrderEnum = "ASC"
	ListAttributeSetsSortOrderDesc ListAttributeSetsSortOrderEnum = "DESC"
)

var mappingListAttributeSetsSortOrderEnum = map[string]ListAttributeSetsSortOrderEnum{
	"ASC":  ListAttributeSetsSortOrderAsc,
	"DESC": ListAttributeSetsSortOrderDesc,
}

var mappingListAttributeSetsSortOrderEnumLowerCase = map[string]ListAttributeSetsSortOrderEnum{
	"asc":  ListAttributeSetsSortOrderAsc,
	"desc": ListAttributeSetsSortOrderDesc,
}

// GetListAttributeSetsSortOrderEnumValues Enumerates the set of values for ListAttributeSetsSortOrderEnum
func GetListAttributeSetsSortOrderEnumValues() []ListAttributeSetsSortOrderEnum {
	values := make([]ListAttributeSetsSortOrderEnum, 0)
	for _, v := range mappingListAttributeSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeSetsSortOrderEnumStringValues Enumerates the set of values in String for ListAttributeSetsSortOrderEnum
func GetListAttributeSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttributeSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeSetsSortOrderEnum(val string) (ListAttributeSetsSortOrderEnum, bool) {
	enum, ok := mappingListAttributeSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeSetsSortByEnum Enum with underlying type: string
type ListAttributeSetsSortByEnum string

// Set of constants representing the allowable values for ListAttributeSetsSortByEnum
const (
	ListAttributeSetsSortByTimecreated ListAttributeSetsSortByEnum = "TIMECREATED"
	ListAttributeSetsSortByDisplayname ListAttributeSetsSortByEnum = "DISPLAYNAME"
)

var mappingListAttributeSetsSortByEnum = map[string]ListAttributeSetsSortByEnum{
	"TIMECREATED": ListAttributeSetsSortByTimecreated,
	"DISPLAYNAME": ListAttributeSetsSortByDisplayname,
}

var mappingListAttributeSetsSortByEnumLowerCase = map[string]ListAttributeSetsSortByEnum{
	"timecreated": ListAttributeSetsSortByTimecreated,
	"displayname": ListAttributeSetsSortByDisplayname,
}

// GetListAttributeSetsSortByEnumValues Enumerates the set of values for ListAttributeSetsSortByEnum
func GetListAttributeSetsSortByEnumValues() []ListAttributeSetsSortByEnum {
	values := make([]ListAttributeSetsSortByEnum, 0)
	for _, v := range mappingListAttributeSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeSetsSortByEnumStringValues Enumerates the set of values in String for ListAttributeSetsSortByEnum
func GetListAttributeSetsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAttributeSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeSetsSortByEnum(val string) (ListAttributeSetsSortByEnum, bool) {
	enum, ok := mappingListAttributeSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeSetsInUseEnum Enum with underlying type: string
type ListAttributeSetsInUseEnum string

// Set of constants representing the allowable values for ListAttributeSetsInUseEnum
const (
	ListAttributeSetsInUseYes ListAttributeSetsInUseEnum = "YES"
	ListAttributeSetsInUseNo  ListAttributeSetsInUseEnum = "NO"
)

var mappingListAttributeSetsInUseEnum = map[string]ListAttributeSetsInUseEnum{
	"YES": ListAttributeSetsInUseYes,
	"NO":  ListAttributeSetsInUseNo,
}

var mappingListAttributeSetsInUseEnumLowerCase = map[string]ListAttributeSetsInUseEnum{
	"yes": ListAttributeSetsInUseYes,
	"no":  ListAttributeSetsInUseNo,
}

// GetListAttributeSetsInUseEnumValues Enumerates the set of values for ListAttributeSetsInUseEnum
func GetListAttributeSetsInUseEnumValues() []ListAttributeSetsInUseEnum {
	values := make([]ListAttributeSetsInUseEnum, 0)
	for _, v := range mappingListAttributeSetsInUseEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeSetsInUseEnumStringValues Enumerates the set of values in String for ListAttributeSetsInUseEnum
func GetListAttributeSetsInUseEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingListAttributeSetsInUseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeSetsInUseEnum(val string) (ListAttributeSetsInUseEnum, bool) {
	enum, ok := mappingListAttributeSetsInUseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
