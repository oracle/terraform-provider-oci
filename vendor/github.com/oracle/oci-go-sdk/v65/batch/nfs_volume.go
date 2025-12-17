// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NfsVolume A description of a NFS type of batch task environment volume.
type NfsVolume struct {

	// The name of the NfsVolume.
	Name *string `mandatory:"true" json:"name"`

	// The FQDN of the NFS server to connect to.
	MountTargetFqdn *string `mandatory:"true" json:"mountTargetFqdn"`

	// The path to the directory on the NFS server to be mounted.
	MountTargetExportPath *string `mandatory:"true" json:"mountTargetExportPath"`

	// The local path to mount the NFS share to.
	LocalMountDirectoryPath *string `mandatory:"true" json:"localMountDirectoryPath"`
}

func (m NfsVolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NfsVolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NfsVolume) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNfsVolume NfsVolume
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNfsVolume
	}{
		"NFS",
		(MarshalTypeNfsVolume)(m),
	}

	return json.Marshal(&s)
}
