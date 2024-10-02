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

// ListZprPolicyWorkRequestErrorsRequest wrapper for the ListZprPolicyWorkRequestErrors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequestErrors.go.html to see an example of how to use ListZprPolicyWorkRequestErrorsRequest.
type ListZprPolicyWorkRequestErrorsRequest struct {

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
	SortBy ListZprPolicyWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListZprPolicyWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprPolicyWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprPolicyWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprPolicyWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListZprPolicyWorkRequestErrorsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
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
func (request ListZprPolicyWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprPolicyWorkRequestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprPolicyWorkRequestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprPolicyWorkRequestErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPolicyWorkRequestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprPolicyWorkRequestErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprPolicyWorkRequestErrorsResponse wrapper for the ListZprPolicyWorkRequestErrors operation
type ListZprPolicyWorkRequestErrorsResponse struct {

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

func (response ListZprPolicyWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprPolicyWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprPolicyWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListZprPolicyWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestErrorsSortByEnum
const (
	ListZprPolicyWorkRequestErrorsSortByTimestamp ListZprPolicyWorkRequestErrorsSortByEnum = "timestamp"
)

var mappingListZprPolicyWorkRequestErrorsSortByEnum = map[string]ListZprPolicyWorkRequestErrorsSortByEnum{
	"timestamp": ListZprPolicyWorkRequestErrorsSortByTimestamp,
}

var mappingListZprPolicyWorkRequestErrorsSortByEnumLowerCase = map[string]ListZprPolicyWorkRequestErrorsSortByEnum{
	"timestamp": ListZprPolicyWorkRequestErrorsSortByTimestamp,
}

// GetListZprPolicyWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListZprPolicyWorkRequestErrorsSortByEnum
func GetListZprPolicyWorkRequestErrorsSortByEnumValues() []ListZprPolicyWorkRequestErrorsSortByEnum {
	values := make([]ListZprPolicyWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestErrorsSortByEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestErrorsSortByEnum
func GetListZprPolicyWorkRequestErrorsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListZprPolicyWorkRequestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestErrorsSortByEnum(val string) (ListZprPolicyWorkRequestErrorsSortByEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprPolicyWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListZprPolicyWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestErrorsSortOrderEnum
const (
	ListZprPolicyWorkRequestErrorsSortOrderAsc  ListZprPolicyWorkRequestErrorsSortOrderEnum = "ASC"
	ListZprPolicyWorkRequestErrorsSortOrderDesc ListZprPolicyWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListZprPolicyWorkRequestErrorsSortOrderEnum = map[string]ListZprPolicyWorkRequestErrorsSortOrderEnum{
	"ASC":  ListZprPolicyWorkRequestErrorsSortOrderAsc,
	"DESC": ListZprPolicyWorkRequestErrorsSortOrderDesc,
}

var mappingListZprPolicyWorkRequestErrorsSortOrderEnumLowerCase = map[string]ListZprPolicyWorkRequestErrorsSortOrderEnum{
	"asc":  ListZprPolicyWorkRequestErrorsSortOrderAsc,
	"desc": ListZprPolicyWorkRequestErrorsSortOrderDesc,
}

// GetListZprPolicyWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListZprPolicyWorkRequestErrorsSortOrderEnum
func GetListZprPolicyWorkRequestErrorsSortOrderEnumValues() []ListZprPolicyWorkRequestErrorsSortOrderEnum {
	values := make([]ListZprPolicyWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestErrorsSortOrderEnum
func GetListZprPolicyWorkRequestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprPolicyWorkRequestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestErrorsSortOrderEnum(val string) (ListZprPolicyWorkRequestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
