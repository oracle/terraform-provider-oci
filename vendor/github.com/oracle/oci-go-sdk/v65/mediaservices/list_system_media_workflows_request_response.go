// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSystemMediaWorkflowsRequest wrapper for the ListSystemMediaWorkflows operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListSystemMediaWorkflows.go.html to see an example of how to use ListSystemMediaWorkflowsRequest.
type ListSystemMediaWorkflowsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources with their system defined, unique name matching the given name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSystemMediaWorkflowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSystemMediaWorkflowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSystemMediaWorkflowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSystemMediaWorkflowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSystemMediaWorkflowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSystemMediaWorkflowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSystemMediaWorkflowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSystemMediaWorkflowsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSystemMediaWorkflowsResponse wrapper for the ListSystemMediaWorkflows operation
type ListSystemMediaWorkflowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SystemMediaWorkflowCollection instances
	SystemMediaWorkflowCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSystemMediaWorkflowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSystemMediaWorkflowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSystemMediaWorkflowsSortOrderEnum Enum with underlying type: string
type ListSystemMediaWorkflowsSortOrderEnum string

// Set of constants representing the allowable values for ListSystemMediaWorkflowsSortOrderEnum
const (
	ListSystemMediaWorkflowsSortOrderAsc  ListSystemMediaWorkflowsSortOrderEnum = "ASC"
	ListSystemMediaWorkflowsSortOrderDesc ListSystemMediaWorkflowsSortOrderEnum = "DESC"
)

var mappingListSystemMediaWorkflowsSortOrderEnum = map[string]ListSystemMediaWorkflowsSortOrderEnum{
	"ASC":  ListSystemMediaWorkflowsSortOrderAsc,
	"DESC": ListSystemMediaWorkflowsSortOrderDesc,
}

var mappingListSystemMediaWorkflowsSortOrderEnumLowerCase = map[string]ListSystemMediaWorkflowsSortOrderEnum{
	"asc":  ListSystemMediaWorkflowsSortOrderAsc,
	"desc": ListSystemMediaWorkflowsSortOrderDesc,
}

// GetListSystemMediaWorkflowsSortOrderEnumValues Enumerates the set of values for ListSystemMediaWorkflowsSortOrderEnum
func GetListSystemMediaWorkflowsSortOrderEnumValues() []ListSystemMediaWorkflowsSortOrderEnum {
	values := make([]ListSystemMediaWorkflowsSortOrderEnum, 0)
	for _, v := range mappingListSystemMediaWorkflowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSystemMediaWorkflowsSortOrderEnumStringValues Enumerates the set of values in String for ListSystemMediaWorkflowsSortOrderEnum
func GetListSystemMediaWorkflowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSystemMediaWorkflowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSystemMediaWorkflowsSortOrderEnum(val string) (ListSystemMediaWorkflowsSortOrderEnum, bool) {
	enum, ok := mappingListSystemMediaWorkflowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
