// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAgreementRecordsRequest wrapper for the ListAgreementRecords operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListAgreementRecords.go.html to see an example of how to use ListAgreementRecordsRequest.
type ListAgreementRecordsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListAgreementRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListAgreementRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgreementRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgreementRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgreementRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgreementRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgreementRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAgreementRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgreementRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgreementRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAgreementRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgreementRecordsResponse wrapper for the ListAgreementRecords operation
type ListAgreementRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgreementRecordCollection instances
	AgreementRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgreementRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgreementRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgreementRecordsSortOrderEnum Enum with underlying type: string
type ListAgreementRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListAgreementRecordsSortOrderEnum
const (
	ListAgreementRecordsSortOrderAsc  ListAgreementRecordsSortOrderEnum = "ASC"
	ListAgreementRecordsSortOrderDesc ListAgreementRecordsSortOrderEnum = "DESC"
)

var mappingListAgreementRecordsSortOrderEnum = map[string]ListAgreementRecordsSortOrderEnum{
	"ASC":  ListAgreementRecordsSortOrderAsc,
	"DESC": ListAgreementRecordsSortOrderDesc,
}

var mappingListAgreementRecordsSortOrderEnumLowerCase = map[string]ListAgreementRecordsSortOrderEnum{
	"asc":  ListAgreementRecordsSortOrderAsc,
	"desc": ListAgreementRecordsSortOrderDesc,
}

// GetListAgreementRecordsSortOrderEnumValues Enumerates the set of values for ListAgreementRecordsSortOrderEnum
func GetListAgreementRecordsSortOrderEnumValues() []ListAgreementRecordsSortOrderEnum {
	values := make([]ListAgreementRecordsSortOrderEnum, 0)
	for _, v := range mappingListAgreementRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgreementRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListAgreementRecordsSortOrderEnum
func GetListAgreementRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgreementRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgreementRecordsSortOrderEnum(val string) (ListAgreementRecordsSortOrderEnum, bool) {
	enum, ok := mappingListAgreementRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgreementRecordsSortByEnum Enum with underlying type: string
type ListAgreementRecordsSortByEnum string

// Set of constants representing the allowable values for ListAgreementRecordsSortByEnum
const (
	ListAgreementRecordsSortByTimecreated ListAgreementRecordsSortByEnum = "timeCreated"
	ListAgreementRecordsSortByDisplayname ListAgreementRecordsSortByEnum = "displayName"
)

var mappingListAgreementRecordsSortByEnum = map[string]ListAgreementRecordsSortByEnum{
	"timeCreated": ListAgreementRecordsSortByTimecreated,
	"displayName": ListAgreementRecordsSortByDisplayname,
}

var mappingListAgreementRecordsSortByEnumLowerCase = map[string]ListAgreementRecordsSortByEnum{
	"timecreated": ListAgreementRecordsSortByTimecreated,
	"displayname": ListAgreementRecordsSortByDisplayname,
}

// GetListAgreementRecordsSortByEnumValues Enumerates the set of values for ListAgreementRecordsSortByEnum
func GetListAgreementRecordsSortByEnumValues() []ListAgreementRecordsSortByEnum {
	values := make([]ListAgreementRecordsSortByEnum, 0)
	for _, v := range mappingListAgreementRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgreementRecordsSortByEnumStringValues Enumerates the set of values in String for ListAgreementRecordsSortByEnum
func GetListAgreementRecordsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAgreementRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgreementRecordsSortByEnum(val string) (ListAgreementRecordsSortByEnum, bool) {
	enum, ok := mappingListAgreementRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
