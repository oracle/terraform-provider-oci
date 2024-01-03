// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCertificateAuthorityVersionsRequest wrapper for the ListCertificateAuthorityVersions operation
type ListCertificateAuthorityVersionsRequest struct {

	// The OCID of the certificate authority (CA).
	CertificateAuthorityId *string `mandatory:"true" contributesTo:"path" name:"certificateAuthorityId"`

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the specified version number. The default value is 0, which means that this filter is not applied.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can specify only one sort order. The default order for 'VERSION_NUMBER' is ascending.
	SortBy ListCertificateAuthorityVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateAuthorityVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateAuthorityVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateAuthorityVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateAuthorityVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateAuthorityVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificateAuthorityVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificateAuthorityVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificateAuthorityVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateAuthorityVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificateAuthorityVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificateAuthorityVersionsResponse wrapper for the ListCertificateAuthorityVersions operation
type ListCertificateAuthorityVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CertificateAuthorityVersionCollection instances
	CertificateAuthorityVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCertificateAuthorityVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateAuthorityVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateAuthorityVersionsSortByEnum Enum with underlying type: string
type ListCertificateAuthorityVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityVersionsSortByEnum
const (
	ListCertificateAuthorityVersionsSortByVersionNumber ListCertificateAuthorityVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateAuthorityVersionsSortByEnum = map[string]ListCertificateAuthorityVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateAuthorityVersionsSortByVersionNumber,
}

var mappingListCertificateAuthorityVersionsSortByEnumLowerCase = map[string]ListCertificateAuthorityVersionsSortByEnum{
	"version_number": ListCertificateAuthorityVersionsSortByVersionNumber,
}

// GetListCertificateAuthorityVersionsSortByEnumValues Enumerates the set of values for ListCertificateAuthorityVersionsSortByEnum
func GetListCertificateAuthorityVersionsSortByEnumValues() []ListCertificateAuthorityVersionsSortByEnum {
	values := make([]ListCertificateAuthorityVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateAuthorityVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthorityVersionsSortByEnumStringValues Enumerates the set of values in String for ListCertificateAuthorityVersionsSortByEnum
func GetListCertificateAuthorityVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION_NUMBER",
	}
}

// GetMappingListCertificateAuthorityVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthorityVersionsSortByEnum(val string) (ListCertificateAuthorityVersionsSortByEnum, bool) {
	enum, ok := mappingListCertificateAuthorityVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateAuthorityVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateAuthorityVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityVersionsSortOrderEnum
const (
	ListCertificateAuthorityVersionsSortOrderAsc  ListCertificateAuthorityVersionsSortOrderEnum = "ASC"
	ListCertificateAuthorityVersionsSortOrderDesc ListCertificateAuthorityVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateAuthorityVersionsSortOrderEnum = map[string]ListCertificateAuthorityVersionsSortOrderEnum{
	"ASC":  ListCertificateAuthorityVersionsSortOrderAsc,
	"DESC": ListCertificateAuthorityVersionsSortOrderDesc,
}

var mappingListCertificateAuthorityVersionsSortOrderEnumLowerCase = map[string]ListCertificateAuthorityVersionsSortOrderEnum{
	"asc":  ListCertificateAuthorityVersionsSortOrderAsc,
	"desc": ListCertificateAuthorityVersionsSortOrderDesc,
}

// GetListCertificateAuthorityVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateAuthorityVersionsSortOrderEnum
func GetListCertificateAuthorityVersionsSortOrderEnumValues() []ListCertificateAuthorityVersionsSortOrderEnum {
	values := make([]ListCertificateAuthorityVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateAuthorityVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthorityVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListCertificateAuthorityVersionsSortOrderEnum
func GetListCertificateAuthorityVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificateAuthorityVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthorityVersionsSortOrderEnum(val string) (ListCertificateAuthorityVersionsSortOrderEnum, bool) {
	enum, ok := mappingListCertificateAuthorityVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
