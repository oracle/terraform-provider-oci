// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbConnectionBundlesRequest wrapper for the ListDbConnectionBundles operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbConnectionBundles.go.html to see an example of how to use ListDbConnectionBundlesRequest.
type ListDbConnectionBundlesRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbConnectionBundlesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `TIMECREATED` is descending. The default order for `TIMEREFRESHED`
	// is descending.
	SortBy ListDbConnectionBundlesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState ListDbConnectionBundlesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter that returns only resources that match the specified database connection bundle type.
	DbConnectionBundleType ListDbConnectionBundlesDbConnectionBundleTypeEnum `mandatory:"false" contributesTo:"query" name:"dbConnectionBundleType" omitEmpty:"true"`

	// The OCID of the VM cluster associated with the connection bundle. If the parameter is set to null, all bundles are returned.
	AssociatedResourceId *string `mandatory:"false" contributesTo:"query" name:"associatedResourceId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbConnectionBundlesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbConnectionBundlesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbConnectionBundlesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbConnectionBundlesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbConnectionBundlesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbConnectionBundlesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbConnectionBundlesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbConnectionBundlesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbConnectionBundlesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbConnectionBundlesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDbConnectionBundlesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbConnectionBundlesDbConnectionBundleTypeEnum(string(request.DbConnectionBundleType)); !ok && request.DbConnectionBundleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbConnectionBundleType: %s. Supported values are: %s.", request.DbConnectionBundleType, strings.Join(GetListDbConnectionBundlesDbConnectionBundleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbConnectionBundlesResponse wrapper for the ListDbConnectionBundles operation
type ListDbConnectionBundlesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbConnectionBundleSummary instances
	Items []DbConnectionBundleSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbConnectionBundlesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbConnectionBundlesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbConnectionBundlesSortOrderEnum Enum with underlying type: string
type ListDbConnectionBundlesSortOrderEnum string

// Set of constants representing the allowable values for ListDbConnectionBundlesSortOrderEnum
const (
	ListDbConnectionBundlesSortOrderAsc  ListDbConnectionBundlesSortOrderEnum = "ASC"
	ListDbConnectionBundlesSortOrderDesc ListDbConnectionBundlesSortOrderEnum = "DESC"
)

var mappingListDbConnectionBundlesSortOrderEnum = map[string]ListDbConnectionBundlesSortOrderEnum{
	"ASC":  ListDbConnectionBundlesSortOrderAsc,
	"DESC": ListDbConnectionBundlesSortOrderDesc,
}

var mappingListDbConnectionBundlesSortOrderEnumLowerCase = map[string]ListDbConnectionBundlesSortOrderEnum{
	"asc":  ListDbConnectionBundlesSortOrderAsc,
	"desc": ListDbConnectionBundlesSortOrderDesc,
}

// GetListDbConnectionBundlesSortOrderEnumValues Enumerates the set of values for ListDbConnectionBundlesSortOrderEnum
func GetListDbConnectionBundlesSortOrderEnumValues() []ListDbConnectionBundlesSortOrderEnum {
	values := make([]ListDbConnectionBundlesSortOrderEnum, 0)
	for _, v := range mappingListDbConnectionBundlesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbConnectionBundlesSortOrderEnumStringValues Enumerates the set of values in String for ListDbConnectionBundlesSortOrderEnum
func GetListDbConnectionBundlesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbConnectionBundlesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbConnectionBundlesSortOrderEnum(val string) (ListDbConnectionBundlesSortOrderEnum, bool) {
	enum, ok := mappingListDbConnectionBundlesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbConnectionBundlesSortByEnum Enum with underlying type: string
type ListDbConnectionBundlesSortByEnum string

// Set of constants representing the allowable values for ListDbConnectionBundlesSortByEnum
const (
	ListDbConnectionBundlesSortByTimecreated   ListDbConnectionBundlesSortByEnum = "TIMECREATED"
	ListDbConnectionBundlesSortByTimerefreshed ListDbConnectionBundlesSortByEnum = "TIMEREFRESHED"
)

var mappingListDbConnectionBundlesSortByEnum = map[string]ListDbConnectionBundlesSortByEnum{
	"TIMECREATED":   ListDbConnectionBundlesSortByTimecreated,
	"TIMEREFRESHED": ListDbConnectionBundlesSortByTimerefreshed,
}

var mappingListDbConnectionBundlesSortByEnumLowerCase = map[string]ListDbConnectionBundlesSortByEnum{
	"timecreated":   ListDbConnectionBundlesSortByTimecreated,
	"timerefreshed": ListDbConnectionBundlesSortByTimerefreshed,
}

// GetListDbConnectionBundlesSortByEnumValues Enumerates the set of values for ListDbConnectionBundlesSortByEnum
func GetListDbConnectionBundlesSortByEnumValues() []ListDbConnectionBundlesSortByEnum {
	values := make([]ListDbConnectionBundlesSortByEnum, 0)
	for _, v := range mappingListDbConnectionBundlesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbConnectionBundlesSortByEnumStringValues Enumerates the set of values in String for ListDbConnectionBundlesSortByEnum
func GetListDbConnectionBundlesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMEREFRESHED",
	}
}

// GetMappingListDbConnectionBundlesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbConnectionBundlesSortByEnum(val string) (ListDbConnectionBundlesSortByEnum, bool) {
	enum, ok := mappingListDbConnectionBundlesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbConnectionBundlesLifecycleStateEnum Enum with underlying type: string
type ListDbConnectionBundlesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDbConnectionBundlesLifecycleStateEnum
const (
	ListDbConnectionBundlesLifecycleStateCreating ListDbConnectionBundlesLifecycleStateEnum = "CREATING"
	ListDbConnectionBundlesLifecycleStateActive   ListDbConnectionBundlesLifecycleStateEnum = "ACTIVE"
	ListDbConnectionBundlesLifecycleStateInactive ListDbConnectionBundlesLifecycleStateEnum = "INACTIVE"
	ListDbConnectionBundlesLifecycleStateUpdating ListDbConnectionBundlesLifecycleStateEnum = "UPDATING"
	ListDbConnectionBundlesLifecycleStateDeleting ListDbConnectionBundlesLifecycleStateEnum = "DELETING"
	ListDbConnectionBundlesLifecycleStateDeleted  ListDbConnectionBundlesLifecycleStateEnum = "DELETED"
	ListDbConnectionBundlesLifecycleStateFailed   ListDbConnectionBundlesLifecycleStateEnum = "FAILED"
)

var mappingListDbConnectionBundlesLifecycleStateEnum = map[string]ListDbConnectionBundlesLifecycleStateEnum{
	"CREATING": ListDbConnectionBundlesLifecycleStateCreating,
	"ACTIVE":   ListDbConnectionBundlesLifecycleStateActive,
	"INACTIVE": ListDbConnectionBundlesLifecycleStateInactive,
	"UPDATING": ListDbConnectionBundlesLifecycleStateUpdating,
	"DELETING": ListDbConnectionBundlesLifecycleStateDeleting,
	"DELETED":  ListDbConnectionBundlesLifecycleStateDeleted,
	"FAILED":   ListDbConnectionBundlesLifecycleStateFailed,
}

var mappingListDbConnectionBundlesLifecycleStateEnumLowerCase = map[string]ListDbConnectionBundlesLifecycleStateEnum{
	"creating": ListDbConnectionBundlesLifecycleStateCreating,
	"active":   ListDbConnectionBundlesLifecycleStateActive,
	"inactive": ListDbConnectionBundlesLifecycleStateInactive,
	"updating": ListDbConnectionBundlesLifecycleStateUpdating,
	"deleting": ListDbConnectionBundlesLifecycleStateDeleting,
	"deleted":  ListDbConnectionBundlesLifecycleStateDeleted,
	"failed":   ListDbConnectionBundlesLifecycleStateFailed,
}

// GetListDbConnectionBundlesLifecycleStateEnumValues Enumerates the set of values for ListDbConnectionBundlesLifecycleStateEnum
func GetListDbConnectionBundlesLifecycleStateEnumValues() []ListDbConnectionBundlesLifecycleStateEnum {
	values := make([]ListDbConnectionBundlesLifecycleStateEnum, 0)
	for _, v := range mappingListDbConnectionBundlesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbConnectionBundlesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDbConnectionBundlesLifecycleStateEnum
func GetListDbConnectionBundlesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDbConnectionBundlesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbConnectionBundlesLifecycleStateEnum(val string) (ListDbConnectionBundlesLifecycleStateEnum, bool) {
	enum, ok := mappingListDbConnectionBundlesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbConnectionBundlesDbConnectionBundleTypeEnum Enum with underlying type: string
type ListDbConnectionBundlesDbConnectionBundleTypeEnum string

// Set of constants representing the allowable values for ListDbConnectionBundlesDbConnectionBundleTypeEnum
const (
	ListDbConnectionBundlesDbConnectionBundleTypeTls  ListDbConnectionBundlesDbConnectionBundleTypeEnum = "TLS"
	ListDbConnectionBundlesDbConnectionBundleTypeMtls ListDbConnectionBundlesDbConnectionBundleTypeEnum = "MTLS"
)

var mappingListDbConnectionBundlesDbConnectionBundleTypeEnum = map[string]ListDbConnectionBundlesDbConnectionBundleTypeEnum{
	"TLS":  ListDbConnectionBundlesDbConnectionBundleTypeTls,
	"MTLS": ListDbConnectionBundlesDbConnectionBundleTypeMtls,
}

var mappingListDbConnectionBundlesDbConnectionBundleTypeEnumLowerCase = map[string]ListDbConnectionBundlesDbConnectionBundleTypeEnum{
	"tls":  ListDbConnectionBundlesDbConnectionBundleTypeTls,
	"mtls": ListDbConnectionBundlesDbConnectionBundleTypeMtls,
}

// GetListDbConnectionBundlesDbConnectionBundleTypeEnumValues Enumerates the set of values for ListDbConnectionBundlesDbConnectionBundleTypeEnum
func GetListDbConnectionBundlesDbConnectionBundleTypeEnumValues() []ListDbConnectionBundlesDbConnectionBundleTypeEnum {
	values := make([]ListDbConnectionBundlesDbConnectionBundleTypeEnum, 0)
	for _, v := range mappingListDbConnectionBundlesDbConnectionBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbConnectionBundlesDbConnectionBundleTypeEnumStringValues Enumerates the set of values in String for ListDbConnectionBundlesDbConnectionBundleTypeEnum
func GetListDbConnectionBundlesDbConnectionBundleTypeEnumStringValues() []string {
	return []string{
		"TLS",
		"MTLS",
	}
}

// GetMappingListDbConnectionBundlesDbConnectionBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbConnectionBundlesDbConnectionBundleTypeEnum(val string) (ListDbConnectionBundlesDbConnectionBundleTypeEnum, bool) {
	enum, ok := mappingListDbConnectionBundlesDbConnectionBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
