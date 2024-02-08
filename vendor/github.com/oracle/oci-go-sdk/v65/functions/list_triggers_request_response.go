// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTriggersRequest wrapper for the ListTriggers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListTriggers.go.html to see an example of how to use ListTriggersRequest.
type ListTriggersRequest struct {

	// A filter to return only resources that match the service trigger source of a PBF.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListTriggersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTriggersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTriggersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTriggersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTriggersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTriggersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTriggersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTriggersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTriggersResponse wrapper for the ListTriggers operation
type ListTriggersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TriggersCollection instances
	TriggersCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTriggersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTriggersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTriggersSortOrderEnum Enum with underlying type: string
type ListTriggersSortOrderEnum string

// Set of constants representing the allowable values for ListTriggersSortOrderEnum
const (
	ListTriggersSortOrderAsc  ListTriggersSortOrderEnum = "ASC"
	ListTriggersSortOrderDesc ListTriggersSortOrderEnum = "DESC"
)

var mappingListTriggersSortOrderEnum = map[string]ListTriggersSortOrderEnum{
	"ASC":  ListTriggersSortOrderAsc,
	"DESC": ListTriggersSortOrderDesc,
}

var mappingListTriggersSortOrderEnumLowerCase = map[string]ListTriggersSortOrderEnum{
	"asc":  ListTriggersSortOrderAsc,
	"desc": ListTriggersSortOrderDesc,
}

// GetListTriggersSortOrderEnumValues Enumerates the set of values for ListTriggersSortOrderEnum
func GetListTriggersSortOrderEnumValues() []ListTriggersSortOrderEnum {
	values := make([]ListTriggersSortOrderEnum, 0)
	for _, v := range mappingListTriggersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTriggersSortOrderEnumStringValues Enumerates the set of values in String for ListTriggersSortOrderEnum
func GetListTriggersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTriggersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTriggersSortOrderEnum(val string) (ListTriggersSortOrderEnum, bool) {
	enum, ok := mappingListTriggersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
