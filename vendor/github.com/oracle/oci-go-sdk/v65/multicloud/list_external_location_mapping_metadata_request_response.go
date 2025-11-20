// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalLocationMappingMetadataRequest wrapper for the ListExternalLocationMappingMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationMappingMetadata.go.html to see an example of how to use ListExternalLocationMappingMetadataRequest.
type ListExternalLocationMappingMetadataRequest struct {

	// The subscription type of the Cloud Service Provider.
	SubscriptionServiceName []SubscriptionTypeEnum `contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment in which to list resources.
	// A Multicloud base compartment is an OCI compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalLocationMappingMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListExternalLocationMappingMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalLocationMappingMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalLocationMappingMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalLocationMappingMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalLocationMappingMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalLocationMappingMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.SubscriptionServiceName {
		if _, ok := GetMappingSubscriptionTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", val, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListExternalLocationMappingMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalLocationMappingMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationMappingMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalLocationMappingMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalLocationMappingMetadataResponse wrapper for the ListExternalLocationMappingMetadata operation
type ListExternalLocationMappingMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalLocationMappingMetadatumSummaryCollection instances
	ExternalLocationMappingMetadatumSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalLocationMappingMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalLocationMappingMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalLocationMappingMetadataSortOrderEnum Enum with underlying type: string
type ListExternalLocationMappingMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListExternalLocationMappingMetadataSortOrderEnum
const (
	ListExternalLocationMappingMetadataSortOrderAsc  ListExternalLocationMappingMetadataSortOrderEnum = "ASC"
	ListExternalLocationMappingMetadataSortOrderDesc ListExternalLocationMappingMetadataSortOrderEnum = "DESC"
)

var mappingListExternalLocationMappingMetadataSortOrderEnum = map[string]ListExternalLocationMappingMetadataSortOrderEnum{
	"ASC":  ListExternalLocationMappingMetadataSortOrderAsc,
	"DESC": ListExternalLocationMappingMetadataSortOrderDesc,
}

var mappingListExternalLocationMappingMetadataSortOrderEnumLowerCase = map[string]ListExternalLocationMappingMetadataSortOrderEnum{
	"asc":  ListExternalLocationMappingMetadataSortOrderAsc,
	"desc": ListExternalLocationMappingMetadataSortOrderDesc,
}

// GetListExternalLocationMappingMetadataSortOrderEnumValues Enumerates the set of values for ListExternalLocationMappingMetadataSortOrderEnum
func GetListExternalLocationMappingMetadataSortOrderEnumValues() []ListExternalLocationMappingMetadataSortOrderEnum {
	values := make([]ListExternalLocationMappingMetadataSortOrderEnum, 0)
	for _, v := range mappingListExternalLocationMappingMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationMappingMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListExternalLocationMappingMetadataSortOrderEnum
func GetListExternalLocationMappingMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalLocationMappingMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationMappingMetadataSortOrderEnum(val string) (ListExternalLocationMappingMetadataSortOrderEnum, bool) {
	enum, ok := mappingListExternalLocationMappingMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationMappingMetadataSortByEnum Enum with underlying type: string
type ListExternalLocationMappingMetadataSortByEnum string

// Set of constants representing the allowable values for ListExternalLocationMappingMetadataSortByEnum
const (
	ListExternalLocationMappingMetadataSortByTimecreated ListExternalLocationMappingMetadataSortByEnum = "timeCreated"
	ListExternalLocationMappingMetadataSortByDisplayname ListExternalLocationMappingMetadataSortByEnum = "displayName"
)

var mappingListExternalLocationMappingMetadataSortByEnum = map[string]ListExternalLocationMappingMetadataSortByEnum{
	"timeCreated": ListExternalLocationMappingMetadataSortByTimecreated,
	"displayName": ListExternalLocationMappingMetadataSortByDisplayname,
}

var mappingListExternalLocationMappingMetadataSortByEnumLowerCase = map[string]ListExternalLocationMappingMetadataSortByEnum{
	"timecreated": ListExternalLocationMappingMetadataSortByTimecreated,
	"displayname": ListExternalLocationMappingMetadataSortByDisplayname,
}

// GetListExternalLocationMappingMetadataSortByEnumValues Enumerates the set of values for ListExternalLocationMappingMetadataSortByEnum
func GetListExternalLocationMappingMetadataSortByEnumValues() []ListExternalLocationMappingMetadataSortByEnum {
	values := make([]ListExternalLocationMappingMetadataSortByEnum, 0)
	for _, v := range mappingListExternalLocationMappingMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationMappingMetadataSortByEnumStringValues Enumerates the set of values in String for ListExternalLocationMappingMetadataSortByEnum
func GetListExternalLocationMappingMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListExternalLocationMappingMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationMappingMetadataSortByEnum(val string) (ListExternalLocationMappingMetadataSortByEnum, bool) {
	enum, ok := mappingListExternalLocationMappingMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
