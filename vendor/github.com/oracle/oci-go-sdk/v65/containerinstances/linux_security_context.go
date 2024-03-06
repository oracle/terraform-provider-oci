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

// LinuxSecurityContext Security context for Linux container.
type LinuxSecurityContext struct {

	// The user ID (UID) to run the entrypoint process of the container. Defaults to user specified UID in container image metadata if not provided. This must be provided if runAsGroup is provided.
	RunAsUser *int `mandatory:"false" json:"runAsUser"`

	// The group ID (GID) to run the entrypoint process of the container. Uses runtime default if not provided.
	RunAsGroup *int `mandatory:"false" json:"runAsGroup"`

	// Indicates if the container must run as a non-root user. If true, the service validates the container image at runtime to ensure that it is not going to run with UID 0 (root) and fails the container instance creation if the validation fails.
	IsNonRootUserCheckEnabled *bool `mandatory:"false" json:"isNonRootUserCheckEnabled"`

	// Determines if the container will have a read-only root file system. Default value is false.
	IsRootFileSystemReadonly *bool `mandatory:"false" json:"isRootFileSystemReadonly"`

	Capabilities *ContainerCapabilities `mandatory:"false" json:"capabilities"`
}

func (m LinuxSecurityContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LinuxSecurityContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LinuxSecurityContext) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLinuxSecurityContext LinuxSecurityContext
	s := struct {
		DiscriminatorParam string `json:"securityContextType"`
		MarshalTypeLinuxSecurityContext
	}{
		"LINUX",
		(MarshalTypeLinuxSecurityContext)(m),
	}

	return json.Marshal(&s)
}
