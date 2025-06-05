// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SnapSummary Provides summary information for a snap.
type SnapSummary struct {

	// The list of snap channel collections.
	Channels []SnapChannelCollection `mandatory:"true" json:"channels"`

	// The name of the snap.
	Name *string `mandatory:"true" json:"name"`

	// The publisher of the snap.
	Publisher *string `mandatory:"true" json:"publisher"`

	// The revision number of the snap channel.
	Revision *string `mandatory:"true" json:"revision"`

	// The tracking channel of the snap.
	Tracking *string `mandatory:"true" json:"tracking"`

	// The version of the snap.
	Version *string `mandatory:"true" json:"version"`

	// The description of of snap.
	Description *string `mandatory:"false" json:"description"`

	// If false, denotes the snap is not signed by the Snap Store.
	IsDangerous *bool `mandatory:"false" json:"isDangerous"`

	// The confinement mode for the snap.
	Mode SnapModesEnum `mandatory:"false" json:"mode,omitempty"`

	// The snap's store url.
	StoreUrl *string `mandatory:"false" json:"storeUrl"`

	// The date and time of the snap's last refresh.
	TimeRefresh *common.SDKTime `mandatory:"false" json:"timeRefresh"`
}

func (m SnapSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSnapModesEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetSnapModesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
