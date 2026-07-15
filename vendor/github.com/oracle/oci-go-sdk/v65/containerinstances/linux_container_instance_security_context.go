// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// LinuxContainerInstanceSecurityContext Security context for all containers in a container instance.
type LinuxContainerInstanceSecurityContext struct {

	// A special supplemental group that applies to all containers in the container instance. Some volume types allow the container instance to change ownership of the volume. The owning GID will be the fsGroup, the setgid bit will be set (new files will be owned by the fsGroup), and the permission bits are OR'd with rw-rw----. If unset, the container instance will not modify the ownership and permissions of volumes.
	FsGroup *int `mandatory:"false" json:"fsGroup"`

	// Defines behavior of changing ownership and permission of the volume before being exposed inside the containers. This only applies to volumes which support fsGroup ownership and permissions, and will have no effect on ephemeral volumes. ON_ROOT_MISMATCH only changes permissions and ownership if the permission and ownership of the root directory does not match the expected permissions and ownership of the volume. This can improve container instance start times. ALWAYS  changes permission and ownership of the volume when it is mounted. If unset, ALWAYS is used.
	FsGroupChangePolicy ContainerFsGroupChangePolicyTypeEnum `mandatory:"false" json:"fsGroupChangePolicy,omitempty"`
}

func (m LinuxContainerInstanceSecurityContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LinuxContainerInstanceSecurityContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerFsGroupChangePolicyTypeEnum(string(m.FsGroupChangePolicy)); !ok && m.FsGroupChangePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FsGroupChangePolicy: %s. Supported values are: %s.", m.FsGroupChangePolicy, strings.Join(GetContainerFsGroupChangePolicyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LinuxContainerInstanceSecurityContext) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLinuxContainerInstanceSecurityContext LinuxContainerInstanceSecurityContext
	s := struct {
		DiscriminatorParam string `json:"securityContextType"`
		MarshalTypeLinuxContainerInstanceSecurityContext
	}{
		"LINUX",
		(MarshalTypeLinuxContainerInstanceSecurityContext)(m),
	}

	return json.Marshal(&s)
}
