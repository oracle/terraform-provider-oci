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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CycleClusterNodeDetails The properties that define a node cycle action.
type CycleClusterNodeDetails struct {

	// The version of Kubernetes to which the node should be updated to.
	KubernetesVersion *string `mandatory:"false" json:"kubernetesVersion"`

	// A list of key/value pairs to update on the underlying OCI instance.
	NodeMetadata map[string]string `mandatory:"false" json:"nodeMetadata"`

	// Specify the source to update. Currently, image is the only supported source.
	NodeSourceDetails NodeSourceDetails `mandatory:"false" json:"nodeSourceDetails"`

	// The SSH public key to update on the underlying OCI instance.
	SshPublicKey *string `mandatory:"false" json:"sshPublicKey"`

	NodeEvictionSettings *NodeEvictionSettings `mandatory:"false" json:"nodeEvictionSettings"`

	// The mode of cycling that should be performed on the node.
	// Currently, BOOT_VOLUME_REPLACE is the only supported mode. It does the in-place update through
	// boot volume replacement.
	CycleMode CycleClusterNodeDetailsCycleModeEnum `mandatory:"false" json:"cycleMode,omitempty"`

	// If the cycling action should be performed when the node is already in-sync with this state
	IsCycleInSyncNode *bool `mandatory:"false" json:"isCycleInSyncNode"`
}

func (m CycleClusterNodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CycleClusterNodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCycleClusterNodeDetailsCycleModeEnum(string(m.CycleMode)); !ok && m.CycleMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CycleMode: %s. Supported values are: %s.", m.CycleMode, strings.Join(GetCycleClusterNodeDetailsCycleModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CycleClusterNodeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		KubernetesVersion    *string                              `json:"kubernetesVersion"`
		NodeMetadata         map[string]string                    `json:"nodeMetadata"`
		NodeSourceDetails    nodesourcedetails                    `json:"nodeSourceDetails"`
		SshPublicKey         *string                              `json:"sshPublicKey"`
		NodeEvictionSettings *NodeEvictionSettings                `json:"nodeEvictionSettings"`
		CycleMode            CycleClusterNodeDetailsCycleModeEnum `json:"cycleMode"`
		IsCycleInSyncNode    *bool                                `json:"isCycleInSyncNode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.KubernetesVersion = model.KubernetesVersion

	m.NodeMetadata = model.NodeMetadata

	nn, e = model.NodeSourceDetails.UnmarshalPolymorphicJSON(model.NodeSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NodeSourceDetails = nn.(NodeSourceDetails)
	} else {
		m.NodeSourceDetails = nil
	}

	m.SshPublicKey = model.SshPublicKey

	m.NodeEvictionSettings = model.NodeEvictionSettings

	m.CycleMode = model.CycleMode

	m.IsCycleInSyncNode = model.IsCycleInSyncNode

	return
}

// CycleClusterNodeDetailsCycleModeEnum Enum with underlying type: string
type CycleClusterNodeDetailsCycleModeEnum string

// Set of constants representing the allowable values for CycleClusterNodeDetailsCycleModeEnum
const (
	CycleClusterNodeDetailsCycleModeBootVolumeReplace CycleClusterNodeDetailsCycleModeEnum = "BOOT_VOLUME_REPLACE"
)

var mappingCycleClusterNodeDetailsCycleModeEnum = map[string]CycleClusterNodeDetailsCycleModeEnum{
	"BOOT_VOLUME_REPLACE": CycleClusterNodeDetailsCycleModeBootVolumeReplace,
}

var mappingCycleClusterNodeDetailsCycleModeEnumLowerCase = map[string]CycleClusterNodeDetailsCycleModeEnum{
	"boot_volume_replace": CycleClusterNodeDetailsCycleModeBootVolumeReplace,
}

// GetCycleClusterNodeDetailsCycleModeEnumValues Enumerates the set of values for CycleClusterNodeDetailsCycleModeEnum
func GetCycleClusterNodeDetailsCycleModeEnumValues() []CycleClusterNodeDetailsCycleModeEnum {
	values := make([]CycleClusterNodeDetailsCycleModeEnum, 0)
	for _, v := range mappingCycleClusterNodeDetailsCycleModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCycleClusterNodeDetailsCycleModeEnumStringValues Enumerates the set of values in String for CycleClusterNodeDetailsCycleModeEnum
func GetCycleClusterNodeDetailsCycleModeEnumStringValues() []string {
	return []string{
		"BOOT_VOLUME_REPLACE",
	}
}

// GetMappingCycleClusterNodeDetailsCycleModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCycleClusterNodeDetailsCycleModeEnum(val string) (CycleClusterNodeDetailsCycleModeEnum, bool) {
	enum, ok := mappingCycleClusterNodeDetailsCycleModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
