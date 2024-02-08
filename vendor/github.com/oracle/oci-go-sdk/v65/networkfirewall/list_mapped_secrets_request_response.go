// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMappedSecretsRequest wrapper for the ListMappedSecrets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListMappedSecrets.go.html to see an example of how to use ListMappedSecretsRequest.
type ListMappedSecretsRequest struct {

	// Unique Network Firewall Policy identifier
	NetworkFirewallPolicyId *string `mandatory:"true" contributesTo:"path" name:"networkFirewallPolicyId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMappedSecretsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMappedSecretsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMappedSecretsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMappedSecretsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMappedSecretsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMappedSecretsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMappedSecretsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMappedSecretsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMappedSecretsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMappedSecretsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMappedSecretsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMappedSecretsResponse wrapper for the ListMappedSecrets operation
type ListMappedSecretsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MappedSecretSummaryCollection instances
	MappedSecretSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. This is to get the page counts overall.
	OpcPageCount *string `presentIn:"header" name:"opc-page-count"`

	// For pagination of a list of items. This provides the count of total items across pages.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListMappedSecretsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMappedSecretsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMappedSecretsSortOrderEnum Enum with underlying type: string
type ListMappedSecretsSortOrderEnum string

// Set of constants representing the allowable values for ListMappedSecretsSortOrderEnum
const (
	ListMappedSecretsSortOrderAsc  ListMappedSecretsSortOrderEnum = "ASC"
	ListMappedSecretsSortOrderDesc ListMappedSecretsSortOrderEnum = "DESC"
)

var mappingListMappedSecretsSortOrderEnum = map[string]ListMappedSecretsSortOrderEnum{
	"ASC":  ListMappedSecretsSortOrderAsc,
	"DESC": ListMappedSecretsSortOrderDesc,
}

var mappingListMappedSecretsSortOrderEnumLowerCase = map[string]ListMappedSecretsSortOrderEnum{
	"asc":  ListMappedSecretsSortOrderAsc,
	"desc": ListMappedSecretsSortOrderDesc,
}

// GetListMappedSecretsSortOrderEnumValues Enumerates the set of values for ListMappedSecretsSortOrderEnum
func GetListMappedSecretsSortOrderEnumValues() []ListMappedSecretsSortOrderEnum {
	values := make([]ListMappedSecretsSortOrderEnum, 0)
	for _, v := range mappingListMappedSecretsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMappedSecretsSortOrderEnumStringValues Enumerates the set of values in String for ListMappedSecretsSortOrderEnum
func GetListMappedSecretsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMappedSecretsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMappedSecretsSortOrderEnum(val string) (ListMappedSecretsSortOrderEnum, bool) {
	enum, ok := mappingListMappedSecretsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMappedSecretsSortByEnum Enum with underlying type: string
type ListMappedSecretsSortByEnum string

// Set of constants representing the allowable values for ListMappedSecretsSortByEnum
const (
	ListMappedSecretsSortByTimecreated ListMappedSecretsSortByEnum = "timeCreated"
	ListMappedSecretsSortByDisplayname ListMappedSecretsSortByEnum = "displayName"
)

var mappingListMappedSecretsSortByEnum = map[string]ListMappedSecretsSortByEnum{
	"timeCreated": ListMappedSecretsSortByTimecreated,
	"displayName": ListMappedSecretsSortByDisplayname,
}

var mappingListMappedSecretsSortByEnumLowerCase = map[string]ListMappedSecretsSortByEnum{
	"timecreated": ListMappedSecretsSortByTimecreated,
	"displayname": ListMappedSecretsSortByDisplayname,
}

// GetListMappedSecretsSortByEnumValues Enumerates the set of values for ListMappedSecretsSortByEnum
func GetListMappedSecretsSortByEnumValues() []ListMappedSecretsSortByEnum {
	values := make([]ListMappedSecretsSortByEnum, 0)
	for _, v := range mappingListMappedSecretsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMappedSecretsSortByEnumStringValues Enumerates the set of values in String for ListMappedSecretsSortByEnum
func GetListMappedSecretsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMappedSecretsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMappedSecretsSortByEnum(val string) (ListMappedSecretsSortByEnum, bool) {
	enum, ok := mappingListMappedSecretsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
