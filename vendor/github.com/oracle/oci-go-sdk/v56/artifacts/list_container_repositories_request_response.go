// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListContainerRepositoriesRequest wrapper for the ListContainerRepositories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerRepositories.go.html to see an example of how to use ListContainerRepositoriesRequest.
type ListContainerRepositoriesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// inspected depending on the the setting of `accessLevel`.
	// Default is false. Can only be set to true when calling the API
	// on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return container images only for the specified container repository OCID.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that match the isPublic value.
	IsPublic *bool `mandatory:"false" contributesTo:"query" name:"isPublic"`

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
	SortBy ListContainerRepositoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListContainerRepositoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainerRepositoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainerRepositoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainerRepositoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainerRepositoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListContainerRepositoriesResponse wrapper for the ListContainerRepositories operation
type ListContainerRepositoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerRepositoryCollection instances
	ContainerRepositoryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListContainerRepositoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainerRepositoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainerRepositoriesSortByEnum Enum with underlying type: string
type ListContainerRepositoriesSortByEnum string

// Set of constants representing the allowable values for ListContainerRepositoriesSortByEnum
const (
	ListContainerRepositoriesSortByTimecreated ListContainerRepositoriesSortByEnum = "TIMECREATED"
	ListContainerRepositoriesSortByDisplayname ListContainerRepositoriesSortByEnum = "DISPLAYNAME"
)

var mappingListContainerRepositoriesSortBy = map[string]ListContainerRepositoriesSortByEnum{
	"TIMECREATED": ListContainerRepositoriesSortByTimecreated,
	"DISPLAYNAME": ListContainerRepositoriesSortByDisplayname,
}

// GetListContainerRepositoriesSortByEnumValues Enumerates the set of values for ListContainerRepositoriesSortByEnum
func GetListContainerRepositoriesSortByEnumValues() []ListContainerRepositoriesSortByEnum {
	values := make([]ListContainerRepositoriesSortByEnum, 0)
	for _, v := range mappingListContainerRepositoriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListContainerRepositoriesSortOrderEnum Enum with underlying type: string
type ListContainerRepositoriesSortOrderEnum string

// Set of constants representing the allowable values for ListContainerRepositoriesSortOrderEnum
const (
	ListContainerRepositoriesSortOrderAsc  ListContainerRepositoriesSortOrderEnum = "ASC"
	ListContainerRepositoriesSortOrderDesc ListContainerRepositoriesSortOrderEnum = "DESC"
)

var mappingListContainerRepositoriesSortOrder = map[string]ListContainerRepositoriesSortOrderEnum{
	"ASC":  ListContainerRepositoriesSortOrderAsc,
	"DESC": ListContainerRepositoriesSortOrderDesc,
}

// GetListContainerRepositoriesSortOrderEnumValues Enumerates the set of values for ListContainerRepositoriesSortOrderEnum
func GetListContainerRepositoriesSortOrderEnumValues() []ListContainerRepositoriesSortOrderEnum {
	values := make([]ListContainerRepositoriesSortOrderEnum, 0)
	for _, v := range mappingListContainerRepositoriesSortOrder {
		values = append(values, v)
	}
	return values
}
