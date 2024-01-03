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

// ListCodeSearchResultsRequest wrapper for the ListCodeSearchResults operation
type ListCodeSearchResultsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Search query provided by the user as per the search query syntax.
	//   Example:
	//     Workflow AdminServiceApiConfiguration retryConfig  -  Search for files with occurrence of all of these keywords.
	//     hello AND world  -  Search for files that has both 'hello' and 'world'.
	//     hello OR world  -  Search for  files that has 'hello' or 'world' or both.
	//     hello NOT world  -  Search for files that has 'hello' but not 'world'.
	//     "hello world"  -  Search for files that has words 'hello' and 'world' in same order.
	//     project:project1 repo:repo1 wfaas  -  Search in repository 'repo1' in project 'project1' for keyword 'wfaas'.
	//     hello path:readme.md  -  Search for files that contain word 'hello' and the file path matches 'readme.md'.
	//     hello ext:c  -  Search for files that has 'hello' within files with the '.c' extension.
	Query *string `mandatory:"true" contributesTo:"query" name:"query"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Value of this is always "ACCESSIBLE" and any other value is not supported.
	// When set to any other value, search will return no results.
	AccessLevel *string `mandatory:"false" contributesTo:"query" name:"accessLevel"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCodeSearchResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCodeSearchResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCodeSearchResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCodeSearchResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCodeSearchResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCodeSearchResultsResponse wrapper for the ListCodeSearchResults operation
type ListCodeSearchResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CodeSearchResultCollection instances
	CodeSearchResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCodeSearchResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCodeSearchResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
