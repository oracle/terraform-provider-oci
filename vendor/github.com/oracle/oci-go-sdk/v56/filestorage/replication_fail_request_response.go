// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
	"strings"
)

// ReplicationFailRequest wrapper for the ReplicationFail operation
type ReplicationFailRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	ReplicationId *string `mandatory:"true" contributesTo:"query" name:"replicationId"`

	// The `objectNum` of the associated replicationTarget.
	ReplicationTargetNum *int `mandatory:"false" contributesTo:"query" name:"replicationTargetNum"`

	// The `deltaState` of the snapshot in-transit.
	DeltaState ReplicationFailDeltaStateEnum `mandatory:"false" contributesTo:"query" name:"deltaState" omitEmpty:"true"`

	// The flag to represent if it is a replication failover message.
	IsFailover *bool `mandatory:"false" contributesTo:"query" name:"isFailover"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ReplicationFailRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ReplicationFailRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ReplicationFailRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ReplicationFailRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ReplicationFailRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingReplicationFailDeltaStateEnum[string(request.DeltaState)]; !ok && request.DeltaState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaState: %s. Supported values are: %s.", request.DeltaState, strings.Join(GetReplicationFailDeltaStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationFailResponse wrapper for the ReplicationFail operation
type ReplicationFailResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Replication instance
	Replication `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ReplicationFailResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ReplicationFailResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ReplicationFailDeltaStateEnum Enum with underlying type: string
type ReplicationFailDeltaStateEnum string

// Set of constants representing the allowable values for ReplicationFailDeltaStateEnum
const (
	ReplicationFailDeltaStateReadyToReplicate     ReplicationFailDeltaStateEnum = "READY_TO_REPLICATE"
	ReplicationFailDeltaStateReplicating          ReplicationFailDeltaStateEnum = "REPLICATING"
	ReplicationFailDeltaStateReplicated           ReplicationFailDeltaStateEnum = "REPLICATED"
	ReplicationFailDeltaStateReplicatingFailed    ReplicationFailDeltaStateEnum = "REPLICATING_FAILED"
	ReplicationFailDeltaStateAbortReplication     ReplicationFailDeltaStateEnum = "ABORT_REPLICATION"
	ReplicationFailDeltaStateAbortReplicationDone ReplicationFailDeltaStateEnum = "ABORT_REPLICATION_DONE"
	ReplicationFailDeltaStateDone                 ReplicationFailDeltaStateEnum = "DONE"
	ReplicationFailDeltaStateReadyToGc            ReplicationFailDeltaStateEnum = "READY_TO_GC"
	ReplicationFailDeltaStateDeleted              ReplicationFailDeltaStateEnum = "DELETED"
)

var mappingReplicationFailDeltaStateEnum = map[string]ReplicationFailDeltaStateEnum{
	"READY_TO_REPLICATE":     ReplicationFailDeltaStateReadyToReplicate,
	"REPLICATING":            ReplicationFailDeltaStateReplicating,
	"REPLICATED":             ReplicationFailDeltaStateReplicated,
	"REPLICATING_FAILED":     ReplicationFailDeltaStateReplicatingFailed,
	"ABORT_REPLICATION":      ReplicationFailDeltaStateAbortReplication,
	"ABORT_REPLICATION_DONE": ReplicationFailDeltaStateAbortReplicationDone,
	"DONE":                   ReplicationFailDeltaStateDone,
	"READY_TO_GC":            ReplicationFailDeltaStateReadyToGc,
	"DELETED":                ReplicationFailDeltaStateDeleted,
}

// GetReplicationFailDeltaStateEnumValues Enumerates the set of values for ReplicationFailDeltaStateEnum
func GetReplicationFailDeltaStateEnumValues() []ReplicationFailDeltaStateEnum {
	values := make([]ReplicationFailDeltaStateEnum, 0)
	for _, v := range mappingReplicationFailDeltaStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationFailDeltaStateEnumStringValues Enumerates the set of values in String for ReplicationFailDeltaStateEnum
func GetReplicationFailDeltaStateEnumStringValues() []string {
	return []string{
		"READY_TO_REPLICATE",
		"REPLICATING",
		"REPLICATED",
		"REPLICATING_FAILED",
		"ABORT_REPLICATION",
		"ABORT_REPLICATION_DONE",
		"DONE",
		"READY_TO_GC",
		"DELETED",
	}
}
