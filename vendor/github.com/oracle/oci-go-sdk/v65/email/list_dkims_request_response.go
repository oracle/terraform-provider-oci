// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDkimsRequest wrapper for the ListDkims operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListDkims.go.html to see an example of how to use ListDkimsRequest.
type ListDkimsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain to which this DKIM belongs.
	EmailDomainId *string `mandatory:"true" contributesTo:"query" name:"emailDomainId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListDkimsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState DkimLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the DKIMs.
	// Default: `TIMECREATED`
	// * **TIMECREATED:** Sorts by timeCreated.
	// * **NAME:** Sorts by name.
	// * **ID:** Sorts by id.
	SortBy ListDkimsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDkimsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDkimsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDkimsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDkimsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDkimsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDkimsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDkimsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDkimLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDkimLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDkimsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDkimsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDkimsResponse wrapper for the ListDkims operation
type ListDkimsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DkimCollection instances
	DkimCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDkimsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDkimsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDkimsSortOrderEnum Enum with underlying type: string
type ListDkimsSortOrderEnum string

// Set of constants representing the allowable values for ListDkimsSortOrderEnum
const (
	ListDkimsSortOrderAsc  ListDkimsSortOrderEnum = "ASC"
	ListDkimsSortOrderDesc ListDkimsSortOrderEnum = "DESC"
)

var mappingListDkimsSortOrderEnum = map[string]ListDkimsSortOrderEnum{
	"ASC":  ListDkimsSortOrderAsc,
	"DESC": ListDkimsSortOrderDesc,
}

var mappingListDkimsSortOrderEnumLowerCase = map[string]ListDkimsSortOrderEnum{
	"asc":  ListDkimsSortOrderAsc,
	"desc": ListDkimsSortOrderDesc,
}

// GetListDkimsSortOrderEnumValues Enumerates the set of values for ListDkimsSortOrderEnum
func GetListDkimsSortOrderEnumValues() []ListDkimsSortOrderEnum {
	values := make([]ListDkimsSortOrderEnum, 0)
	for _, v := range mappingListDkimsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDkimsSortOrderEnumStringValues Enumerates the set of values in String for ListDkimsSortOrderEnum
func GetListDkimsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDkimsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDkimsSortOrderEnum(val string) (ListDkimsSortOrderEnum, bool) {
	enum, ok := mappingListDkimsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDkimsSortByEnum Enum with underlying type: string
type ListDkimsSortByEnum string

// Set of constants representing the allowable values for ListDkimsSortByEnum
const (
	ListDkimsSortByTimecreated ListDkimsSortByEnum = "TIMECREATED"
	ListDkimsSortById          ListDkimsSortByEnum = "ID"
	ListDkimsSortByName        ListDkimsSortByEnum = "NAME"
)

var mappingListDkimsSortByEnum = map[string]ListDkimsSortByEnum{
	"TIMECREATED": ListDkimsSortByTimecreated,
	"ID":          ListDkimsSortById,
	"NAME":        ListDkimsSortByName,
}

var mappingListDkimsSortByEnumLowerCase = map[string]ListDkimsSortByEnum{
	"timecreated": ListDkimsSortByTimecreated,
	"id":          ListDkimsSortById,
	"name":        ListDkimsSortByName,
}

// GetListDkimsSortByEnumValues Enumerates the set of values for ListDkimsSortByEnum
func GetListDkimsSortByEnumValues() []ListDkimsSortByEnum {
	values := make([]ListDkimsSortByEnum, 0)
	for _, v := range mappingListDkimsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDkimsSortByEnumStringValues Enumerates the set of values in String for ListDkimsSortByEnum
func GetListDkimsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"ID",
		"NAME",
	}
}

// GetMappingListDkimsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDkimsSortByEnum(val string) (ListDkimsSortByEnum, bool) {
	enum, ok := mappingListDkimsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
