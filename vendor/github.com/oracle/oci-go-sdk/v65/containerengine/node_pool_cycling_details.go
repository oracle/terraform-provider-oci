// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodePoolCyclingDetails Node Pool Cycling Details
type NodePoolCyclingDetails struct {

	// Maximum active nodes that would be terminated from nodepool during the cycling nodepool process.
	// OKE supports both integer and percentage input.
	// Defaults to 0, Ranges from 0 to Nodepool size or 0% to 100%
	MaximumUnavailable *string `mandatory:"false" json:"maximumUnavailable"`

	// Maximum additional new compute instances that would be temporarily created and added to nodepool during the cycling nodepool process.
	// OKE supports both integer and percentage input.
	// Defaults to 1, Ranges from 0 to Nodepool size or 0% to 100%
	MaximumSurge *string `mandatory:"false" json:"maximumSurge"`

	// If cycling operation should be performed on the nodes in the node pool.
	IsNodeCyclingEnabled *bool `mandatory:"false" json:"isNodeCyclingEnabled"`

	// The mode of cycling that should be performed on the OKE nodes.
	// BOOT_VOLUME_REPLACE cycling applies the updates in place, whereas INSTANCE_REPLACE cycling deletes and recreates a new node with the changes applied.
	// BOOT_VOLUME_REPLACE can update only a subset of node properties that are possible to update using INSTANCE_REPLACE.
	CycleMode NodePoolCyclingDetailsCycleModeEnum `mandatory:"false" json:"cycleMode,omitempty"`

	// The scope of cycling operation against nodes.
	// OUT_OF_SYNC_NODES means cycling will only target at out-of-sync nodes whose current state is different than the specified state, whereas ALL_NODES will target at all nodes.
	Scope NodePoolCyclingDetailsScopeEnum `mandatory:"false" json:"scope,omitempty"`
}

func (m NodePoolCyclingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodePoolCyclingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodePoolCyclingDetailsCycleModeEnum(string(m.CycleMode)); !ok && m.CycleMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CycleMode: %s. Supported values are: %s.", m.CycleMode, strings.Join(GetNodePoolCyclingDetailsCycleModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodePoolCyclingDetailsScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetNodePoolCyclingDetailsScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodePoolCyclingDetailsCycleModeEnum Enum with underlying type: string
type NodePoolCyclingDetailsCycleModeEnum string

// Set of constants representing the allowable values for NodePoolCyclingDetailsCycleModeEnum
const (
	NodePoolCyclingDetailsCycleModeBootVolumeReplace NodePoolCyclingDetailsCycleModeEnum = "BOOT_VOLUME_REPLACE"
	NodePoolCyclingDetailsCycleModeInstanceReplace   NodePoolCyclingDetailsCycleModeEnum = "INSTANCE_REPLACE"
)

var mappingNodePoolCyclingDetailsCycleModeEnum = map[string]NodePoolCyclingDetailsCycleModeEnum{
	"BOOT_VOLUME_REPLACE": NodePoolCyclingDetailsCycleModeBootVolumeReplace,
	"INSTANCE_REPLACE":    NodePoolCyclingDetailsCycleModeInstanceReplace,
}

var mappingNodePoolCyclingDetailsCycleModeEnumLowerCase = map[string]NodePoolCyclingDetailsCycleModeEnum{
	"boot_volume_replace": NodePoolCyclingDetailsCycleModeBootVolumeReplace,
	"instance_replace":    NodePoolCyclingDetailsCycleModeInstanceReplace,
}

// GetNodePoolCyclingDetailsCycleModeEnumValues Enumerates the set of values for NodePoolCyclingDetailsCycleModeEnum
func GetNodePoolCyclingDetailsCycleModeEnumValues() []NodePoolCyclingDetailsCycleModeEnum {
	values := make([]NodePoolCyclingDetailsCycleModeEnum, 0)
	for _, v := range mappingNodePoolCyclingDetailsCycleModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodePoolCyclingDetailsCycleModeEnumStringValues Enumerates the set of values in String for NodePoolCyclingDetailsCycleModeEnum
func GetNodePoolCyclingDetailsCycleModeEnumStringValues() []string {
	return []string{
		"BOOT_VOLUME_REPLACE",
		"INSTANCE_REPLACE",
	}
}

// GetMappingNodePoolCyclingDetailsCycleModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodePoolCyclingDetailsCycleModeEnum(val string) (NodePoolCyclingDetailsCycleModeEnum, bool) {
	enum, ok := mappingNodePoolCyclingDetailsCycleModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NodePoolCyclingDetailsScopeEnum Enum with underlying type: string
type NodePoolCyclingDetailsScopeEnum string

// Set of constants representing the allowable values for NodePoolCyclingDetailsScopeEnum
const (
	NodePoolCyclingDetailsScopeOutOfSyncNodes NodePoolCyclingDetailsScopeEnum = "OUT_OF_SYNC_NODES"
	NodePoolCyclingDetailsScopeAllNodes       NodePoolCyclingDetailsScopeEnum = "ALL_NODES"
)

var mappingNodePoolCyclingDetailsScopeEnum = map[string]NodePoolCyclingDetailsScopeEnum{
	"OUT_OF_SYNC_NODES": NodePoolCyclingDetailsScopeOutOfSyncNodes,
	"ALL_NODES":         NodePoolCyclingDetailsScopeAllNodes,
}

var mappingNodePoolCyclingDetailsScopeEnumLowerCase = map[string]NodePoolCyclingDetailsScopeEnum{
	"out_of_sync_nodes": NodePoolCyclingDetailsScopeOutOfSyncNodes,
	"all_nodes":         NodePoolCyclingDetailsScopeAllNodes,
}

// GetNodePoolCyclingDetailsScopeEnumValues Enumerates the set of values for NodePoolCyclingDetailsScopeEnum
func GetNodePoolCyclingDetailsScopeEnumValues() []NodePoolCyclingDetailsScopeEnum {
	values := make([]NodePoolCyclingDetailsScopeEnum, 0)
	for _, v := range mappingNodePoolCyclingDetailsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodePoolCyclingDetailsScopeEnumStringValues Enumerates the set of values in String for NodePoolCyclingDetailsScopeEnum
func GetNodePoolCyclingDetailsScopeEnumStringValues() []string {
	return []string{
		"OUT_OF_SYNC_NODES",
		"ALL_NODES",
	}
}

// GetMappingNodePoolCyclingDetailsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodePoolCyclingDetailsScopeEnum(val string) (NodePoolCyclingDetailsScopeEnum, bool) {
	enum, ok := mappingNodePoolCyclingDetailsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
