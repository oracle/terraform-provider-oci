// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgRouteLimitUpdate Data plane information about DRG import/export route limits.
type DrgRouteLimitUpdate struct {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	Id *string `mandatory:"false" json:"id"`

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	IsDelete *bool `mandatory:"false" json:"isDelete"`

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The label given to the Drg Route Table that has the updated import or export limits.
	VrfLabel *int `mandatory:"false" json:"vrfLabel"`

	// The OCID for the Drg Route Table's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID for the Drg that this route table belongs to
	DrgId *string `mandatory:"false" json:"drgId"`

	// Limit on the number of dynamic route rules this DRG Route Table can export. This is
	// calculated based on the tenancy's per-DRG limit, the number of Route Tables in the DRG,
	// and the number of DRG Attachments on the DRG.
	ExportLimit *int `mandatory:"false" json:"exportLimit"`

	// Limit on the number of dynamic route rules this DRG Route Table can contain. This is
	// calculated based on the tenancy's per-DRG limit and the number of Route Tables in the DRG.
	ImportLimit *int `mandatory:"false" json:"importLimit"`

	// Max ECMP width
	EcmpLimit *int `mandatory:"false" json:"ecmpLimit"`

	// Common Export route target to use for the DRG Attachment instead of per-attachment export route target.
	// This is applicable to DRG Attachments that are assigned to a DRG Route Table which is whitelisted for high
	// throughput mode.
	CommonExportRouteTargetVC *int `mandatory:"false" json:"commonExportRouteTargetVC"`

	// to indicate whether the DRG Route Table is whitelisted for high throughput mode.
	IsHighThroughputModeEnabled *bool `mandatory:"false" json:"isHighThroughputModeEnabled"`
}

// GetId returns Id
func (m DrgRouteLimitUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgRouteLimitUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgRouteLimitUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgRouteLimitUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgRouteLimitUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgRouteLimitUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgRouteLimitUpdate DrgRouteLimitUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgRouteLimitUpdate
	}{
		"DrgRouteLimitUpdate",
		(MarshalTypeDrgRouteLimitUpdate)(m),
	}

	return json.Marshal(&s)
}
