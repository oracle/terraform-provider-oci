// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsagesAndQuotasRequest wrapper for the ListUsagesAndQuotas operation
type ListUsagesAndQuotasRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can choose any value.
	// By default, when you sort by default user type, results are shown
	// in descending order.
	SortBy ListUsagesAndQuotasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListUsagesAndQuotasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// An option to only display the users or groups that violate their quota rules.
	// If `areViolatorsOnly` is false, the list result will display all the quota and usage report.
	// If `areViolatorsOnly` is true, the list result will only display the quota and usage report for
	// the users or groups that violate their quota rules.
	AreViolatorsOnly *bool `mandatory:"false" contributesTo:"query" name:"areViolatorsOnly"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsagesAndQuotasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsagesAndQuotasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsagesAndQuotasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsagesAndQuotasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsagesAndQuotasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUsagesAndQuotasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsagesAndQuotasSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsagesAndQuotasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsagesAndQuotasSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsagesAndQuotasResponse wrapper for the ListUsagesAndQuotas operation
type ListUsagesAndQuotasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UsageAndQuotaSummary instances
	Items []UsageAndQuotaSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListUsagesAndQuotasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsagesAndQuotasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsagesAndQuotasSortByEnum Enum with underlying type: string
type ListUsagesAndQuotasSortByEnum string

// Set of constants representing the allowable values for ListUsagesAndQuotasSortByEnum
const (
	ListUsagesAndQuotasSortByFileSystemLevel ListUsagesAndQuotasSortByEnum = "FILE_SYSTEM_LEVEL"
	ListUsagesAndQuotasSortByDefaultGroup    ListUsagesAndQuotasSortByEnum = "DEFAULT_GROUP"
	ListUsagesAndQuotasSortByDefaultUser     ListUsagesAndQuotasSortByEnum = "DEFAULT_USER"
	ListUsagesAndQuotasSortByIndividualGroup ListUsagesAndQuotasSortByEnum = "INDIVIDUAL_GROUP"
	ListUsagesAndQuotasSortByIndividualUser  ListUsagesAndQuotasSortByEnum = "INDIVIDUAL_USER"
)

var mappingListUsagesAndQuotasSortByEnum = map[string]ListUsagesAndQuotasSortByEnum{
	"FILE_SYSTEM_LEVEL": ListUsagesAndQuotasSortByFileSystemLevel,
	"DEFAULT_GROUP":     ListUsagesAndQuotasSortByDefaultGroup,
	"DEFAULT_USER":      ListUsagesAndQuotasSortByDefaultUser,
	"INDIVIDUAL_GROUP":  ListUsagesAndQuotasSortByIndividualGroup,
	"INDIVIDUAL_USER":   ListUsagesAndQuotasSortByIndividualUser,
}

var mappingListUsagesAndQuotasSortByEnumLowerCase = map[string]ListUsagesAndQuotasSortByEnum{
	"file_system_level": ListUsagesAndQuotasSortByFileSystemLevel,
	"default_group":     ListUsagesAndQuotasSortByDefaultGroup,
	"default_user":      ListUsagesAndQuotasSortByDefaultUser,
	"individual_group":  ListUsagesAndQuotasSortByIndividualGroup,
	"individual_user":   ListUsagesAndQuotasSortByIndividualUser,
}

// GetListUsagesAndQuotasSortByEnumValues Enumerates the set of values for ListUsagesAndQuotasSortByEnum
func GetListUsagesAndQuotasSortByEnumValues() []ListUsagesAndQuotasSortByEnum {
	values := make([]ListUsagesAndQuotasSortByEnum, 0)
	for _, v := range mappingListUsagesAndQuotasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsagesAndQuotasSortByEnumStringValues Enumerates the set of values in String for ListUsagesAndQuotasSortByEnum
func GetListUsagesAndQuotasSortByEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingListUsagesAndQuotasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsagesAndQuotasSortByEnum(val string) (ListUsagesAndQuotasSortByEnum, bool) {
	enum, ok := mappingListUsagesAndQuotasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsagesAndQuotasSortOrderEnum Enum with underlying type: string
type ListUsagesAndQuotasSortOrderEnum string

// Set of constants representing the allowable values for ListUsagesAndQuotasSortOrderEnum
const (
	ListUsagesAndQuotasSortOrderAsc  ListUsagesAndQuotasSortOrderEnum = "ASC"
	ListUsagesAndQuotasSortOrderDesc ListUsagesAndQuotasSortOrderEnum = "DESC"
)

var mappingListUsagesAndQuotasSortOrderEnum = map[string]ListUsagesAndQuotasSortOrderEnum{
	"ASC":  ListUsagesAndQuotasSortOrderAsc,
	"DESC": ListUsagesAndQuotasSortOrderDesc,
}

var mappingListUsagesAndQuotasSortOrderEnumLowerCase = map[string]ListUsagesAndQuotasSortOrderEnum{
	"asc":  ListUsagesAndQuotasSortOrderAsc,
	"desc": ListUsagesAndQuotasSortOrderDesc,
}

// GetListUsagesAndQuotasSortOrderEnumValues Enumerates the set of values for ListUsagesAndQuotasSortOrderEnum
func GetListUsagesAndQuotasSortOrderEnumValues() []ListUsagesAndQuotasSortOrderEnum {
	values := make([]ListUsagesAndQuotasSortOrderEnum, 0)
	for _, v := range mappingListUsagesAndQuotasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsagesAndQuotasSortOrderEnumStringValues Enumerates the set of values in String for ListUsagesAndQuotasSortOrderEnum
func GetListUsagesAndQuotasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsagesAndQuotasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsagesAndQuotasSortOrderEnum(val string) (ListUsagesAndQuotasSortOrderEnum, bool) {
	enum, ok := mappingListUsagesAndQuotasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
