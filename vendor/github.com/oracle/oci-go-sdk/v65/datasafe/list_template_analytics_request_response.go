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

// ListTemplateAnalyticsRequest wrapper for the ListTemplateAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTemplateAnalytics.go.html to see an example of how to use ListTemplateAnalyticsRequest.
type ListTemplateAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTemplateAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The OCID of the security assessment of type TEMPLATE.
	TemplateAssessmentId *string `mandatory:"false" contributesTo:"query" name:"templateAssessmentId"`

	// The OCID of the security assessment of type TEMPLATE_BASELINE.
	TemplateBaselineAssessmentId *string `mandatory:"false" contributesTo:"query" name:"templateBaselineAssessmentId"`

	// A filter to return only the target group related information if the OCID belongs to a target group.
	IsGroup *bool `mandatory:"false" contributesTo:"query" name:"isGroup"`

	// A filter to return only the statistics where the comparison between the latest assessment and the template baseline assessment is done.
	IsCompared *bool `mandatory:"false" contributesTo:"query" name:"isCompared"`

	// A filter to return only the statistics where the latest assessment is compliant with the template baseline assessment.
	IsCompliant *bool `mandatory:"false" contributesTo:"query" name:"isCompliant"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTemplateAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTemplateAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTemplateAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTemplateAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTemplateAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTemplateAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTemplateAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTemplateAnalyticsResponse wrapper for the ListTemplateAnalytics operation
type ListTemplateAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TemplateAnalyticsCollection instances
	TemplateAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTemplateAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTemplateAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTemplateAnalyticsAccessLevelEnum Enum with underlying type: string
type ListTemplateAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListTemplateAnalyticsAccessLevelEnum
const (
	ListTemplateAnalyticsAccessLevelRestricted ListTemplateAnalyticsAccessLevelEnum = "RESTRICTED"
	ListTemplateAnalyticsAccessLevelAccessible ListTemplateAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTemplateAnalyticsAccessLevelEnum = map[string]ListTemplateAnalyticsAccessLevelEnum{
	"RESTRICTED": ListTemplateAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListTemplateAnalyticsAccessLevelAccessible,
}

var mappingListTemplateAnalyticsAccessLevelEnumLowerCase = map[string]ListTemplateAnalyticsAccessLevelEnum{
	"restricted": ListTemplateAnalyticsAccessLevelRestricted,
	"accessible": ListTemplateAnalyticsAccessLevelAccessible,
}

// GetListTemplateAnalyticsAccessLevelEnumValues Enumerates the set of values for ListTemplateAnalyticsAccessLevelEnum
func GetListTemplateAnalyticsAccessLevelEnumValues() []ListTemplateAnalyticsAccessLevelEnum {
	values := make([]ListTemplateAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListTemplateAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTemplateAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListTemplateAnalyticsAccessLevelEnum
func GetListTemplateAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTemplateAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplateAnalyticsAccessLevelEnum(val string) (ListTemplateAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListTemplateAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
