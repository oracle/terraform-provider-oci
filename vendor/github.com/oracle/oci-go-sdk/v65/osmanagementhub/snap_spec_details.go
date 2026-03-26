// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SnapSpecDetails Details about a specific snap.
type SnapSpecDetails struct {

	// The name of the snap to install.
	Name *string `mandatory:"true" json:"name"`

	// The channel to install from (e.g. stable, edge, beta, candidate, or a custom channel).
	Channel *string `mandatory:"false" json:"channel"`

	// If true, allows installing snaps not signed by the Snap Store. E.g., snaps from local file. Use with caution.
	IsSigned *bool `mandatory:"false" json:"isSigned"`

	// The modes for the snap.
	Mode SnapModesEnum `mandatory:"false" json:"mode,omitempty"`

	// The version of the snap.
	Version *string `mandatory:"false" json:"version"`

	// The revision to install.
	Revision *string `mandatory:"false" json:"revision"`
}

func (m SnapSpecDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapSpecDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSnapModesEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetSnapModesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
