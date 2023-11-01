// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceUpdatablePackagesRequest wrapper for the ListManagedInstanceUpdatablePackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceUpdatablePackages.go.html to see an example of how to use ListManagedInstanceUpdatablePackagesRequest.
type ListManagedInstanceUpdatablePackagesRequest struct {

	// The OCID of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	AdvisoryName []string `contributesTo:"query" name:"advisoryName" collectionFormat:"multi"`

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceUpdatablePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceUpdatablePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceUpdatablePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceUpdatablePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceUpdatablePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceUpdatablePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceUpdatablePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListManagedInstanceUpdatablePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceUpdatablePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceUpdatablePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceUpdatablePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceUpdatablePackagesResponse wrapper for the ListManagedInstanceUpdatablePackages operation
type ListManagedInstanceUpdatablePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UpdatablePackageCollection instances
	UpdatablePackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceUpdatablePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceUpdatablePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceUpdatablePackagesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceUpdatablePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceUpdatablePackagesSortOrderEnum
const (
	ListManagedInstanceUpdatablePackagesSortOrderAsc  ListManagedInstanceUpdatablePackagesSortOrderEnum = "ASC"
	ListManagedInstanceUpdatablePackagesSortOrderDesc ListManagedInstanceUpdatablePackagesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceUpdatablePackagesSortOrderEnum = map[string]ListManagedInstanceUpdatablePackagesSortOrderEnum{
	"ASC":  ListManagedInstanceUpdatablePackagesSortOrderAsc,
	"DESC": ListManagedInstanceUpdatablePackagesSortOrderDesc,
}

var mappingListManagedInstanceUpdatablePackagesSortOrderEnumLowerCase = map[string]ListManagedInstanceUpdatablePackagesSortOrderEnum{
	"asc":  ListManagedInstanceUpdatablePackagesSortOrderAsc,
	"desc": ListManagedInstanceUpdatablePackagesSortOrderDesc,
}

// GetListManagedInstanceUpdatablePackagesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceUpdatablePackagesSortOrderEnum
func GetListManagedInstanceUpdatablePackagesSortOrderEnumValues() []ListManagedInstanceUpdatablePackagesSortOrderEnum {
	values := make([]ListManagedInstanceUpdatablePackagesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceUpdatablePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceUpdatablePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceUpdatablePackagesSortOrderEnum
func GetListManagedInstanceUpdatablePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceUpdatablePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceUpdatablePackagesSortOrderEnum(val string) (ListManagedInstanceUpdatablePackagesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceUpdatablePackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceUpdatablePackagesSortByEnum Enum with underlying type: string
type ListManagedInstanceUpdatablePackagesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceUpdatablePackagesSortByEnum
const (
	ListManagedInstanceUpdatablePackagesSortByTimecreated ListManagedInstanceUpdatablePackagesSortByEnum = "timeCreated"
	ListManagedInstanceUpdatablePackagesSortByDisplayname ListManagedInstanceUpdatablePackagesSortByEnum = "displayName"
)

var mappingListManagedInstanceUpdatablePackagesSortByEnum = map[string]ListManagedInstanceUpdatablePackagesSortByEnum{
	"timeCreated": ListManagedInstanceUpdatablePackagesSortByTimecreated,
	"displayName": ListManagedInstanceUpdatablePackagesSortByDisplayname,
}

var mappingListManagedInstanceUpdatablePackagesSortByEnumLowerCase = map[string]ListManagedInstanceUpdatablePackagesSortByEnum{
	"timecreated": ListManagedInstanceUpdatablePackagesSortByTimecreated,
	"displayname": ListManagedInstanceUpdatablePackagesSortByDisplayname,
}

// GetListManagedInstanceUpdatablePackagesSortByEnumValues Enumerates the set of values for ListManagedInstanceUpdatablePackagesSortByEnum
func GetListManagedInstanceUpdatablePackagesSortByEnumValues() []ListManagedInstanceUpdatablePackagesSortByEnum {
	values := make([]ListManagedInstanceUpdatablePackagesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceUpdatablePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceUpdatablePackagesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceUpdatablePackagesSortByEnum
func GetListManagedInstanceUpdatablePackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceUpdatablePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceUpdatablePackagesSortByEnum(val string) (ListManagedInstanceUpdatablePackagesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceUpdatablePackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
