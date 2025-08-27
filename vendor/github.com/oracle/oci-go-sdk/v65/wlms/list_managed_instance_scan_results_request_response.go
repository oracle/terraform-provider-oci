// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceScanResultsRequest wrapper for the ListManagedInstanceScanResults operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceScanResults.go.html to see an example of how to use ListManagedInstanceScanResultsRequest.
type ListManagedInstanceScanResultsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceScanResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeOfServerCheck_ is **descending**.
	// Default order for _serverName_ is **ascending**.
	// If no value is specified, _timeOfServerCheck_ is default.
	SortBy ListManagedInstanceScanResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"false" contributesTo:"query" name:"wlsDomainId"`

	// The name of the server.
	ServerName *string `mandatory:"false" contributesTo:"query" name:"serverName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceScanResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceScanResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceScanResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceScanResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceScanResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceScanResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceScanResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceScanResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceScanResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceScanResultsResponse wrapper for the ListManagedInstanceScanResults operation
type ListManagedInstanceScanResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScanResultCollection instances
	ScanResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceScanResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceScanResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceScanResultsSortOrderEnum Enum with underlying type: string
type ListManagedInstanceScanResultsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceScanResultsSortOrderEnum
const (
	ListManagedInstanceScanResultsSortOrderAsc  ListManagedInstanceScanResultsSortOrderEnum = "ASC"
	ListManagedInstanceScanResultsSortOrderDesc ListManagedInstanceScanResultsSortOrderEnum = "DESC"
)

var mappingListManagedInstanceScanResultsSortOrderEnum = map[string]ListManagedInstanceScanResultsSortOrderEnum{
	"ASC":  ListManagedInstanceScanResultsSortOrderAsc,
	"DESC": ListManagedInstanceScanResultsSortOrderDesc,
}

var mappingListManagedInstanceScanResultsSortOrderEnumLowerCase = map[string]ListManagedInstanceScanResultsSortOrderEnum{
	"asc":  ListManagedInstanceScanResultsSortOrderAsc,
	"desc": ListManagedInstanceScanResultsSortOrderDesc,
}

// GetListManagedInstanceScanResultsSortOrderEnumValues Enumerates the set of values for ListManagedInstanceScanResultsSortOrderEnum
func GetListManagedInstanceScanResultsSortOrderEnumValues() []ListManagedInstanceScanResultsSortOrderEnum {
	values := make([]ListManagedInstanceScanResultsSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceScanResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceScanResultsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceScanResultsSortOrderEnum
func GetListManagedInstanceScanResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceScanResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceScanResultsSortOrderEnum(val string) (ListManagedInstanceScanResultsSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceScanResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceScanResultsSortByEnum Enum with underlying type: string
type ListManagedInstanceScanResultsSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceScanResultsSortByEnum
const (
	ListManagedInstanceScanResultsSortByTimeofservercheck ListManagedInstanceScanResultsSortByEnum = "timeOfServerCheck"
	ListManagedInstanceScanResultsSortByServername        ListManagedInstanceScanResultsSortByEnum = "serverName"
)

var mappingListManagedInstanceScanResultsSortByEnum = map[string]ListManagedInstanceScanResultsSortByEnum{
	"timeOfServerCheck": ListManagedInstanceScanResultsSortByTimeofservercheck,
	"serverName":        ListManagedInstanceScanResultsSortByServername,
}

var mappingListManagedInstanceScanResultsSortByEnumLowerCase = map[string]ListManagedInstanceScanResultsSortByEnum{
	"timeofservercheck": ListManagedInstanceScanResultsSortByTimeofservercheck,
	"servername":        ListManagedInstanceScanResultsSortByServername,
}

// GetListManagedInstanceScanResultsSortByEnumValues Enumerates the set of values for ListManagedInstanceScanResultsSortByEnum
func GetListManagedInstanceScanResultsSortByEnumValues() []ListManagedInstanceScanResultsSortByEnum {
	values := make([]ListManagedInstanceScanResultsSortByEnum, 0)
	for _, v := range mappingListManagedInstanceScanResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceScanResultsSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceScanResultsSortByEnum
func GetListManagedInstanceScanResultsSortByEnumStringValues() []string {
	return []string{
		"timeOfServerCheck",
		"serverName",
	}
}

// GetMappingListManagedInstanceScanResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceScanResultsSortByEnum(val string) (ListManagedInstanceScanResultsSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceScanResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
