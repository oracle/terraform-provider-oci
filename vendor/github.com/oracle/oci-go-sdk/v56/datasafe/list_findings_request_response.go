// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListFindingsRequest wrapper for the ListFindings operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindings.go.html to see an example of how to use ListFindingsRequest.
type ListFindingsRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only findings of a particular risk level.
	Severity ListFindingsSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListFindingsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Each finding has a key. This key is same for the finding across targets
	FindingKey *string `mandatory:"false" contributesTo:"query" name:"findingKey"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFindingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFindingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFindingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFindingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFindingsResponse wrapper for the ListFindings operation
type ListFindingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FindingSummary instances
	Items []FindingSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListFindingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFindingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFindingsSeverityEnum Enum with underlying type: string
type ListFindingsSeverityEnum string

// Set of constants representing the allowable values for ListFindingsSeverityEnum
const (
	ListFindingsSeverityHigh     ListFindingsSeverityEnum = "HIGH"
	ListFindingsSeverityMedium   ListFindingsSeverityEnum = "MEDIUM"
	ListFindingsSeverityLow      ListFindingsSeverityEnum = "LOW"
	ListFindingsSeverityEvaluate ListFindingsSeverityEnum = "EVALUATE"
	ListFindingsSeverityAdvisory ListFindingsSeverityEnum = "ADVISORY"
	ListFindingsSeverityPass     ListFindingsSeverityEnum = "PASS"
)

var mappingListFindingsSeverity = map[string]ListFindingsSeverityEnum{
	"HIGH":     ListFindingsSeverityHigh,
	"MEDIUM":   ListFindingsSeverityMedium,
	"LOW":      ListFindingsSeverityLow,
	"EVALUATE": ListFindingsSeverityEvaluate,
	"ADVISORY": ListFindingsSeverityAdvisory,
	"PASS":     ListFindingsSeverityPass,
}

// GetListFindingsSeverityEnumValues Enumerates the set of values for ListFindingsSeverityEnum
func GetListFindingsSeverityEnumValues() []ListFindingsSeverityEnum {
	values := make([]ListFindingsSeverityEnum, 0)
	for _, v := range mappingListFindingsSeverity {
		values = append(values, v)
	}
	return values
}

// ListFindingsAccessLevelEnum Enum with underlying type: string
type ListFindingsAccessLevelEnum string

// Set of constants representing the allowable values for ListFindingsAccessLevelEnum
const (
	ListFindingsAccessLevelRestricted ListFindingsAccessLevelEnum = "RESTRICTED"
	ListFindingsAccessLevelAccessible ListFindingsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListFindingsAccessLevel = map[string]ListFindingsAccessLevelEnum{
	"RESTRICTED": ListFindingsAccessLevelRestricted,
	"ACCESSIBLE": ListFindingsAccessLevelAccessible,
}

// GetListFindingsAccessLevelEnumValues Enumerates the set of values for ListFindingsAccessLevelEnum
func GetListFindingsAccessLevelEnumValues() []ListFindingsAccessLevelEnum {
	values := make([]ListFindingsAccessLevelEnum, 0)
	for _, v := range mappingListFindingsAccessLevel {
		values = append(values, v)
	}
	return values
}
