// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package licensemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLicenseRecordsRequest wrapper for the ListLicenseRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListLicenseRecords.go.html to see an example of how to use ListLicenseRecordsRequest.
type ListLicenseRecordsRequest struct {

	// Unique product license identifier.
	ProductLicenseId *string `mandatory:"true" contributesTo:"query" name:"productLicenseId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, whether `ASC` or `DESC`.
	SortOrder ListLicenseRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `expirationDate`
	// * **expirationDate:** Sorts by expiration date of the license record.
	SortBy ListLicenseRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLicenseRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLicenseRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLicenseRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLicenseRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLicenseRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLicenseRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLicenseRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLicenseRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLicenseRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLicenseRecordsResponse wrapper for the ListLicenseRecords operation
type ListLicenseRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LicenseRecordCollection instances
	LicenseRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLicenseRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLicenseRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLicenseRecordsSortOrderEnum Enum with underlying type: string
type ListLicenseRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListLicenseRecordsSortOrderEnum
const (
	ListLicenseRecordsSortOrderAsc  ListLicenseRecordsSortOrderEnum = "ASC"
	ListLicenseRecordsSortOrderDesc ListLicenseRecordsSortOrderEnum = "DESC"
)

var mappingListLicenseRecordsSortOrderEnum = map[string]ListLicenseRecordsSortOrderEnum{
	"ASC":  ListLicenseRecordsSortOrderAsc,
	"DESC": ListLicenseRecordsSortOrderDesc,
}

var mappingListLicenseRecordsSortOrderEnumLowerCase = map[string]ListLicenseRecordsSortOrderEnum{
	"asc":  ListLicenseRecordsSortOrderAsc,
	"desc": ListLicenseRecordsSortOrderDesc,
}

// GetListLicenseRecordsSortOrderEnumValues Enumerates the set of values for ListLicenseRecordsSortOrderEnum
func GetListLicenseRecordsSortOrderEnumValues() []ListLicenseRecordsSortOrderEnum {
	values := make([]ListLicenseRecordsSortOrderEnum, 0)
	for _, v := range mappingListLicenseRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLicenseRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListLicenseRecordsSortOrderEnum
func GetListLicenseRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLicenseRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLicenseRecordsSortOrderEnum(val string) (ListLicenseRecordsSortOrderEnum, bool) {
	enum, ok := mappingListLicenseRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLicenseRecordsSortByEnum Enum with underlying type: string
type ListLicenseRecordsSortByEnum string

// Set of constants representing the allowable values for ListLicenseRecordsSortByEnum
const (
	ListLicenseRecordsSortByExpirationdate ListLicenseRecordsSortByEnum = "expirationDate"
)

var mappingListLicenseRecordsSortByEnum = map[string]ListLicenseRecordsSortByEnum{
	"expirationDate": ListLicenseRecordsSortByExpirationdate,
}

var mappingListLicenseRecordsSortByEnumLowerCase = map[string]ListLicenseRecordsSortByEnum{
	"expirationdate": ListLicenseRecordsSortByExpirationdate,
}

// GetListLicenseRecordsSortByEnumValues Enumerates the set of values for ListLicenseRecordsSortByEnum
func GetListLicenseRecordsSortByEnumValues() []ListLicenseRecordsSortByEnum {
	values := make([]ListLicenseRecordsSortByEnum, 0)
	for _, v := range mappingListLicenseRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLicenseRecordsSortByEnumStringValues Enumerates the set of values in String for ListLicenseRecordsSortByEnum
func GetListLicenseRecordsSortByEnumStringValues() []string {
	return []string{
		"expirationDate",
	}
}

// GetMappingListLicenseRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLicenseRecordsSortByEnum(val string) (ListLicenseRecordsSortByEnum, bool) {
	enum, ok := mappingListLicenseRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
