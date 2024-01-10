// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLifecycleStageInstalledPackagesRequest wrapper for the ListLifecycleStageInstalledPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleStageInstalledPackages.go.html to see an example of how to use ListLifecycleStageInstalledPackagesRequest.
type ListLifecycleStageInstalledPackagesRequest struct {

	// The OCID of the lifecycle stage.
	LifecycleStageId *string `mandatory:"true" contributesTo:"path" name:"lifecycleStageId"`

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only lifecycle stage whose lifecycle state matches the given lifecycle state.
	LifecycleState LifecycleStageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLifecycleStageInstalledPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListLifecycleStageInstalledPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLifecycleStageInstalledPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLifecycleStageInstalledPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLifecycleStageInstalledPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLifecycleStageInstalledPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLifecycleStageInstalledPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLifecycleStageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleStageInstalledPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLifecycleStageInstalledPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleStageInstalledPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLifecycleStageInstalledPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLifecycleStageInstalledPackagesResponse wrapper for the ListLifecycleStageInstalledPackages operation
type ListLifecycleStageInstalledPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledPackageCollection instances
	InstalledPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLifecycleStageInstalledPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLifecycleStageInstalledPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLifecycleStageInstalledPackagesSortOrderEnum Enum with underlying type: string
type ListLifecycleStageInstalledPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListLifecycleStageInstalledPackagesSortOrderEnum
const (
	ListLifecycleStageInstalledPackagesSortOrderAsc  ListLifecycleStageInstalledPackagesSortOrderEnum = "ASC"
	ListLifecycleStageInstalledPackagesSortOrderDesc ListLifecycleStageInstalledPackagesSortOrderEnum = "DESC"
)

var mappingListLifecycleStageInstalledPackagesSortOrderEnum = map[string]ListLifecycleStageInstalledPackagesSortOrderEnum{
	"ASC":  ListLifecycleStageInstalledPackagesSortOrderAsc,
	"DESC": ListLifecycleStageInstalledPackagesSortOrderDesc,
}

var mappingListLifecycleStageInstalledPackagesSortOrderEnumLowerCase = map[string]ListLifecycleStageInstalledPackagesSortOrderEnum{
	"asc":  ListLifecycleStageInstalledPackagesSortOrderAsc,
	"desc": ListLifecycleStageInstalledPackagesSortOrderDesc,
}

// GetListLifecycleStageInstalledPackagesSortOrderEnumValues Enumerates the set of values for ListLifecycleStageInstalledPackagesSortOrderEnum
func GetListLifecycleStageInstalledPackagesSortOrderEnumValues() []ListLifecycleStageInstalledPackagesSortOrderEnum {
	values := make([]ListLifecycleStageInstalledPackagesSortOrderEnum, 0)
	for _, v := range mappingListLifecycleStageInstalledPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStageInstalledPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListLifecycleStageInstalledPackagesSortOrderEnum
func GetListLifecycleStageInstalledPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLifecycleStageInstalledPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStageInstalledPackagesSortOrderEnum(val string) (ListLifecycleStageInstalledPackagesSortOrderEnum, bool) {
	enum, ok := mappingListLifecycleStageInstalledPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleStageInstalledPackagesSortByEnum Enum with underlying type: string
type ListLifecycleStageInstalledPackagesSortByEnum string

// Set of constants representing the allowable values for ListLifecycleStageInstalledPackagesSortByEnum
const (
	ListLifecycleStageInstalledPackagesSortByTimecreated ListLifecycleStageInstalledPackagesSortByEnum = "timeCreated"
	ListLifecycleStageInstalledPackagesSortByDisplayname ListLifecycleStageInstalledPackagesSortByEnum = "displayName"
)

var mappingListLifecycleStageInstalledPackagesSortByEnum = map[string]ListLifecycleStageInstalledPackagesSortByEnum{
	"timeCreated": ListLifecycleStageInstalledPackagesSortByTimecreated,
	"displayName": ListLifecycleStageInstalledPackagesSortByDisplayname,
}

var mappingListLifecycleStageInstalledPackagesSortByEnumLowerCase = map[string]ListLifecycleStageInstalledPackagesSortByEnum{
	"timecreated": ListLifecycleStageInstalledPackagesSortByTimecreated,
	"displayname": ListLifecycleStageInstalledPackagesSortByDisplayname,
}

// GetListLifecycleStageInstalledPackagesSortByEnumValues Enumerates the set of values for ListLifecycleStageInstalledPackagesSortByEnum
func GetListLifecycleStageInstalledPackagesSortByEnumValues() []ListLifecycleStageInstalledPackagesSortByEnum {
	values := make([]ListLifecycleStageInstalledPackagesSortByEnum, 0)
	for _, v := range mappingListLifecycleStageInstalledPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStageInstalledPackagesSortByEnumStringValues Enumerates the set of values in String for ListLifecycleStageInstalledPackagesSortByEnum
func GetListLifecycleStageInstalledPackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLifecycleStageInstalledPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStageInstalledPackagesSortByEnum(val string) (ListLifecycleStageInstalledPackagesSortByEnum, bool) {
	enum, ok := mappingListLifecycleStageInstalledPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
