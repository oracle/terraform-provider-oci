// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccProvisionedPackagesRequest wrapper for the ListCccProvisionedPackages operation
type ListCccProvisionedPackagesRequest struct {

	// Compute Cloud@Customer disaster recovery project configuration
	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CccProvisionedPackageId *string `mandatory:"false" contributesTo:"query" name:"cccProvisionedPackageId"`

	// A filter used to return only resources that match the given infrastructureId.
	CccInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"cccInfrastructureId"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no
	// partial results are displayed.
	AccessLevel ListCccProvisionedPackagesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return resources only when their lifecycleState matches the given lifecycleState.
	LifecycleState CccProvisionedPackageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccProvisionedPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCccProvisionedPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccProvisionedPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccProvisionedPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccProvisionedPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccProvisionedPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccProvisionedPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccProvisionedPackagesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCccProvisionedPackagesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCccProvisionedPackageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCccProvisionedPackageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccProvisionedPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccProvisionedPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccProvisionedPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccProvisionedPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccProvisionedPackagesResponse wrapper for the ListCccProvisionedPackages operation
type ListCccProvisionedPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccProvisionedPackageCollection instances
	CccProvisionedPackageCollection `presentIn:"body"`

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

func (response ListCccProvisionedPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccProvisionedPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccProvisionedPackagesAccessLevelEnum Enum with underlying type: string
type ListCccProvisionedPackagesAccessLevelEnum string

// Set of constants representing the allowable values for ListCccProvisionedPackagesAccessLevelEnum
const (
	ListCccProvisionedPackagesAccessLevelRestricted ListCccProvisionedPackagesAccessLevelEnum = "RESTRICTED"
	ListCccProvisionedPackagesAccessLevelAccessible ListCccProvisionedPackagesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCccProvisionedPackagesAccessLevelEnum = map[string]ListCccProvisionedPackagesAccessLevelEnum{
	"RESTRICTED": ListCccProvisionedPackagesAccessLevelRestricted,
	"ACCESSIBLE": ListCccProvisionedPackagesAccessLevelAccessible,
}

var mappingListCccProvisionedPackagesAccessLevelEnumLowerCase = map[string]ListCccProvisionedPackagesAccessLevelEnum{
	"restricted": ListCccProvisionedPackagesAccessLevelRestricted,
	"accessible": ListCccProvisionedPackagesAccessLevelAccessible,
}

// GetListCccProvisionedPackagesAccessLevelEnumValues Enumerates the set of values for ListCccProvisionedPackagesAccessLevelEnum
func GetListCccProvisionedPackagesAccessLevelEnumValues() []ListCccProvisionedPackagesAccessLevelEnum {
	values := make([]ListCccProvisionedPackagesAccessLevelEnum, 0)
	for _, v := range mappingListCccProvisionedPackagesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccProvisionedPackagesAccessLevelEnumStringValues Enumerates the set of values in String for ListCccProvisionedPackagesAccessLevelEnum
func GetListCccProvisionedPackagesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCccProvisionedPackagesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccProvisionedPackagesAccessLevelEnum(val string) (ListCccProvisionedPackagesAccessLevelEnum, bool) {
	enum, ok := mappingListCccProvisionedPackagesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccProvisionedPackagesSortOrderEnum Enum with underlying type: string
type ListCccProvisionedPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListCccProvisionedPackagesSortOrderEnum
const (
	ListCccProvisionedPackagesSortOrderAsc  ListCccProvisionedPackagesSortOrderEnum = "ASC"
	ListCccProvisionedPackagesSortOrderDesc ListCccProvisionedPackagesSortOrderEnum = "DESC"
)

var mappingListCccProvisionedPackagesSortOrderEnum = map[string]ListCccProvisionedPackagesSortOrderEnum{
	"ASC":  ListCccProvisionedPackagesSortOrderAsc,
	"DESC": ListCccProvisionedPackagesSortOrderDesc,
}

var mappingListCccProvisionedPackagesSortOrderEnumLowerCase = map[string]ListCccProvisionedPackagesSortOrderEnum{
	"asc":  ListCccProvisionedPackagesSortOrderAsc,
	"desc": ListCccProvisionedPackagesSortOrderDesc,
}

// GetListCccProvisionedPackagesSortOrderEnumValues Enumerates the set of values for ListCccProvisionedPackagesSortOrderEnum
func GetListCccProvisionedPackagesSortOrderEnumValues() []ListCccProvisionedPackagesSortOrderEnum {
	values := make([]ListCccProvisionedPackagesSortOrderEnum, 0)
	for _, v := range mappingListCccProvisionedPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccProvisionedPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListCccProvisionedPackagesSortOrderEnum
func GetListCccProvisionedPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccProvisionedPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccProvisionedPackagesSortOrderEnum(val string) (ListCccProvisionedPackagesSortOrderEnum, bool) {
	enum, ok := mappingListCccProvisionedPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccProvisionedPackagesSortByEnum Enum with underlying type: string
type ListCccProvisionedPackagesSortByEnum string

// Set of constants representing the allowable values for ListCccProvisionedPackagesSortByEnum
const (
	ListCccProvisionedPackagesSortByTimecreated ListCccProvisionedPackagesSortByEnum = "timeCreated"
	ListCccProvisionedPackagesSortByDisplayname ListCccProvisionedPackagesSortByEnum = "displayName"
)

var mappingListCccProvisionedPackagesSortByEnum = map[string]ListCccProvisionedPackagesSortByEnum{
	"timeCreated": ListCccProvisionedPackagesSortByTimecreated,
	"displayName": ListCccProvisionedPackagesSortByDisplayname,
}

var mappingListCccProvisionedPackagesSortByEnumLowerCase = map[string]ListCccProvisionedPackagesSortByEnum{
	"timecreated": ListCccProvisionedPackagesSortByTimecreated,
	"displayname": ListCccProvisionedPackagesSortByDisplayname,
}

// GetListCccProvisionedPackagesSortByEnumValues Enumerates the set of values for ListCccProvisionedPackagesSortByEnum
func GetListCccProvisionedPackagesSortByEnumValues() []ListCccProvisionedPackagesSortByEnum {
	values := make([]ListCccProvisionedPackagesSortByEnum, 0)
	for _, v := range mappingListCccProvisionedPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccProvisionedPackagesSortByEnumStringValues Enumerates the set of values in String for ListCccProvisionedPackagesSortByEnum
func GetListCccProvisionedPackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCccProvisionedPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccProvisionedPackagesSortByEnum(val string) (ListCccProvisionedPackagesSortByEnum, bool) {
	enum, ok := mappingListCccProvisionedPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
