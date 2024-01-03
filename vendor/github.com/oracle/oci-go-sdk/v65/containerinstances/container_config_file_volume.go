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

// ContainerConfigFileVolume The volume based on configuration files received during container creation.
type ContainerConfigFileVolume struct {

	// The name of the volume. This must be unique within a single container instance.
	Name *string `mandatory:"true" json:"name"`

	// Contains string key value pairs which can be mounted as individual files inside the container. The value needs to be base64 encoded. It is decoded to plain text before the mount.
	Configs []ContainerConfigFile `mandatory:"false" json:"configs"`
}

// GetName returns Name
func (m ContainerConfigFileVolume) GetName() *string {
	return m.Name
}

func (m ContainerConfigFileVolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerConfigFileVolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ContainerConfigFileVolume) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerConfigFileVolume ContainerConfigFileVolume
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeContainerConfigFileVolume
	}{
		"CONFIGFILE",
		(MarshalTypeContainerConfigFileVolume)(m),
	}

	return json.Marshal(&s)
}
