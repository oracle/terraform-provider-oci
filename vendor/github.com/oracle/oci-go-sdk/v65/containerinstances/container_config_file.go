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

// ContainerConfigFile The file that is mounted on a container instance through a volume mount.
type ContainerConfigFile struct {

	// The name of the file. The fileName should be unique across the volume.
	FileName *string `mandatory:"true" json:"fileName"`

	// The base64 encoded contents of the file. The contents are decoded to plain text before mounted as a file to a container inside container instance.
	Data []byte `mandatory:"true" json:"data"`

	// (Optional) Relative path for this file inside the volume mount directory. By default, the file is presented at the root of the volume mount path.
	Path *string `mandatory:"false" json:"path"`
}

func (m ContainerConfigFile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerConfigFile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
