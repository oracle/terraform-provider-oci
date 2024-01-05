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

// CreateContainerEmptyDirVolumeDetails The empty directory for the container instance.
type CreateContainerEmptyDirVolumeDetails struct {

	// The name of the volume. This must be unique within a single container instance.
	Name *string `mandatory:"true" json:"name"`

	// The volume type of the empty directory, can be either File Storage or Memory.
	BackingStore ContainerEmptyDirVolumeBackingStoreEnum `mandatory:"false" json:"backingStore,omitempty"`
}

// GetName returns Name
func (m CreateContainerEmptyDirVolumeDetails) GetName() *string {
	return m.Name
}

func (m CreateContainerEmptyDirVolumeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerEmptyDirVolumeDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateContainerEmptyDirVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateContainerEmptyDirVolumeDetails CreateContainerEmptyDirVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeCreateContainerEmptyDirVolumeDetails
	}{
		"EMPTYDIR",
		(MarshalTypeCreateContainerEmptyDirVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
