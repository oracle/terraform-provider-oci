// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"net/http"
	"strings"
)

// ReplicationStatusUpdateRequest wrapper for the ReplicationStatusUpdate operation
type ReplicationStatusUpdateRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	ReplicationId *string `mandatory:"true" contributesTo:"query" name:"replicationId"`

	// The `deltaStatus` of the snapshot during replication operations.
	DeltaStatus ReplicationStatusUpdateDeltaStatusEnum `mandatory:"false" contributesTo:"query" name:"deltaStatus" omitEmpty:"true"`

	// The `deltaState` of the snapshot in-transit.
	DeltaState ReplicationStatusUpdateDeltaStateEnum `mandatory:"false" contributesTo:"query" name:"deltaState" omitEmpty:"true"`

	// The `objectNum` of the associated replicationTarget.
	ReplicationTargetNum *string `mandatory:"false" contributesTo:"query" name:"replicationTargetNum"`

	// The `objectNum` of the start point of the snapshot during replication operations.
	LastSnapshotNum *string `mandatory:"false" contributesTo:"query" name:"lastSnapshotNum"`

	// The `objectNum` of the end point of the snapshot during replication operations.
	NewSnapshotNum *string `mandatory:"false" contributesTo:"query" name:"newSnapshotNum"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ReplicationStatusUpdateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ReplicationStatusUpdateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ReplicationStatusUpdateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ReplicationStatusUpdateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ReplicationStatusUpdateRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationStatusUpdateDeltaStatusEnum(string(request.DeltaStatus)); !ok && request.DeltaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaStatus: %s. Supported values are: %s.", request.DeltaStatus, strings.Join(GetReplicationStatusUpdateDeltaStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReplicationStatusUpdateDeltaStateEnum(string(request.DeltaState)); !ok && request.DeltaState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaState: %s. Supported values are: %s.", request.DeltaState, strings.Join(GetReplicationStatusUpdateDeltaStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationStatusUpdateResponse wrapper for the ReplicationStatusUpdate operation
type ReplicationStatusUpdateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Replication instance
	Replication `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ReplicationStatusUpdateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ReplicationStatusUpdateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ReplicationStatusUpdateDeltaStatusEnum Enum with underlying type: string
type ReplicationStatusUpdateDeltaStatusEnum string

// Set of constants representing the allowable values for ReplicationStatusUpdateDeltaStatusEnum
const (
	ReplicationStatusUpdateDeltaStatusIdle         ReplicationStatusUpdateDeltaStatusEnum = "IDLE"
	ReplicationStatusUpdateDeltaStatusCapturing    ReplicationStatusUpdateDeltaStatusEnum = "CAPTURING"
	ReplicationStatusUpdateDeltaStatusApplying     ReplicationStatusUpdateDeltaStatusEnum = "APPLYING"
	ReplicationStatusUpdateDeltaStatusServiceError ReplicationStatusUpdateDeltaStatusEnum = "SERVICE_ERROR"
	ReplicationStatusUpdateDeltaStatusUserError    ReplicationStatusUpdateDeltaStatusEnum = "USER_ERROR"
	ReplicationStatusUpdateDeltaStatusFailed       ReplicationStatusUpdateDeltaStatusEnum = "FAILED"
)

var mappingReplicationStatusUpdateDeltaStatusEnum = map[string]ReplicationStatusUpdateDeltaStatusEnum{
	"IDLE":          ReplicationStatusUpdateDeltaStatusIdle,
	"CAPTURING":     ReplicationStatusUpdateDeltaStatusCapturing,
	"APPLYING":      ReplicationStatusUpdateDeltaStatusApplying,
	"SERVICE_ERROR": ReplicationStatusUpdateDeltaStatusServiceError,
	"USER_ERROR":    ReplicationStatusUpdateDeltaStatusUserError,
	"FAILED":        ReplicationStatusUpdateDeltaStatusFailed,
}

var mappingReplicationStatusUpdateDeltaStatusEnumLowerCase = map[string]ReplicationStatusUpdateDeltaStatusEnum{
	"idle":          ReplicationStatusUpdateDeltaStatusIdle,
	"capturing":     ReplicationStatusUpdateDeltaStatusCapturing,
	"applying":      ReplicationStatusUpdateDeltaStatusApplying,
	"service_error": ReplicationStatusUpdateDeltaStatusServiceError,
	"user_error":    ReplicationStatusUpdateDeltaStatusUserError,
	"failed":        ReplicationStatusUpdateDeltaStatusFailed,
}

// GetReplicationStatusUpdateDeltaStatusEnumValues Enumerates the set of values for ReplicationStatusUpdateDeltaStatusEnum
func GetReplicationStatusUpdateDeltaStatusEnumValues() []ReplicationStatusUpdateDeltaStatusEnum {
	values := make([]ReplicationStatusUpdateDeltaStatusEnum, 0)
	for _, v := range mappingReplicationStatusUpdateDeltaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationStatusUpdateDeltaStatusEnumStringValues Enumerates the set of values in String for ReplicationStatusUpdateDeltaStatusEnum
func GetReplicationStatusUpdateDeltaStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"CAPTURING",
		"APPLYING",
		"SERVICE_ERROR",
		"USER_ERROR",
		"FAILED",
	}
}

// GetMappingReplicationStatusUpdateDeltaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationStatusUpdateDeltaStatusEnum(val string) (ReplicationStatusUpdateDeltaStatusEnum, bool) {
	enum, ok := mappingReplicationStatusUpdateDeltaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReplicationStatusUpdateDeltaStateEnum Enum with underlying type: string
type ReplicationStatusUpdateDeltaStateEnum string

// Set of constants representing the allowable values for ReplicationStatusUpdateDeltaStateEnum
const (
	ReplicationStatusUpdateDeltaStateReadyToReplicate     ReplicationStatusUpdateDeltaStateEnum = "READY_TO_REPLICATE"
	ReplicationStatusUpdateDeltaStateReplicating          ReplicationStatusUpdateDeltaStateEnum = "REPLICATING"
	ReplicationStatusUpdateDeltaStateReplicated           ReplicationStatusUpdateDeltaStateEnum = "REPLICATED"
	ReplicationStatusUpdateDeltaStateReplicatingFailed    ReplicationStatusUpdateDeltaStateEnum = "REPLICATING_FAILED"
	ReplicationStatusUpdateDeltaStateAbortReplication     ReplicationStatusUpdateDeltaStateEnum = "ABORT_REPLICATION"
	ReplicationStatusUpdateDeltaStateAbortReplicationDone ReplicationStatusUpdateDeltaStateEnum = "ABORT_REPLICATION_DONE"
	ReplicationStatusUpdateDeltaStateDone                 ReplicationStatusUpdateDeltaStateEnum = "DONE"
	ReplicationStatusUpdateDeltaStateReadyToGc            ReplicationStatusUpdateDeltaStateEnum = "READY_TO_GC"
	ReplicationStatusUpdateDeltaStateDeleted              ReplicationStatusUpdateDeltaStateEnum = "DELETED"
)

var mappingReplicationStatusUpdateDeltaStateEnum = map[string]ReplicationStatusUpdateDeltaStateEnum{
	"READY_TO_REPLICATE":     ReplicationStatusUpdateDeltaStateReadyToReplicate,
	"REPLICATING":            ReplicationStatusUpdateDeltaStateReplicating,
	"REPLICATED":             ReplicationStatusUpdateDeltaStateReplicated,
	"REPLICATING_FAILED":     ReplicationStatusUpdateDeltaStateReplicatingFailed,
	"ABORT_REPLICATION":      ReplicationStatusUpdateDeltaStateAbortReplication,
	"ABORT_REPLICATION_DONE": ReplicationStatusUpdateDeltaStateAbortReplicationDone,
	"DONE":                   ReplicationStatusUpdateDeltaStateDone,
	"READY_TO_GC":            ReplicationStatusUpdateDeltaStateReadyToGc,
	"DELETED":                ReplicationStatusUpdateDeltaStateDeleted,
}

var mappingReplicationStatusUpdateDeltaStateEnumLowerCase = map[string]ReplicationStatusUpdateDeltaStateEnum{
	"ready_to_replicate":     ReplicationStatusUpdateDeltaStateReadyToReplicate,
	"replicating":            ReplicationStatusUpdateDeltaStateReplicating,
	"replicated":             ReplicationStatusUpdateDeltaStateReplicated,
	"replicating_failed":     ReplicationStatusUpdateDeltaStateReplicatingFailed,
	"abort_replication":      ReplicationStatusUpdateDeltaStateAbortReplication,
	"abort_replication_done": ReplicationStatusUpdateDeltaStateAbortReplicationDone,
	"done":                   ReplicationStatusUpdateDeltaStateDone,
	"ready_to_gc":            ReplicationStatusUpdateDeltaStateReadyToGc,
	"deleted":                ReplicationStatusUpdateDeltaStateDeleted,
}

// GetReplicationStatusUpdateDeltaStateEnumValues Enumerates the set of values for ReplicationStatusUpdateDeltaStateEnum
func GetReplicationStatusUpdateDeltaStateEnumValues() []ReplicationStatusUpdateDeltaStateEnum {
	values := make([]ReplicationStatusUpdateDeltaStateEnum, 0)
	for _, v := range mappingReplicationStatusUpdateDeltaStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationStatusUpdateDeltaStateEnumStringValues Enumerates the set of values in String for ReplicationStatusUpdateDeltaStateEnum
func GetReplicationStatusUpdateDeltaStateEnumStringValues() []string {
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

// GetMappingReplicationStatusUpdateDeltaStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationStatusUpdateDeltaStateEnum(val string) (ReplicationStatusUpdateDeltaStateEnum, bool) {
	enum, ok := mappingReplicationStatusUpdateDeltaStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}