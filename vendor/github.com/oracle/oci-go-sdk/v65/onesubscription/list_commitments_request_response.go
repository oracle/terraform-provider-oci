// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCommitmentsRequest wrapper for the ListCommitments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListCommitments.go.html to see an example of how to use ListCommitmentsRequest.
type ListCommitmentsRequest struct {

	// This param is used to get the commitments for a particular subscribed service
	SubscribedServiceId *string `mandatory:"true" contributesTo:"query" name:"subscribedServiceId"`

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListCommitmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListCommitmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCommitmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCommitmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCommitmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCommitmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCommitmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCommitmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCommitmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCommitmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCommitmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCommitmentsResponse wrapper for the ListCommitments operation
type ListCommitmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CommitmentSummary instances
	Items []CommitmentSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCommitmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCommitmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCommitmentsSortOrderEnum Enum with underlying type: string
type ListCommitmentsSortOrderEnum string

// Set of constants representing the allowable values for ListCommitmentsSortOrderEnum
const (
	ListCommitmentsSortOrderAsc  ListCommitmentsSortOrderEnum = "ASC"
	ListCommitmentsSortOrderDesc ListCommitmentsSortOrderEnum = "DESC"
)

var mappingListCommitmentsSortOrderEnum = map[string]ListCommitmentsSortOrderEnum{
	"ASC":  ListCommitmentsSortOrderAsc,
	"DESC": ListCommitmentsSortOrderDesc,
}

var mappingListCommitmentsSortOrderEnumLowerCase = map[string]ListCommitmentsSortOrderEnum{
	"asc":  ListCommitmentsSortOrderAsc,
	"desc": ListCommitmentsSortOrderDesc,
}

// GetListCommitmentsSortOrderEnumValues Enumerates the set of values for ListCommitmentsSortOrderEnum
func GetListCommitmentsSortOrderEnumValues() []ListCommitmentsSortOrderEnum {
	values := make([]ListCommitmentsSortOrderEnum, 0)
	for _, v := range mappingListCommitmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCommitmentsSortOrderEnumStringValues Enumerates the set of values in String for ListCommitmentsSortOrderEnum
func GetListCommitmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCommitmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCommitmentsSortOrderEnum(val string) (ListCommitmentsSortOrderEnum, bool) {
	enum, ok := mappingListCommitmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCommitmentsSortByEnum Enum with underlying type: string
type ListCommitmentsSortByEnum string

// Set of constants representing the allowable values for ListCommitmentsSortByEnum
const (
	ListCommitmentsSortByOrdernumber   ListCommitmentsSortByEnum = "ORDERNUMBER"
	ListCommitmentsSortByTimeinvoicing ListCommitmentsSortByEnum = "TIMEINVOICING"
)

var mappingListCommitmentsSortByEnum = map[string]ListCommitmentsSortByEnum{
	"ORDERNUMBER":   ListCommitmentsSortByOrdernumber,
	"TIMEINVOICING": ListCommitmentsSortByTimeinvoicing,
}

var mappingListCommitmentsSortByEnumLowerCase = map[string]ListCommitmentsSortByEnum{
	"ordernumber":   ListCommitmentsSortByOrdernumber,
	"timeinvoicing": ListCommitmentsSortByTimeinvoicing,
}

// GetListCommitmentsSortByEnumValues Enumerates the set of values for ListCommitmentsSortByEnum
func GetListCommitmentsSortByEnumValues() []ListCommitmentsSortByEnum {
	values := make([]ListCommitmentsSortByEnum, 0)
	for _, v := range mappingListCommitmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCommitmentsSortByEnumStringValues Enumerates the set of values in String for ListCommitmentsSortByEnum
func GetListCommitmentsSortByEnumStringValues() []string {
	return []string{
		"ORDERNUMBER",
		"TIMEINVOICING",
	}
}

// GetMappingListCommitmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCommitmentsSortByEnum(val string) (ListCommitmentsSortByEnum, bool) {
	enum, ok := mappingListCommitmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
