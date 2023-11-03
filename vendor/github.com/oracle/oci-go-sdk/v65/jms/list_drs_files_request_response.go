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

// ListDrsFilesRequest wrapper for the ListDrsFiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListDrsFiles.go.html to see an example of how to use ListDrsFilesRequest.
type ListDrsFilesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListDrsFilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field that sorts the DRS details results. Only one sort order can be provided.
	// The default order for _drsFileKey_ is **descending**.
	// If no value is specified, then _drsFileKey_ is default.
	SortBy ListDrsFilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDrsFilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDrsFilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDrsFilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDrsFilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDrsFilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDrsFilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDrsFilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrsFilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDrsFilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDrsFilesResponse wrapper for the ListDrsFiles operation
type ListDrsFilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DrsFileCollection instances
	DrsFileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDrsFilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDrsFilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDrsFilesSortOrderEnum Enum with underlying type: string
type ListDrsFilesSortOrderEnum string

// Set of constants representing the allowable values for ListDrsFilesSortOrderEnum
const (
	ListDrsFilesSortOrderAsc  ListDrsFilesSortOrderEnum = "ASC"
	ListDrsFilesSortOrderDesc ListDrsFilesSortOrderEnum = "DESC"
)

var mappingListDrsFilesSortOrderEnum = map[string]ListDrsFilesSortOrderEnum{
	"ASC":  ListDrsFilesSortOrderAsc,
	"DESC": ListDrsFilesSortOrderDesc,
}

var mappingListDrsFilesSortOrderEnumLowerCase = map[string]ListDrsFilesSortOrderEnum{
	"asc":  ListDrsFilesSortOrderAsc,
	"desc": ListDrsFilesSortOrderDesc,
}

// GetListDrsFilesSortOrderEnumValues Enumerates the set of values for ListDrsFilesSortOrderEnum
func GetListDrsFilesSortOrderEnumValues() []ListDrsFilesSortOrderEnum {
	values := make([]ListDrsFilesSortOrderEnum, 0)
	for _, v := range mappingListDrsFilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrsFilesSortOrderEnumStringValues Enumerates the set of values in String for ListDrsFilesSortOrderEnum
func GetListDrsFilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDrsFilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrsFilesSortOrderEnum(val string) (ListDrsFilesSortOrderEnum, bool) {
	enum, ok := mappingListDrsFilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrsFilesSortByEnum Enum with underlying type: string
type ListDrsFilesSortByEnum string

// Set of constants representing the allowable values for ListDrsFilesSortByEnum
const (
	ListDrsFilesSortByBucketname   ListDrsFilesSortByEnum = "bucketName"
	ListDrsFilesSortByNamespace    ListDrsFilesSortByEnum = "namespace"
	ListDrsFilesSortByDrsfilekey   ListDrsFilesSortByEnum = "drsFileKey"
	ListDrsFilesSortByDrsfilename  ListDrsFilesSortByEnum = "drsFileName"
	ListDrsFilesSortByChecksumtype ListDrsFilesSortByEnum = "checksumType"
	ListDrsFilesSortByIsdefault    ListDrsFilesSortByEnum = "isDefault"
)

var mappingListDrsFilesSortByEnum = map[string]ListDrsFilesSortByEnum{
	"bucketName":   ListDrsFilesSortByBucketname,
	"namespace":    ListDrsFilesSortByNamespace,
	"drsFileKey":   ListDrsFilesSortByDrsfilekey,
	"drsFileName":  ListDrsFilesSortByDrsfilename,
	"checksumType": ListDrsFilesSortByChecksumtype,
	"isDefault":    ListDrsFilesSortByIsdefault,
}

var mappingListDrsFilesSortByEnumLowerCase = map[string]ListDrsFilesSortByEnum{
	"bucketname":   ListDrsFilesSortByBucketname,
	"namespace":    ListDrsFilesSortByNamespace,
	"drsfilekey":   ListDrsFilesSortByDrsfilekey,
	"drsfilename":  ListDrsFilesSortByDrsfilename,
	"checksumtype": ListDrsFilesSortByChecksumtype,
	"isdefault":    ListDrsFilesSortByIsdefault,
}

// GetListDrsFilesSortByEnumValues Enumerates the set of values for ListDrsFilesSortByEnum
func GetListDrsFilesSortByEnumValues() []ListDrsFilesSortByEnum {
	values := make([]ListDrsFilesSortByEnum, 0)
	for _, v := range mappingListDrsFilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrsFilesSortByEnumStringValues Enumerates the set of values in String for ListDrsFilesSortByEnum
func GetListDrsFilesSortByEnumStringValues() []string {
	return []string{
		"bucketName",
		"namespace",
		"drsFileKey",
		"drsFileName",
		"checksumType",
		"isDefault",
	}
}

// GetMappingListDrsFilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrsFilesSortByEnum(val string) (ListDrsFilesSortByEnum, bool) {
	enum, ok := mappingListDrsFilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
