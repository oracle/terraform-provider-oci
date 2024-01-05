// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContainerImagesRequest wrapper for the ListContainerImages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerImages.go.html to see an example of how to use ListContainerImagesRequest.
type ListContainerImagesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// inspected depending on the the setting of `accessLevel`.
	// Default is false. Can only be set to true when calling the API
	// on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return a container image summary only for the specified container image OCID.
	ImageId *string `mandatory:"false" contributesTo:"query" name:"imageId"`

	// A filter to return container images based on whether there are any associated versions.
	IsVersioned *bool `mandatory:"false" contributesTo:"query" name:"isVersioned"`

	// A filter to return container images only for the specified container repository OCID.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return container images or container image signatures that match the repository name.
	// Example: `foo` or `foo*`
	RepositoryName *string `mandatory:"false" contributesTo:"query" name:"repositoryName"`

	// A filter to return container images that match the version.
	// Example: `foo` or `foo*`
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// A filter to return only resources that match the given lifecycle state name exactly.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListContainerImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListContainerImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainerImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainerImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainerImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainerImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainerImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListContainerImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContainerImagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainerImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContainerImagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContainerImagesResponse wrapper for the ListContainerImages operation
type ListContainerImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerImageCollection instances
	ContainerImageCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListContainerImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainerImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainerImagesSortByEnum Enum with underlying type: string
type ListContainerImagesSortByEnum string

// Set of constants representing the allowable values for ListContainerImagesSortByEnum
const (
	ListContainerImagesSortByTimecreated ListContainerImagesSortByEnum = "TIMECREATED"
	ListContainerImagesSortByDisplayname ListContainerImagesSortByEnum = "DISPLAYNAME"
)

var mappingListContainerImagesSortByEnum = map[string]ListContainerImagesSortByEnum{
	"TIMECREATED": ListContainerImagesSortByTimecreated,
	"DISPLAYNAME": ListContainerImagesSortByDisplayname,
}

var mappingListContainerImagesSortByEnumLowerCase = map[string]ListContainerImagesSortByEnum{
	"timecreated": ListContainerImagesSortByTimecreated,
	"displayname": ListContainerImagesSortByDisplayname,
}

// GetListContainerImagesSortByEnumValues Enumerates the set of values for ListContainerImagesSortByEnum
func GetListContainerImagesSortByEnumValues() []ListContainerImagesSortByEnum {
	values := make([]ListContainerImagesSortByEnum, 0)
	for _, v := range mappingListContainerImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerImagesSortByEnumStringValues Enumerates the set of values in String for ListContainerImagesSortByEnum
func GetListContainerImagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListContainerImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerImagesSortByEnum(val string) (ListContainerImagesSortByEnum, bool) {
	enum, ok := mappingListContainerImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainerImagesSortOrderEnum Enum with underlying type: string
type ListContainerImagesSortOrderEnum string

// Set of constants representing the allowable values for ListContainerImagesSortOrderEnum
const (
	ListContainerImagesSortOrderAsc  ListContainerImagesSortOrderEnum = "ASC"
	ListContainerImagesSortOrderDesc ListContainerImagesSortOrderEnum = "DESC"
)

var mappingListContainerImagesSortOrderEnum = map[string]ListContainerImagesSortOrderEnum{
	"ASC":  ListContainerImagesSortOrderAsc,
	"DESC": ListContainerImagesSortOrderDesc,
}

var mappingListContainerImagesSortOrderEnumLowerCase = map[string]ListContainerImagesSortOrderEnum{
	"asc":  ListContainerImagesSortOrderAsc,
	"desc": ListContainerImagesSortOrderDesc,
}

// GetListContainerImagesSortOrderEnumValues Enumerates the set of values for ListContainerImagesSortOrderEnum
func GetListContainerImagesSortOrderEnumValues() []ListContainerImagesSortOrderEnum {
	values := make([]ListContainerImagesSortOrderEnum, 0)
	for _, v := range mappingListContainerImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerImagesSortOrderEnumStringValues Enumerates the set of values in String for ListContainerImagesSortOrderEnum
func GetListContainerImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContainerImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerImagesSortOrderEnum(val string) (ListContainerImagesSortOrderEnum, bool) {
	enum, ok := mappingListContainerImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
