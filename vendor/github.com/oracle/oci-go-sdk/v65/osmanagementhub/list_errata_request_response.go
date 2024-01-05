// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListErrataRequest wrapper for the ListErrata operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListErrata.go.html to see an example of how to use ListErrataRequest.
type ListErrataRequest struct {

	// The OCID of the compartment that contains the resources to list. This parameter is required.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that may partially match the erratum name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only profiles that match the given osFamily.
	OsFamily ListErrataOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// The advisory severity.
	AdvisorySeverity []AdvisorySeverityEnum `contributesTo:"query" name:"advisorySeverity" omitEmpty:"true" collectionFormat:"multi"`

	// The issue date after which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateStart"`

	// The issue date before which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateEnd"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListErrataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort errata by. Only one sort order may be provided. Default order for timeIssued is descending. Default order for name is ascending. If no value is specified timeIssued is default.
	SortBy ListErrataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListErrataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListErrataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListErrataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListErrataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListErrataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListErrataOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListErrataOsFamilyEnumStringValues(), ",")))
	}
	for _, val := range request.AdvisorySeverity {
		if _, ok := GetMappingAdvisorySeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisorySeverity: %s. Supported values are: %s.", val, strings.Join(GetAdvisorySeverityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListErrataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListErrataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListErrataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListErrataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListErrataResponse wrapper for the ListErrata operation
type ListErrataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ErratumCollection instances
	ErratumCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListErrataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListErrataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListErrataOsFamilyEnum Enum with underlying type: string
type ListErrataOsFamilyEnum string

// Set of constants representing the allowable values for ListErrataOsFamilyEnum
const (
	ListErrataOsFamily9 ListErrataOsFamilyEnum = "ORACLE_LINUX_9"
	ListErrataOsFamily8 ListErrataOsFamilyEnum = "ORACLE_LINUX_8"
	ListErrataOsFamily7 ListErrataOsFamilyEnum = "ORACLE_LINUX_7"
)

var mappingListErrataOsFamilyEnum = map[string]ListErrataOsFamilyEnum{
	"ORACLE_LINUX_9": ListErrataOsFamily9,
	"ORACLE_LINUX_8": ListErrataOsFamily8,
	"ORACLE_LINUX_7": ListErrataOsFamily7,
}

var mappingListErrataOsFamilyEnumLowerCase = map[string]ListErrataOsFamilyEnum{
	"oracle_linux_9": ListErrataOsFamily9,
	"oracle_linux_8": ListErrataOsFamily8,
	"oracle_linux_7": ListErrataOsFamily7,
}

// GetListErrataOsFamilyEnumValues Enumerates the set of values for ListErrataOsFamilyEnum
func GetListErrataOsFamilyEnumValues() []ListErrataOsFamilyEnum {
	values := make([]ListErrataOsFamilyEnum, 0)
	for _, v := range mappingListErrataOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListErrataOsFamilyEnumStringValues Enumerates the set of values in String for ListErrataOsFamilyEnum
func GetListErrataOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
	}
}

// GetMappingListErrataOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListErrataOsFamilyEnum(val string) (ListErrataOsFamilyEnum, bool) {
	enum, ok := mappingListErrataOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListErrataSortOrderEnum Enum with underlying type: string
type ListErrataSortOrderEnum string

// Set of constants representing the allowable values for ListErrataSortOrderEnum
const (
	ListErrataSortOrderAsc  ListErrataSortOrderEnum = "ASC"
	ListErrataSortOrderDesc ListErrataSortOrderEnum = "DESC"
)

var mappingListErrataSortOrderEnum = map[string]ListErrataSortOrderEnum{
	"ASC":  ListErrataSortOrderAsc,
	"DESC": ListErrataSortOrderDesc,
}

var mappingListErrataSortOrderEnumLowerCase = map[string]ListErrataSortOrderEnum{
	"asc":  ListErrataSortOrderAsc,
	"desc": ListErrataSortOrderDesc,
}

// GetListErrataSortOrderEnumValues Enumerates the set of values for ListErrataSortOrderEnum
func GetListErrataSortOrderEnumValues() []ListErrataSortOrderEnum {
	values := make([]ListErrataSortOrderEnum, 0)
	for _, v := range mappingListErrataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListErrataSortOrderEnumStringValues Enumerates the set of values in String for ListErrataSortOrderEnum
func GetListErrataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListErrataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListErrataSortOrderEnum(val string) (ListErrataSortOrderEnum, bool) {
	enum, ok := mappingListErrataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListErrataSortByEnum Enum with underlying type: string
type ListErrataSortByEnum string

// Set of constants representing the allowable values for ListErrataSortByEnum
const (
	ListErrataSortByTimeissued ListErrataSortByEnum = "timeIssued"
	ListErrataSortByName       ListErrataSortByEnum = "name"
)

var mappingListErrataSortByEnum = map[string]ListErrataSortByEnum{
	"timeIssued": ListErrataSortByTimeissued,
	"name":       ListErrataSortByName,
}

var mappingListErrataSortByEnumLowerCase = map[string]ListErrataSortByEnum{
	"timeissued": ListErrataSortByTimeissued,
	"name":       ListErrataSortByName,
}

// GetListErrataSortByEnumValues Enumerates the set of values for ListErrataSortByEnum
func GetListErrataSortByEnumValues() []ListErrataSortByEnum {
	values := make([]ListErrataSortByEnum, 0)
	for _, v := range mappingListErrataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListErrataSortByEnumStringValues Enumerates the set of values in String for ListErrataSortByEnum
func GetListErrataSortByEnumStringValues() []string {
	return []string{
		"timeIssued",
		"name",
	}
}

// GetMappingListErrataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListErrataSortByEnum(val string) (ListErrataSortByEnum, bool) {
	enum, ok := mappingListErrataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
