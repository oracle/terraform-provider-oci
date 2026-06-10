// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddonSummary Summary information about a KafkaClusterAddon.
type AddonSummary struct {

	// A user-friendly name.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the KafkaClusterAddon.
	LifecycleState KafkaClusterAddonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Addon Type
	AddonType KafkaClusterAddonAddonTypeEnum `mandatory:"true" json:"addonType"`

	// The date and time the KafkaClusterAddon was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the KafkaClusterAddon was updated, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the KafkaClusterAddon in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AddonSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddonSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterAddonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKafkaClusterAddonLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaClusterAddonAddonTypeEnum(string(m.AddonType)); !ok && m.AddonType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AddonType: %s. Supported values are: %s.", m.AddonType, strings.Join(GetKafkaClusterAddonAddonTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
