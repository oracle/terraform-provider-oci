// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCapacityReservationInfosRequest wrapper for the ListCapacityReservationInfos operation
type ListCapacityReservationInfosRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the capacity reservation.
	CapacityReservationId *string `mandatory:"false" contributesTo:"query" name:"capacityReservationId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer CPG.
	CustomerCpgId *string `mandatory:"false" contributesTo:"query" name:"customerCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer tenancy.
	CustomerTenancyId *string `mandatory:"false" contributesTo:"query" name:"customerTenancyId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCapacityReservationInfosSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCapacityReservationInfosRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCapacityReservationInfosRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCapacityReservationInfosRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCapacityReservationInfosRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCapacityReservationInfosRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCapacityReservationInfosSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCapacityReservationInfosSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCapacityReservationInfosResponse wrapper for the ListCapacityReservationInfos operation
type ListCapacityReservationInfosResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CapacityReservationInfoCollection instances
	CapacityReservationInfoCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCapacityReservationInfosResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCapacityReservationInfosResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCapacityReservationInfosSortOrderEnum Enum with underlying type: string
type ListCapacityReservationInfosSortOrderEnum string

// Set of constants representing the allowable values for ListCapacityReservationInfosSortOrderEnum
const (
	ListCapacityReservationInfosSortOrderAsc  ListCapacityReservationInfosSortOrderEnum = "ASC"
	ListCapacityReservationInfosSortOrderDesc ListCapacityReservationInfosSortOrderEnum = "DESC"
)

var mappingListCapacityReservationInfosSortOrderEnum = map[string]ListCapacityReservationInfosSortOrderEnum{
	"ASC":  ListCapacityReservationInfosSortOrderAsc,
	"DESC": ListCapacityReservationInfosSortOrderDesc,
}

var mappingListCapacityReservationInfosSortOrderEnumLowerCase = map[string]ListCapacityReservationInfosSortOrderEnum{
	"asc":  ListCapacityReservationInfosSortOrderAsc,
	"desc": ListCapacityReservationInfosSortOrderDesc,
}

// GetListCapacityReservationInfosSortOrderEnumValues Enumerates the set of values for ListCapacityReservationInfosSortOrderEnum
func GetListCapacityReservationInfosSortOrderEnumValues() []ListCapacityReservationInfosSortOrderEnum {
	values := make([]ListCapacityReservationInfosSortOrderEnum, 0)
	for _, v := range mappingListCapacityReservationInfosSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCapacityReservationInfosSortOrderEnumStringValues Enumerates the set of values in String for ListCapacityReservationInfosSortOrderEnum
func GetListCapacityReservationInfosSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCapacityReservationInfosSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCapacityReservationInfosSortOrderEnum(val string) (ListCapacityReservationInfosSortOrderEnum, bool) {
	enum, ok := mappingListCapacityReservationInfosSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
