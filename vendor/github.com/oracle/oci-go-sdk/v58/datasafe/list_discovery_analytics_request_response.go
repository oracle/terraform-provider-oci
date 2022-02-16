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

// ListDiscoveryAnalyticsRequest wrapper for the ListDiscoveryAnalytics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryAnalytics.go.html to see an example of how to use ListDiscoveryAnalyticsRequest.
type ListDiscoveryAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Attribute by which the discovery analytics data should be grouped.
	GroupBy ListDiscoveryAnalyticsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoveryAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListDiscoveryAnalyticsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryAnalyticsResponse wrapper for the ListDiscoveryAnalytics operation
type ListDiscoveryAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryAnalyticsCollection instances
	DiscoveryAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDiscoveryAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryAnalyticsGroupByEnum Enum with underlying type: string
type ListDiscoveryAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListDiscoveryAnalyticsGroupByEnum
const (
	ListDiscoveryAnalyticsGroupByTargetid             ListDiscoveryAnalyticsGroupByEnum = "targetId"
	ListDiscoveryAnalyticsGroupBySensitivedatamodelid ListDiscoveryAnalyticsGroupByEnum = "sensitiveDataModelId"
)

var mappingListDiscoveryAnalyticsGroupByEnum = map[string]ListDiscoveryAnalyticsGroupByEnum{
	"targetId":             ListDiscoveryAnalyticsGroupByTargetid,
	"sensitiveDataModelId": ListDiscoveryAnalyticsGroupBySensitivedatamodelid,
}

// GetListDiscoveryAnalyticsGroupByEnumValues Enumerates the set of values for ListDiscoveryAnalyticsGroupByEnum
func GetListDiscoveryAnalyticsGroupByEnumValues() []ListDiscoveryAnalyticsGroupByEnum {
	values := make([]ListDiscoveryAnalyticsGroupByEnum, 0)
	for _, v := range mappingListDiscoveryAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListDiscoveryAnalyticsGroupByEnum
func GetListDiscoveryAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"sensitiveDataModelId",
	}
}

// GetMappingListDiscoveryAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryAnalyticsGroupByEnum(val string) (ListDiscoveryAnalyticsGroupByEnum, bool) {
	mappingListDiscoveryAnalyticsGroupByEnumIgnoreCase := make(map[string]ListDiscoveryAnalyticsGroupByEnum)
	for k, v := range mappingListDiscoveryAnalyticsGroupByEnum {
		mappingListDiscoveryAnalyticsGroupByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDiscoveryAnalyticsGroupByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
