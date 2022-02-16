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

// ListAuditTrailAnalyticsRequest wrapper for the ListAuditTrailAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrailAnalytics.go.html to see an example of how to use ListAuditTrailAnalyticsRequest.
type ListAuditTrailAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditTrailAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The group by parameter for summarize operation on audit trail.
	GroupBy []ListAuditTrailAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditTrailAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditTrailAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditTrailAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditTrailAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditTrailAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditTrailAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditTrailAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListAuditTrailAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListAuditTrailAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditTrailAnalyticsResponse wrapper for the ListAuditTrailAnalytics operation
type ListAuditTrailAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditTrailAnalyticCollection instances
	AuditTrailAnalyticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditTrailAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditTrailAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditTrailAnalyticsAccessLevelEnum Enum with underlying type: string
type ListAuditTrailAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditTrailAnalyticsAccessLevelEnum
const (
	ListAuditTrailAnalyticsAccessLevelRestricted ListAuditTrailAnalyticsAccessLevelEnum = "RESTRICTED"
	ListAuditTrailAnalyticsAccessLevelAccessible ListAuditTrailAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditTrailAnalyticsAccessLevelEnum = map[string]ListAuditTrailAnalyticsAccessLevelEnum{
	"RESTRICTED": ListAuditTrailAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditTrailAnalyticsAccessLevelAccessible,
}

// GetListAuditTrailAnalyticsAccessLevelEnumValues Enumerates the set of values for ListAuditTrailAnalyticsAccessLevelEnum
func GetListAuditTrailAnalyticsAccessLevelEnumValues() []ListAuditTrailAnalyticsAccessLevelEnum {
	values := make([]ListAuditTrailAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListAuditTrailAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditTrailAnalyticsAccessLevelEnum
func GetListAuditTrailAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditTrailAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailAnalyticsAccessLevelEnum(val string) (ListAuditTrailAnalyticsAccessLevelEnum, bool) {
	mappingListAuditTrailAnalyticsAccessLevelEnumIgnoreCase := make(map[string]ListAuditTrailAnalyticsAccessLevelEnum)
	for k, v := range mappingListAuditTrailAnalyticsAccessLevelEnum {
		mappingListAuditTrailAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditTrailAnalyticsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditTrailAnalyticsGroupByEnum Enum with underlying type: string
type ListAuditTrailAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListAuditTrailAnalyticsGroupByEnum
const (
	ListAuditTrailAnalyticsGroupByLocation       ListAuditTrailAnalyticsGroupByEnum = "location"
	ListAuditTrailAnalyticsGroupByLifecyclestate ListAuditTrailAnalyticsGroupByEnum = "lifecycleState"
	ListAuditTrailAnalyticsGroupByStatus         ListAuditTrailAnalyticsGroupByEnum = "status"
	ListAuditTrailAnalyticsGroupByTargetid       ListAuditTrailAnalyticsGroupByEnum = "targetId"
)

var mappingListAuditTrailAnalyticsGroupByEnum = map[string]ListAuditTrailAnalyticsGroupByEnum{
	"location":       ListAuditTrailAnalyticsGroupByLocation,
	"lifecycleState": ListAuditTrailAnalyticsGroupByLifecyclestate,
	"status":         ListAuditTrailAnalyticsGroupByStatus,
	"targetId":       ListAuditTrailAnalyticsGroupByTargetid,
}

// GetListAuditTrailAnalyticsGroupByEnumValues Enumerates the set of values for ListAuditTrailAnalyticsGroupByEnum
func GetListAuditTrailAnalyticsGroupByEnumValues() []ListAuditTrailAnalyticsGroupByEnum {
	values := make([]ListAuditTrailAnalyticsGroupByEnum, 0)
	for _, v := range mappingListAuditTrailAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListAuditTrailAnalyticsGroupByEnum
func GetListAuditTrailAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"location",
		"lifecycleState",
		"status",
		"targetId",
	}
}

// GetMappingListAuditTrailAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailAnalyticsGroupByEnum(val string) (ListAuditTrailAnalyticsGroupByEnum, bool) {
	mappingListAuditTrailAnalyticsGroupByEnumIgnoreCase := make(map[string]ListAuditTrailAnalyticsGroupByEnum)
	for k, v := range mappingListAuditTrailAnalyticsGroupByEnum {
		mappingListAuditTrailAnalyticsGroupByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditTrailAnalyticsGroupByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
