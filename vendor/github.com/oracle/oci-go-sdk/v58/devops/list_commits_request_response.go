// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListCommitsRequest wrapper for the ListCommits operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListCommits.go.html to see an example of how to use ListCommitsRequest.
type ListCommitsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// A filter to return only resources that match the given reference name.
	RefName *string `mandatory:"false" contributesTo:"query" name:"refName"`

	// A filter to exclude commits that match the given reference name.
	ExcludeRefName *string `mandatory:"false" contributesTo:"query" name:"excludeRefName"`

	// A filter to return only commits that affect any of the specified paths.
	FilePath *string `mandatory:"false" contributesTo:"query" name:"filePath"`

	// A filter to return commits only created after the specified timestamp value.
	TimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampGreaterThanOrEqualTo"`

	// A filter to return commits only created before the specified timestamp value.
	TimestampLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampLessThanOrEqualTo"`

	// A filter to return any commits that contains the given message.
	CommitMessage *string `mandatory:"false" contributesTo:"query" name:"commitMessage"`

	// A filter to return any commits that are pushed by the requested author.
	AuthorName *string `mandatory:"false" contributesTo:"query" name:"authorName"`

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

func (request ListCommitsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCommitsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCommitsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCommitsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCommitsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCommitsResponse wrapper for the ListCommits operation
type ListCommitsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryCommitCollection instances
	RepositoryCommitCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCommitsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCommitsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
