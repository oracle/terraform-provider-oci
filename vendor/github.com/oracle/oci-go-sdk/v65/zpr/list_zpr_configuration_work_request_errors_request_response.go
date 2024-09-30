// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprConfigurationWorkRequestErrorsRequest wrapper for the ListZprConfigurationWorkRequestErrors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprConfigurationWorkRequestErrors.go.html to see an example of how to use ListZprConfigurationWorkRequestErrorsRequest.
type ListZprConfigurationWorkRequestErrorsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for `timestamp` is descending.
	SortBy ListZprConfigurationWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListZprConfigurationWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprConfigurationWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprConfigurationWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprConfigurationWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListZprConfigurationWorkRequestErrorsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["workRequestId"] != nil {
		templateParam := mandatoryParamMap["workRequestId"]
		for _, template := range templateParam {
			replacementParam := *request.WorkRequestId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprConfigurationWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprConfigurationWorkRequestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprConfigurationWorkRequestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprConfigurationWorkRequestErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprConfigurationWorkRequestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprConfigurationWorkRequestErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprConfigurationWorkRequestErrorsResponse wrapper for the ListZprConfigurationWorkRequestErrors operation
type ListZprConfigurationWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestErrorCollection instances
	WorkRequestErrorCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListZprConfigurationWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprConfigurationWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprConfigurationWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListZprConfigurationWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListZprConfigurationWorkRequestErrorsSortByEnum
const (
	ListZprConfigurationWorkRequestErrorsSortByTimestamp ListZprConfigurationWorkRequestErrorsSortByEnum = "timestamp"
)

var mappingListZprConfigurationWorkRequestErrorsSortByEnum = map[string]ListZprConfigurationWorkRequestErrorsSortByEnum{
	"timestamp": ListZprConfigurationWorkRequestErrorsSortByTimestamp,
}

var mappingListZprConfigurationWorkRequestErrorsSortByEnumLowerCase = map[string]ListZprConfigurationWorkRequestErrorsSortByEnum{
	"timestamp": ListZprConfigurationWorkRequestErrorsSortByTimestamp,
}

// GetListZprConfigurationWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListZprConfigurationWorkRequestErrorsSortByEnum
func GetListZprConfigurationWorkRequestErrorsSortByEnumValues() []ListZprConfigurationWorkRequestErrorsSortByEnum {
	values := make([]ListZprConfigurationWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListZprConfigurationWorkRequestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprConfigurationWorkRequestErrorsSortByEnumStringValues Enumerates the set of values in String for ListZprConfigurationWorkRequestErrorsSortByEnum
func GetListZprConfigurationWorkRequestErrorsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListZprConfigurationWorkRequestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprConfigurationWorkRequestErrorsSortByEnum(val string) (ListZprConfigurationWorkRequestErrorsSortByEnum, bool) {
	enum, ok := mappingListZprConfigurationWorkRequestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprConfigurationWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListZprConfigurationWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListZprConfigurationWorkRequestErrorsSortOrderEnum
const (
	ListZprConfigurationWorkRequestErrorsSortOrderAsc  ListZprConfigurationWorkRequestErrorsSortOrderEnum = "ASC"
	ListZprConfigurationWorkRequestErrorsSortOrderDesc ListZprConfigurationWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListZprConfigurationWorkRequestErrorsSortOrderEnum = map[string]ListZprConfigurationWorkRequestErrorsSortOrderEnum{
	"ASC":  ListZprConfigurationWorkRequestErrorsSortOrderAsc,
	"DESC": ListZprConfigurationWorkRequestErrorsSortOrderDesc,
}

var mappingListZprConfigurationWorkRequestErrorsSortOrderEnumLowerCase = map[string]ListZprConfigurationWorkRequestErrorsSortOrderEnum{
	"asc":  ListZprConfigurationWorkRequestErrorsSortOrderAsc,
	"desc": ListZprConfigurationWorkRequestErrorsSortOrderDesc,
}

// GetListZprConfigurationWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListZprConfigurationWorkRequestErrorsSortOrderEnum
func GetListZprConfigurationWorkRequestErrorsSortOrderEnumValues() []ListZprConfigurationWorkRequestErrorsSortOrderEnum {
	values := make([]ListZprConfigurationWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListZprConfigurationWorkRequestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprConfigurationWorkRequestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListZprConfigurationWorkRequestErrorsSortOrderEnum
func GetListZprConfigurationWorkRequestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprConfigurationWorkRequestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprConfigurationWorkRequestErrorsSortOrderEnum(val string) (ListZprConfigurationWorkRequestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListZprConfigurationWorkRequestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
