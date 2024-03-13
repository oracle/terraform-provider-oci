// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeReplaceConfiguration The information about the NodeReplaceConfiguration.
type NodeReplaceConfiguration struct {

	// The unique identifier for the NodeReplaceConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the bdsInstance which is the parent resource id.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	LevelTypeDetails LevelTypeDetails `mandatory:"true" json:"levelTypeDetails"`

	// The state of the NodeReplaceConfiguration.
	LifecycleState NodeReplaceConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the NodeReplaceConfiguration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the NodeReplaceConfiguration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Type of compute instance health metric to use for node replacement
	MetricType NodeReplaceConfigurationMetricTypeEnum `mandatory:"true" json:"metricType"`

	// This value is the minimum period of time to wait for metric emission before triggering node replacement. The value is in minutes.
	DurationInMinutes *int `mandatory:"true" json:"durationInMinutes"`
}

func (m NodeReplaceConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeReplaceConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeReplaceConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeReplaceConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeReplaceConfigurationMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetNodeReplaceConfigurationMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *NodeReplaceConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                *string                                    `json:"id"`
		BdsInstanceId     *string                                    `json:"bdsInstanceId"`
		DisplayName       *string                                    `json:"displayName"`
		LevelTypeDetails  leveltypedetails                           `json:"levelTypeDetails"`
		LifecycleState    NodeReplaceConfigurationLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated       *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated       *common.SDKTime                            `json:"timeUpdated"`
		MetricType        NodeReplaceConfigurationMetricTypeEnum     `json:"metricType"`
		DurationInMinutes *int                                       `json:"durationInMinutes"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.BdsInstanceId = model.BdsInstanceId

	m.DisplayName = model.DisplayName

	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.MetricType = model.MetricType

	m.DurationInMinutes = model.DurationInMinutes

	return
}

// NodeReplaceConfigurationLifecycleStateEnum Enum with underlying type: string
type NodeReplaceConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for NodeReplaceConfigurationLifecycleStateEnum
const (
	NodeReplaceConfigurationLifecycleStateCreating NodeReplaceConfigurationLifecycleStateEnum = "CREATING"
	NodeReplaceConfigurationLifecycleStateActive   NodeReplaceConfigurationLifecycleStateEnum = "ACTIVE"
	NodeReplaceConfigurationLifecycleStateUpdating NodeReplaceConfigurationLifecycleStateEnum = "UPDATING"
	NodeReplaceConfigurationLifecycleStateDeleting NodeReplaceConfigurationLifecycleStateEnum = "DELETING"
	NodeReplaceConfigurationLifecycleStateDeleted  NodeReplaceConfigurationLifecycleStateEnum = "DELETED"
	NodeReplaceConfigurationLifecycleStateFailed   NodeReplaceConfigurationLifecycleStateEnum = "FAILED"
)

var mappingNodeReplaceConfigurationLifecycleStateEnum = map[string]NodeReplaceConfigurationLifecycleStateEnum{
	"CREATING": NodeReplaceConfigurationLifecycleStateCreating,
	"ACTIVE":   NodeReplaceConfigurationLifecycleStateActive,
	"UPDATING": NodeReplaceConfigurationLifecycleStateUpdating,
	"DELETING": NodeReplaceConfigurationLifecycleStateDeleting,
	"DELETED":  NodeReplaceConfigurationLifecycleStateDeleted,
	"FAILED":   NodeReplaceConfigurationLifecycleStateFailed,
}

var mappingNodeReplaceConfigurationLifecycleStateEnumLowerCase = map[string]NodeReplaceConfigurationLifecycleStateEnum{
	"creating": NodeReplaceConfigurationLifecycleStateCreating,
	"active":   NodeReplaceConfigurationLifecycleStateActive,
	"updating": NodeReplaceConfigurationLifecycleStateUpdating,
	"deleting": NodeReplaceConfigurationLifecycleStateDeleting,
	"deleted":  NodeReplaceConfigurationLifecycleStateDeleted,
	"failed":   NodeReplaceConfigurationLifecycleStateFailed,
}

// GetNodeReplaceConfigurationLifecycleStateEnumValues Enumerates the set of values for NodeReplaceConfigurationLifecycleStateEnum
func GetNodeReplaceConfigurationLifecycleStateEnumValues() []NodeReplaceConfigurationLifecycleStateEnum {
	values := make([]NodeReplaceConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingNodeReplaceConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeReplaceConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for NodeReplaceConfigurationLifecycleStateEnum
func GetNodeReplaceConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingNodeReplaceConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeReplaceConfigurationLifecycleStateEnum(val string) (NodeReplaceConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingNodeReplaceConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NodeReplaceConfigurationMetricTypeEnum Enum with underlying type: string
type NodeReplaceConfigurationMetricTypeEnum string

// Set of constants representing the allowable values for NodeReplaceConfigurationMetricTypeEnum
const (
	NodeReplaceConfigurationMetricTypeStatus              NodeReplaceConfigurationMetricTypeEnum = "INSTANCE_STATUS"
	NodeReplaceConfigurationMetricTypeAccessibilityStatus NodeReplaceConfigurationMetricTypeEnum = "INSTANCE_ACCESSIBILITY_STATUS"
)

var mappingNodeReplaceConfigurationMetricTypeEnum = map[string]NodeReplaceConfigurationMetricTypeEnum{
	"INSTANCE_STATUS":               NodeReplaceConfigurationMetricTypeStatus,
	"INSTANCE_ACCESSIBILITY_STATUS": NodeReplaceConfigurationMetricTypeAccessibilityStatus,
}

var mappingNodeReplaceConfigurationMetricTypeEnumLowerCase = map[string]NodeReplaceConfigurationMetricTypeEnum{
	"instance_status":               NodeReplaceConfigurationMetricTypeStatus,
	"instance_accessibility_status": NodeReplaceConfigurationMetricTypeAccessibilityStatus,
}

// GetNodeReplaceConfigurationMetricTypeEnumValues Enumerates the set of values for NodeReplaceConfigurationMetricTypeEnum
func GetNodeReplaceConfigurationMetricTypeEnumValues() []NodeReplaceConfigurationMetricTypeEnum {
	values := make([]NodeReplaceConfigurationMetricTypeEnum, 0)
	for _, v := range mappingNodeReplaceConfigurationMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeReplaceConfigurationMetricTypeEnumStringValues Enumerates the set of values in String for NodeReplaceConfigurationMetricTypeEnum
func GetNodeReplaceConfigurationMetricTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_STATUS",
		"INSTANCE_ACCESSIBILITY_STATUS",
	}
}

// GetMappingNodeReplaceConfigurationMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeReplaceConfigurationMetricTypeEnum(val string) (NodeReplaceConfigurationMetricTypeEnum, bool) {
	enum, ok := mappingNodeReplaceConfigurationMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
