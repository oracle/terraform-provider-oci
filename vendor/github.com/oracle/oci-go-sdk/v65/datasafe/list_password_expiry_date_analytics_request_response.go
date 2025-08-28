// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPasswordExpiryDateAnalyticsRequest wrapper for the ListPasswordExpiryDateAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListPasswordExpiryDateAnalytics.go.html to see an example of how to use ListPasswordExpiryDateAnalyticsRequest.
type ListPasswordExpiryDateAnalyticsRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListPasswordExpiryDateAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items that match the specified user category.
	UserCategory *string `mandatory:"false" contributesTo:"query" name:"userCategory"`

	// A filter to return users whose password expiry date in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimePasswordExpiryLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePasswordExpiryLessThan"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPasswordExpiryDateAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPasswordExpiryDateAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPasswordExpiryDateAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPasswordExpiryDateAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPasswordExpiryDateAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPasswordExpiryDateAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListPasswordExpiryDateAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPasswordExpiryDateAnalyticsResponse wrapper for the ListPasswordExpiryDateAnalytics operation
type ListPasswordExpiryDateAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []UserAggregation instance
	Items []UserAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPasswordExpiryDateAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPasswordExpiryDateAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPasswordExpiryDateAnalyticsAccessLevelEnum Enum with underlying type: string
type ListPasswordExpiryDateAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListPasswordExpiryDateAnalyticsAccessLevelEnum
const (
	ListPasswordExpiryDateAnalyticsAccessLevelRestricted ListPasswordExpiryDateAnalyticsAccessLevelEnum = "RESTRICTED"
	ListPasswordExpiryDateAnalyticsAccessLevelAccessible ListPasswordExpiryDateAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListPasswordExpiryDateAnalyticsAccessLevelEnum = map[string]ListPasswordExpiryDateAnalyticsAccessLevelEnum{
	"RESTRICTED": ListPasswordExpiryDateAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListPasswordExpiryDateAnalyticsAccessLevelAccessible,
}

var mappingListPasswordExpiryDateAnalyticsAccessLevelEnumLowerCase = map[string]ListPasswordExpiryDateAnalyticsAccessLevelEnum{
	"restricted": ListPasswordExpiryDateAnalyticsAccessLevelRestricted,
	"accessible": ListPasswordExpiryDateAnalyticsAccessLevelAccessible,
}

// GetListPasswordExpiryDateAnalyticsAccessLevelEnumValues Enumerates the set of values for ListPasswordExpiryDateAnalyticsAccessLevelEnum
func GetListPasswordExpiryDateAnalyticsAccessLevelEnumValues() []ListPasswordExpiryDateAnalyticsAccessLevelEnum {
	values := make([]ListPasswordExpiryDateAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListPasswordExpiryDateAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListPasswordExpiryDateAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListPasswordExpiryDateAnalyticsAccessLevelEnum
func GetListPasswordExpiryDateAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListPasswordExpiryDateAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPasswordExpiryDateAnalyticsAccessLevelEnum(val string) (ListPasswordExpiryDateAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListPasswordExpiryDateAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
