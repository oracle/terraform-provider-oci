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

// SnapSummary Provides summary information for a snap.
type SnapSummary struct {

	// The name of the snap.
	Name *string `mandatory:"true" json:"name"`

	// The publisher of the snap.
	Publisher *string `mandatory:"true" json:"publisher"`

	// The revision number of the snap channel.
	Revision *string `mandatory:"true" json:"revision"`

	// The version of the snap.
	Version *string `mandatory:"true" json:"version"`

	// The track this snap is following.
	Tracking *string `mandatory:"true" json:"tracking"`

	// The description of of snap.
	Description *string `mandatory:"false" json:"description"`

	// The snap's store url.
	StoreUrl *string `mandatory:"false" json:"storeUrl"`

	// The date and time of the snap's last refresh in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format.
	TimeRefreshed *common.SDKTime `mandatory:"false" json:"timeRefreshed"`
}

func (m SnapSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
