// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamSummary Summary representation of a stream.
type StreamSummary struct {

	// The name of the stream.
	// Example: `TelemetryEvents`
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the stream.
	Id *string `mandatory:"true" json:"id"`

	// The number of partitions in the stream.
	Partitions *int `mandatory:"true" json:"partitions"`

	// The OCID of the compartment that contains the stream.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the stream pool that contains the stream.
	StreamPoolId *string `mandatory:"true" json:"streamPoolId"`

	// The current state of the stream.
	LifecycleState StreamSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the stream was created, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The endpoint to use when creating the StreamClient to consume or publish messages in the stream.
	// If the associated stream pool is private, the endpoint is also private and can only be accessed from inside the stream pool's associated subnet.
	MessagesEndpoint *string `mandatory:"true" json:"messagesEndpoint"`

	// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StreamSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamSummaryLifecycleStateEnum Enum with underlying type: string
type StreamSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for StreamSummaryLifecycleStateEnum
const (
	StreamSummaryLifecycleStateCreating StreamSummaryLifecycleStateEnum = "CREATING"
	StreamSummaryLifecycleStateActive   StreamSummaryLifecycleStateEnum = "ACTIVE"
	StreamSummaryLifecycleStateDeleting StreamSummaryLifecycleStateEnum = "DELETING"
	StreamSummaryLifecycleStateDeleted  StreamSummaryLifecycleStateEnum = "DELETED"
	StreamSummaryLifecycleStateFailed   StreamSummaryLifecycleStateEnum = "FAILED"
	StreamSummaryLifecycleStateUpdating StreamSummaryLifecycleStateEnum = "UPDATING"
)

var mappingStreamSummaryLifecycleStateEnum = map[string]StreamSummaryLifecycleStateEnum{
	"CREATING": StreamSummaryLifecycleStateCreating,
	"ACTIVE":   StreamSummaryLifecycleStateActive,
	"DELETING": StreamSummaryLifecycleStateDeleting,
	"DELETED":  StreamSummaryLifecycleStateDeleted,
	"FAILED":   StreamSummaryLifecycleStateFailed,
	"UPDATING": StreamSummaryLifecycleStateUpdating,
}

var mappingStreamSummaryLifecycleStateEnumLowerCase = map[string]StreamSummaryLifecycleStateEnum{
	"creating": StreamSummaryLifecycleStateCreating,
	"active":   StreamSummaryLifecycleStateActive,
	"deleting": StreamSummaryLifecycleStateDeleting,
	"deleted":  StreamSummaryLifecycleStateDeleted,
	"failed":   StreamSummaryLifecycleStateFailed,
	"updating": StreamSummaryLifecycleStateUpdating,
}

// GetStreamSummaryLifecycleStateEnumValues Enumerates the set of values for StreamSummaryLifecycleStateEnum
func GetStreamSummaryLifecycleStateEnumValues() []StreamSummaryLifecycleStateEnum {
	values := make([]StreamSummaryLifecycleStateEnum, 0)
	for _, v := range mappingStreamSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for StreamSummaryLifecycleStateEnum
func GetStreamSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingStreamSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamSummaryLifecycleStateEnum(val string) (StreamSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingStreamSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
