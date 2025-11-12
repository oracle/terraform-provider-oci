// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCatalogItemsRequest wrapper for the ListCatalogItems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListCatalogItems.go.html to see an example of how to use ListCatalogItemsRequest.
type ListCatalogItemsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ConfigSourceType (https://docs.oracle.com/iaas/definitions/CatalogItem/configSourceType) Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, URL_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE.
	ConfigSourceType *string `mandatory:"false" contributesTo:"query" name:"configSourceType"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState CatalogItemLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCatalogItemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort by. Default value for `timeCreated`
	// is descending. Default order for `displayName` is ascending
	SortBy ListCatalogItemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// catalogListingId of the package. This is an integer whose min and max length are specified.
	CatalogListingId *string `mandatory:"false" contributesTo:"query" name:"catalogListingId"`

	// Parameter to list all catalog items only with latest version or list all catalog items with all versions.
	CatalogListingVersionCriteria ListCatalogItemsCatalogListingVersionCriteriaEnum `mandatory:"false" contributesTo:"query" name:"catalogListingVersionCriteria" omitEmpty:"true"`

	// A filter to return only resources that match the given package type. The
	// state value is case-insensitive.
	PackageType CatalogItemPackageTypeEnum `mandatory:"false" contributesTo:"query" name:"packageType" omitEmpty:"true"`

	// The indicator to append Public Items from the root compartment to any query, when set to TRUE.
	ShouldListPublicItems *bool `mandatory:"false" contributesTo:"query" name:"shouldListPublicItems"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCatalogItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCatalogItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCatalogItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCatalogItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCatalogItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogItemLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCatalogItemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogItemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCatalogItemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogItemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCatalogItemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogItemsCatalogListingVersionCriteriaEnum(string(request.CatalogListingVersionCriteria)); !ok && request.CatalogListingVersionCriteria != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CatalogListingVersionCriteria: %s. Supported values are: %s.", request.CatalogListingVersionCriteria, strings.Join(GetListCatalogItemsCatalogListingVersionCriteriaEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCatalogItemPackageTypeEnum(string(request.PackageType)); !ok && request.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", request.PackageType, strings.Join(GetCatalogItemPackageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCatalogItemsResponse wrapper for the ListCatalogItems operation
type ListCatalogItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CatalogItemCollection instances
	CatalogItemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCatalogItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogItemsSortOrderEnum Enum with underlying type: string
type ListCatalogItemsSortOrderEnum string

// Set of constants representing the allowable values for ListCatalogItemsSortOrderEnum
const (
	ListCatalogItemsSortOrderAsc  ListCatalogItemsSortOrderEnum = "ASC"
	ListCatalogItemsSortOrderDesc ListCatalogItemsSortOrderEnum = "DESC"
)

var mappingListCatalogItemsSortOrderEnum = map[string]ListCatalogItemsSortOrderEnum{
	"ASC":  ListCatalogItemsSortOrderAsc,
	"DESC": ListCatalogItemsSortOrderDesc,
}

var mappingListCatalogItemsSortOrderEnumLowerCase = map[string]ListCatalogItemsSortOrderEnum{
	"asc":  ListCatalogItemsSortOrderAsc,
	"desc": ListCatalogItemsSortOrderDesc,
}

// GetListCatalogItemsSortOrderEnumValues Enumerates the set of values for ListCatalogItemsSortOrderEnum
func GetListCatalogItemsSortOrderEnumValues() []ListCatalogItemsSortOrderEnum {
	values := make([]ListCatalogItemsSortOrderEnum, 0)
	for _, v := range mappingListCatalogItemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogItemsSortOrderEnumStringValues Enumerates the set of values in String for ListCatalogItemsSortOrderEnum
func GetListCatalogItemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCatalogItemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogItemsSortOrderEnum(val string) (ListCatalogItemsSortOrderEnum, bool) {
	enum, ok := mappingListCatalogItemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCatalogItemsSortByEnum Enum with underlying type: string
type ListCatalogItemsSortByEnum string

// Set of constants representing the allowable values for ListCatalogItemsSortByEnum
const (
	ListCatalogItemsSortByTimecreated             ListCatalogItemsSortByEnum = "timeCreated"
	ListCatalogItemsSortByDisplayname             ListCatalogItemsSortByEnum = "displayName"
	ListCatalogItemsSortByTimebackfilllastchecked ListCatalogItemsSortByEnum = "timeBackfillLastChecked"
)

var mappingListCatalogItemsSortByEnum = map[string]ListCatalogItemsSortByEnum{
	"timeCreated":             ListCatalogItemsSortByTimecreated,
	"displayName":             ListCatalogItemsSortByDisplayname,
	"timeBackfillLastChecked": ListCatalogItemsSortByTimebackfilllastchecked,
}

var mappingListCatalogItemsSortByEnumLowerCase = map[string]ListCatalogItemsSortByEnum{
	"timecreated":             ListCatalogItemsSortByTimecreated,
	"displayname":             ListCatalogItemsSortByDisplayname,
	"timebackfilllastchecked": ListCatalogItemsSortByTimebackfilllastchecked,
}

// GetListCatalogItemsSortByEnumValues Enumerates the set of values for ListCatalogItemsSortByEnum
func GetListCatalogItemsSortByEnumValues() []ListCatalogItemsSortByEnum {
	values := make([]ListCatalogItemsSortByEnum, 0)
	for _, v := range mappingListCatalogItemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogItemsSortByEnumStringValues Enumerates the set of values in String for ListCatalogItemsSortByEnum
func GetListCatalogItemsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"timeBackfillLastChecked",
	}
}

// GetMappingListCatalogItemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogItemsSortByEnum(val string) (ListCatalogItemsSortByEnum, bool) {
	enum, ok := mappingListCatalogItemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCatalogItemsCatalogListingVersionCriteriaEnum Enum with underlying type: string
type ListCatalogItemsCatalogListingVersionCriteriaEnum string

// Set of constants representing the allowable values for ListCatalogItemsCatalogListingVersionCriteriaEnum
const (
	ListCatalogItemsCatalogListingVersionCriteriaAllVersions         ListCatalogItemsCatalogListingVersionCriteriaEnum = "LIST_ALL_VERSIONS"
	ListCatalogItemsCatalogListingVersionCriteriaEarliestVersionOnly ListCatalogItemsCatalogListingVersionCriteriaEnum = "LIST_EARLIEST_VERSION_ONLY"
	ListCatalogItemsCatalogListingVersionCriteriaLatestVersionOnly   ListCatalogItemsCatalogListingVersionCriteriaEnum = "LIST_LATEST_VERSION_ONLY"
)

var mappingListCatalogItemsCatalogListingVersionCriteriaEnum = map[string]ListCatalogItemsCatalogListingVersionCriteriaEnum{
	"LIST_ALL_VERSIONS":          ListCatalogItemsCatalogListingVersionCriteriaAllVersions,
	"LIST_EARLIEST_VERSION_ONLY": ListCatalogItemsCatalogListingVersionCriteriaEarliestVersionOnly,
	"LIST_LATEST_VERSION_ONLY":   ListCatalogItemsCatalogListingVersionCriteriaLatestVersionOnly,
}

var mappingListCatalogItemsCatalogListingVersionCriteriaEnumLowerCase = map[string]ListCatalogItemsCatalogListingVersionCriteriaEnum{
	"list_all_versions":          ListCatalogItemsCatalogListingVersionCriteriaAllVersions,
	"list_earliest_version_only": ListCatalogItemsCatalogListingVersionCriteriaEarliestVersionOnly,
	"list_latest_version_only":   ListCatalogItemsCatalogListingVersionCriteriaLatestVersionOnly,
}

// GetListCatalogItemsCatalogListingVersionCriteriaEnumValues Enumerates the set of values for ListCatalogItemsCatalogListingVersionCriteriaEnum
func GetListCatalogItemsCatalogListingVersionCriteriaEnumValues() []ListCatalogItemsCatalogListingVersionCriteriaEnum {
	values := make([]ListCatalogItemsCatalogListingVersionCriteriaEnum, 0)
	for _, v := range mappingListCatalogItemsCatalogListingVersionCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogItemsCatalogListingVersionCriteriaEnumStringValues Enumerates the set of values in String for ListCatalogItemsCatalogListingVersionCriteriaEnum
func GetListCatalogItemsCatalogListingVersionCriteriaEnumStringValues() []string {
	return []string{
		"LIST_ALL_VERSIONS",
		"LIST_EARLIEST_VERSION_ONLY",
		"LIST_LATEST_VERSION_ONLY",
	}
}

// GetMappingListCatalogItemsCatalogListingVersionCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogItemsCatalogListingVersionCriteriaEnum(val string) (ListCatalogItemsCatalogListingVersionCriteriaEnum, bool) {
	enum, ok := mappingListCatalogItemsCatalogListingVersionCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
