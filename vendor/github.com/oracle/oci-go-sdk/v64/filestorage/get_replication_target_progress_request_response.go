// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"net/http"
	"strings"
)

// GetReplicationTargetProgressRequest wrapper for the GetReplicationTargetProgress operation
type GetReplicationTargetProgressRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication target.
	ReplicationTargetId *string `mandatory:"true" contributesTo:"query" name:"replicationTargetId"`

	// The `progressPercentage` of the associated replication source job.
	SourceProgressPercentage *int `mandatory:"true" contributesTo:"query" name:"sourceProgressPercentage"`

	// The `objectNum` of the associated replication.
	ReplicationNum *string `mandatory:"false" contributesTo:"query" name:"replicationNum"`

	// The `objectNum` of the end point of the snapshot during replication operations.
	NewSnapshotNum *string `mandatory:"false" contributesTo:"query" name:"newSnapshotNum"`

	// The `objectNum` of the start point of the snapshot during replication operations.
	LastSnapshotNum *string `mandatory:"false" contributesTo:"query" name:"lastSnapshotNum"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetReplicationTargetProgressRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetReplicationTargetProgressRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetReplicationTargetProgressRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetReplicationTargetProgressRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetReplicationTargetProgressRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetReplicationTargetProgressResponse wrapper for the GetReplicationTargetProgress operation
type GetReplicationTargetProgressResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ReplicationTargetProgress instance
	ReplicationTargetProgress `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetReplicationTargetProgressResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetReplicationTargetProgressResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
