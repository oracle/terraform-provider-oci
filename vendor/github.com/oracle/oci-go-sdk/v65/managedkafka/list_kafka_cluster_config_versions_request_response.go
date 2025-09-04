// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListKafkaClusterConfigVersionsRequest wrapper for the ListKafkaClusterConfigVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managedkafka/ListKafkaClusterConfigVersions.go.html to see an example of how to use ListKafkaClusterConfigVersionsRequest.
type ListKafkaClusterConfigVersionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
	KafkaClusterConfigId *string `mandatory:"true" contributesTo:"path" name:"kafkaClusterConfigId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListKafkaClusterConfigVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `versionNumber`
	// is descending.
	SortBy ListKafkaClusterConfigVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKafkaClusterConfigVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKafkaClusterConfigVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKafkaClusterConfigVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKafkaClusterConfigVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKafkaClusterConfigVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListKafkaClusterConfigVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKafkaClusterConfigVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKafkaClusterConfigVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKafkaClusterConfigVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKafkaClusterConfigVersionsResponse wrapper for the ListKafkaClusterConfigVersions operation
type ListKafkaClusterConfigVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KafkaClusterConfigVersionCollection instances
	KafkaClusterConfigVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListKafkaClusterConfigVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKafkaClusterConfigVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKafkaClusterConfigVersionsSortOrderEnum Enum with underlying type: string
type ListKafkaClusterConfigVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListKafkaClusterConfigVersionsSortOrderEnum
const (
	ListKafkaClusterConfigVersionsSortOrderAsc  ListKafkaClusterConfigVersionsSortOrderEnum = "ASC"
	ListKafkaClusterConfigVersionsSortOrderDesc ListKafkaClusterConfigVersionsSortOrderEnum = "DESC"
)

var mappingListKafkaClusterConfigVersionsSortOrderEnum = map[string]ListKafkaClusterConfigVersionsSortOrderEnum{
	"ASC":  ListKafkaClusterConfigVersionsSortOrderAsc,
	"DESC": ListKafkaClusterConfigVersionsSortOrderDesc,
}

var mappingListKafkaClusterConfigVersionsSortOrderEnumLowerCase = map[string]ListKafkaClusterConfigVersionsSortOrderEnum{
	"asc":  ListKafkaClusterConfigVersionsSortOrderAsc,
	"desc": ListKafkaClusterConfigVersionsSortOrderDesc,
}

// GetListKafkaClusterConfigVersionsSortOrderEnumValues Enumerates the set of values for ListKafkaClusterConfigVersionsSortOrderEnum
func GetListKafkaClusterConfigVersionsSortOrderEnumValues() []ListKafkaClusterConfigVersionsSortOrderEnum {
	values := make([]ListKafkaClusterConfigVersionsSortOrderEnum, 0)
	for _, v := range mappingListKafkaClusterConfigVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClusterConfigVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListKafkaClusterConfigVersionsSortOrderEnum
func GetListKafkaClusterConfigVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKafkaClusterConfigVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClusterConfigVersionsSortOrderEnum(val string) (ListKafkaClusterConfigVersionsSortOrderEnum, bool) {
	enum, ok := mappingListKafkaClusterConfigVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKafkaClusterConfigVersionsSortByEnum Enum with underlying type: string
type ListKafkaClusterConfigVersionsSortByEnum string

// Set of constants representing the allowable values for ListKafkaClusterConfigVersionsSortByEnum
const (
	ListKafkaClusterConfigVersionsSortByVersionnumber ListKafkaClusterConfigVersionsSortByEnum = "versionNumber"
)

var mappingListKafkaClusterConfigVersionsSortByEnum = map[string]ListKafkaClusterConfigVersionsSortByEnum{
	"versionNumber": ListKafkaClusterConfigVersionsSortByVersionnumber,
}

var mappingListKafkaClusterConfigVersionsSortByEnumLowerCase = map[string]ListKafkaClusterConfigVersionsSortByEnum{
	"versionnumber": ListKafkaClusterConfigVersionsSortByVersionnumber,
}

// GetListKafkaClusterConfigVersionsSortByEnumValues Enumerates the set of values for ListKafkaClusterConfigVersionsSortByEnum
func GetListKafkaClusterConfigVersionsSortByEnumValues() []ListKafkaClusterConfigVersionsSortByEnum {
	values := make([]ListKafkaClusterConfigVersionsSortByEnum, 0)
	for _, v := range mappingListKafkaClusterConfigVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClusterConfigVersionsSortByEnumStringValues Enumerates the set of values in String for ListKafkaClusterConfigVersionsSortByEnum
func GetListKafkaClusterConfigVersionsSortByEnumStringValues() []string {
	return []string{
		"versionNumber",
	}
}

// GetMappingListKafkaClusterConfigVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClusterConfigVersionsSortByEnum(val string) (ListKafkaClusterConfigVersionsSortByEnum, bool) {
	enum, ok := mappingListKafkaClusterConfigVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
