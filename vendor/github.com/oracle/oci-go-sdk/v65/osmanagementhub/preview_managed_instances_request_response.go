// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// PreviewManagedInstancesRequest wrapper for the PreviewManagedInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/PreviewManagedInstances.go.html to see an example of how to use PreviewManagedInstancesRequest.
type PreviewManagedInstancesRequest struct {

	// Provides the information used to Preview the dynamic set.
	PreviewManagedInstancesDetails `contributesTo:"body"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Indicates whether to include subcompartments in the returned results. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder PreviewManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy PreviewManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PreviewManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PreviewManagedInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PreviewManagedInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PreviewManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PreviewManagedInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPreviewManagedInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetPreviewManagedInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPreviewManagedInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetPreviewManagedInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PreviewManagedInstancesResponse wrapper for the PreviewManagedInstances operation
type PreviewManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceCollection instances
	ManagedInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items in the result. Used for pagination of a list of items.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response PreviewManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PreviewManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// PreviewManagedInstancesSortOrderEnum Enum with underlying type: string
type PreviewManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for PreviewManagedInstancesSortOrderEnum
const (
	PreviewManagedInstancesSortOrderAsc  PreviewManagedInstancesSortOrderEnum = "ASC"
	PreviewManagedInstancesSortOrderDesc PreviewManagedInstancesSortOrderEnum = "DESC"
)

var mappingPreviewManagedInstancesSortOrderEnum = map[string]PreviewManagedInstancesSortOrderEnum{
	"ASC":  PreviewManagedInstancesSortOrderAsc,
	"DESC": PreviewManagedInstancesSortOrderDesc,
}

var mappingPreviewManagedInstancesSortOrderEnumLowerCase = map[string]PreviewManagedInstancesSortOrderEnum{
	"asc":  PreviewManagedInstancesSortOrderAsc,
	"desc": PreviewManagedInstancesSortOrderDesc,
}

// GetPreviewManagedInstancesSortOrderEnumValues Enumerates the set of values for PreviewManagedInstancesSortOrderEnum
func GetPreviewManagedInstancesSortOrderEnumValues() []PreviewManagedInstancesSortOrderEnum {
	values := make([]PreviewManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingPreviewManagedInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetPreviewManagedInstancesSortOrderEnumStringValues Enumerates the set of values in String for PreviewManagedInstancesSortOrderEnum
func GetPreviewManagedInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingPreviewManagedInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPreviewManagedInstancesSortOrderEnum(val string) (PreviewManagedInstancesSortOrderEnum, bool) {
	enum, ok := mappingPreviewManagedInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PreviewManagedInstancesSortByEnum Enum with underlying type: string
type PreviewManagedInstancesSortByEnum string

// Set of constants representing the allowable values for PreviewManagedInstancesSortByEnum
const (
	PreviewManagedInstancesSortByTimecreated PreviewManagedInstancesSortByEnum = "timeCreated"
	PreviewManagedInstancesSortByDisplayname PreviewManagedInstancesSortByEnum = "displayName"
)

var mappingPreviewManagedInstancesSortByEnum = map[string]PreviewManagedInstancesSortByEnum{
	"timeCreated": PreviewManagedInstancesSortByTimecreated,
	"displayName": PreviewManagedInstancesSortByDisplayname,
}

var mappingPreviewManagedInstancesSortByEnumLowerCase = map[string]PreviewManagedInstancesSortByEnum{
	"timecreated": PreviewManagedInstancesSortByTimecreated,
	"displayname": PreviewManagedInstancesSortByDisplayname,
}

// GetPreviewManagedInstancesSortByEnumValues Enumerates the set of values for PreviewManagedInstancesSortByEnum
func GetPreviewManagedInstancesSortByEnumValues() []PreviewManagedInstancesSortByEnum {
	values := make([]PreviewManagedInstancesSortByEnum, 0)
	for _, v := range mappingPreviewManagedInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPreviewManagedInstancesSortByEnumStringValues Enumerates the set of values in String for PreviewManagedInstancesSortByEnum
func GetPreviewManagedInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingPreviewManagedInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPreviewManagedInstancesSortByEnum(val string) (PreviewManagedInstancesSortByEnum, bool) {
	enum, ok := mappingPreviewManagedInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
