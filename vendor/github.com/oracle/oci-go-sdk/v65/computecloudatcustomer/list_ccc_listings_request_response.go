// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccListingsRequest wrapper for the ListCccListings operation
type ListCccListingsRequest struct {

	// The name of the listing.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The unique identifier for the listing.
	CccListingId *string `mandatory:"false" contributesTo:"query" name:"cccListingId"`

	// A filter to return only packages that match the given package type exactly.
	PackageType ListCccListingsPackageTypeEnum `mandatory:"false" contributesTo:"query" name:"packageType" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `timeReleased` displays results in descending order by default.
	// You can change your preference by specifying a different sort order.
	SortBy ListCccListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Name of the product category or categories. If you specify multiple categories, then Marketplace returns any listing with
	// one or more matching categories.
	Category []string `contributesTo:"query" name:"category" collectionFormat:"multi"`

	// Name of the pricing type. If multiple pricing types are provided, then any listing with
	// one or more matching pricing models will be returned.
	Pricing []ListCccListingsPricingEnum `contributesTo:"query" name:"pricing" omitEmpty:"true" collectionFormat:"multi"`

	// The operating system of the listing.
	OperatingSystems []string `contributesTo:"query" name:"operatingSystems" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccListingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccListingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccListingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccListingsPackageTypeEnum(string(request.PackageType)); !ok && request.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", request.PackageType, strings.Join(GetListCccListingsPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccListingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccListingsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccListingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccListingsSortOrderEnumStringValues(), ",")))
	}
	for _, val := range request.Pricing {
		if _, ok := GetMappingListCccListingsPricingEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pricing: %s. Supported values are: %s.", val, strings.Join(GetListCccListingsPricingEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccListingsResponse wrapper for the ListCccListings operation
type ListCccListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccListingCollection instances
	CccListingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCccListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccListingsPackageTypeEnum Enum with underlying type: string
type ListCccListingsPackageTypeEnum string

// Set of constants representing the allowable values for ListCccListingsPackageTypeEnum
const (
	ListCccListingsPackageTypeOrchestration ListCccListingsPackageTypeEnum = "Orchestration"
	ListCccListingsPackageTypeImage         ListCccListingsPackageTypeEnum = "Image"
	ListCccListingsPackageTypeContainer     ListCccListingsPackageTypeEnum = "Container"
	ListCccListingsPackageTypeKubernetes    ListCccListingsPackageTypeEnum = "Kubernetes"
	ListCccListingsPackageTypeSaas          ListCccListingsPackageTypeEnum = "Saas"
)

var mappingListCccListingsPackageTypeEnum = map[string]ListCccListingsPackageTypeEnum{
	"Orchestration": ListCccListingsPackageTypeOrchestration,
	"Image":         ListCccListingsPackageTypeImage,
	"Container":     ListCccListingsPackageTypeContainer,
	"Kubernetes":    ListCccListingsPackageTypeKubernetes,
	"Saas":          ListCccListingsPackageTypeSaas,
}

var mappingListCccListingsPackageTypeEnumLowerCase = map[string]ListCccListingsPackageTypeEnum{
	"orchestration": ListCccListingsPackageTypeOrchestration,
	"image":         ListCccListingsPackageTypeImage,
	"container":     ListCccListingsPackageTypeContainer,
	"kubernetes":    ListCccListingsPackageTypeKubernetes,
	"saas":          ListCccListingsPackageTypeSaas,
}

// GetListCccListingsPackageTypeEnumValues Enumerates the set of values for ListCccListingsPackageTypeEnum
func GetListCccListingsPackageTypeEnumValues() []ListCccListingsPackageTypeEnum {
	values := make([]ListCccListingsPackageTypeEnum, 0)
	for _, v := range mappingListCccListingsPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsPackageTypeEnumStringValues Enumerates the set of values in String for ListCccListingsPackageTypeEnum
func GetListCccListingsPackageTypeEnumStringValues() []string {
	return []string{
		"Orchestration",
		"Image",
		"Container",
		"Kubernetes",
		"Saas",
	}
}

// GetMappingListCccListingsPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsPackageTypeEnum(val string) (ListCccListingsPackageTypeEnum, bool) {
	enum, ok := mappingListCccListingsPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccListingsSortByEnum Enum with underlying type: string
type ListCccListingsSortByEnum string

// Set of constants representing the allowable values for ListCccListingsSortByEnum
const (
	ListCccListingsSortByTimereleased ListCccListingsSortByEnum = "timeReleased"
)

var mappingListCccListingsSortByEnum = map[string]ListCccListingsSortByEnum{
	"timeReleased": ListCccListingsSortByTimereleased,
}

var mappingListCccListingsSortByEnumLowerCase = map[string]ListCccListingsSortByEnum{
	"timereleased": ListCccListingsSortByTimereleased,
}

// GetListCccListingsSortByEnumValues Enumerates the set of values for ListCccListingsSortByEnum
func GetListCccListingsSortByEnumValues() []ListCccListingsSortByEnum {
	values := make([]ListCccListingsSortByEnum, 0)
	for _, v := range mappingListCccListingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsSortByEnumStringValues Enumerates the set of values in String for ListCccListingsSortByEnum
func GetListCccListingsSortByEnumStringValues() []string {
	return []string{
		"timeReleased",
	}
}

// GetMappingListCccListingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsSortByEnum(val string) (ListCccListingsSortByEnum, bool) {
	enum, ok := mappingListCccListingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccListingsSortOrderEnum Enum with underlying type: string
type ListCccListingsSortOrderEnum string

// Set of constants representing the allowable values for ListCccListingsSortOrderEnum
const (
	ListCccListingsSortOrderAsc  ListCccListingsSortOrderEnum = "ASC"
	ListCccListingsSortOrderDesc ListCccListingsSortOrderEnum = "DESC"
)

var mappingListCccListingsSortOrderEnum = map[string]ListCccListingsSortOrderEnum{
	"ASC":  ListCccListingsSortOrderAsc,
	"DESC": ListCccListingsSortOrderDesc,
}

var mappingListCccListingsSortOrderEnumLowerCase = map[string]ListCccListingsSortOrderEnum{
	"asc":  ListCccListingsSortOrderAsc,
	"desc": ListCccListingsSortOrderDesc,
}

// GetListCccListingsSortOrderEnumValues Enumerates the set of values for ListCccListingsSortOrderEnum
func GetListCccListingsSortOrderEnumValues() []ListCccListingsSortOrderEnum {
	values := make([]ListCccListingsSortOrderEnum, 0)
	for _, v := range mappingListCccListingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsSortOrderEnumStringValues Enumerates the set of values in String for ListCccListingsSortOrderEnum
func GetListCccListingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccListingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsSortOrderEnum(val string) (ListCccListingsSortOrderEnum, bool) {
	enum, ok := mappingListCccListingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccListingsPricingEnum Enum with underlying type: string
type ListCccListingsPricingEnum string

// Set of constants representing the allowable values for ListCccListingsPricingEnum
const (
	ListCccListingsPricingFree  ListCccListingsPricingEnum = "Free"
	ListCccListingsPricingByol  ListCccListingsPricingEnum = "BYOL"
	ListCccListingsPricingPaygo ListCccListingsPricingEnum = "PayGo"
)

var mappingListCccListingsPricingEnum = map[string]ListCccListingsPricingEnum{
	"Free":  ListCccListingsPricingFree,
	"BYOL":  ListCccListingsPricingByol,
	"PayGo": ListCccListingsPricingPaygo,
}

var mappingListCccListingsPricingEnumLowerCase = map[string]ListCccListingsPricingEnum{
	"free":  ListCccListingsPricingFree,
	"byol":  ListCccListingsPricingByol,
	"paygo": ListCccListingsPricingPaygo,
}

// GetListCccListingsPricingEnumValues Enumerates the set of values for ListCccListingsPricingEnum
func GetListCccListingsPricingEnumValues() []ListCccListingsPricingEnum {
	values := make([]ListCccListingsPricingEnum, 0)
	for _, v := range mappingListCccListingsPricingEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsPricingEnumStringValues Enumerates the set of values in String for ListCccListingsPricingEnum
func GetListCccListingsPricingEnumStringValues() []string {
	return []string{
		"Free",
		"BYOL",
		"PayGo",
	}
}

// GetMappingListCccListingsPricingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsPricingEnum(val string) (ListCccListingsPricingEnum, bool) {
	enum, ok := mappingListCccListingsPricingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
