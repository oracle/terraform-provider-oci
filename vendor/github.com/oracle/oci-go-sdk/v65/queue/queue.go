// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// A description of the Queue API
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Queue Description of Queue.
type Queue struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the Queue was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Queue was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Queue.
	LifecycleState QueueLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The endpoint to use to consume or publish messages in the queue.
	MessagesEndpoint *string `mandatory:"true" json:"messagesEndpoint"`

	// The retention period of the messages in the queue, in seconds.
	RetentionInSeconds *int `mandatory:"true" json:"retentionInSeconds"`

	// The default visibility of the messages consumed from the queue.
	VisibilityInSeconds *int `mandatory:"true" json:"visibilityInSeconds"`

	// The default polling timeout of the messages in the queue, in seconds.
	TimeoutInSeconds *int `mandatory:"true" json:"timeoutInSeconds"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
	DeadLetterQueueDeliveryCount *int `mandatory:"true" json:"deadLetterQueueDeliveryCount"`

	// Queue Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Id of the custom master encryption key which will be used to encrypt messages content
	CustomEncryptionKeyId *string `mandatory:"false" json:"customEncryptionKeyId"`

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

func (m Queue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Queue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueueLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetQueueLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueueLifecycleStateEnum Enum with underlying type: string
type QueueLifecycleStateEnum string

// Set of constants representing the allowable values for QueueLifecycleStateEnum
const (
	QueueLifecycleStateCreating QueueLifecycleStateEnum = "CREATING"
	QueueLifecycleStateUpdating QueueLifecycleStateEnum = "UPDATING"
	QueueLifecycleStateActive   QueueLifecycleStateEnum = "ACTIVE"
	QueueLifecycleStateDeleting QueueLifecycleStateEnum = "DELETING"
	QueueLifecycleStateDeleted  QueueLifecycleStateEnum = "DELETED"
	QueueLifecycleStateFailed   QueueLifecycleStateEnum = "FAILED"
)

var mappingQueueLifecycleStateEnum = map[string]QueueLifecycleStateEnum{
	"CREATING": QueueLifecycleStateCreating,
	"UPDATING": QueueLifecycleStateUpdating,
	"ACTIVE":   QueueLifecycleStateActive,
	"DELETING": QueueLifecycleStateDeleting,
	"DELETED":  QueueLifecycleStateDeleted,
	"FAILED":   QueueLifecycleStateFailed,
}

var mappingQueueLifecycleStateEnumLowerCase = map[string]QueueLifecycleStateEnum{
	"creating": QueueLifecycleStateCreating,
	"updating": QueueLifecycleStateUpdating,
	"active":   QueueLifecycleStateActive,
	"deleting": QueueLifecycleStateDeleting,
	"deleted":  QueueLifecycleStateDeleted,
	"failed":   QueueLifecycleStateFailed,
}

// GetQueueLifecycleStateEnumValues Enumerates the set of values for QueueLifecycleStateEnum
func GetQueueLifecycleStateEnumValues() []QueueLifecycleStateEnum {
	values := make([]QueueLifecycleStateEnum, 0)
	for _, v := range mappingQueueLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetQueueLifecycleStateEnumStringValues Enumerates the set of values in String for QueueLifecycleStateEnum
func GetQueueLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingQueueLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueueLifecycleStateEnum(val string) (QueueLifecycleStateEnum, bool) {
	enum, ok := mappingQueueLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
