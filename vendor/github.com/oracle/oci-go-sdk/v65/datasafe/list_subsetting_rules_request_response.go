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

// ListSubsettingRulesRequest wrapper for the ListSubsettingRules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingRules.go.html to see an example of how to use ListSubsettingRulesRequest.
type ListSubsettingRulesRequest struct {

	// The OCID of the subsetting policy.
	SubsettingPolicyId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for other fields is ascending.
	SortBy ListSubsettingRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingRulesResponse wrapper for the ListSubsettingRules operation
type ListSubsettingRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingRuleCollection instances
	SubsettingRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingRulesSortOrderEnum Enum with underlying type: string
type ListSubsettingRulesSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingRulesSortOrderEnum
const (
	ListSubsettingRulesSortOrderAsc  ListSubsettingRulesSortOrderEnum = "ASC"
	ListSubsettingRulesSortOrderDesc ListSubsettingRulesSortOrderEnum = "DESC"
)

var mappingListSubsettingRulesSortOrderEnum = map[string]ListSubsettingRulesSortOrderEnum{
	"ASC":  ListSubsettingRulesSortOrderAsc,
	"DESC": ListSubsettingRulesSortOrderDesc,
}

var mappingListSubsettingRulesSortOrderEnumLowerCase = map[string]ListSubsettingRulesSortOrderEnum{
	"asc":  ListSubsettingRulesSortOrderAsc,
	"desc": ListSubsettingRulesSortOrderDesc,
}

// GetListSubsettingRulesSortOrderEnumValues Enumerates the set of values for ListSubsettingRulesSortOrderEnum
func GetListSubsettingRulesSortOrderEnumValues() []ListSubsettingRulesSortOrderEnum {
	values := make([]ListSubsettingRulesSortOrderEnum, 0)
	for _, v := range mappingListSubsettingRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingRulesSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingRulesSortOrderEnum
func GetListSubsettingRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingRulesSortOrderEnum(val string) (ListSubsettingRulesSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingRulesSortByEnum Enum with underlying type: string
type ListSubsettingRulesSortByEnum string

// Set of constants representing the allowable values for ListSubsettingRulesSortByEnum
const (
	ListSubsettingRulesSortByTimecreated ListSubsettingRulesSortByEnum = "timeCreated"
	ListSubsettingRulesSortBySchemaname  ListSubsettingRulesSortByEnum = "schemaName"
	ListSubsettingRulesSortByObjectname  ListSubsettingRulesSortByEnum = "objectName"
)

var mappingListSubsettingRulesSortByEnum = map[string]ListSubsettingRulesSortByEnum{
	"timeCreated": ListSubsettingRulesSortByTimecreated,
	"schemaName":  ListSubsettingRulesSortBySchemaname,
	"objectName":  ListSubsettingRulesSortByObjectname,
}

var mappingListSubsettingRulesSortByEnumLowerCase = map[string]ListSubsettingRulesSortByEnum{
	"timecreated": ListSubsettingRulesSortByTimecreated,
	"schemaname":  ListSubsettingRulesSortBySchemaname,
	"objectname":  ListSubsettingRulesSortByObjectname,
}

// GetListSubsettingRulesSortByEnumValues Enumerates the set of values for ListSubsettingRulesSortByEnum
func GetListSubsettingRulesSortByEnumValues() []ListSubsettingRulesSortByEnum {
	values := make([]ListSubsettingRulesSortByEnum, 0)
	for _, v := range mappingListSubsettingRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingRulesSortByEnumStringValues Enumerates the set of values in String for ListSubsettingRulesSortByEnum
func GetListSubsettingRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"schemaName",
		"objectName",
	}
}

// GetMappingListSubsettingRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingRulesSortByEnum(val string) (ListSubsettingRulesSortByEnum, bool) {
	enum, ok := mappingListSubsettingRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
