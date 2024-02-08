// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaskingAnalyticsRequest wrapper for the ListMaskingAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingAnalytics.go.html to see an example of how to use ListMaskingAnalyticsRequest.
type ListMaskingAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Attribute by which the masking analytics data should be grouped.
	GroupBy ListMaskingAnalyticsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

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

func (request ListMaskingAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListMaskingAnalyticsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingAnalyticsResponse wrapper for the ListMaskingAnalytics operation
type ListMaskingAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingAnalyticsCollection instances
	MaskingAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingAnalyticsGroupByEnum Enum with underlying type: string
type ListMaskingAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListMaskingAnalyticsGroupByEnum
const (
	ListMaskingAnalyticsGroupByTargetid ListMaskingAnalyticsGroupByEnum = "targetId"
	ListMaskingAnalyticsGroupByPolicyid ListMaskingAnalyticsGroupByEnum = "policyId"
)

var mappingListMaskingAnalyticsGroupByEnum = map[string]ListMaskingAnalyticsGroupByEnum{
	"targetId": ListMaskingAnalyticsGroupByTargetid,
	"policyId": ListMaskingAnalyticsGroupByPolicyid,
}

var mappingListMaskingAnalyticsGroupByEnumLowerCase = map[string]ListMaskingAnalyticsGroupByEnum{
	"targetid": ListMaskingAnalyticsGroupByTargetid,
	"policyid": ListMaskingAnalyticsGroupByPolicyid,
}

// GetListMaskingAnalyticsGroupByEnumValues Enumerates the set of values for ListMaskingAnalyticsGroupByEnum
func GetListMaskingAnalyticsGroupByEnumValues() []ListMaskingAnalyticsGroupByEnum {
	values := make([]ListMaskingAnalyticsGroupByEnum, 0)
	for _, v := range mappingListMaskingAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListMaskingAnalyticsGroupByEnum
func GetListMaskingAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"policyId",
	}
}

// GetMappingListMaskingAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingAnalyticsGroupByEnum(val string) (ListMaskingAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListMaskingAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
