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

// SyncRepositoryRequest wrapper for the SyncRepository operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/SyncRepository.go.html to see an example of how to use SyncRepositoryRequest.
type SyncRepositoryRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// If fetch and merge is passed in, the changes from the upstream will be fetched and merged into the destination branch.
	// If discard is passed in, the changes in the fork will be overwritten with the changes brought in from the upstream.
	SyncMergeStrategy SyncRepositorySyncMergeStrategyEnum `mandatory:"true" contributesTo:"query" name:"syncMergeStrategy" omitEmpty:"true"`

	// Details required for syncing a repository with its upstream.
	SyncRepositoryDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the etag from a previous GET or POST response for that resource. The resource will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated earlier due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SyncRepositoryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SyncRepositoryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SyncRepositoryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SyncRepositoryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SyncRepositoryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSyncRepositorySyncMergeStrategyEnum(string(request.SyncMergeStrategy)); !ok && request.SyncMergeStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyncMergeStrategy: %s. Supported values are: %s.", request.SyncMergeStrategy, strings.Join(GetSyncRepositorySyncMergeStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SyncRepositoryResponse wrapper for the SyncRepository operation
type SyncRepositoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SyncRepositoryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SyncRepositoryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SyncRepositorySyncMergeStrategyEnum Enum with underlying type: string
type SyncRepositorySyncMergeStrategyEnum string

// Set of constants representing the allowable values for SyncRepositorySyncMergeStrategyEnum
const (
	SyncRepositorySyncMergeStrategyFetchAndMerge SyncRepositorySyncMergeStrategyEnum = "FETCH_AND_MERGE"
	SyncRepositorySyncMergeStrategyDiscard       SyncRepositorySyncMergeStrategyEnum = "DISCARD"
)

var mappingSyncRepositorySyncMergeStrategyEnum = map[string]SyncRepositorySyncMergeStrategyEnum{
	"FETCH_AND_MERGE": SyncRepositorySyncMergeStrategyFetchAndMerge,
	"DISCARD":         SyncRepositorySyncMergeStrategyDiscard,
}

var mappingSyncRepositorySyncMergeStrategyEnumLowerCase = map[string]SyncRepositorySyncMergeStrategyEnum{
	"fetch_and_merge": SyncRepositorySyncMergeStrategyFetchAndMerge,
	"discard":         SyncRepositorySyncMergeStrategyDiscard,
}

// GetSyncRepositorySyncMergeStrategyEnumValues Enumerates the set of values for SyncRepositorySyncMergeStrategyEnum
func GetSyncRepositorySyncMergeStrategyEnumValues() []SyncRepositorySyncMergeStrategyEnum {
	values := make([]SyncRepositorySyncMergeStrategyEnum, 0)
	for _, v := range mappingSyncRepositorySyncMergeStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetSyncRepositorySyncMergeStrategyEnumStringValues Enumerates the set of values in String for SyncRepositorySyncMergeStrategyEnum
func GetSyncRepositorySyncMergeStrategyEnumStringValues() []string {
	return []string{
		"FETCH_AND_MERGE",
		"DISCARD",
	}
}

// GetMappingSyncRepositorySyncMergeStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSyncRepositorySyncMergeStrategyEnum(val string) (SyncRepositorySyncMergeStrategyEnum, bool) {
	enum, ok := mappingSyncRepositorySyncMergeStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
