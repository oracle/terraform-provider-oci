// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerEmptyDirVolume The empty directory volume of a container instance. You can create up to 64 EmptyDir per container instance.
type ContainerEmptyDirVolume struct {

	// The name of the volume. This must be unique within a single container instance.
	Name *string `mandatory:"true" json:"name"`

	// The volume type of the empty directory, can be either File Storage or Memory.
	BackingStore ContainerEmptyDirVolumeBackingStoreEnum `mandatory:"false" json:"backingStore,omitempty"`
}

// GetName returns Name
func (m ContainerEmptyDirVolume) GetName() *string {
	return m.Name
}

func (m ContainerEmptyDirVolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerEmptyDirVolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerEmptyDirVolumeBackingStoreEnum(string(m.BackingStore)); !ok && m.BackingStore != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackingStore: %s. Supported values are: %s.", m.BackingStore, strings.Join(GetContainerEmptyDirVolumeBackingStoreEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ContainerEmptyDirVolume) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerEmptyDirVolume ContainerEmptyDirVolume
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeContainerEmptyDirVolume
	}{
		"EMPTYDIR",
		(MarshalTypeContainerEmptyDirVolume)(m),
	}

	return json.Marshal(&s)
}

// ContainerEmptyDirVolumeBackingStoreEnum Enum with underlying type: string
type ContainerEmptyDirVolumeBackingStoreEnum string

// Set of constants representing the allowable values for ContainerEmptyDirVolumeBackingStoreEnum
const (
	ContainerEmptyDirVolumeBackingStoreEphemeralStorage ContainerEmptyDirVolumeBackingStoreEnum = "EPHEMERAL_STORAGE"
	ContainerEmptyDirVolumeBackingStoreMemory           ContainerEmptyDirVolumeBackingStoreEnum = "MEMORY"
)

var mappingContainerEmptyDirVolumeBackingStoreEnum = map[string]ContainerEmptyDirVolumeBackingStoreEnum{
	"EPHEMERAL_STORAGE": ContainerEmptyDirVolumeBackingStoreEphemeralStorage,
	"MEMORY":            ContainerEmptyDirVolumeBackingStoreMemory,
}

var mappingContainerEmptyDirVolumeBackingStoreEnumLowerCase = map[string]ContainerEmptyDirVolumeBackingStoreEnum{
	"ephemeral_storage": ContainerEmptyDirVolumeBackingStoreEphemeralStorage,
	"memory":            ContainerEmptyDirVolumeBackingStoreMemory,
}

// GetContainerEmptyDirVolumeBackingStoreEnumValues Enumerates the set of values for ContainerEmptyDirVolumeBackingStoreEnum
func GetContainerEmptyDirVolumeBackingStoreEnumValues() []ContainerEmptyDirVolumeBackingStoreEnum {
	values := make([]ContainerEmptyDirVolumeBackingStoreEnum, 0)
	for _, v := range mappingContainerEmptyDirVolumeBackingStoreEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerEmptyDirVolumeBackingStoreEnumStringValues Enumerates the set of values in String for ContainerEmptyDirVolumeBackingStoreEnum
func GetContainerEmptyDirVolumeBackingStoreEnumStringValues() []string {
	return []string{
		"EPHEMERAL_STORAGE",
		"MEMORY",
	}
}

// GetMappingContainerEmptyDirVolumeBackingStoreEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerEmptyDirVolumeBackingStoreEnum(val string) (ContainerEmptyDirVolumeBackingStoreEnum, bool) {
	enum, ok := mappingContainerEmptyDirVolumeBackingStoreEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
