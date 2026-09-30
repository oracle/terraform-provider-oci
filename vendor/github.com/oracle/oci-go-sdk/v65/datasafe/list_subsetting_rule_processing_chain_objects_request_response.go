// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingRuleProcessingChainObjectsRequest wrapper for the ListSubsettingRuleProcessingChainObjects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingRuleProcessingChainObjects.go.html to see an example of how to use ListSubsettingRuleProcessingChainObjectsRequest.
type ListSubsettingRuleProcessingChainObjectsRequest struct {

	// The OCID of the subsetting policy.
	SubsettingPolicyId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyId"`

	// The unique key that identifies the subsetting rule. It's numeric and unique within a subsetting policy.
	SubsettingRuleKey *string `mandatory:"true" contributesTo:"path" name:"subsettingRuleKey"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for key is ascending.
	// The default order for other fields is ascending.
	SortBy ListSubsettingRuleProcessingChainObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return the processing chain objects which are enabled for processing.
	IsEnabledForProcessing *bool `mandatory:"false" contributesTo:"query" name:"isEnabledForProcessing"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingRuleProcessingChainObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingRuleProcessingChainObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingRuleProcessingChainObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingRuleProcessingChainObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingRuleProcessingChainObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingRuleProcessingChainObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingRuleProcessingChainObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingRuleProcessingChainObjectsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingRuleProcessingChainObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingRuleProcessingChainObjectsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingRuleProcessingChainObjectsResponse wrapper for the ListSubsettingRuleProcessingChainObjects operation
type ListSubsettingRuleProcessingChainObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingRuleProcessingChainObjectsCollection instances
	SubsettingRuleProcessingChainObjectsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingRuleProcessingChainObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingRuleProcessingChainObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingRuleProcessingChainObjectsSortByEnum Enum with underlying type: string
type ListSubsettingRuleProcessingChainObjectsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingRuleProcessingChainObjectsSortByEnum
const (
	ListSubsettingRuleProcessingChainObjectsSortByKey ListSubsettingRuleProcessingChainObjectsSortByEnum = "key"
)

var mappingListSubsettingRuleProcessingChainObjectsSortByEnum = map[string]ListSubsettingRuleProcessingChainObjectsSortByEnum{
	"key": ListSubsettingRuleProcessingChainObjectsSortByKey,
}

var mappingListSubsettingRuleProcessingChainObjectsSortByEnumLowerCase = map[string]ListSubsettingRuleProcessingChainObjectsSortByEnum{
	"key": ListSubsettingRuleProcessingChainObjectsSortByKey,
}

// GetListSubsettingRuleProcessingChainObjectsSortByEnumValues Enumerates the set of values for ListSubsettingRuleProcessingChainObjectsSortByEnum
func GetListSubsettingRuleProcessingChainObjectsSortByEnumValues() []ListSubsettingRuleProcessingChainObjectsSortByEnum {
	values := make([]ListSubsettingRuleProcessingChainObjectsSortByEnum, 0)
	for _, v := range mappingListSubsettingRuleProcessingChainObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingRuleProcessingChainObjectsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingRuleProcessingChainObjectsSortByEnum
func GetListSubsettingRuleProcessingChainObjectsSortByEnumStringValues() []string {
	return []string{
		"key",
	}
}

// GetMappingListSubsettingRuleProcessingChainObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingRuleProcessingChainObjectsSortByEnum(val string) (ListSubsettingRuleProcessingChainObjectsSortByEnum, bool) {
	enum, ok := mappingListSubsettingRuleProcessingChainObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingRuleProcessingChainObjectsSortOrderEnum Enum with underlying type: string
type ListSubsettingRuleProcessingChainObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingRuleProcessingChainObjectsSortOrderEnum
const (
	ListSubsettingRuleProcessingChainObjectsSortOrderAsc  ListSubsettingRuleProcessingChainObjectsSortOrderEnum = "ASC"
	ListSubsettingRuleProcessingChainObjectsSortOrderDesc ListSubsettingRuleProcessingChainObjectsSortOrderEnum = "DESC"
)

var mappingListSubsettingRuleProcessingChainObjectsSortOrderEnum = map[string]ListSubsettingRuleProcessingChainObjectsSortOrderEnum{
	"ASC":  ListSubsettingRuleProcessingChainObjectsSortOrderAsc,
	"DESC": ListSubsettingRuleProcessingChainObjectsSortOrderDesc,
}

var mappingListSubsettingRuleProcessingChainObjectsSortOrderEnumLowerCase = map[string]ListSubsettingRuleProcessingChainObjectsSortOrderEnum{
	"asc":  ListSubsettingRuleProcessingChainObjectsSortOrderAsc,
	"desc": ListSubsettingRuleProcessingChainObjectsSortOrderDesc,
}

// GetListSubsettingRuleProcessingChainObjectsSortOrderEnumValues Enumerates the set of values for ListSubsettingRuleProcessingChainObjectsSortOrderEnum
func GetListSubsettingRuleProcessingChainObjectsSortOrderEnumValues() []ListSubsettingRuleProcessingChainObjectsSortOrderEnum {
	values := make([]ListSubsettingRuleProcessingChainObjectsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingRuleProcessingChainObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingRuleProcessingChainObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingRuleProcessingChainObjectsSortOrderEnum
func GetListSubsettingRuleProcessingChainObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingRuleProcessingChainObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingRuleProcessingChainObjectsSortOrderEnum(val string) (ListSubsettingRuleProcessingChainObjectsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingRuleProcessingChainObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
