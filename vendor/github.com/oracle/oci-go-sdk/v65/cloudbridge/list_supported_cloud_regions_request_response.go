// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSupportedCloudRegionsRequest wrapper for the ListSupportedCloudRegions operation
type ListSupportedCloudRegionsRequest struct {

	// The asset source type.
	AssetSourceType ListSupportedCloudRegionsAssetSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"assetSourceType" omitEmpty:"true"`

	// A filter to return only supported cloud regions which name contains given nameContains as sub-string.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The field to sort by. Only one sort order may be provided. By default, name is in ascending order.
	SortBy ListSupportedCloudRegionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSupportedCloudRegionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSupportedCloudRegionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSupportedCloudRegionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSupportedCloudRegionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSupportedCloudRegionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSupportedCloudRegionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSupportedCloudRegionsAssetSourceTypeEnum(string(request.AssetSourceType)); !ok && request.AssetSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetSourceType: %s. Supported values are: %s.", request.AssetSourceType, strings.Join(GetListSupportedCloudRegionsAssetSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSupportedCloudRegionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSupportedCloudRegionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSupportedCloudRegionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSupportedCloudRegionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSupportedCloudRegionsResponse wrapper for the ListSupportedCloudRegions operation
type ListSupportedCloudRegionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SupportedCloudRegionCollection instances
	SupportedCloudRegionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSupportedCloudRegionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSupportedCloudRegionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSupportedCloudRegionsAssetSourceTypeEnum Enum with underlying type: string
type ListSupportedCloudRegionsAssetSourceTypeEnum string

// Set of constants representing the allowable values for ListSupportedCloudRegionsAssetSourceTypeEnum
const (
	ListSupportedCloudRegionsAssetSourceTypeVmware   ListSupportedCloudRegionsAssetSourceTypeEnum = "VMWARE"
	ListSupportedCloudRegionsAssetSourceTypeAws      ListSupportedCloudRegionsAssetSourceTypeEnum = "AWS"
	ListSupportedCloudRegionsAssetSourceTypeOracleDb ListSupportedCloudRegionsAssetSourceTypeEnum = "ORACLE_DB"
)

var mappingListSupportedCloudRegionsAssetSourceTypeEnum = map[string]ListSupportedCloudRegionsAssetSourceTypeEnum{
	"VMWARE":    ListSupportedCloudRegionsAssetSourceTypeVmware,
	"AWS":       ListSupportedCloudRegionsAssetSourceTypeAws,
	"ORACLE_DB": ListSupportedCloudRegionsAssetSourceTypeOracleDb,
}

var mappingListSupportedCloudRegionsAssetSourceTypeEnumLowerCase = map[string]ListSupportedCloudRegionsAssetSourceTypeEnum{
	"vmware":    ListSupportedCloudRegionsAssetSourceTypeVmware,
	"aws":       ListSupportedCloudRegionsAssetSourceTypeAws,
	"oracle_db": ListSupportedCloudRegionsAssetSourceTypeOracleDb,
}

// GetListSupportedCloudRegionsAssetSourceTypeEnumValues Enumerates the set of values for ListSupportedCloudRegionsAssetSourceTypeEnum
func GetListSupportedCloudRegionsAssetSourceTypeEnumValues() []ListSupportedCloudRegionsAssetSourceTypeEnum {
	values := make([]ListSupportedCloudRegionsAssetSourceTypeEnum, 0)
	for _, v := range mappingListSupportedCloudRegionsAssetSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedCloudRegionsAssetSourceTypeEnumStringValues Enumerates the set of values in String for ListSupportedCloudRegionsAssetSourceTypeEnum
func GetListSupportedCloudRegionsAssetSourceTypeEnumStringValues() []string {
	return []string{
		"VMWARE",
		"AWS",
		"ORACLE_DB",
	}
}

// GetMappingListSupportedCloudRegionsAssetSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedCloudRegionsAssetSourceTypeEnum(val string) (ListSupportedCloudRegionsAssetSourceTypeEnum, bool) {
	enum, ok := mappingListSupportedCloudRegionsAssetSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSupportedCloudRegionsSortByEnum Enum with underlying type: string
type ListSupportedCloudRegionsSortByEnum string

// Set of constants representing the allowable values for ListSupportedCloudRegionsSortByEnum
const (
	ListSupportedCloudRegionsSortByName ListSupportedCloudRegionsSortByEnum = "name"
)

var mappingListSupportedCloudRegionsSortByEnum = map[string]ListSupportedCloudRegionsSortByEnum{
	"name": ListSupportedCloudRegionsSortByName,
}

var mappingListSupportedCloudRegionsSortByEnumLowerCase = map[string]ListSupportedCloudRegionsSortByEnum{
	"name": ListSupportedCloudRegionsSortByName,
}

// GetListSupportedCloudRegionsSortByEnumValues Enumerates the set of values for ListSupportedCloudRegionsSortByEnum
func GetListSupportedCloudRegionsSortByEnumValues() []ListSupportedCloudRegionsSortByEnum {
	values := make([]ListSupportedCloudRegionsSortByEnum, 0)
	for _, v := range mappingListSupportedCloudRegionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedCloudRegionsSortByEnumStringValues Enumerates the set of values in String for ListSupportedCloudRegionsSortByEnum
func GetListSupportedCloudRegionsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListSupportedCloudRegionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedCloudRegionsSortByEnum(val string) (ListSupportedCloudRegionsSortByEnum, bool) {
	enum, ok := mappingListSupportedCloudRegionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSupportedCloudRegionsSortOrderEnum Enum with underlying type: string
type ListSupportedCloudRegionsSortOrderEnum string

// Set of constants representing the allowable values for ListSupportedCloudRegionsSortOrderEnum
const (
	ListSupportedCloudRegionsSortOrderAsc  ListSupportedCloudRegionsSortOrderEnum = "ASC"
	ListSupportedCloudRegionsSortOrderDesc ListSupportedCloudRegionsSortOrderEnum = "DESC"
)

var mappingListSupportedCloudRegionsSortOrderEnum = map[string]ListSupportedCloudRegionsSortOrderEnum{
	"ASC":  ListSupportedCloudRegionsSortOrderAsc,
	"DESC": ListSupportedCloudRegionsSortOrderDesc,
}

var mappingListSupportedCloudRegionsSortOrderEnumLowerCase = map[string]ListSupportedCloudRegionsSortOrderEnum{
	"asc":  ListSupportedCloudRegionsSortOrderAsc,
	"desc": ListSupportedCloudRegionsSortOrderDesc,
}

// GetListSupportedCloudRegionsSortOrderEnumValues Enumerates the set of values for ListSupportedCloudRegionsSortOrderEnum
func GetListSupportedCloudRegionsSortOrderEnumValues() []ListSupportedCloudRegionsSortOrderEnum {
	values := make([]ListSupportedCloudRegionsSortOrderEnum, 0)
	for _, v := range mappingListSupportedCloudRegionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSupportedCloudRegionsSortOrderEnumStringValues Enumerates the set of values in String for ListSupportedCloudRegionsSortOrderEnum
func GetListSupportedCloudRegionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSupportedCloudRegionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSupportedCloudRegionsSortOrderEnum(val string) (ListSupportedCloudRegionsSortOrderEnum, bool) {
	enum, ok := mappingListSupportedCloudRegionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
