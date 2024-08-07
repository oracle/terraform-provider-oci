// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCommitDiffsRequest wrapper for the ListCommitDiffs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListCommitDiffs.go.html to see an example of how to use ListCommitDiffsRequest.
type ListCommitDiffsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// The commit or reference name to compare changes against.
	BaseVersion *string `mandatory:"true" contributesTo:"query" name:"baseVersion"`

	// The commit or reference name where changes are coming from.
	TargetVersion *string `mandatory:"true" contributesTo:"query" name:"targetVersion"`

	// The target repository identifier
	TargetRepositoryId *string `mandatory:"false" contributesTo:"query" name:"targetRepositoryId"`

	// Boolean value to indicate whether to use merge base or most recent revision.
	IsComparisonFromMergeBase *bool `mandatory:"false" contributesTo:"query" name:"isComparisonFromMergeBase"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCommitDiffsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCommitDiffsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCommitDiffsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCommitDiffsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCommitDiffsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCommitDiffsResponse wrapper for the ListCommitDiffs operation
type ListCommitDiffsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiffCollection instances
	DiffCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCommitDiffsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCommitDiffsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
