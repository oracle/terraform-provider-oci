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

// ListTemplateAssociationAnalyticsRequest wrapper for the ListTemplateAssociationAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTemplateAssociationAnalytics.go.html to see an example of how to use ListTemplateAssociationAnalyticsRequest.
type ListTemplateAssociationAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTemplateAssociationAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The OCID of the security assessment of type TEMPLATE.
	TemplateAssessmentId *string `mandatory:"false" contributesTo:"query" name:"templateAssessmentId"`

	// The OCID of the security assessment of type TEMPLATE_BASELINE.
	TemplateBaselineAssessmentId *string `mandatory:"false" contributesTo:"query" name:"templateBaselineAssessmentId"`

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

func (request ListTemplateAssociationAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTemplateAssociationAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTemplateAssociationAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTemplateAssociationAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTemplateAssociationAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTemplateAssociationAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTemplateAssociationAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTemplateAssociationAnalyticsResponse wrapper for the ListTemplateAssociationAnalytics operation
type ListTemplateAssociationAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TemplateAssociationAnalyticsCollection instances
	TemplateAssociationAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTemplateAssociationAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTemplateAssociationAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTemplateAssociationAnalyticsAccessLevelEnum Enum with underlying type: string
type ListTemplateAssociationAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListTemplateAssociationAnalyticsAccessLevelEnum
const (
	ListTemplateAssociationAnalyticsAccessLevelRestricted ListTemplateAssociationAnalyticsAccessLevelEnum = "RESTRICTED"
	ListTemplateAssociationAnalyticsAccessLevelAccessible ListTemplateAssociationAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTemplateAssociationAnalyticsAccessLevelEnum = map[string]ListTemplateAssociationAnalyticsAccessLevelEnum{
	"RESTRICTED": ListTemplateAssociationAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListTemplateAssociationAnalyticsAccessLevelAccessible,
}

var mappingListTemplateAssociationAnalyticsAccessLevelEnumLowerCase = map[string]ListTemplateAssociationAnalyticsAccessLevelEnum{
	"restricted": ListTemplateAssociationAnalyticsAccessLevelRestricted,
	"accessible": ListTemplateAssociationAnalyticsAccessLevelAccessible,
}

// GetListTemplateAssociationAnalyticsAccessLevelEnumValues Enumerates the set of values for ListTemplateAssociationAnalyticsAccessLevelEnum
func GetListTemplateAssociationAnalyticsAccessLevelEnumValues() []ListTemplateAssociationAnalyticsAccessLevelEnum {
	values := make([]ListTemplateAssociationAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListTemplateAssociationAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTemplateAssociationAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListTemplateAssociationAnalyticsAccessLevelEnum
func GetListTemplateAssociationAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTemplateAssociationAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplateAssociationAnalyticsAccessLevelEnum(val string) (ListTemplateAssociationAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListTemplateAssociationAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
