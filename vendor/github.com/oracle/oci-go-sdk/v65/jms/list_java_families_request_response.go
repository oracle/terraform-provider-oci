// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaFamiliesRequest wrapper for the ListJavaFamilies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaFamilies.go.html to see an example of how to use ListJavaFamiliesRequest.
type ListJavaFamiliesRequest struct {

	// The version identifier for the Java family.
	FamilyVersion *string `mandatory:"false" contributesTo:"query" name:"familyVersion"`

	// The display name for the Java family.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter the Java Release Family versions by support status.
	IsSupportedVersion *bool `mandatory:"false" contributesTo:"query" name:"isSupportedVersion"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaFamiliesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If no value is specified _familyVersion_ is default.
	SortBy ListJavaFamiliesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaFamiliesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaFamiliesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaFamiliesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaFamiliesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaFamiliesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaFamiliesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaFamiliesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaFamiliesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaFamiliesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaFamiliesResponse wrapper for the ListJavaFamilies operation
type ListJavaFamiliesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaFamilyCollection instances
	JavaFamilyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaFamiliesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaFamiliesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaFamiliesSortOrderEnum Enum with underlying type: string
type ListJavaFamiliesSortOrderEnum string

// Set of constants representing the allowable values for ListJavaFamiliesSortOrderEnum
const (
	ListJavaFamiliesSortOrderAsc  ListJavaFamiliesSortOrderEnum = "ASC"
	ListJavaFamiliesSortOrderDesc ListJavaFamiliesSortOrderEnum = "DESC"
)

var mappingListJavaFamiliesSortOrderEnum = map[string]ListJavaFamiliesSortOrderEnum{
	"ASC":  ListJavaFamiliesSortOrderAsc,
	"DESC": ListJavaFamiliesSortOrderDesc,
}

var mappingListJavaFamiliesSortOrderEnumLowerCase = map[string]ListJavaFamiliesSortOrderEnum{
	"asc":  ListJavaFamiliesSortOrderAsc,
	"desc": ListJavaFamiliesSortOrderDesc,
}

// GetListJavaFamiliesSortOrderEnumValues Enumerates the set of values for ListJavaFamiliesSortOrderEnum
func GetListJavaFamiliesSortOrderEnumValues() []ListJavaFamiliesSortOrderEnum {
	values := make([]ListJavaFamiliesSortOrderEnum, 0)
	for _, v := range mappingListJavaFamiliesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaFamiliesSortOrderEnumStringValues Enumerates the set of values in String for ListJavaFamiliesSortOrderEnum
func GetListJavaFamiliesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaFamiliesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaFamiliesSortOrderEnum(val string) (ListJavaFamiliesSortOrderEnum, bool) {
	enum, ok := mappingListJavaFamiliesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaFamiliesSortByEnum Enum with underlying type: string
type ListJavaFamiliesSortByEnum string

// Set of constants representing the allowable values for ListJavaFamiliesSortByEnum
const (
	ListJavaFamiliesSortByFamilyversion        ListJavaFamiliesSortByEnum = "familyVersion"
	ListJavaFamiliesSortByEndofsupportlifedate ListJavaFamiliesSortByEnum = "endOfSupportLifeDate"
	ListJavaFamiliesSortBySupporttype          ListJavaFamiliesSortByEnum = "supportType"
)

var mappingListJavaFamiliesSortByEnum = map[string]ListJavaFamiliesSortByEnum{
	"familyVersion":        ListJavaFamiliesSortByFamilyversion,
	"endOfSupportLifeDate": ListJavaFamiliesSortByEndofsupportlifedate,
	"supportType":          ListJavaFamiliesSortBySupporttype,
}

var mappingListJavaFamiliesSortByEnumLowerCase = map[string]ListJavaFamiliesSortByEnum{
	"familyversion":        ListJavaFamiliesSortByFamilyversion,
	"endofsupportlifedate": ListJavaFamiliesSortByEndofsupportlifedate,
	"supporttype":          ListJavaFamiliesSortBySupporttype,
}

// GetListJavaFamiliesSortByEnumValues Enumerates the set of values for ListJavaFamiliesSortByEnum
func GetListJavaFamiliesSortByEnumValues() []ListJavaFamiliesSortByEnum {
	values := make([]ListJavaFamiliesSortByEnum, 0)
	for _, v := range mappingListJavaFamiliesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaFamiliesSortByEnumStringValues Enumerates the set of values in String for ListJavaFamiliesSortByEnum
func GetListJavaFamiliesSortByEnumStringValues() []string {
	return []string{
		"familyVersion",
		"endOfSupportLifeDate",
		"supportType",
	}
}

// GetMappingListJavaFamiliesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaFamiliesSortByEnum(val string) (ListJavaFamiliesSortByEnum, bool) {
	enum, ok := mappingListJavaFamiliesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
