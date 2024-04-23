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

// ListManagedInstanceErrataRequest wrapper for the ListManagedInstanceErrata operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceErrata.go.html to see an example of how to use ListManagedInstanceErrataRequest.
type ListManagedInstanceErrataRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that may partially match the erratum name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceErrataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort errata by. Only one sort order may be provided. Default order for timeIssued is descending. Default order for name is ascending. If no value is specified timeIssued is default.
	SortBy ListManagedInstanceErrataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceErrataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceErrataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceErrataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceErrataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceErrataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListManagedInstanceErrataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceErrataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceErrataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceErrataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceErrataResponse wrapper for the ListManagedInstanceErrata operation
type ListManagedInstanceErrataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceErratumSummaryCollection instances
	ManagedInstanceErratumSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceErrataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceErrataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceErrataSortOrderEnum Enum with underlying type: string
type ListManagedInstanceErrataSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceErrataSortOrderEnum
const (
	ListManagedInstanceErrataSortOrderAsc  ListManagedInstanceErrataSortOrderEnum = "ASC"
	ListManagedInstanceErrataSortOrderDesc ListManagedInstanceErrataSortOrderEnum = "DESC"
)

var mappingListManagedInstanceErrataSortOrderEnum = map[string]ListManagedInstanceErrataSortOrderEnum{
	"ASC":  ListManagedInstanceErrataSortOrderAsc,
	"DESC": ListManagedInstanceErrataSortOrderDesc,
}

var mappingListManagedInstanceErrataSortOrderEnumLowerCase = map[string]ListManagedInstanceErrataSortOrderEnum{
	"asc":  ListManagedInstanceErrataSortOrderAsc,
	"desc": ListManagedInstanceErrataSortOrderDesc,
}

// GetListManagedInstanceErrataSortOrderEnumValues Enumerates the set of values for ListManagedInstanceErrataSortOrderEnum
func GetListManagedInstanceErrataSortOrderEnumValues() []ListManagedInstanceErrataSortOrderEnum {
	values := make([]ListManagedInstanceErrataSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceErrataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceErrataSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceErrataSortOrderEnum
func GetListManagedInstanceErrataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceErrataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceErrataSortOrderEnum(val string) (ListManagedInstanceErrataSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceErrataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceErrataSortByEnum Enum with underlying type: string
type ListManagedInstanceErrataSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceErrataSortByEnum
const (
	ListManagedInstanceErrataSortByTimeissued ListManagedInstanceErrataSortByEnum = "timeIssued"
	ListManagedInstanceErrataSortByName       ListManagedInstanceErrataSortByEnum = "name"
)

var mappingListManagedInstanceErrataSortByEnum = map[string]ListManagedInstanceErrataSortByEnum{
	"timeIssued": ListManagedInstanceErrataSortByTimeissued,
	"name":       ListManagedInstanceErrataSortByName,
}

var mappingListManagedInstanceErrataSortByEnumLowerCase = map[string]ListManagedInstanceErrataSortByEnum{
	"timeissued": ListManagedInstanceErrataSortByTimeissued,
	"name":       ListManagedInstanceErrataSortByName,
}

// GetListManagedInstanceErrataSortByEnumValues Enumerates the set of values for ListManagedInstanceErrataSortByEnum
func GetListManagedInstanceErrataSortByEnumValues() []ListManagedInstanceErrataSortByEnum {
	values := make([]ListManagedInstanceErrataSortByEnum, 0)
	for _, v := range mappingListManagedInstanceErrataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceErrataSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceErrataSortByEnum
func GetListManagedInstanceErrataSortByEnumStringValues() []string {
	return []string{
		"timeIssued",
		"name",
	}
}

// GetMappingListManagedInstanceErrataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceErrataSortByEnum(val string) (ListManagedInstanceErrataSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceErrataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
