// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListImportableComputeEntitiesRequest wrapper for the ListImportableComputeEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableComputeEntities.go.html to see an example of how to use ListImportableComputeEntitiesRequest.
type ListImportableComputeEntitiesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListImportableComputeEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Compute entity list sort options.
	SortBy ListImportableComputeEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportableComputeEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportableComputeEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportableComputeEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportableComputeEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportableComputeEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListImportableComputeEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListImportableComputeEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportableComputeEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListImportableComputeEntitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportableComputeEntitiesResponse wrapper for the ListImportableComputeEntities operation
type ListImportableComputeEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImportableComputeEntitySummaryCollection instances
	ImportableComputeEntitySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportableComputeEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportableComputeEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImportableComputeEntitiesSortOrderEnum Enum with underlying type: string
type ListImportableComputeEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListImportableComputeEntitiesSortOrderEnum
const (
	ListImportableComputeEntitiesSortOrderAsc  ListImportableComputeEntitiesSortOrderEnum = "ASC"
	ListImportableComputeEntitiesSortOrderDesc ListImportableComputeEntitiesSortOrderEnum = "DESC"
)

var mappingListImportableComputeEntitiesSortOrderEnum = map[string]ListImportableComputeEntitiesSortOrderEnum{
	"ASC":  ListImportableComputeEntitiesSortOrderAsc,
	"DESC": ListImportableComputeEntitiesSortOrderDesc,
}

var mappingListImportableComputeEntitiesSortOrderEnumLowerCase = map[string]ListImportableComputeEntitiesSortOrderEnum{
	"asc":  ListImportableComputeEntitiesSortOrderAsc,
	"desc": ListImportableComputeEntitiesSortOrderDesc,
}

// GetListImportableComputeEntitiesSortOrderEnumValues Enumerates the set of values for ListImportableComputeEntitiesSortOrderEnum
func GetListImportableComputeEntitiesSortOrderEnumValues() []ListImportableComputeEntitiesSortOrderEnum {
	values := make([]ListImportableComputeEntitiesSortOrderEnum, 0)
	for _, v := range mappingListImportableComputeEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportableComputeEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListImportableComputeEntitiesSortOrderEnum
func GetListImportableComputeEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListImportableComputeEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportableComputeEntitiesSortOrderEnum(val string) (ListImportableComputeEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListImportableComputeEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportableComputeEntitiesSortByEnum Enum with underlying type: string
type ListImportableComputeEntitiesSortByEnum string

// Set of constants representing the allowable values for ListImportableComputeEntitiesSortByEnum
const (
	ListImportableComputeEntitiesSortByComputeid          ListImportableComputeEntitiesSortByEnum = "computeId"
	ListImportableComputeEntitiesSortByComputedisplayname ListImportableComputeEntitiesSortByEnum = "computeDisplayName"
	ListImportableComputeEntitiesSortByPlatformtype       ListImportableComputeEntitiesSortByEnum = "platformType"
	ListImportableComputeEntitiesSortByHostname           ListImportableComputeEntitiesSortByEnum = "hostName"
)

var mappingListImportableComputeEntitiesSortByEnum = map[string]ListImportableComputeEntitiesSortByEnum{
	"computeId":          ListImportableComputeEntitiesSortByComputeid,
	"computeDisplayName": ListImportableComputeEntitiesSortByComputedisplayname,
	"platformType":       ListImportableComputeEntitiesSortByPlatformtype,
	"hostName":           ListImportableComputeEntitiesSortByHostname,
}

var mappingListImportableComputeEntitiesSortByEnumLowerCase = map[string]ListImportableComputeEntitiesSortByEnum{
	"computeid":          ListImportableComputeEntitiesSortByComputeid,
	"computedisplayname": ListImportableComputeEntitiesSortByComputedisplayname,
	"platformtype":       ListImportableComputeEntitiesSortByPlatformtype,
	"hostname":           ListImportableComputeEntitiesSortByHostname,
}

// GetListImportableComputeEntitiesSortByEnumValues Enumerates the set of values for ListImportableComputeEntitiesSortByEnum
func GetListImportableComputeEntitiesSortByEnumValues() []ListImportableComputeEntitiesSortByEnum {
	values := make([]ListImportableComputeEntitiesSortByEnum, 0)
	for _, v := range mappingListImportableComputeEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportableComputeEntitiesSortByEnumStringValues Enumerates the set of values in String for ListImportableComputeEntitiesSortByEnum
func GetListImportableComputeEntitiesSortByEnumStringValues() []string {
	return []string{
		"computeId",
		"computeDisplayName",
		"platformType",
		"hostName",
	}
}

// GetMappingListImportableComputeEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportableComputeEntitiesSortByEnum(val string) (ListImportableComputeEntitiesSortByEnum, bool) {
	enum, ok := mappingListImportableComputeEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
