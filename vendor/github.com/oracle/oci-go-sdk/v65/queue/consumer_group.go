// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConsumerGroup A detailed representation of a consumer group.
type ConsumerGroup struct {

	// A unique identifier for the consumer group that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the consumer group. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time that the consumer group was created, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the consumer group was updated, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the consumer group.
	LifecycleState ConsumerGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The filter used by the consumer group. Only messages matching the filter will be available by consumers of the group.
	Filter *string `mandatory:"true" json:"filter"`

	// The OCID of the associated queue.
	QueueId *string `mandatory:"true" json:"queueId"`

	// Any additional details about the current state of the consumer group.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue.
	// A value of 0 indicates that the DLQ is not used.
	// If the value isn't set, it will be using the value defined at the queue level.
	DeadLetterQueueDeliveryCount *int `mandatory:"false" json:"deadLetterQueueDeliveryCount"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ConsumerGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsumerGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConsumerGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConsumerGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConsumerGroupLifecycleStateEnum Enum with underlying type: string
type ConsumerGroupLifecycleStateEnum string

// Set of constants representing the allowable values for ConsumerGroupLifecycleStateEnum
const (
	ConsumerGroupLifecycleStateCreating ConsumerGroupLifecycleStateEnum = "CREATING"
	ConsumerGroupLifecycleStateUpdating ConsumerGroupLifecycleStateEnum = "UPDATING"
	ConsumerGroupLifecycleStateActive   ConsumerGroupLifecycleStateEnum = "ACTIVE"
	ConsumerGroupLifecycleStateDeleting ConsumerGroupLifecycleStateEnum = "DELETING"
	ConsumerGroupLifecycleStateDeleted  ConsumerGroupLifecycleStateEnum = "DELETED"
	ConsumerGroupLifecycleStateFailed   ConsumerGroupLifecycleStateEnum = "FAILED"
	ConsumerGroupLifecycleStateInactive ConsumerGroupLifecycleStateEnum = "INACTIVE"
)

var mappingConsumerGroupLifecycleStateEnum = map[string]ConsumerGroupLifecycleStateEnum{
	"CREATING": ConsumerGroupLifecycleStateCreating,
	"UPDATING": ConsumerGroupLifecycleStateUpdating,
	"ACTIVE":   ConsumerGroupLifecycleStateActive,
	"DELETING": ConsumerGroupLifecycleStateDeleting,
	"DELETED":  ConsumerGroupLifecycleStateDeleted,
	"FAILED":   ConsumerGroupLifecycleStateFailed,
	"INACTIVE": ConsumerGroupLifecycleStateInactive,
}

var mappingConsumerGroupLifecycleStateEnumLowerCase = map[string]ConsumerGroupLifecycleStateEnum{
	"creating": ConsumerGroupLifecycleStateCreating,
	"updating": ConsumerGroupLifecycleStateUpdating,
	"active":   ConsumerGroupLifecycleStateActive,
	"deleting": ConsumerGroupLifecycleStateDeleting,
	"deleted":  ConsumerGroupLifecycleStateDeleted,
	"failed":   ConsumerGroupLifecycleStateFailed,
	"inactive": ConsumerGroupLifecycleStateInactive,
}

// GetConsumerGroupLifecycleStateEnumValues Enumerates the set of values for ConsumerGroupLifecycleStateEnum
func GetConsumerGroupLifecycleStateEnumValues() []ConsumerGroupLifecycleStateEnum {
	values := make([]ConsumerGroupLifecycleStateEnum, 0)
	for _, v := range mappingConsumerGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConsumerGroupLifecycleStateEnumStringValues Enumerates the set of values in String for ConsumerGroupLifecycleStateEnum
func GetConsumerGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingConsumerGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsumerGroupLifecycleStateEnum(val string) (ConsumerGroupLifecycleStateEnum, bool) {
	enum, ok := mappingConsumerGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
