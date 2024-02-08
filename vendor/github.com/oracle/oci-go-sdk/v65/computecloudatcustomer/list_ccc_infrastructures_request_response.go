// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccInfrastructuresRequest wrapper for the ListCccInfrastructures operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ListCccInfrastructures.go.html to see an example of how to use ListCccInfrastructuresRequest.
type ListCccInfrastructuresRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and
	// sub-compartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no
	// partial results are displayed.
	AccessLevel ListCccInfrastructuresAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter used to return only resources that match the given lifecycleState.
	LifecycleState CccInfrastructureLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// An OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a
	// Compute Cloud@Customer Infrastructure.
	CccInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"cccInfrastructureId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccInfrastructuresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCccInfrastructuresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccInfrastructuresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccInfrastructuresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccInfrastructuresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccInfrastructuresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccInfrastructuresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccInfrastructuresAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCccInfrastructuresAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCccInfrastructureLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCccInfrastructureLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccInfrastructuresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccInfrastructuresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccInfrastructuresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccInfrastructuresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccInfrastructuresResponse wrapper for the ListCccInfrastructures operation
type ListCccInfrastructuresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccInfrastructureCollection instances
	CccInfrastructureCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCccInfrastructuresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccInfrastructuresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccInfrastructuresAccessLevelEnum Enum with underlying type: string
type ListCccInfrastructuresAccessLevelEnum string

// Set of constants representing the allowable values for ListCccInfrastructuresAccessLevelEnum
const (
	ListCccInfrastructuresAccessLevelRestricted ListCccInfrastructuresAccessLevelEnum = "RESTRICTED"
	ListCccInfrastructuresAccessLevelAccessible ListCccInfrastructuresAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCccInfrastructuresAccessLevelEnum = map[string]ListCccInfrastructuresAccessLevelEnum{
	"RESTRICTED": ListCccInfrastructuresAccessLevelRestricted,
	"ACCESSIBLE": ListCccInfrastructuresAccessLevelAccessible,
}

var mappingListCccInfrastructuresAccessLevelEnumLowerCase = map[string]ListCccInfrastructuresAccessLevelEnum{
	"restricted": ListCccInfrastructuresAccessLevelRestricted,
	"accessible": ListCccInfrastructuresAccessLevelAccessible,
}

// GetListCccInfrastructuresAccessLevelEnumValues Enumerates the set of values for ListCccInfrastructuresAccessLevelEnum
func GetListCccInfrastructuresAccessLevelEnumValues() []ListCccInfrastructuresAccessLevelEnum {
	values := make([]ListCccInfrastructuresAccessLevelEnum, 0)
	for _, v := range mappingListCccInfrastructuresAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccInfrastructuresAccessLevelEnumStringValues Enumerates the set of values in String for ListCccInfrastructuresAccessLevelEnum
func GetListCccInfrastructuresAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCccInfrastructuresAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccInfrastructuresAccessLevelEnum(val string) (ListCccInfrastructuresAccessLevelEnum, bool) {
	enum, ok := mappingListCccInfrastructuresAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccInfrastructuresSortOrderEnum Enum with underlying type: string
type ListCccInfrastructuresSortOrderEnum string

// Set of constants representing the allowable values for ListCccInfrastructuresSortOrderEnum
const (
	ListCccInfrastructuresSortOrderAsc  ListCccInfrastructuresSortOrderEnum = "ASC"
	ListCccInfrastructuresSortOrderDesc ListCccInfrastructuresSortOrderEnum = "DESC"
)

var mappingListCccInfrastructuresSortOrderEnum = map[string]ListCccInfrastructuresSortOrderEnum{
	"ASC":  ListCccInfrastructuresSortOrderAsc,
	"DESC": ListCccInfrastructuresSortOrderDesc,
}

var mappingListCccInfrastructuresSortOrderEnumLowerCase = map[string]ListCccInfrastructuresSortOrderEnum{
	"asc":  ListCccInfrastructuresSortOrderAsc,
	"desc": ListCccInfrastructuresSortOrderDesc,
}

// GetListCccInfrastructuresSortOrderEnumValues Enumerates the set of values for ListCccInfrastructuresSortOrderEnum
func GetListCccInfrastructuresSortOrderEnumValues() []ListCccInfrastructuresSortOrderEnum {
	values := make([]ListCccInfrastructuresSortOrderEnum, 0)
	for _, v := range mappingListCccInfrastructuresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccInfrastructuresSortOrderEnumStringValues Enumerates the set of values in String for ListCccInfrastructuresSortOrderEnum
func GetListCccInfrastructuresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccInfrastructuresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccInfrastructuresSortOrderEnum(val string) (ListCccInfrastructuresSortOrderEnum, bool) {
	enum, ok := mappingListCccInfrastructuresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccInfrastructuresSortByEnum Enum with underlying type: string
type ListCccInfrastructuresSortByEnum string

// Set of constants representing the allowable values for ListCccInfrastructuresSortByEnum
const (
	ListCccInfrastructuresSortByTimecreated ListCccInfrastructuresSortByEnum = "timeCreated"
	ListCccInfrastructuresSortByDisplayname ListCccInfrastructuresSortByEnum = "displayName"
)

var mappingListCccInfrastructuresSortByEnum = map[string]ListCccInfrastructuresSortByEnum{
	"timeCreated": ListCccInfrastructuresSortByTimecreated,
	"displayName": ListCccInfrastructuresSortByDisplayname,
}

var mappingListCccInfrastructuresSortByEnumLowerCase = map[string]ListCccInfrastructuresSortByEnum{
	"timecreated": ListCccInfrastructuresSortByTimecreated,
	"displayname": ListCccInfrastructuresSortByDisplayname,
}

// GetListCccInfrastructuresSortByEnumValues Enumerates the set of values for ListCccInfrastructuresSortByEnum
func GetListCccInfrastructuresSortByEnumValues() []ListCccInfrastructuresSortByEnum {
	values := make([]ListCccInfrastructuresSortByEnum, 0)
	for _, v := range mappingListCccInfrastructuresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccInfrastructuresSortByEnumStringValues Enumerates the set of values in String for ListCccInfrastructuresSortByEnum
func GetListCccInfrastructuresSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCccInfrastructuresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccInfrastructuresSortByEnum(val string) (ListCccInfrastructuresSortByEnum, bool) {
	enum, ok := mappingListCccInfrastructuresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
