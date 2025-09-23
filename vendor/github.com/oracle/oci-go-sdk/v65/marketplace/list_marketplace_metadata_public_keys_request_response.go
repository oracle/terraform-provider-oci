// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMarketplaceMetadataPublicKeysRequest wrapper for the ListMarketplaceMetadataPublicKeys operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListMarketplaceMetadataPublicKeys.go.html to see an example of how to use ListMarketplaceMetadataPublicKeysRequest.
type ListMarketplaceMetadataPublicKeysRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListMarketplaceMetadataPublicKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for keyId is descending.
	SortBy ListMarketplaceMetadataPublicKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMarketplaceMetadataPublicKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMarketplaceMetadataPublicKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMarketplaceMetadataPublicKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMarketplaceMetadataPublicKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMarketplaceMetadataPublicKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMarketplaceMetadataPublicKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMarketplaceMetadataPublicKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMarketplaceMetadataPublicKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMarketplaceMetadataPublicKeysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMarketplaceMetadataPublicKeysResponse wrapper for the ListMarketplaceMetadataPublicKeys operation
type ListMarketplaceMetadataPublicKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MarketplaceMetadataPublicKeySummary instances
	Items []MarketplaceMetadataPublicKeySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMarketplaceMetadataPublicKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMarketplaceMetadataPublicKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMarketplaceMetadataPublicKeysSortOrderEnum Enum with underlying type: string
type ListMarketplaceMetadataPublicKeysSortOrderEnum string

// Set of constants representing the allowable values for ListMarketplaceMetadataPublicKeysSortOrderEnum
const (
	ListMarketplaceMetadataPublicKeysSortOrderAsc  ListMarketplaceMetadataPublicKeysSortOrderEnum = "ASC"
	ListMarketplaceMetadataPublicKeysSortOrderDesc ListMarketplaceMetadataPublicKeysSortOrderEnum = "DESC"
)

var mappingListMarketplaceMetadataPublicKeysSortOrderEnum = map[string]ListMarketplaceMetadataPublicKeysSortOrderEnum{
	"ASC":  ListMarketplaceMetadataPublicKeysSortOrderAsc,
	"DESC": ListMarketplaceMetadataPublicKeysSortOrderDesc,
}

var mappingListMarketplaceMetadataPublicKeysSortOrderEnumLowerCase = map[string]ListMarketplaceMetadataPublicKeysSortOrderEnum{
	"asc":  ListMarketplaceMetadataPublicKeysSortOrderAsc,
	"desc": ListMarketplaceMetadataPublicKeysSortOrderDesc,
}

// GetListMarketplaceMetadataPublicKeysSortOrderEnumValues Enumerates the set of values for ListMarketplaceMetadataPublicKeysSortOrderEnum
func GetListMarketplaceMetadataPublicKeysSortOrderEnumValues() []ListMarketplaceMetadataPublicKeysSortOrderEnum {
	values := make([]ListMarketplaceMetadataPublicKeysSortOrderEnum, 0)
	for _, v := range mappingListMarketplaceMetadataPublicKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketplaceMetadataPublicKeysSortOrderEnumStringValues Enumerates the set of values in String for ListMarketplaceMetadataPublicKeysSortOrderEnum
func GetListMarketplaceMetadataPublicKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMarketplaceMetadataPublicKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketplaceMetadataPublicKeysSortOrderEnum(val string) (ListMarketplaceMetadataPublicKeysSortOrderEnum, bool) {
	enum, ok := mappingListMarketplaceMetadataPublicKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMarketplaceMetadataPublicKeysSortByEnum Enum with underlying type: string
type ListMarketplaceMetadataPublicKeysSortByEnum string

// Set of constants representing the allowable values for ListMarketplaceMetadataPublicKeysSortByEnum
const (
	ListMarketplaceMetadataPublicKeysSortByKeyid ListMarketplaceMetadataPublicKeysSortByEnum = "keyId"
)

var mappingListMarketplaceMetadataPublicKeysSortByEnum = map[string]ListMarketplaceMetadataPublicKeysSortByEnum{
	"keyId": ListMarketplaceMetadataPublicKeysSortByKeyid,
}

var mappingListMarketplaceMetadataPublicKeysSortByEnumLowerCase = map[string]ListMarketplaceMetadataPublicKeysSortByEnum{
	"keyid": ListMarketplaceMetadataPublicKeysSortByKeyid,
}

// GetListMarketplaceMetadataPublicKeysSortByEnumValues Enumerates the set of values for ListMarketplaceMetadataPublicKeysSortByEnum
func GetListMarketplaceMetadataPublicKeysSortByEnumValues() []ListMarketplaceMetadataPublicKeysSortByEnum {
	values := make([]ListMarketplaceMetadataPublicKeysSortByEnum, 0)
	for _, v := range mappingListMarketplaceMetadataPublicKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketplaceMetadataPublicKeysSortByEnumStringValues Enumerates the set of values in String for ListMarketplaceMetadataPublicKeysSortByEnum
func GetListMarketplaceMetadataPublicKeysSortByEnumStringValues() []string {
	return []string{
		"keyId",
	}
}

// GetMappingListMarketplaceMetadataPublicKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketplaceMetadataPublicKeysSortByEnum(val string) (ListMarketplaceMetadataPublicKeysSortByEnum, bool) {
	enum, ok := mappingListMarketplaceMetadataPublicKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
