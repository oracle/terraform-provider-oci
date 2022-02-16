// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAuditProfileAnalyticsRequest wrapper for the ListAuditProfileAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfileAnalytics.go.html to see an example of how to use ListAuditProfileAnalyticsRequest.
type ListAuditProfileAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditProfileAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The group by parameter for summarize operation on audit.
	GroupBy []ListAuditProfileAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditProfileAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditProfileAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditProfileAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditProfileAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditProfileAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditProfileAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditProfileAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListAuditProfileAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListAuditProfileAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditProfileAnalyticsResponse wrapper for the ListAuditProfileAnalytics operation
type ListAuditProfileAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditProfileAnalyticCollection instances
	AuditProfileAnalyticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditProfileAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditProfileAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditProfileAnalyticsAccessLevelEnum Enum with underlying type: string
type ListAuditProfileAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditProfileAnalyticsAccessLevelEnum
const (
	ListAuditProfileAnalyticsAccessLevelRestricted ListAuditProfileAnalyticsAccessLevelEnum = "RESTRICTED"
	ListAuditProfileAnalyticsAccessLevelAccessible ListAuditProfileAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditProfileAnalyticsAccessLevelEnum = map[string]ListAuditProfileAnalyticsAccessLevelEnum{
	"RESTRICTED": ListAuditProfileAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditProfileAnalyticsAccessLevelAccessible,
}

// GetListAuditProfileAnalyticsAccessLevelEnumValues Enumerates the set of values for ListAuditProfileAnalyticsAccessLevelEnum
func GetListAuditProfileAnalyticsAccessLevelEnumValues() []ListAuditProfileAnalyticsAccessLevelEnum {
	values := make([]ListAuditProfileAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListAuditProfileAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfileAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditProfileAnalyticsAccessLevelEnum
func GetListAuditProfileAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditProfileAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfileAnalyticsAccessLevelEnum(val string) (ListAuditProfileAnalyticsAccessLevelEnum, bool) {
	mappingListAuditProfileAnalyticsAccessLevelEnumIgnoreCase := make(map[string]ListAuditProfileAnalyticsAccessLevelEnum)
	for k, v := range mappingListAuditProfileAnalyticsAccessLevelEnum {
		mappingListAuditProfileAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditProfileAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditProfileAnalyticsGroupByEnum Enum with underlying type: string
type ListAuditProfileAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListAuditProfileAnalyticsGroupByEnum
const (
	ListAuditProfileAnalyticsGroupByIspaidusageenabled ListAuditProfileAnalyticsGroupByEnum = "isPaidUsageEnabled"
)

var mappingListAuditProfileAnalyticsGroupByEnum = map[string]ListAuditProfileAnalyticsGroupByEnum{
	"isPaidUsageEnabled": ListAuditProfileAnalyticsGroupByIspaidusageenabled,
}

// GetListAuditProfileAnalyticsGroupByEnumValues Enumerates the set of values for ListAuditProfileAnalyticsGroupByEnum
func GetListAuditProfileAnalyticsGroupByEnumValues() []ListAuditProfileAnalyticsGroupByEnum {
	values := make([]ListAuditProfileAnalyticsGroupByEnum, 0)
	for _, v := range mappingListAuditProfileAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfileAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListAuditProfileAnalyticsGroupByEnum
func GetListAuditProfileAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"isPaidUsageEnabled",
	}
}

// GetMappingListAuditProfileAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfileAnalyticsGroupByEnum(val string) (ListAuditProfileAnalyticsGroupByEnum, bool) {
	mappingListAuditProfileAnalyticsGroupByEnumIgnoreCase := make(map[string]ListAuditProfileAnalyticsGroupByEnum)
	for k, v := range mappingListAuditProfileAnalyticsGroupByEnum {
		mappingListAuditProfileAnalyticsGroupByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditProfileAnalyticsGroupByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
