// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelDeploymentModelStatesRequest wrapper for the ListModelDeploymentModelStates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelDeploymentModelStates.go.html to see an example of how to use ListModelDeploymentModelStatesRequest.
type ListModelDeploymentModelStatesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
	ModelDeploymentId *string `mandatory:"true" contributesTo:"path" name:"modelDeploymentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the inference key.
	InferenceKey *string `mandatory:"false" contributesTo:"query" name:"inferenceKey"`

	// <b>Filter</b> results by the model ocid.
	ModelId *string `mandatory:"false" contributesTo:"query" name:"modelId"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListModelDeploymentModelStatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, results are shown
	// in descending order. When you sort by `displayName`, results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelDeploymentModelStatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelDeploymentModelStatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelDeploymentModelStatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelDeploymentModelStatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelDeploymentModelStatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelDeploymentModelStatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelDeploymentModelStatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelDeploymentModelStatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelDeploymentModelStatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelDeploymentModelStatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelDeploymentModelStatesResponse wrapper for the ListModelDeploymentModelStates operation
type ListModelDeploymentModelStatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelDeploymentModelStateSummary instances
	Items []ModelDeploymentModelStateSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelDeploymentModelStatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelDeploymentModelStatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelDeploymentModelStatesSortOrderEnum Enum with underlying type: string
type ListModelDeploymentModelStatesSortOrderEnum string

// Set of constants representing the allowable values for ListModelDeploymentModelStatesSortOrderEnum
const (
	ListModelDeploymentModelStatesSortOrderAsc  ListModelDeploymentModelStatesSortOrderEnum = "ASC"
	ListModelDeploymentModelStatesSortOrderDesc ListModelDeploymentModelStatesSortOrderEnum = "DESC"
)

var mappingListModelDeploymentModelStatesSortOrderEnum = map[string]ListModelDeploymentModelStatesSortOrderEnum{
	"ASC":  ListModelDeploymentModelStatesSortOrderAsc,
	"DESC": ListModelDeploymentModelStatesSortOrderDesc,
}

var mappingListModelDeploymentModelStatesSortOrderEnumLowerCase = map[string]ListModelDeploymentModelStatesSortOrderEnum{
	"asc":  ListModelDeploymentModelStatesSortOrderAsc,
	"desc": ListModelDeploymentModelStatesSortOrderDesc,
}

// GetListModelDeploymentModelStatesSortOrderEnumValues Enumerates the set of values for ListModelDeploymentModelStatesSortOrderEnum
func GetListModelDeploymentModelStatesSortOrderEnumValues() []ListModelDeploymentModelStatesSortOrderEnum {
	values := make([]ListModelDeploymentModelStatesSortOrderEnum, 0)
	for _, v := range mappingListModelDeploymentModelStatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelDeploymentModelStatesSortOrderEnumStringValues Enumerates the set of values in String for ListModelDeploymentModelStatesSortOrderEnum
func GetListModelDeploymentModelStatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelDeploymentModelStatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelDeploymentModelStatesSortOrderEnum(val string) (ListModelDeploymentModelStatesSortOrderEnum, bool) {
	enum, ok := mappingListModelDeploymentModelStatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelDeploymentModelStatesSortByEnum Enum with underlying type: string
type ListModelDeploymentModelStatesSortByEnum string

// Set of constants representing the allowable values for ListModelDeploymentModelStatesSortByEnum
const (
	ListModelDeploymentModelStatesSortByTimecreated ListModelDeploymentModelStatesSortByEnum = "timeCreated"
	ListModelDeploymentModelStatesSortByDisplayname ListModelDeploymentModelStatesSortByEnum = "displayName"
)

var mappingListModelDeploymentModelStatesSortByEnum = map[string]ListModelDeploymentModelStatesSortByEnum{
	"timeCreated": ListModelDeploymentModelStatesSortByTimecreated,
	"displayName": ListModelDeploymentModelStatesSortByDisplayname,
}

var mappingListModelDeploymentModelStatesSortByEnumLowerCase = map[string]ListModelDeploymentModelStatesSortByEnum{
	"timecreated": ListModelDeploymentModelStatesSortByTimecreated,
	"displayname": ListModelDeploymentModelStatesSortByDisplayname,
}

// GetListModelDeploymentModelStatesSortByEnumValues Enumerates the set of values for ListModelDeploymentModelStatesSortByEnum
func GetListModelDeploymentModelStatesSortByEnumValues() []ListModelDeploymentModelStatesSortByEnum {
	values := make([]ListModelDeploymentModelStatesSortByEnum, 0)
	for _, v := range mappingListModelDeploymentModelStatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelDeploymentModelStatesSortByEnumStringValues Enumerates the set of values in String for ListModelDeploymentModelStatesSortByEnum
func GetListModelDeploymentModelStatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListModelDeploymentModelStatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelDeploymentModelStatesSortByEnum(val string) (ListModelDeploymentModelStatesSortByEnum, bool) {
	enum, ok := mappingListModelDeploymentModelStatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
