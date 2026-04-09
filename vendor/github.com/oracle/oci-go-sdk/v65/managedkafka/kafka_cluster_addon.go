// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KafkaClusterAddon The base data object to represent a KafkaClusterAddon.
type KafkaClusterAddon interface {

	// A unique user-friendly name.
	GetName() *string

	// The current state of the KafkaCluster.
	GetLifecycleState() KafkaClusterAddonLifecycleStateEnum

	// Description of the add on
	GetDescription() *string

	// The time the addon was created.
	GetTimeCreated() *common.SDKTime

	// The time the addon was updated.
	GetTimeUpdated() *common.SDKTime
}

type kafkaclusteraddon struct {
	JsonData       []byte
	Description    *string                             `mandatory:"false" json:"description"`
	TimeCreated    *common.SDKTime                     `mandatory:"false" json:"timeCreated"`
	TimeUpdated    *common.SDKTime                     `mandatory:"false" json:"timeUpdated"`
	Name           *string                             `mandatory:"true" json:"name"`
	LifecycleState KafkaClusterAddonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	AddonType      string                              `json:"addonType"`
}

// UnmarshalJSON unmarshals json
func (m *kafkaclusteraddon) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkafkaclusteraddon kafkaclusteraddon
	s := struct {
		Model Unmarshalerkafkaclusteraddon
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.LifecycleState = s.Model.LifecycleState
	m.Description = s.Model.Description
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.AddonType = s.Model.AddonType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *kafkaclusteraddon) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AddonType {
	case "PUBLICCONNECTIVITY":
		mm := PublicConnectivityAddon{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KafkaClusterAddon: %s.", m.AddonType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m kafkaclusteraddon) GetDescription() *string {
	return m.Description
}

// GetTimeCreated returns TimeCreated
func (m kafkaclusteraddon) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m kafkaclusteraddon) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetName returns Name
func (m kafkaclusteraddon) GetName() *string {
	return m.Name
}

// GetLifecycleState returns LifecycleState
func (m kafkaclusteraddon) GetLifecycleState() KafkaClusterAddonLifecycleStateEnum {
	return m.LifecycleState
}

func (m kafkaclusteraddon) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m kafkaclusteraddon) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterAddonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetKafkaClusterAddonLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KafkaClusterAddonLifecycleStateEnum Enum with underlying type: string
type KafkaClusterAddonLifecycleStateEnum string

// Set of constants representing the allowable values for KafkaClusterAddonLifecycleStateEnum
const (
	KafkaClusterAddonLifecycleStateCreating KafkaClusterAddonLifecycleStateEnum = "CREATING"
	KafkaClusterAddonLifecycleStateUpdating KafkaClusterAddonLifecycleStateEnum = "UPDATING"
	KafkaClusterAddonLifecycleStateActive   KafkaClusterAddonLifecycleStateEnum = "ACTIVE"
	KafkaClusterAddonLifecycleStateDeleting KafkaClusterAddonLifecycleStateEnum = "DELETING"
	KafkaClusterAddonLifecycleStateDeleted  KafkaClusterAddonLifecycleStateEnum = "DELETED"
	KafkaClusterAddonLifecycleStateFailed   KafkaClusterAddonLifecycleStateEnum = "FAILED"
)

var mappingKafkaClusterAddonLifecycleStateEnum = map[string]KafkaClusterAddonLifecycleStateEnum{
	"CREATING": KafkaClusterAddonLifecycleStateCreating,
	"UPDATING": KafkaClusterAddonLifecycleStateUpdating,
	"ACTIVE":   KafkaClusterAddonLifecycleStateActive,
	"DELETING": KafkaClusterAddonLifecycleStateDeleting,
	"DELETED":  KafkaClusterAddonLifecycleStateDeleted,
	"FAILED":   KafkaClusterAddonLifecycleStateFailed,
}

var mappingKafkaClusterAddonLifecycleStateEnumLowerCase = map[string]KafkaClusterAddonLifecycleStateEnum{
	"creating": KafkaClusterAddonLifecycleStateCreating,
	"updating": KafkaClusterAddonLifecycleStateUpdating,
	"active":   KafkaClusterAddonLifecycleStateActive,
	"deleting": KafkaClusterAddonLifecycleStateDeleting,
	"deleted":  KafkaClusterAddonLifecycleStateDeleted,
	"failed":   KafkaClusterAddonLifecycleStateFailed,
}

// GetKafkaClusterAddonLifecycleStateEnumValues Enumerates the set of values for KafkaClusterAddonLifecycleStateEnum
func GetKafkaClusterAddonLifecycleStateEnumValues() []KafkaClusterAddonLifecycleStateEnum {
	values := make([]KafkaClusterAddonLifecycleStateEnum, 0)
	for _, v := range mappingKafkaClusterAddonLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterAddonLifecycleStateEnumStringValues Enumerates the set of values in String for KafkaClusterAddonLifecycleStateEnum
func GetKafkaClusterAddonLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingKafkaClusterAddonLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterAddonLifecycleStateEnum(val string) (KafkaClusterAddonLifecycleStateEnum, bool) {
	enum, ok := mappingKafkaClusterAddonLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// KafkaClusterAddonAddonTypeEnum Enum with underlying type: string
type KafkaClusterAddonAddonTypeEnum string

// Set of constants representing the allowable values for KafkaClusterAddonAddonTypeEnum
const (
	KafkaClusterAddonAddonTypePublicconnectivity KafkaClusterAddonAddonTypeEnum = "PUBLICCONNECTIVITY"
)

var mappingKafkaClusterAddonAddonTypeEnum = map[string]KafkaClusterAddonAddonTypeEnum{
	"PUBLICCONNECTIVITY": KafkaClusterAddonAddonTypePublicconnectivity,
}

var mappingKafkaClusterAddonAddonTypeEnumLowerCase = map[string]KafkaClusterAddonAddonTypeEnum{
	"publicconnectivity": KafkaClusterAddonAddonTypePublicconnectivity,
}

// GetKafkaClusterAddonAddonTypeEnumValues Enumerates the set of values for KafkaClusterAddonAddonTypeEnum
func GetKafkaClusterAddonAddonTypeEnumValues() []KafkaClusterAddonAddonTypeEnum {
	values := make([]KafkaClusterAddonAddonTypeEnum, 0)
	for _, v := range mappingKafkaClusterAddonAddonTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaClusterAddonAddonTypeEnumStringValues Enumerates the set of values in String for KafkaClusterAddonAddonTypeEnum
func GetKafkaClusterAddonAddonTypeEnumStringValues() []string {
	return []string{
		"PUBLICCONNECTIVITY",
	}
}

// GetMappingKafkaClusterAddonAddonTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaClusterAddonAddonTypeEnum(val string) (KafkaClusterAddonAddonTypeEnum, bool) {
	enum, ok := mappingKafkaClusterAddonAddonTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
