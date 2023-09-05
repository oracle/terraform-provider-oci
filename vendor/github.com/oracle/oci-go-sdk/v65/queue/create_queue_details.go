// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateQueueDetails The information about a new queue.
type CreateQueueDetails struct {

	// The user-friendly name of the queue.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the queue.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The retention period of messages in the queue, in seconds.
	RetentionInSeconds *int `mandatory:"false" json:"retentionInSeconds"`

	// The default visibility timeout of the messages consumed from the queue, in seconds.
	VisibilityInSeconds *int `mandatory:"false" json:"visibilityInSeconds"`

	// The default polling timeout of the messages in the queue, in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// The percentage of allocated queue resources that can be consumed by a single channel. For example, if a queue has a storage limit of 2Gb, and a single channel consumption limit is 0.1 (10%), that means data size of a single channel  can't exceed 200Mb. Consumption limit of 100% (default) means that a single channel can consume up-to all allocated queue's resources.
	ChannelConsumptionLimit *int `mandatory:"false" json:"channelConsumptionLimit"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue. A value of 0 indicates that the DLQ is not used.
	DeadLetterQueueDeliveryCount *int `mandatory:"false" json:"deadLetterQueueDeliveryCount"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom encryption key to be used to encrypt messages content.
	CustomEncryptionKeyId *string `mandatory:"false" json:"customEncryptionKeyId"`

	// The capability to add on the queue
	Capabilities []CapabilityDetails `mandatory:"false" json:"capabilities"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateQueueDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateQueueDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateQueueDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RetentionInSeconds           *int                              `json:"retentionInSeconds"`
		VisibilityInSeconds          *int                              `json:"visibilityInSeconds"`
		TimeoutInSeconds             *int                              `json:"timeoutInSeconds"`
		ChannelConsumptionLimit      *int                              `json:"channelConsumptionLimit"`
		DeadLetterQueueDeliveryCount *int                              `json:"deadLetterQueueDeliveryCount"`
		CustomEncryptionKeyId        *string                           `json:"customEncryptionKeyId"`
		Capabilities                 []capabilitydetails               `json:"capabilities"`
		FreeformTags                 map[string]string                 `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{} `json:"definedTags"`
		DisplayName                  *string                           `json:"displayName"`
		CompartmentId                *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.RetentionInSeconds = model.RetentionInSeconds

	m.VisibilityInSeconds = model.VisibilityInSeconds

	m.TimeoutInSeconds = model.TimeoutInSeconds

	m.ChannelConsumptionLimit = model.ChannelConsumptionLimit

	m.DeadLetterQueueDeliveryCount = model.DeadLetterQueueDeliveryCount

	m.CustomEncryptionKeyId = model.CustomEncryptionKeyId

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
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	return
}
