// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRegistrationPoliciesRequest wrapper for the ListRegistrationPolicies operation
type ListRegistrationPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Field to sort the registration policies by. Only one sort order (sortOrder) can be specified. Defaults to sorting by `timeCreated` in descending order
	SortBy ListRegistrationPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter registration policies by their lifecycle state.
	LifecycleState RegistrationPolicyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter registration policies by resource type.
	EnablementLevel RegistrationPolicyEnablementLevelEnum `mandatory:"false" contributesTo:"query" name:"enablementLevel" omitEmpty:"true"`

	// Filter to return the registration policy matching the specified resource OCID.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListRegistrationPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListRegistrationPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter to return the registration policy matching the specified OCID.
	RegistrationPolicyId *string `mandatory:"false" contributesTo:"query" name:"registrationPolicyId"`

	// Filter to return the registration policies matching the specified connectionType i.e ONPREM_CONNECTOR or PRIVATE_ENDPOINT.
	ConnectionType PolicyConnectionOptionConnectionTypeEnum `mandatory:"false" contributesTo:"query" name:"connectionType" omitEmpty:"true"`

	// Filter to return the registration policies matching the specified connection ID.
	ConnectionId *string `mandatory:"false" contributesTo:"query" name:"connectionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRegistrationPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRegistrationPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRegistrationPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRegistrationPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRegistrationPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRegistrationPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRegistrationPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRegistrationPolicyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRegistrationPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRegistrationPolicyEnablementLevelEnum(string(request.EnablementLevel)); !ok && request.EnablementLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnablementLevel: %s. Supported values are: %s.", request.EnablementLevel, strings.Join(GetRegistrationPolicyEnablementLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRegistrationPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListRegistrationPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRegistrationPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRegistrationPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPolicyConnectionOptionConnectionTypeEnum(string(request.ConnectionType)); !ok && request.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", request.ConnectionType, strings.Join(GetPolicyConnectionOptionConnectionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRegistrationPoliciesResponse wrapper for the ListRegistrationPolicies operation
type ListRegistrationPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RegistrationPolicyCollection instances
	RegistrationPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListRegistrationPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRegistrationPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRegistrationPoliciesSortByEnum Enum with underlying type: string
type ListRegistrationPoliciesSortByEnum string

// Set of constants representing the allowable values for ListRegistrationPoliciesSortByEnum
const (
	ListRegistrationPoliciesSortByTimecreated     ListRegistrationPoliciesSortByEnum = "timeCreated"
	ListRegistrationPoliciesSortByDisplayname     ListRegistrationPoliciesSortByEnum = "displayName"
	ListRegistrationPoliciesSortByEnablementlevel ListRegistrationPoliciesSortByEnum = "enablementLevel"
	ListRegistrationPoliciesSortByLifecyclestate  ListRegistrationPoliciesSortByEnum = "lifecycleState"
)

var mappingListRegistrationPoliciesSortByEnum = map[string]ListRegistrationPoliciesSortByEnum{
	"timeCreated":     ListRegistrationPoliciesSortByTimecreated,
	"displayName":     ListRegistrationPoliciesSortByDisplayname,
	"enablementLevel": ListRegistrationPoliciesSortByEnablementlevel,
	"lifecycleState":  ListRegistrationPoliciesSortByLifecyclestate,
}

var mappingListRegistrationPoliciesSortByEnumLowerCase = map[string]ListRegistrationPoliciesSortByEnum{
	"timecreated":     ListRegistrationPoliciesSortByTimecreated,
	"displayname":     ListRegistrationPoliciesSortByDisplayname,
	"enablementlevel": ListRegistrationPoliciesSortByEnablementlevel,
	"lifecyclestate":  ListRegistrationPoliciesSortByLifecyclestate,
}

// GetListRegistrationPoliciesSortByEnumValues Enumerates the set of values for ListRegistrationPoliciesSortByEnum
func GetListRegistrationPoliciesSortByEnumValues() []ListRegistrationPoliciesSortByEnum {
	values := make([]ListRegistrationPoliciesSortByEnum, 0)
	for _, v := range mappingListRegistrationPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRegistrationPoliciesSortByEnumStringValues Enumerates the set of values in String for ListRegistrationPoliciesSortByEnum
func GetListRegistrationPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"enablementLevel",
		"lifecycleState",
	}
}

// GetMappingListRegistrationPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRegistrationPoliciesSortByEnum(val string) (ListRegistrationPoliciesSortByEnum, bool) {
	enum, ok := mappingListRegistrationPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRegistrationPoliciesAccessLevelEnum Enum with underlying type: string
type ListRegistrationPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListRegistrationPoliciesAccessLevelEnum
const (
	ListRegistrationPoliciesAccessLevelRestricted ListRegistrationPoliciesAccessLevelEnum = "RESTRICTED"
	ListRegistrationPoliciesAccessLevelAccessible ListRegistrationPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListRegistrationPoliciesAccessLevelEnum = map[string]ListRegistrationPoliciesAccessLevelEnum{
	"RESTRICTED": ListRegistrationPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListRegistrationPoliciesAccessLevelAccessible,
}

var mappingListRegistrationPoliciesAccessLevelEnumLowerCase = map[string]ListRegistrationPoliciesAccessLevelEnum{
	"restricted": ListRegistrationPoliciesAccessLevelRestricted,
	"accessible": ListRegistrationPoliciesAccessLevelAccessible,
}

// GetListRegistrationPoliciesAccessLevelEnumValues Enumerates the set of values for ListRegistrationPoliciesAccessLevelEnum
func GetListRegistrationPoliciesAccessLevelEnumValues() []ListRegistrationPoliciesAccessLevelEnum {
	values := make([]ListRegistrationPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListRegistrationPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListRegistrationPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListRegistrationPoliciesAccessLevelEnum
func GetListRegistrationPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListRegistrationPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRegistrationPoliciesAccessLevelEnum(val string) (ListRegistrationPoliciesAccessLevelEnum, bool) {
	enum, ok := mappingListRegistrationPoliciesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRegistrationPoliciesSortOrderEnum Enum with underlying type: string
type ListRegistrationPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListRegistrationPoliciesSortOrderEnum
const (
	ListRegistrationPoliciesSortOrderAsc  ListRegistrationPoliciesSortOrderEnum = "ASC"
	ListRegistrationPoliciesSortOrderDesc ListRegistrationPoliciesSortOrderEnum = "DESC"
)

var mappingListRegistrationPoliciesSortOrderEnum = map[string]ListRegistrationPoliciesSortOrderEnum{
	"ASC":  ListRegistrationPoliciesSortOrderAsc,
	"DESC": ListRegistrationPoliciesSortOrderDesc,
}

var mappingListRegistrationPoliciesSortOrderEnumLowerCase = map[string]ListRegistrationPoliciesSortOrderEnum{
	"asc":  ListRegistrationPoliciesSortOrderAsc,
	"desc": ListRegistrationPoliciesSortOrderDesc,
}

// GetListRegistrationPoliciesSortOrderEnumValues Enumerates the set of values for ListRegistrationPoliciesSortOrderEnum
func GetListRegistrationPoliciesSortOrderEnumValues() []ListRegistrationPoliciesSortOrderEnum {
	values := make([]ListRegistrationPoliciesSortOrderEnum, 0)
	for _, v := range mappingListRegistrationPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRegistrationPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListRegistrationPoliciesSortOrderEnum
func GetListRegistrationPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRegistrationPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRegistrationPoliciesSortOrderEnum(val string) (ListRegistrationPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListRegistrationPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
