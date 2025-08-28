// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InventoryRecordPatchDetails Details about an installed patch on the target
type InventoryRecordPatchDetails struct {

	// Name of the installed patch
	PatchName *string `mandatory:"true" json:"patchName"`

	// Description for the installed patch
	PatchDescription *string `mandatory:"true" json:"patchDescription"`

	// Date on which the patch was applied to the target
	TimeApplied *common.SDKTime `mandatory:"true" json:"timeApplied"`

	// Type of patch applied
	PatchType *string `mandatory:"true" json:"patchType"`

	// OCID of the installed patch
	PatchId *string `mandatory:"false" json:"patchId"`
}

func (m InventoryRecordPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InventoryRecordPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
