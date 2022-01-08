// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"io"
	"net/http"
	"strings"
)

// ReplicationTargetStatusUpdateRequest wrapper for the ReplicationTargetStatusUpdate operation
type ReplicationTargetStatusUpdateRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication target.
	ReplicationTargetId *string `mandatory:"true" contributesTo:"query" name:"replicationTargetId"`

	// The `deltaStatus` of the snapshot during replication operations.
	DeltaStatus ReplicationTargetStatusUpdateDeltaStatusEnum `mandatory:"false" contributesTo:"query" name:"deltaStatus" omitEmpty:"true"`

	// The `deltaState` of the snapshot in-transit.
	DeltaState ReplicationTargetStatusUpdateDeltaStateEnum `mandatory:"false" contributesTo:"query" name:"deltaState" omitEmpty:"true"`

	// The `objectNum` of the associated replication.
	ReplicationNum *int `mandatory:"false" contributesTo:"query" name:"replicationNum"`

	// The `objectNum` of the start point of the snapshot during replication operations.
	LastSnapshotNum *int `mandatory:"false" contributesTo:"query" name:"lastSnapshotNum"`

	// The `objectNum` of the end point of the snapshot during replication operations.
	NewSnapshotNum *int `mandatory:"false" contributesTo:"query" name:"newSnapshotNum"`

	// The `snapshotTime` of the most recent recoverable replication snapshot.
	RecoveryPointTime *int `mandatory:"false" contributesTo:"query" name:"recoveryPointTime"`

	// The kmsKeyOcid for the Snapshot in-flight.
	KmsKeyOcid *string `mandatory:"false" contributesTo:"query" name:"kmsKeyOcid"`

	// The total number of bytes in the Snapshot in-flight.
	DeltaByteCount *int `mandatory:"false" contributesTo:"query" name:"deltaByteCount"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The cipherText of the ReplicationDeltaTransferKey for the Snapshot in-flight.
	CipherTextDetails io.ReadCloser `mandatory:"false" contributesTo:"body" encoding:"binary"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ReplicationTargetStatusUpdateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ReplicationTargetStatusUpdateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request ReplicationTargetStatusUpdateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.CipherTextDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ReplicationTargetStatusUpdateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ReplicationTargetStatusUpdateRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingReplicationTargetStatusUpdateDeltaStatusEnum[string(request.DeltaStatus)]; !ok && request.DeltaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaStatus: %s. Supported values are: %s.", request.DeltaStatus, strings.Join(GetReplicationTargetStatusUpdateDeltaStatusEnumStringValues(), ",")))
	}
	if _, ok := mappingReplicationTargetStatusUpdateDeltaStateEnum[string(request.DeltaState)]; !ok && request.DeltaState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaState: %s. Supported values are: %s.", request.DeltaState, strings.Join(GetReplicationTargetStatusUpdateDeltaStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationTargetStatusUpdateResponse wrapper for the ReplicationTargetStatusUpdate operation
type ReplicationTargetStatusUpdateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ReplicationTarget instance
	ReplicationTarget `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ReplicationTargetStatusUpdateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ReplicationTargetStatusUpdateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ReplicationTargetStatusUpdateDeltaStatusEnum Enum with underlying type: string
type ReplicationTargetStatusUpdateDeltaStatusEnum string

// Set of constants representing the allowable values for ReplicationTargetStatusUpdateDeltaStatusEnum
const (
	ReplicationTargetStatusUpdateDeltaStatusIdle         ReplicationTargetStatusUpdateDeltaStatusEnum = "IDLE"
	ReplicationTargetStatusUpdateDeltaStatusCapturing    ReplicationTargetStatusUpdateDeltaStatusEnum = "CAPTURING"
	ReplicationTargetStatusUpdateDeltaStatusApplying     ReplicationTargetStatusUpdateDeltaStatusEnum = "APPLYING"
	ReplicationTargetStatusUpdateDeltaStatusServiceError ReplicationTargetStatusUpdateDeltaStatusEnum = "SERVICE_ERROR"
	ReplicationTargetStatusUpdateDeltaStatusUserError    ReplicationTargetStatusUpdateDeltaStatusEnum = "USER_ERROR"
	ReplicationTargetStatusUpdateDeltaStatusFailed       ReplicationTargetStatusUpdateDeltaStatusEnum = "FAILED"
)

var mappingReplicationTargetStatusUpdateDeltaStatusEnum = map[string]ReplicationTargetStatusUpdateDeltaStatusEnum{
	"IDLE":          ReplicationTargetStatusUpdateDeltaStatusIdle,
	"CAPTURING":     ReplicationTargetStatusUpdateDeltaStatusCapturing,
	"APPLYING":      ReplicationTargetStatusUpdateDeltaStatusApplying,
	"SERVICE_ERROR": ReplicationTargetStatusUpdateDeltaStatusServiceError,
	"USER_ERROR":    ReplicationTargetStatusUpdateDeltaStatusUserError,
	"FAILED":        ReplicationTargetStatusUpdateDeltaStatusFailed,
}

// GetReplicationTargetStatusUpdateDeltaStatusEnumValues Enumerates the set of values for ReplicationTargetStatusUpdateDeltaStatusEnum
func GetReplicationTargetStatusUpdateDeltaStatusEnumValues() []ReplicationTargetStatusUpdateDeltaStatusEnum {
	values := make([]ReplicationTargetStatusUpdateDeltaStatusEnum, 0)
	for _, v := range mappingReplicationTargetStatusUpdateDeltaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationTargetStatusUpdateDeltaStatusEnumStringValues Enumerates the set of values in String for ReplicationTargetStatusUpdateDeltaStatusEnum
func GetReplicationTargetStatusUpdateDeltaStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"CAPTURING",
		"APPLYING",
		"SERVICE_ERROR",
		"USER_ERROR",
		"FAILED",
	}
}

// ReplicationTargetStatusUpdateDeltaStateEnum Enum with underlying type: string
type ReplicationTargetStatusUpdateDeltaStateEnum string

// Set of constants representing the allowable values for ReplicationTargetStatusUpdateDeltaStateEnum
const (
	ReplicationTargetStatusUpdateDeltaStateReadyToReplicate     ReplicationTargetStatusUpdateDeltaStateEnum = "READY_TO_REPLICATE"
	ReplicationTargetStatusUpdateDeltaStateReplicating          ReplicationTargetStatusUpdateDeltaStateEnum = "REPLICATING"
	ReplicationTargetStatusUpdateDeltaStateReplicated           ReplicationTargetStatusUpdateDeltaStateEnum = "REPLICATED"
	ReplicationTargetStatusUpdateDeltaStateReplicatingFailed    ReplicationTargetStatusUpdateDeltaStateEnum = "REPLICATING_FAILED"
	ReplicationTargetStatusUpdateDeltaStateAbortReplication     ReplicationTargetStatusUpdateDeltaStateEnum = "ABORT_REPLICATION"
	ReplicationTargetStatusUpdateDeltaStateAbortReplicationDone ReplicationTargetStatusUpdateDeltaStateEnum = "ABORT_REPLICATION_DONE"
	ReplicationTargetStatusUpdateDeltaStateDone                 ReplicationTargetStatusUpdateDeltaStateEnum = "DONE"
	ReplicationTargetStatusUpdateDeltaStateReadyToGc            ReplicationTargetStatusUpdateDeltaStateEnum = "READY_TO_GC"
	ReplicationTargetStatusUpdateDeltaStateDeleted              ReplicationTargetStatusUpdateDeltaStateEnum = "DELETED"
)

var mappingReplicationTargetStatusUpdateDeltaStateEnum = map[string]ReplicationTargetStatusUpdateDeltaStateEnum{
	"READY_TO_REPLICATE":     ReplicationTargetStatusUpdateDeltaStateReadyToReplicate,
	"REPLICATING":            ReplicationTargetStatusUpdateDeltaStateReplicating,
	"REPLICATED":             ReplicationTargetStatusUpdateDeltaStateReplicated,
	"REPLICATING_FAILED":     ReplicationTargetStatusUpdateDeltaStateReplicatingFailed,
	"ABORT_REPLICATION":      ReplicationTargetStatusUpdateDeltaStateAbortReplication,
	"ABORT_REPLICATION_DONE": ReplicationTargetStatusUpdateDeltaStateAbortReplicationDone,
	"DONE":                   ReplicationTargetStatusUpdateDeltaStateDone,
	"READY_TO_GC":            ReplicationTargetStatusUpdateDeltaStateReadyToGc,
	"DELETED":                ReplicationTargetStatusUpdateDeltaStateDeleted,
}

// GetReplicationTargetStatusUpdateDeltaStateEnumValues Enumerates the set of values for ReplicationTargetStatusUpdateDeltaStateEnum
func GetReplicationTargetStatusUpdateDeltaStateEnumValues() []ReplicationTargetStatusUpdateDeltaStateEnum {
	values := make([]ReplicationTargetStatusUpdateDeltaStateEnum, 0)
	for _, v := range mappingReplicationTargetStatusUpdateDeltaStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationTargetStatusUpdateDeltaStateEnumStringValues Enumerates the set of values in String for ReplicationTargetStatusUpdateDeltaStateEnum
func GetReplicationTargetStatusUpdateDeltaStateEnumStringValues() []string {
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
