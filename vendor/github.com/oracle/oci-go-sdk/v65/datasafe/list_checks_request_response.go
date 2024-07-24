// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChecksRequest wrapper for the ListChecks operation
type ListChecksRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListChecksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(key eq 'USER.INACTIVE') and (title contains 'Owner Account')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Each check in security assessment has an associated key (think of key as a check's name).
	// For a given check, the key will be the same across targets. The user can use these keys to filter the checks.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChecksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChecksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChecksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChecksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChecksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListChecksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChecksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChecksResponse wrapper for the ListChecks operation
type ListChecksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CheckSummary instances
	Items []CheckSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListChecksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChecksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChecksSortOrderEnum Enum with underlying type: string
type ListChecksSortOrderEnum string

// Set of constants representing the allowable values for ListChecksSortOrderEnum
const (
	ListChecksSortOrderAsc  ListChecksSortOrderEnum = "ASC"
	ListChecksSortOrderDesc ListChecksSortOrderEnum = "DESC"
)

var mappingListChecksSortOrderEnum = map[string]ListChecksSortOrderEnum{
	"ASC":  ListChecksSortOrderAsc,
	"DESC": ListChecksSortOrderDesc,
}

var mappingListChecksSortOrderEnumLowerCase = map[string]ListChecksSortOrderEnum{
	"asc":  ListChecksSortOrderAsc,
	"desc": ListChecksSortOrderDesc,
}

// GetListChecksSortOrderEnumValues Enumerates the set of values for ListChecksSortOrderEnum
func GetListChecksSortOrderEnumValues() []ListChecksSortOrderEnum {
	values := make([]ListChecksSortOrderEnum, 0)
	for _, v := range mappingListChecksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksSortOrderEnumStringValues Enumerates the set of values in String for ListChecksSortOrderEnum
func GetListChecksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChecksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksSortOrderEnum(val string) (ListChecksSortOrderEnum, bool) {
	enum, ok := mappingListChecksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
