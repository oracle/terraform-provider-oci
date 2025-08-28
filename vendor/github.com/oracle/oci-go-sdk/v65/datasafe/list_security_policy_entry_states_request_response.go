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

// ListSecurityPolicyEntryStatesRequest wrapper for the ListSecurityPolicyEntryStates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyEntryStates.go.html to see an example of how to use ListSecurityPolicyEntryStatesRequest.
type ListSecurityPolicyEntryStatesRequest struct {

	// The OCID of the security policy deployment resource.
	SecurityPolicyDeploymentId *string `mandatory:"true" contributesTo:"path" name:"securityPolicyDeploymentId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the security policy deployment.
	DeploymentStatus ListSecurityPolicyEntryStatesDeploymentStatusEnum `mandatory:"false" contributesTo:"query" name:"deploymentStatus" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified security policy entry OCID.
	SecurityPolicyEntryId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyEntryId"`

	// The type of the security policy deployment.
	SecurityPolicyEntryType SecurityPolicyEntryStateSummaryEntryTypeEnum `mandatory:"false" contributesTo:"query" name:"securityPolicyEntryType" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified target id.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPolicyEntryStatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPolicyEntryStatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPolicyEntryStatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPolicyEntryStatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPolicyEntryStatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityPolicyEntryStatesDeploymentStatusEnum(string(request.DeploymentStatus)); !ok && request.DeploymentStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentStatus: %s. Supported values are: %s.", request.DeploymentStatus, strings.Join(GetListSecurityPolicyEntryStatesDeploymentStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityPolicyEntryStateSummaryEntryTypeEnum(string(request.SecurityPolicyEntryType)); !ok && request.SecurityPolicyEntryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityPolicyEntryType: %s. Supported values are: %s.", request.SecurityPolicyEntryType, strings.Join(GetSecurityPolicyEntryStateSummaryEntryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPolicyEntryStatesResponse wrapper for the ListSecurityPolicyEntryStates operation
type ListSecurityPolicyEntryStatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyEntryStateCollection instances
	SecurityPolicyEntryStateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityPolicyEntryStatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPolicyEntryStatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPolicyEntryStatesDeploymentStatusEnum Enum with underlying type: string
type ListSecurityPolicyEntryStatesDeploymentStatusEnum string

// Set of constants representing the allowable values for ListSecurityPolicyEntryStatesDeploymentStatusEnum
const (
	ListSecurityPolicyEntryStatesDeploymentStatusCreated           ListSecurityPolicyEntryStatesDeploymentStatusEnum = "CREATED"
	ListSecurityPolicyEntryStatesDeploymentStatusModified          ListSecurityPolicyEntryStatesDeploymentStatusEnum = "MODIFIED"
	ListSecurityPolicyEntryStatesDeploymentStatusConflict          ListSecurityPolicyEntryStatesDeploymentStatusEnum = "CONFLICT"
	ListSecurityPolicyEntryStatesDeploymentStatusConnectivityIssue ListSecurityPolicyEntryStatesDeploymentStatusEnum = "CONNECTIVITY_ISSUE"
	ListSecurityPolicyEntryStatesDeploymentStatusUnsupportedSyntax ListSecurityPolicyEntryStatesDeploymentStatusEnum = "UNSUPPORTED_SYNTAX"
	ListSecurityPolicyEntryStatesDeploymentStatusUnknownError      ListSecurityPolicyEntryStatesDeploymentStatusEnum = "UNKNOWN_ERROR"
	ListSecurityPolicyEntryStatesDeploymentStatusUnauthorized      ListSecurityPolicyEntryStatesDeploymentStatusEnum = "UNAUTHORIZED"
	ListSecurityPolicyEntryStatesDeploymentStatusDeleted           ListSecurityPolicyEntryStatesDeploymentStatusEnum = "DELETED"
)

var mappingListSecurityPolicyEntryStatesDeploymentStatusEnum = map[string]ListSecurityPolicyEntryStatesDeploymentStatusEnum{
	"CREATED":            ListSecurityPolicyEntryStatesDeploymentStatusCreated,
	"MODIFIED":           ListSecurityPolicyEntryStatesDeploymentStatusModified,
	"CONFLICT":           ListSecurityPolicyEntryStatesDeploymentStatusConflict,
	"CONNECTIVITY_ISSUE": ListSecurityPolicyEntryStatesDeploymentStatusConnectivityIssue,
	"UNSUPPORTED_SYNTAX": ListSecurityPolicyEntryStatesDeploymentStatusUnsupportedSyntax,
	"UNKNOWN_ERROR":      ListSecurityPolicyEntryStatesDeploymentStatusUnknownError,
	"UNAUTHORIZED":       ListSecurityPolicyEntryStatesDeploymentStatusUnauthorized,
	"DELETED":            ListSecurityPolicyEntryStatesDeploymentStatusDeleted,
}

var mappingListSecurityPolicyEntryStatesDeploymentStatusEnumLowerCase = map[string]ListSecurityPolicyEntryStatesDeploymentStatusEnum{
	"created":            ListSecurityPolicyEntryStatesDeploymentStatusCreated,
	"modified":           ListSecurityPolicyEntryStatesDeploymentStatusModified,
	"conflict":           ListSecurityPolicyEntryStatesDeploymentStatusConflict,
	"connectivity_issue": ListSecurityPolicyEntryStatesDeploymentStatusConnectivityIssue,
	"unsupported_syntax": ListSecurityPolicyEntryStatesDeploymentStatusUnsupportedSyntax,
	"unknown_error":      ListSecurityPolicyEntryStatesDeploymentStatusUnknownError,
	"unauthorized":       ListSecurityPolicyEntryStatesDeploymentStatusUnauthorized,
	"deleted":            ListSecurityPolicyEntryStatesDeploymentStatusDeleted,
}

// GetListSecurityPolicyEntryStatesDeploymentStatusEnumValues Enumerates the set of values for ListSecurityPolicyEntryStatesDeploymentStatusEnum
func GetListSecurityPolicyEntryStatesDeploymentStatusEnumValues() []ListSecurityPolicyEntryStatesDeploymentStatusEnum {
	values := make([]ListSecurityPolicyEntryStatesDeploymentStatusEnum, 0)
	for _, v := range mappingListSecurityPolicyEntryStatesDeploymentStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyEntryStatesDeploymentStatusEnumStringValues Enumerates the set of values in String for ListSecurityPolicyEntryStatesDeploymentStatusEnum
func GetListSecurityPolicyEntryStatesDeploymentStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"MODIFIED",
		"CONFLICT",
		"CONNECTIVITY_ISSUE",
		"UNSUPPORTED_SYNTAX",
		"UNKNOWN_ERROR",
		"UNAUTHORIZED",
		"DELETED",
	}
}

// GetMappingListSecurityPolicyEntryStatesDeploymentStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyEntryStatesDeploymentStatusEnum(val string) (ListSecurityPolicyEntryStatesDeploymentStatusEnum, bool) {
	enum, ok := mappingListSecurityPolicyEntryStatesDeploymentStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
