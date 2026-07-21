// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPartnersRequest wrapper for the ListPartners operation
type ListPartnersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListPartnersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPartnersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPartnersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPartnersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPartnersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPartnersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPartnersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPartnersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPartnersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPartnersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPartnersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPartnersResponse wrapper for the ListPartners operation
type ListPartnersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PartnerCollection instances
	PartnerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPartnersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPartnersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPartnersSortByEnum Enum with underlying type: string
type ListPartnersSortByEnum string

// Set of constants representing the allowable values for ListPartnersSortByEnum
const (
	ListPartnersSortByTimecreated ListPartnersSortByEnum = "timeCreated"
	ListPartnersSortByDisplayname ListPartnersSortByEnum = "displayName"
	ListPartnersSortByProductid   ListPartnersSortByEnum = "productId"
)

var mappingListPartnersSortByEnum = map[string]ListPartnersSortByEnum{
	"timeCreated": ListPartnersSortByTimecreated,
	"displayName": ListPartnersSortByDisplayname,
	"productId":   ListPartnersSortByProductid,
}

var mappingListPartnersSortByEnumLowerCase = map[string]ListPartnersSortByEnum{
	"timecreated": ListPartnersSortByTimecreated,
	"displayname": ListPartnersSortByDisplayname,
	"productid":   ListPartnersSortByProductid,
}

// GetListPartnersSortByEnumValues Enumerates the set of values for ListPartnersSortByEnum
func GetListPartnersSortByEnumValues() []ListPartnersSortByEnum {
	values := make([]ListPartnersSortByEnum, 0)
	for _, v := range mappingListPartnersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPartnersSortByEnumStringValues Enumerates the set of values in String for ListPartnersSortByEnum
func GetListPartnersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"productId",
	}
}

// GetMappingListPartnersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPartnersSortByEnum(val string) (ListPartnersSortByEnum, bool) {
	enum, ok := mappingListPartnersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPartnersSortOrderEnum Enum with underlying type: string
type ListPartnersSortOrderEnum string

// Set of constants representing the allowable values for ListPartnersSortOrderEnum
const (
	ListPartnersSortOrderAsc  ListPartnersSortOrderEnum = "ASC"
	ListPartnersSortOrderDesc ListPartnersSortOrderEnum = "DESC"
)

var mappingListPartnersSortOrderEnum = map[string]ListPartnersSortOrderEnum{
	"ASC":  ListPartnersSortOrderAsc,
	"DESC": ListPartnersSortOrderDesc,
}

var mappingListPartnersSortOrderEnumLowerCase = map[string]ListPartnersSortOrderEnum{
	"asc":  ListPartnersSortOrderAsc,
	"desc": ListPartnersSortOrderDesc,
}

// GetListPartnersSortOrderEnumValues Enumerates the set of values for ListPartnersSortOrderEnum
func GetListPartnersSortOrderEnumValues() []ListPartnersSortOrderEnum {
	values := make([]ListPartnersSortOrderEnum, 0)
	for _, v := range mappingListPartnersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPartnersSortOrderEnumStringValues Enumerates the set of values in String for ListPartnersSortOrderEnum
func GetListPartnersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPartnersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPartnersSortOrderEnum(val string) (ListPartnersSortOrderEnum, bool) {
	enum, ok := mappingListPartnersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
