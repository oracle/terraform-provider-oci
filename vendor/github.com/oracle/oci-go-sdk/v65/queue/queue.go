// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Queue A detailed representation of a queue and its configuration.
type Queue struct {

	// A unique identifier for the queue that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the queue.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time that the queue was created, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the queue was updated, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the queue.
	LifecycleState QueueLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The endpoint to use to consume or publish messages in the queue.
	MessagesEndpoint *string `mandatory:"true" json:"messagesEndpoint"`

	// The retention period of the messages in the queue, in seconds.
	RetentionInSeconds *int `mandatory:"true" json:"retentionInSeconds"`

	// The default visibility timeout of the messages consumed from the queue, in seconds.
	VisibilityInSeconds *int `mandatory:"true" json:"visibilityInSeconds"`

	// The default polling timeout of the messages in the queue, in seconds.
	TimeoutInSeconds *int `mandatory:"true" json:"timeoutInSeconds"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
	DeadLetterQueueDeliveryCount *int `mandatory:"true" json:"deadLetterQueueDeliveryCount"`

	// A user-friendly name for the queue. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Any additional details about the current state of the queue.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom encryption key to be used to encrypt messages content.
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

	// The percentage of allocated queue resources that can be consumed by a single channel. For example, if a queue has a storage limit of 2Gb, and a single channel consumption limit is 0.1 (10%), that means data size of a single channel  can't exceed 200Mb. Consumption limit of 100% (default) means that a single channel can consume up-to all allocated queue's resources.
	ChannelConsumptionLimit *int `mandatory:"false" json:"channelConsumptionLimit"`

	// The list of capabilities enabled on the queue
	Capabilities []CapabilityDetails `mandatory:"false" json:"capabilities"`
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

// UnmarshalJSON unmarshals from json
func (m *Queue) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                  *string                           `json:"displayName"`
		LifecycleDetails             *string                           `json:"lifecycleDetails"`
		CustomEncryptionKeyId        *string                           `json:"customEncryptionKeyId"`
		FreeformTags                 map[string]string                 `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                   map[string]map[string]interface{} `json:"systemTags"`
		ChannelConsumptionLimit      *int                              `json:"channelConsumptionLimit"`
		Capabilities                 []capabilitydetails               `json:"capabilities"`
		Id                           *string                           `json:"id"`
		CompartmentId                *string                           `json:"compartmentId"`
		TimeCreated                  *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                  *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState               QueueLifecycleStateEnum           `json:"lifecycleState"`
		MessagesEndpoint             *string                           `json:"messagesEndpoint"`
		RetentionInSeconds           *int                              `json:"retentionInSeconds"`
		VisibilityInSeconds          *int                              `json:"visibilityInSeconds"`
		TimeoutInSeconds             *int                              `json:"timeoutInSeconds"`
		DeadLetterQueueDeliveryCount *int                              `json:"deadLetterQueueDeliveryCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.LifecycleDetails = model.LifecycleDetails

	m.CustomEncryptionKeyId = model.CustomEncryptionKeyId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ChannelConsumptionLimit = model.ChannelConsumptionLimit

	m.Capabilities = make([]CapabilityDetails, len(model.Capabilities))
	for i, n := range model.Capabilities {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Capabilities[i] = nn.(CapabilityDetails)
		} else {
			m.Capabilities[i] = nil
		}
	}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.MessagesEndpoint = model.MessagesEndpoint

	m.RetentionInSeconds = model.RetentionInSeconds

	m.VisibilityInSeconds = model.VisibilityInSeconds

	m.TimeoutInSeconds = model.TimeoutInSeconds

	m.DeadLetterQueueDeliveryCount = model.DeadLetterQueueDeliveryCount

	return
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
	QueueLifecycleStateInactive QueueLifecycleStateEnum = "INACTIVE"
)

var mappingQueueLifecycleStateEnum = map[string]QueueLifecycleStateEnum{
	"CREATING": QueueLifecycleStateCreating,
	"UPDATING": QueueLifecycleStateUpdating,
	"ACTIVE":   QueueLifecycleStateActive,
	"DELETING": QueueLifecycleStateDeleting,
	"DELETED":  QueueLifecycleStateDeleted,
	"FAILED":   QueueLifecycleStateFailed,
	"INACTIVE": QueueLifecycleStateInactive,
}

var mappingQueueLifecycleStateEnumLowerCase = map[string]QueueLifecycleStateEnum{
	"creating": QueueLifecycleStateCreating,
	"updating": QueueLifecycleStateUpdating,
	"active":   QueueLifecycleStateActive,
	"deleting": QueueLifecycleStateDeleting,
	"deleted":  QueueLifecycleStateDeleted,
	"failed":   QueueLifecycleStateFailed,
	"inactive": QueueLifecycleStateInactive,
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
		"INACTIVE",
	}
}

// GetMappingQueueLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueueLifecycleStateEnum(val string) (QueueLifecycleStateEnum, bool) {
	enum, ok := mappingQueueLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
