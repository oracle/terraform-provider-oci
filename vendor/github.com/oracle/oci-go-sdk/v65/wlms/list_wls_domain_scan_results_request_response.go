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

// ListWlsDomainScanResultsRequest wrapper for the ListWlsDomainScanResults operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainScanResults.go.html to see an example of how to use ListWlsDomainScanResultsRequest.
type ListWlsDomainScanResultsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The name of the server.
	ServerName *string `mandatory:"false" contributesTo:"query" name:"serverName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListWlsDomainScanResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeOfServerCheck_ is **descending**.
	// Default order for _serverName_ is **ascending**.
	// If no value is specified, _timeOfServerCheck_ is default.
	SortBy ListWlsDomainScanResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlsDomainScanResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlsDomainScanResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlsDomainScanResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlsDomainScanResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlsDomainScanResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlsDomainScanResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlsDomainScanResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainScanResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlsDomainScanResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlsDomainScanResultsResponse wrapper for the ListWlsDomainScanResults operation
type ListWlsDomainScanResultsResponse struct {

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

func (response ListWlsDomainScanResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlsDomainScanResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlsDomainScanResultsSortOrderEnum Enum with underlying type: string
type ListWlsDomainScanResultsSortOrderEnum string

// Set of constants representing the allowable values for ListWlsDomainScanResultsSortOrderEnum
const (
	ListWlsDomainScanResultsSortOrderAsc  ListWlsDomainScanResultsSortOrderEnum = "ASC"
	ListWlsDomainScanResultsSortOrderDesc ListWlsDomainScanResultsSortOrderEnum = "DESC"
)

var mappingListWlsDomainScanResultsSortOrderEnum = map[string]ListWlsDomainScanResultsSortOrderEnum{
	"ASC":  ListWlsDomainScanResultsSortOrderAsc,
	"DESC": ListWlsDomainScanResultsSortOrderDesc,
}

var mappingListWlsDomainScanResultsSortOrderEnumLowerCase = map[string]ListWlsDomainScanResultsSortOrderEnum{
	"asc":  ListWlsDomainScanResultsSortOrderAsc,
	"desc": ListWlsDomainScanResultsSortOrderDesc,
}

// GetListWlsDomainScanResultsSortOrderEnumValues Enumerates the set of values for ListWlsDomainScanResultsSortOrderEnum
func GetListWlsDomainScanResultsSortOrderEnumValues() []ListWlsDomainScanResultsSortOrderEnum {
	values := make([]ListWlsDomainScanResultsSortOrderEnum, 0)
	for _, v := range mappingListWlsDomainScanResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainScanResultsSortOrderEnumStringValues Enumerates the set of values in String for ListWlsDomainScanResultsSortOrderEnum
func GetListWlsDomainScanResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlsDomainScanResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainScanResultsSortOrderEnum(val string) (ListWlsDomainScanResultsSortOrderEnum, bool) {
	enum, ok := mappingListWlsDomainScanResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainScanResultsSortByEnum Enum with underlying type: string
type ListWlsDomainScanResultsSortByEnum string

// Set of constants representing the allowable values for ListWlsDomainScanResultsSortByEnum
const (
	ListWlsDomainScanResultsSortByTimeofservercheck ListWlsDomainScanResultsSortByEnum = "timeOfServerCheck"
	ListWlsDomainScanResultsSortByServername        ListWlsDomainScanResultsSortByEnum = "serverName"
)

var mappingListWlsDomainScanResultsSortByEnum = map[string]ListWlsDomainScanResultsSortByEnum{
	"timeOfServerCheck": ListWlsDomainScanResultsSortByTimeofservercheck,
	"serverName":        ListWlsDomainScanResultsSortByServername,
}

var mappingListWlsDomainScanResultsSortByEnumLowerCase = map[string]ListWlsDomainScanResultsSortByEnum{
	"timeofservercheck": ListWlsDomainScanResultsSortByTimeofservercheck,
	"servername":        ListWlsDomainScanResultsSortByServername,
}

// GetListWlsDomainScanResultsSortByEnumValues Enumerates the set of values for ListWlsDomainScanResultsSortByEnum
func GetListWlsDomainScanResultsSortByEnumValues() []ListWlsDomainScanResultsSortByEnum {
	values := make([]ListWlsDomainScanResultsSortByEnum, 0)
	for _, v := range mappingListWlsDomainScanResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainScanResultsSortByEnumStringValues Enumerates the set of values in String for ListWlsDomainScanResultsSortByEnum
func GetListWlsDomainScanResultsSortByEnumStringValues() []string {
	return []string{
		"timeOfServerCheck",
		"serverName",
	}
}

// GetMappingListWlsDomainScanResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainScanResultsSortByEnum(val string) (ListWlsDomainScanResultsSortByEnum, bool) {
	enum, ok := mappingListWlsDomainScanResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
