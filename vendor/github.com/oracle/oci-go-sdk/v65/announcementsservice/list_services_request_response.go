// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListServicesRequest wrapper for the ListServices operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListServices.go.html to see an example of how to use ListServicesRequest.
type ListServicesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only announcements affecting a specific platform.
	PlatformType ListServicesPlatformTypeEnum `mandatory:"false" contributesTo:"query" name:"platformType" omitEmpty:"true"`

	// Filter by comms manager name
	CommsManagerName ListServicesCommsManagerNameEnum `mandatory:"false" contributesTo:"query" name:"commsManagerName" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort by service name parameter
	SortBy ListServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether ascending ('ASC') or descending ('DESC').
	SortOrder ListServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServicesPlatformTypeEnum(string(request.PlatformType)); !ok && request.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", request.PlatformType, strings.Join(GetListServicesPlatformTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServicesCommsManagerNameEnum(string(request.CommsManagerName)); !ok && request.CommsManagerName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommsManagerName: %s. Supported values are: %s.", request.CommsManagerName, strings.Join(GetListServicesCommsManagerNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServicesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServicesResponse wrapper for the ListServices operation
type ListServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServicesCollection instances
	ServicesCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServicesPlatformTypeEnum Enum with underlying type: string
type ListServicesPlatformTypeEnum string

// Set of constants representing the allowable values for ListServicesPlatformTypeEnum
const (
	ListServicesPlatformTypeIaas ListServicesPlatformTypeEnum = "IAAS"
	ListServicesPlatformTypeSaas ListServicesPlatformTypeEnum = "SAAS"
	ListServicesPlatformTypePaas ListServicesPlatformTypeEnum = "PAAS"
)

var mappingListServicesPlatformTypeEnum = map[string]ListServicesPlatformTypeEnum{
	"IAAS": ListServicesPlatformTypeIaas,
	"SAAS": ListServicesPlatformTypeSaas,
	"PAAS": ListServicesPlatformTypePaas,
}

var mappingListServicesPlatformTypeEnumLowerCase = map[string]ListServicesPlatformTypeEnum{
	"iaas": ListServicesPlatformTypeIaas,
	"saas": ListServicesPlatformTypeSaas,
	"paas": ListServicesPlatformTypePaas,
}

// GetListServicesPlatformTypeEnumValues Enumerates the set of values for ListServicesPlatformTypeEnum
func GetListServicesPlatformTypeEnumValues() []ListServicesPlatformTypeEnum {
	values := make([]ListServicesPlatformTypeEnum, 0)
	for _, v := range mappingListServicesPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListServicesPlatformTypeEnumStringValues Enumerates the set of values in String for ListServicesPlatformTypeEnum
func GetListServicesPlatformTypeEnumStringValues() []string {
	return []string{
		"IAAS",
		"SAAS",
		"PAAS",
	}
}

// GetMappingListServicesPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServicesPlatformTypeEnum(val string) (ListServicesPlatformTypeEnum, bool) {
	enum, ok := mappingListServicesPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServicesCommsManagerNameEnum Enum with underlying type: string
type ListServicesCommsManagerNameEnum string

// Set of constants representing the allowable values for ListServicesCommsManagerNameEnum
const (
	ListServicesCommsManagerNameCn     ListServicesCommsManagerNameEnum = "CN"
	ListServicesCommsManagerNameFusion ListServicesCommsManagerNameEnum = "FUSION"
	ListServicesCommsManagerNameAs     ListServicesCommsManagerNameEnum = "AS"
	ListServicesCommsManagerNameErf    ListServicesCommsManagerNameEnum = "ERF"
)

var mappingListServicesCommsManagerNameEnum = map[string]ListServicesCommsManagerNameEnum{
	"CN":     ListServicesCommsManagerNameCn,
	"FUSION": ListServicesCommsManagerNameFusion,
	"AS":     ListServicesCommsManagerNameAs,
	"ERF":    ListServicesCommsManagerNameErf,
}

var mappingListServicesCommsManagerNameEnumLowerCase = map[string]ListServicesCommsManagerNameEnum{
	"cn":     ListServicesCommsManagerNameCn,
	"fusion": ListServicesCommsManagerNameFusion,
	"as":     ListServicesCommsManagerNameAs,
	"erf":    ListServicesCommsManagerNameErf,
}

// GetListServicesCommsManagerNameEnumValues Enumerates the set of values for ListServicesCommsManagerNameEnum
func GetListServicesCommsManagerNameEnumValues() []ListServicesCommsManagerNameEnum {
	values := make([]ListServicesCommsManagerNameEnum, 0)
	for _, v := range mappingListServicesCommsManagerNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListServicesCommsManagerNameEnumStringValues Enumerates the set of values in String for ListServicesCommsManagerNameEnum
func GetListServicesCommsManagerNameEnumStringValues() []string {
	return []string{
		"CN",
		"FUSION",
		"AS",
		"ERF",
	}
}

// GetMappingListServicesCommsManagerNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServicesCommsManagerNameEnum(val string) (ListServicesCommsManagerNameEnum, bool) {
	enum, ok := mappingListServicesCommsManagerNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServicesSortByEnum Enum with underlying type: string
type ListServicesSortByEnum string

// Set of constants representing the allowable values for ListServicesSortByEnum
const (
	ListServicesSortByServicename ListServicesSortByEnum = "serviceName"
)

var mappingListServicesSortByEnum = map[string]ListServicesSortByEnum{
	"serviceName": ListServicesSortByServicename,
}

var mappingListServicesSortByEnumLowerCase = map[string]ListServicesSortByEnum{
	"servicename": ListServicesSortByServicename,
}

// GetListServicesSortByEnumValues Enumerates the set of values for ListServicesSortByEnum
func GetListServicesSortByEnumValues() []ListServicesSortByEnum {
	values := make([]ListServicesSortByEnum, 0)
	for _, v := range mappingListServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServicesSortByEnumStringValues Enumerates the set of values in String for ListServicesSortByEnum
func GetListServicesSortByEnumStringValues() []string {
	return []string{
		"serviceName",
	}
}

// GetMappingListServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServicesSortByEnum(val string) (ListServicesSortByEnum, bool) {
	enum, ok := mappingListServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServicesSortOrderEnum Enum with underlying type: string
type ListServicesSortOrderEnum string

// Set of constants representing the allowable values for ListServicesSortOrderEnum
const (
	ListServicesSortOrderAsc  ListServicesSortOrderEnum = "ASC"
	ListServicesSortOrderDesc ListServicesSortOrderEnum = "DESC"
)

var mappingListServicesSortOrderEnum = map[string]ListServicesSortOrderEnum{
	"ASC":  ListServicesSortOrderAsc,
	"DESC": ListServicesSortOrderDesc,
}

var mappingListServicesSortOrderEnumLowerCase = map[string]ListServicesSortOrderEnum{
	"asc":  ListServicesSortOrderAsc,
	"desc": ListServicesSortOrderDesc,
}

// GetListServicesSortOrderEnumValues Enumerates the set of values for ListServicesSortOrderEnum
func GetListServicesSortOrderEnumValues() []ListServicesSortOrderEnum {
	values := make([]ListServicesSortOrderEnum, 0)
	for _, v := range mappingListServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServicesSortOrderEnumStringValues Enumerates the set of values in String for ListServicesSortOrderEnum
func GetListServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServicesSortOrderEnum(val string) (ListServicesSortOrderEnum, bool) {
	enum, ok := mappingListServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
