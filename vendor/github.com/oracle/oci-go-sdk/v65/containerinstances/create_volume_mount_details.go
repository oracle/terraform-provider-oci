// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVolumeMountDetails Defines the mapping from volume to a mount path in a container.
type CreateVolumeMountDetails struct {

	// The volume access path.
	MountPath *string `mandatory:"true" json:"mountPath"`

	// The name of the volume. Avoid entering confidential information.
	VolumeName *string `mandatory:"true" json:"volumeName"`

	// A subpath inside the referenced volume.
	SubPath *string `mandatory:"false" json:"subPath"`

	// Whether the volume was mounted in read-only mode. By default, the volume is not read-only.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// If there is more than one partition in the volume, reference this number of partitions.
	// Here is an example:
	// Number  Start   End     Size    File system  Name                  Flags
	// 1      1049kB  106MB   105MB   fat16        EFI System Partition  boot, esp
	// 2      106MB   1180MB  1074MB  xfs
	// 3      1180MB  50.0GB  48.8GB                                     lvm
	Partition *int `mandatory:"false" json:"partition"`
}

func (m CreateVolumeMountDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVolumeMountDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
