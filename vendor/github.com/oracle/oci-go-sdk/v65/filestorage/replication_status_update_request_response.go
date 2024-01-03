// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
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
	DeltaState *string `mandatory:"false" contributesTo:"query" name:"deltaState"`

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
	ReplicationStatusUpdateDeltaStatusTransferring ReplicationStatusUpdateDeltaStatusEnum = "TRANSFERRING"
)

var mappingReplicationStatusUpdateDeltaStatusEnum = map[string]ReplicationStatusUpdateDeltaStatusEnum{
	"IDLE":          ReplicationStatusUpdateDeltaStatusIdle,
	"CAPTURING":     ReplicationStatusUpdateDeltaStatusCapturing,
	"APPLYING":      ReplicationStatusUpdateDeltaStatusApplying,
	"SERVICE_ERROR": ReplicationStatusUpdateDeltaStatusServiceError,
	"USER_ERROR":    ReplicationStatusUpdateDeltaStatusUserError,
	"FAILED":        ReplicationStatusUpdateDeltaStatusFailed,
	"TRANSFERRING":  ReplicationStatusUpdateDeltaStatusTransferring,
}

var mappingReplicationStatusUpdateDeltaStatusEnumLowerCase = map[string]ReplicationStatusUpdateDeltaStatusEnum{
	"idle":          ReplicationStatusUpdateDeltaStatusIdle,
	"capturing":     ReplicationStatusUpdateDeltaStatusCapturing,
	"applying":      ReplicationStatusUpdateDeltaStatusApplying,
	"service_error": ReplicationStatusUpdateDeltaStatusServiceError,
	"user_error":    ReplicationStatusUpdateDeltaStatusUserError,
	"failed":        ReplicationStatusUpdateDeltaStatusFailed,
	"transferring":  ReplicationStatusUpdateDeltaStatusTransferring,
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
		"TRANSFERRING",
	}
}

// GetMappingReplicationStatusUpdateDeltaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationStatusUpdateDeltaStatusEnum(val string) (ReplicationStatusUpdateDeltaStatusEnum, bool) {
	enum, ok := mappingReplicationStatusUpdateDeltaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
