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

// ListDecryptionProfilesRequest wrapper for the ListDecryptionProfiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListDecryptionProfiles.go.html to see an example of how to use ListDecryptionProfilesRequest.
type ListDecryptionProfilesRequest struct {

	// Unique Network Firewall Policy identifier
	NetworkFirewallPolicyId *string `mandatory:"true" contributesTo:"path" name:"networkFirewallPolicyId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDecryptionProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDecryptionProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDecryptionProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDecryptionProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDecryptionProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDecryptionProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDecryptionProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDecryptionProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDecryptionProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDecryptionProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDecryptionProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDecryptionProfilesResponse wrapper for the ListDecryptionProfiles operation
type ListDecryptionProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DecryptionProfileSummaryCollection instances
	DecryptionProfileSummaryCollection `presentIn:"body"`

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

func (response ListDecryptionProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDecryptionProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDecryptionProfilesSortOrderEnum Enum with underlying type: string
type ListDecryptionProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListDecryptionProfilesSortOrderEnum
const (
	ListDecryptionProfilesSortOrderAsc  ListDecryptionProfilesSortOrderEnum = "ASC"
	ListDecryptionProfilesSortOrderDesc ListDecryptionProfilesSortOrderEnum = "DESC"
)

var mappingListDecryptionProfilesSortOrderEnum = map[string]ListDecryptionProfilesSortOrderEnum{
	"ASC":  ListDecryptionProfilesSortOrderAsc,
	"DESC": ListDecryptionProfilesSortOrderDesc,
}

var mappingListDecryptionProfilesSortOrderEnumLowerCase = map[string]ListDecryptionProfilesSortOrderEnum{
	"asc":  ListDecryptionProfilesSortOrderAsc,
	"desc": ListDecryptionProfilesSortOrderDesc,
}

// GetListDecryptionProfilesSortOrderEnumValues Enumerates the set of values for ListDecryptionProfilesSortOrderEnum
func GetListDecryptionProfilesSortOrderEnumValues() []ListDecryptionProfilesSortOrderEnum {
	values := make([]ListDecryptionProfilesSortOrderEnum, 0)
	for _, v := range mappingListDecryptionProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDecryptionProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListDecryptionProfilesSortOrderEnum
func GetListDecryptionProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDecryptionProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDecryptionProfilesSortOrderEnum(val string) (ListDecryptionProfilesSortOrderEnum, bool) {
	enum, ok := mappingListDecryptionProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDecryptionProfilesSortByEnum Enum with underlying type: string
type ListDecryptionProfilesSortByEnum string

// Set of constants representing the allowable values for ListDecryptionProfilesSortByEnum
const (
	ListDecryptionProfilesSortByTimecreated ListDecryptionProfilesSortByEnum = "timeCreated"
	ListDecryptionProfilesSortByDisplayname ListDecryptionProfilesSortByEnum = "displayName"
)

var mappingListDecryptionProfilesSortByEnum = map[string]ListDecryptionProfilesSortByEnum{
	"timeCreated": ListDecryptionProfilesSortByTimecreated,
	"displayName": ListDecryptionProfilesSortByDisplayname,
}

var mappingListDecryptionProfilesSortByEnumLowerCase = map[string]ListDecryptionProfilesSortByEnum{
	"timecreated": ListDecryptionProfilesSortByTimecreated,
	"displayname": ListDecryptionProfilesSortByDisplayname,
}

// GetListDecryptionProfilesSortByEnumValues Enumerates the set of values for ListDecryptionProfilesSortByEnum
func GetListDecryptionProfilesSortByEnumValues() []ListDecryptionProfilesSortByEnum {
	values := make([]ListDecryptionProfilesSortByEnum, 0)
	for _, v := range mappingListDecryptionProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDecryptionProfilesSortByEnumStringValues Enumerates the set of values in String for ListDecryptionProfilesSortByEnum
func GetListDecryptionProfilesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDecryptionProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDecryptionProfilesSortByEnum(val string) (ListDecryptionProfilesSortByEnum, bool) {
	enum, ok := mappingListDecryptionProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
