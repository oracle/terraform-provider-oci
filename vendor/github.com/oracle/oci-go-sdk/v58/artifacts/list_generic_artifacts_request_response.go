// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListGenericArtifactsRequest wrapper for the ListGenericArtifacts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListGenericArtifacts.go.html to see an example of how to use ListGenericArtifactsRequest.
type ListGenericArtifactsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return the artifacts only for the specified repository OCID.
	RepositoryId *string `mandatory:"true" contributesTo:"query" name:"repositoryId"`

	// A filter to return the resources for the specified OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by a prefix for the `artifactPath` and and return artifacts that begin with the specified prefix in their path.
	ArtifactPath *string `mandatory:"false" contributesTo:"query" name:"artifactPath"`

	// Filter results by a prefix for `version` and return artifacts that that begin with the specified prefix in their version.
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// Filter results by a specified SHA256 digest for the artifact.
	Sha256 *string `mandatory:"false" contributesTo:"query" name:"sha256"`

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
	SortBy ListGenericArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListGenericArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGenericArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGenericArtifactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGenericArtifactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGenericArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGenericArtifactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGenericArtifactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGenericArtifactsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGenericArtifactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGenericArtifactsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGenericArtifactsResponse wrapper for the ListGenericArtifacts operation
type ListGenericArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GenericArtifactCollection instances
	GenericArtifactCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListGenericArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGenericArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGenericArtifactsSortByEnum Enum with underlying type: string
type ListGenericArtifactsSortByEnum string

// Set of constants representing the allowable values for ListGenericArtifactsSortByEnum
const (
	ListGenericArtifactsSortByTimecreated ListGenericArtifactsSortByEnum = "TIMECREATED"
	ListGenericArtifactsSortByDisplayname ListGenericArtifactsSortByEnum = "DISPLAYNAME"
)

var mappingListGenericArtifactsSortByEnum = map[string]ListGenericArtifactsSortByEnum{
	"TIMECREATED": ListGenericArtifactsSortByTimecreated,
	"DISPLAYNAME": ListGenericArtifactsSortByDisplayname,
}

// GetListGenericArtifactsSortByEnumValues Enumerates the set of values for ListGenericArtifactsSortByEnum
func GetListGenericArtifactsSortByEnumValues() []ListGenericArtifactsSortByEnum {
	values := make([]ListGenericArtifactsSortByEnum, 0)
	for _, v := range mappingListGenericArtifactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenericArtifactsSortByEnumStringValues Enumerates the set of values in String for ListGenericArtifactsSortByEnum
func GetListGenericArtifactsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListGenericArtifactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenericArtifactsSortByEnum(val string) (ListGenericArtifactsSortByEnum, bool) {
	mappingListGenericArtifactsSortByEnumIgnoreCase := make(map[string]ListGenericArtifactsSortByEnum)
	for k, v := range mappingListGenericArtifactsSortByEnum {
		mappingListGenericArtifactsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListGenericArtifactsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListGenericArtifactsSortOrderEnum Enum with underlying type: string
type ListGenericArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListGenericArtifactsSortOrderEnum
const (
	ListGenericArtifactsSortOrderAsc  ListGenericArtifactsSortOrderEnum = "ASC"
	ListGenericArtifactsSortOrderDesc ListGenericArtifactsSortOrderEnum = "DESC"
)

var mappingListGenericArtifactsSortOrderEnum = map[string]ListGenericArtifactsSortOrderEnum{
	"ASC":  ListGenericArtifactsSortOrderAsc,
	"DESC": ListGenericArtifactsSortOrderDesc,
}

// GetListGenericArtifactsSortOrderEnumValues Enumerates the set of values for ListGenericArtifactsSortOrderEnum
func GetListGenericArtifactsSortOrderEnumValues() []ListGenericArtifactsSortOrderEnum {
	values := make([]ListGenericArtifactsSortOrderEnum, 0)
	for _, v := range mappingListGenericArtifactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenericArtifactsSortOrderEnumStringValues Enumerates the set of values in String for ListGenericArtifactsSortOrderEnum
func GetListGenericArtifactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGenericArtifactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenericArtifactsSortOrderEnum(val string) (ListGenericArtifactsSortOrderEnum, bool) {
	mappingListGenericArtifactsSortOrderEnumIgnoreCase := make(map[string]ListGenericArtifactsSortOrderEnum)
	for k, v := range mappingListGenericArtifactsSortOrderEnum {
		mappingListGenericArtifactsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListGenericArtifactsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
