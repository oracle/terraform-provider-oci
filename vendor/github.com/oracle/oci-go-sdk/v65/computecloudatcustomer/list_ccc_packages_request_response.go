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

// ListCccPackagesRequest wrapper for the ListCccPackages operation
type ListCccPackagesRequest struct {

	// The unique identifier for the listing.
	CccListingId *string `mandatory:"true" contributesTo:"query" name:"cccListingId"`

	// The version of the package. Package versions are unique within a listing.
	CccPackageId *string `mandatory:"false" contributesTo:"query" name:"cccPackageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only packages that match the given package type exactly.
	PackageType ListCccPackagesPackageTypeEnum `mandatory:"false" contributesTo:"query" name:"packageType" omitEmpty:"true"`

	// The client request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `timeReleased` displays results in descending order by default.
	// You can change your preference by specifying a different sort order.
	SortBy ListCccPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccPackagesPackageTypeEnum(string(request.PackageType)); !ok && request.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", request.PackageType, strings.Join(GetListCccPackagesPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccPackagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccPackagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccPackagesResponse wrapper for the ListCccPackages operation
type ListCccPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccPackageCollection instances
	CccPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCccPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccPackagesPackageTypeEnum Enum with underlying type: string
type ListCccPackagesPackageTypeEnum string

// Set of constants representing the allowable values for ListCccPackagesPackageTypeEnum
const (
	ListCccPackagesPackageTypeOrchestration ListCccPackagesPackageTypeEnum = "Orchestration"
	ListCccPackagesPackageTypeImage         ListCccPackagesPackageTypeEnum = "Image"
	ListCccPackagesPackageTypeContainer     ListCccPackagesPackageTypeEnum = "Container"
	ListCccPackagesPackageTypeKubernetes    ListCccPackagesPackageTypeEnum = "Kubernetes"
	ListCccPackagesPackageTypeSaas          ListCccPackagesPackageTypeEnum = "Saas"
)

var mappingListCccPackagesPackageTypeEnum = map[string]ListCccPackagesPackageTypeEnum{
	"Orchestration": ListCccPackagesPackageTypeOrchestration,
	"Image":         ListCccPackagesPackageTypeImage,
	"Container":     ListCccPackagesPackageTypeContainer,
	"Kubernetes":    ListCccPackagesPackageTypeKubernetes,
	"Saas":          ListCccPackagesPackageTypeSaas,
}

var mappingListCccPackagesPackageTypeEnumLowerCase = map[string]ListCccPackagesPackageTypeEnum{
	"orchestration": ListCccPackagesPackageTypeOrchestration,
	"image":         ListCccPackagesPackageTypeImage,
	"container":     ListCccPackagesPackageTypeContainer,
	"kubernetes":    ListCccPackagesPackageTypeKubernetes,
	"saas":          ListCccPackagesPackageTypeSaas,
}

// GetListCccPackagesPackageTypeEnumValues Enumerates the set of values for ListCccPackagesPackageTypeEnum
func GetListCccPackagesPackageTypeEnumValues() []ListCccPackagesPackageTypeEnum {
	values := make([]ListCccPackagesPackageTypeEnum, 0)
	for _, v := range mappingListCccPackagesPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccPackagesPackageTypeEnumStringValues Enumerates the set of values in String for ListCccPackagesPackageTypeEnum
func GetListCccPackagesPackageTypeEnumStringValues() []string {
	return []string{
		"Orchestration",
		"Image",
		"Container",
		"Kubernetes",
		"Saas",
	}
}

// GetMappingListCccPackagesPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccPackagesPackageTypeEnum(val string) (ListCccPackagesPackageTypeEnum, bool) {
	enum, ok := mappingListCccPackagesPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccPackagesSortByEnum Enum with underlying type: string
type ListCccPackagesSortByEnum string

// Set of constants representing the allowable values for ListCccPackagesSortByEnum
const (
	ListCccPackagesSortByTimereleased ListCccPackagesSortByEnum = "timeReleased"
)

var mappingListCccPackagesSortByEnum = map[string]ListCccPackagesSortByEnum{
	"timeReleased": ListCccPackagesSortByTimereleased,
}

var mappingListCccPackagesSortByEnumLowerCase = map[string]ListCccPackagesSortByEnum{
	"timereleased": ListCccPackagesSortByTimereleased,
}

// GetListCccPackagesSortByEnumValues Enumerates the set of values for ListCccPackagesSortByEnum
func GetListCccPackagesSortByEnumValues() []ListCccPackagesSortByEnum {
	values := make([]ListCccPackagesSortByEnum, 0)
	for _, v := range mappingListCccPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccPackagesSortByEnumStringValues Enumerates the set of values in String for ListCccPackagesSortByEnum
func GetListCccPackagesSortByEnumStringValues() []string {
	return []string{
		"timeReleased",
	}
}

// GetMappingListCccPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccPackagesSortByEnum(val string) (ListCccPackagesSortByEnum, bool) {
	enum, ok := mappingListCccPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccPackagesSortOrderEnum Enum with underlying type: string
type ListCccPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListCccPackagesSortOrderEnum
const (
	ListCccPackagesSortOrderAsc  ListCccPackagesSortOrderEnum = "ASC"
	ListCccPackagesSortOrderDesc ListCccPackagesSortOrderEnum = "DESC"
)

var mappingListCccPackagesSortOrderEnum = map[string]ListCccPackagesSortOrderEnum{
	"ASC":  ListCccPackagesSortOrderAsc,
	"DESC": ListCccPackagesSortOrderDesc,
}

var mappingListCccPackagesSortOrderEnumLowerCase = map[string]ListCccPackagesSortOrderEnum{
	"asc":  ListCccPackagesSortOrderAsc,
	"desc": ListCccPackagesSortOrderDesc,
}

// GetListCccPackagesSortOrderEnumValues Enumerates the set of values for ListCccPackagesSortOrderEnum
func GetListCccPackagesSortOrderEnumValues() []ListCccPackagesSortOrderEnum {
	values := make([]ListCccPackagesSortOrderEnum, 0)
	for _, v := range mappingListCccPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListCccPackagesSortOrderEnum
func GetListCccPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccPackagesSortOrderEnum(val string) (ListCccPackagesSortOrderEnum, bool) {
	enum, ok := mappingListCccPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
