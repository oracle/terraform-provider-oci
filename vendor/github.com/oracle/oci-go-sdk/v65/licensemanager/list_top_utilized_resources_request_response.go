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

// ListTopUtilizedResourcesRequest wrapper for the ListTopUtilizedResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListTopUtilizedResources.go.html to see an example of how to use ListTopUtilizedResourcesRequest.
type ListTopUtilizedResourcesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates if the given compartment is the root compartment.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// A filter to return only resources whose unit matches the given resource unit.
	ResourceUnitType ListTopUtilizedResourcesResourceUnitTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceUnitType" omitEmpty:"true"`

	// The sort order to use, whether `ASC` or `DESC`.
	SortOrder ListTopUtilizedResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `totalUnits`
	// * **totalUnits:** Sorts by totalUnits consumed by resource.
	SortBy ListTopUtilizedResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTopUtilizedResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTopUtilizedResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTopUtilizedResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTopUtilizedResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTopUtilizedResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTopUtilizedResourcesResourceUnitTypeEnum(string(request.ResourceUnitType)); !ok && request.ResourceUnitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceUnitType: %s. Supported values are: %s.", request.ResourceUnitType, strings.Join(GetListTopUtilizedResourcesResourceUnitTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTopUtilizedResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTopUtilizedResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTopUtilizedResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTopUtilizedResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTopUtilizedResourcesResponse wrapper for the ListTopUtilizedResources operation
type ListTopUtilizedResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TopUtilizedResourceCollection instances
	TopUtilizedResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTopUtilizedResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTopUtilizedResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTopUtilizedResourcesResourceUnitTypeEnum Enum with underlying type: string
type ListTopUtilizedResourcesResourceUnitTypeEnum string

// Set of constants representing the allowable values for ListTopUtilizedResourcesResourceUnitTypeEnum
const (
	ListTopUtilizedResourcesResourceUnitTypeOcpu ListTopUtilizedResourcesResourceUnitTypeEnum = "OCPU"
	ListTopUtilizedResourcesResourceUnitTypeEcpu ListTopUtilizedResourcesResourceUnitTypeEnum = "ECPU"
)

var mappingListTopUtilizedResourcesResourceUnitTypeEnum = map[string]ListTopUtilizedResourcesResourceUnitTypeEnum{
	"OCPU": ListTopUtilizedResourcesResourceUnitTypeOcpu,
	"ECPU": ListTopUtilizedResourcesResourceUnitTypeEcpu,
}

var mappingListTopUtilizedResourcesResourceUnitTypeEnumLowerCase = map[string]ListTopUtilizedResourcesResourceUnitTypeEnum{
	"ocpu": ListTopUtilizedResourcesResourceUnitTypeOcpu,
	"ecpu": ListTopUtilizedResourcesResourceUnitTypeEcpu,
}

// GetListTopUtilizedResourcesResourceUnitTypeEnumValues Enumerates the set of values for ListTopUtilizedResourcesResourceUnitTypeEnum
func GetListTopUtilizedResourcesResourceUnitTypeEnumValues() []ListTopUtilizedResourcesResourceUnitTypeEnum {
	values := make([]ListTopUtilizedResourcesResourceUnitTypeEnum, 0)
	for _, v := range mappingListTopUtilizedResourcesResourceUnitTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopUtilizedResourcesResourceUnitTypeEnumStringValues Enumerates the set of values in String for ListTopUtilizedResourcesResourceUnitTypeEnum
func GetListTopUtilizedResourcesResourceUnitTypeEnumStringValues() []string {
	return []string{
		"OCPU",
		"ECPU",
	}
}

// GetMappingListTopUtilizedResourcesResourceUnitTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopUtilizedResourcesResourceUnitTypeEnum(val string) (ListTopUtilizedResourcesResourceUnitTypeEnum, bool) {
	enum, ok := mappingListTopUtilizedResourcesResourceUnitTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTopUtilizedResourcesSortOrderEnum Enum with underlying type: string
type ListTopUtilizedResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListTopUtilizedResourcesSortOrderEnum
const (
	ListTopUtilizedResourcesSortOrderAsc  ListTopUtilizedResourcesSortOrderEnum = "ASC"
	ListTopUtilizedResourcesSortOrderDesc ListTopUtilizedResourcesSortOrderEnum = "DESC"
)

var mappingListTopUtilizedResourcesSortOrderEnum = map[string]ListTopUtilizedResourcesSortOrderEnum{
	"ASC":  ListTopUtilizedResourcesSortOrderAsc,
	"DESC": ListTopUtilizedResourcesSortOrderDesc,
}

var mappingListTopUtilizedResourcesSortOrderEnumLowerCase = map[string]ListTopUtilizedResourcesSortOrderEnum{
	"asc":  ListTopUtilizedResourcesSortOrderAsc,
	"desc": ListTopUtilizedResourcesSortOrderDesc,
}

// GetListTopUtilizedResourcesSortOrderEnumValues Enumerates the set of values for ListTopUtilizedResourcesSortOrderEnum
func GetListTopUtilizedResourcesSortOrderEnumValues() []ListTopUtilizedResourcesSortOrderEnum {
	values := make([]ListTopUtilizedResourcesSortOrderEnum, 0)
	for _, v := range mappingListTopUtilizedResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopUtilizedResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListTopUtilizedResourcesSortOrderEnum
func GetListTopUtilizedResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTopUtilizedResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopUtilizedResourcesSortOrderEnum(val string) (ListTopUtilizedResourcesSortOrderEnum, bool) {
	enum, ok := mappingListTopUtilizedResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTopUtilizedResourcesSortByEnum Enum with underlying type: string
type ListTopUtilizedResourcesSortByEnum string

// Set of constants representing the allowable values for ListTopUtilizedResourcesSortByEnum
const (
	ListTopUtilizedResourcesSortByTotalunits ListTopUtilizedResourcesSortByEnum = "totalUnits"
)

var mappingListTopUtilizedResourcesSortByEnum = map[string]ListTopUtilizedResourcesSortByEnum{
	"totalUnits": ListTopUtilizedResourcesSortByTotalunits,
}

var mappingListTopUtilizedResourcesSortByEnumLowerCase = map[string]ListTopUtilizedResourcesSortByEnum{
	"totalunits": ListTopUtilizedResourcesSortByTotalunits,
}

// GetListTopUtilizedResourcesSortByEnumValues Enumerates the set of values for ListTopUtilizedResourcesSortByEnum
func GetListTopUtilizedResourcesSortByEnumValues() []ListTopUtilizedResourcesSortByEnum {
	values := make([]ListTopUtilizedResourcesSortByEnum, 0)
	for _, v := range mappingListTopUtilizedResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopUtilizedResourcesSortByEnumStringValues Enumerates the set of values in String for ListTopUtilizedResourcesSortByEnum
func GetListTopUtilizedResourcesSortByEnumStringValues() []string {
	return []string{
		"totalUnits",
	}
}

// GetMappingListTopUtilizedResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopUtilizedResourcesSortByEnum(val string) (ListTopUtilizedResourcesSortByEnum, bool) {
	enum, ok := mappingListTopUtilizedResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
