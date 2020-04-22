// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StreamPoolSummary The summary representation of a stream pool.
type StreamPoolSummary struct {

	// The OCID of the stream pool.
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID that the pool belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the stream pool.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the stream pool.
	LifecycleState StreamPoolSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the stream pool was created, expressed in in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// True if the stream pool is private, false otherwise.
	// The associated endpoint and subnetId of a private stream pool can be retrieved through the GetStreamPool API.
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StreamPoolSummary) String() string {
	return common.PointerString(m)
}

// StreamPoolSummaryLifecycleStateEnum Enum with underlying type: string
type StreamPoolSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for StreamPoolSummaryLifecycleStateEnum
const (
	StreamPoolSummaryLifecycleStateCreating StreamPoolSummaryLifecycleStateEnum = "CREATING"
	StreamPoolSummaryLifecycleStateActive   StreamPoolSummaryLifecycleStateEnum = "ACTIVE"
	StreamPoolSummaryLifecycleStateDeleting StreamPoolSummaryLifecycleStateEnum = "DELETING"
	StreamPoolSummaryLifecycleStateDeleted  StreamPoolSummaryLifecycleStateEnum = "DELETED"
	StreamPoolSummaryLifecycleStateFailed   StreamPoolSummaryLifecycleStateEnum = "FAILED"
	StreamPoolSummaryLifecycleStateUpdating StreamPoolSummaryLifecycleStateEnum = "UPDATING"
)

var mappingStreamPoolSummaryLifecycleState = map[string]StreamPoolSummaryLifecycleStateEnum{
	"CREATING": StreamPoolSummaryLifecycleStateCreating,
	"ACTIVE":   StreamPoolSummaryLifecycleStateActive,
	"DELETING": StreamPoolSummaryLifecycleStateDeleting,
	"DELETED":  StreamPoolSummaryLifecycleStateDeleted,
	"FAILED":   StreamPoolSummaryLifecycleStateFailed,
	"UPDATING": StreamPoolSummaryLifecycleStateUpdating,
}

// GetStreamPoolSummaryLifecycleStateEnumValues Enumerates the set of values for StreamPoolSummaryLifecycleStateEnum
func GetStreamPoolSummaryLifecycleStateEnumValues() []StreamPoolSummaryLifecycleStateEnum {
	values := make([]StreamPoolSummaryLifecycleStateEnum, 0)
	for _, v := range mappingStreamPoolSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
