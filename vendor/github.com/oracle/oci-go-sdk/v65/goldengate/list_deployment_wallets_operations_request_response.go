// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentWalletsOperationsRequest wrapper for the ListDeploymentWalletsOperations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentWalletsOperations.go.html to see an example of how to use ListDeploymentWalletsOperationsRequest.
type ListDeploymentWalletsOperationsRequest struct {

	// A unique Deployment identifier.
	DeploymentId *string `mandatory:"true" contributesTo:"path" name:"deploymentId"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeStarted' is
	// descending.
	SortBy ListDeploymentWalletsOperationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentWalletsOperationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentWalletsOperationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentWalletsOperationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentWalletsOperationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentWalletsOperationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentWalletsOperationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentWalletsOperationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentWalletsOperationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentWalletsOperationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentWalletsOperationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentWalletsOperationsResponse wrapper for the ListDeploymentWalletsOperations operation
type ListDeploymentWalletsOperationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentWalletsOperationCollection instances
	DeploymentWalletsOperationCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentWalletsOperationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentWalletsOperationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentWalletsOperationsSortByEnum Enum with underlying type: string
type ListDeploymentWalletsOperationsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentWalletsOperationsSortByEnum
const (
	ListDeploymentWalletsOperationsSortByTimestarted ListDeploymentWalletsOperationsSortByEnum = "timeStarted"
)

var mappingListDeploymentWalletsOperationsSortByEnum = map[string]ListDeploymentWalletsOperationsSortByEnum{
	"timeStarted": ListDeploymentWalletsOperationsSortByTimestarted,
}

var mappingListDeploymentWalletsOperationsSortByEnumLowerCase = map[string]ListDeploymentWalletsOperationsSortByEnum{
	"timestarted": ListDeploymentWalletsOperationsSortByTimestarted,
}

// GetListDeploymentWalletsOperationsSortByEnumValues Enumerates the set of values for ListDeploymentWalletsOperationsSortByEnum
func GetListDeploymentWalletsOperationsSortByEnumValues() []ListDeploymentWalletsOperationsSortByEnum {
	values := make([]ListDeploymentWalletsOperationsSortByEnum, 0)
	for _, v := range mappingListDeploymentWalletsOperationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentWalletsOperationsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentWalletsOperationsSortByEnum
func GetListDeploymentWalletsOperationsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
	}
}

// GetMappingListDeploymentWalletsOperationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentWalletsOperationsSortByEnum(val string) (ListDeploymentWalletsOperationsSortByEnum, bool) {
	enum, ok := mappingListDeploymentWalletsOperationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentWalletsOperationsSortOrderEnum Enum with underlying type: string
type ListDeploymentWalletsOperationsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentWalletsOperationsSortOrderEnum
const (
	ListDeploymentWalletsOperationsSortOrderAsc  ListDeploymentWalletsOperationsSortOrderEnum = "ASC"
	ListDeploymentWalletsOperationsSortOrderDesc ListDeploymentWalletsOperationsSortOrderEnum = "DESC"
)

var mappingListDeploymentWalletsOperationsSortOrderEnum = map[string]ListDeploymentWalletsOperationsSortOrderEnum{
	"ASC":  ListDeploymentWalletsOperationsSortOrderAsc,
	"DESC": ListDeploymentWalletsOperationsSortOrderDesc,
}

var mappingListDeploymentWalletsOperationsSortOrderEnumLowerCase = map[string]ListDeploymentWalletsOperationsSortOrderEnum{
	"asc":  ListDeploymentWalletsOperationsSortOrderAsc,
	"desc": ListDeploymentWalletsOperationsSortOrderDesc,
}

// GetListDeploymentWalletsOperationsSortOrderEnumValues Enumerates the set of values for ListDeploymentWalletsOperationsSortOrderEnum
func GetListDeploymentWalletsOperationsSortOrderEnumValues() []ListDeploymentWalletsOperationsSortOrderEnum {
	values := make([]ListDeploymentWalletsOperationsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentWalletsOperationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentWalletsOperationsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentWalletsOperationsSortOrderEnum
func GetListDeploymentWalletsOperationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentWalletsOperationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentWalletsOperationsSortOrderEnum(val string) (ListDeploymentWalletsOperationsSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentWalletsOperationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
