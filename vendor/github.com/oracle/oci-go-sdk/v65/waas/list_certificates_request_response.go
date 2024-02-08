// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCertificatesRequest wrapper for the ListCertificates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListCertificates.go.html to see an example of how to use ListCertificatesRequest.
type ListCertificatesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value by which certificate summaries are sorted in a paginated 'List' call. If unspecified, defaults to `timeCreated`.
	SortBy ListCertificatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListCertificatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter certificates using a list of certificates OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter certificates using a list of display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// Filter certificates using a list of lifecycle states.
	LifecycleState []LifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches certificates created on or after the specified date-time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches certificates created before the specified date-time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificatesSortOrderEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStatesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificatesResponse wrapper for the ListCertificates operation
type ListCertificatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CertificateSummary instances
	Items []CertificateSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificatesSortByEnum Enum with underlying type: string
type ListCertificatesSortByEnum string

// Set of constants representing the allowable values for ListCertificatesSortByEnum
const (
	ListCertificatesSortById            ListCertificatesSortByEnum = "id"
	ListCertificatesSortByCompartmentid ListCertificatesSortByEnum = "compartmentId"
	ListCertificatesSortByDisplayname   ListCertificatesSortByEnum = "displayName"
	ListCertificatesSortByNotvalidafter ListCertificatesSortByEnum = "notValidAfter"
	ListCertificatesSortByTimecreated   ListCertificatesSortByEnum = "timeCreated"
)

var mappingListCertificatesSortByEnum = map[string]ListCertificatesSortByEnum{
	"id":            ListCertificatesSortById,
	"compartmentId": ListCertificatesSortByCompartmentid,
	"displayName":   ListCertificatesSortByDisplayname,
	"notValidAfter": ListCertificatesSortByNotvalidafter,
	"timeCreated":   ListCertificatesSortByTimecreated,
}

var mappingListCertificatesSortByEnumLowerCase = map[string]ListCertificatesSortByEnum{
	"id":            ListCertificatesSortById,
	"compartmentid": ListCertificatesSortByCompartmentid,
	"displayname":   ListCertificatesSortByDisplayname,
	"notvalidafter": ListCertificatesSortByNotvalidafter,
	"timecreated":   ListCertificatesSortByTimecreated,
}

// GetListCertificatesSortByEnumValues Enumerates the set of values for ListCertificatesSortByEnum
func GetListCertificatesSortByEnumValues() []ListCertificatesSortByEnum {
	values := make([]ListCertificatesSortByEnum, 0)
	for _, v := range mappingListCertificatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificatesSortByEnumStringValues Enumerates the set of values in String for ListCertificatesSortByEnum
func GetListCertificatesSortByEnumStringValues() []string {
	return []string{
		"id",
		"compartmentId",
		"displayName",
		"notValidAfter",
		"timeCreated",
	}
}

// GetMappingListCertificatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificatesSortByEnum(val string) (ListCertificatesSortByEnum, bool) {
	enum, ok := mappingListCertificatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificatesSortOrderEnum Enum with underlying type: string
type ListCertificatesSortOrderEnum string

// Set of constants representing the allowable values for ListCertificatesSortOrderEnum
const (
	ListCertificatesSortOrderAsc  ListCertificatesSortOrderEnum = "ASC"
	ListCertificatesSortOrderDesc ListCertificatesSortOrderEnum = "DESC"
)

var mappingListCertificatesSortOrderEnum = map[string]ListCertificatesSortOrderEnum{
	"ASC":  ListCertificatesSortOrderAsc,
	"DESC": ListCertificatesSortOrderDesc,
}

var mappingListCertificatesSortOrderEnumLowerCase = map[string]ListCertificatesSortOrderEnum{
	"asc":  ListCertificatesSortOrderAsc,
	"desc": ListCertificatesSortOrderDesc,
}

// GetListCertificatesSortOrderEnumValues Enumerates the set of values for ListCertificatesSortOrderEnum
func GetListCertificatesSortOrderEnumValues() []ListCertificatesSortOrderEnum {
	values := make([]ListCertificatesSortOrderEnum, 0)
	for _, v := range mappingListCertificatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificatesSortOrderEnumStringValues Enumerates the set of values in String for ListCertificatesSortOrderEnum
func GetListCertificatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificatesSortOrderEnum(val string) (ListCertificatesSortOrderEnum, bool) {
	enum, ok := mappingListCertificatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
